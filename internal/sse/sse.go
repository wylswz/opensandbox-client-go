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
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)

	var dataBuf strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "data:") {
			data := strings.TrimSpace(strings.TrimPrefix(line, "data:"))
			if data == "[DONE]" || data == "" {
				continue
			}
			dataBuf.Reset()
			dataBuf.WriteString(data)
			for scanner.Scan() {
				next := scanner.Text()
				if next == "" {
					break
				}
				if strings.HasPrefix(next, "data:") {
					if raw := dataBuf.String(); raw != "" {
						events = append(events, json.RawMessage([]byte(raw)))
					}
					dataBuf.Reset()
					dataBuf.WriteString(strings.TrimSpace(strings.TrimPrefix(next, "data:")))
					continue
				}
				dataBuf.WriteString("\n")
				dataBuf.WriteString(next)
			}
			if raw := dataBuf.String(); raw != "" {
				events = append(events, json.RawMessage([]byte(raw)))
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return events, err
	}
	return events, nil
}
