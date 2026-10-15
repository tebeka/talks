# CLI Design for Humans and AI: Building Dual-Audience Tools in Go

What happens when an AI agent runs your CLI tool? If your application relies on interactive prompts or raw stdout dumps, the answer is usually broken loops and hung processes.

In this talk, we'll build a CLI from scratch in Go, engineering it for dual consumption. We’ll cover why Go is an ideal language for building command line tools and how simple architectural tweaks can transform a standard terminal utility into an agent-friendly execution engine.

What you'll learn:
- How to automatically switch between human UI (colors, spinners) and machine output (JSON/NDJSON) using TTY checks.
- Techniques for stream separation (stdout vs. stderr) and token-efficient output filtering.
- How to structure CLI errors so AI agents can self-correct and retry autonomously.
