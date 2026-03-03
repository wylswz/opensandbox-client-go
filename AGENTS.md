# Agents

Guidelines for building the Go client and examples.

## Core Principles

- Keep the client **rich**: include ergonomic, developer-focused APIs rather than only transport/unmarshal wrappers.
- Prefer reusable abstractions when implementing features used across commands, examples, or workflows.
- While developing examples, identify common patterns and promote them into first-class client APIs.

## Implementation Rules

- Put built binaries under `bin/`.
- Prefer named constants over repeated string literals.
- Keep examples realistic and production-oriented so they validate API design quality.
