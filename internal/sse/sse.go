// Package sse provides pure Server-Sent Events parsing over HTTP.
// It has no knowledge of auth, headers, or domain types—the caller builds
// the request and provides the shared HTTP client.
package sse

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Stream executes req and calls fn for each SSE data event as it arrives in real time.
// Unlike Do, events are not buffered — fn is called immediately upon receipt.
// Caller is responsible for building req with method, URL, body, and all headers.
func Stream(ctx context.Context, client *http.Client, req *http.Request, fn func(json.RawMessage)) error {
	if client == nil {
		client = http.DefaultClient
	}
	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%s: %s", resp.Status, string(body))
	}

	ct := resp.Header.Get("Content-Type")
	if !strings.Contains(ct, "text/event-stream") {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected content-type %q: %s", ct, string(body))
	}

	return scanStream(resp.Body, fn)
}

// Do executes req with client and parses the response as text/event-stream.
// Returns the raw JSON payload of each data event. Caller is responsible for
// building req with method, URL, body, and all headers (auth, content-type, etc.).
func Do(ctx context.Context, client *http.Client, req *http.Request) ([]json.RawMessage, error) {
	if client == nil {
		client = http.DefaultClient
	}
	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("%s: %s", resp.Status, string(body))
	}

	ct := resp.Header.Get("Content-Type")
	if !strings.Contains(ct, "text/event-stream") {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected content-type %q: %s", ct, string(body))
	}

	return ParseStream(resp.Body)
}

// ParseStream reads SSE events from r and returns the raw JSON data of each event.
// SSE format: "data: {json}\n\n" per event. Skips empty data and "[DONE]".
func ParseStream(r io.Reader) ([]json.RawMessage, error) {
	var events []json.RawMessage
	err := scanStream(r, func(raw json.RawMessage) {
		events = append(events, raw)
	})
	return events, err
}

// scanStream parses event-stream bodies and emits each payload via emit.
// It supports two wire formats seen in OpenSandbox servers:
// 1) Canonical SSE: lines prefixed by "data:" and separated by blank lines.
// 2) JSON-lines-over-event-stream: raw JSON lines separated by blank lines.
func scanStream(r io.Reader, emit func(json.RawMessage)) error {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)

	var buf strings.Builder
	flush := func() {
		raw := strings.TrimSpace(buf.String())
		buf.Reset()
		if raw == "" || raw == "[DONE]" {
			return
		}
		emit(json.RawMessage([]byte(raw)))
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			flush()
			continue
		}
		if strings.HasPrefix(line, ":") {
			// SSE comment / heartbeat.
			continue
		}
		if strings.HasPrefix(line, "event:") || strings.HasPrefix(line, "id:") || strings.HasPrefix(line, "retry:") {
			continue
		}

		var payload string
		if strings.HasPrefix(line, "data:") {
			payload = strings.TrimSpace(strings.TrimPrefix(line, "data:"))
		} else {
			// Non-standard but observed: JSON line without "data:" prefix.
			payload = line
		}
		if payload == "" || payload == "[DONE]" {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(payload)
	}

	// Flush final event even if stream ends without trailing blank line.
	flush()
	return scanner.Err()
}
