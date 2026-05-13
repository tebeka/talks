# Using AI Agents to write Go code, the Good, the Bad and the Ugly

- About
    - GopherCon Israel sponsors
- What is LLM?
    - Predict next word, not reasoning
    - Randomness and hallucination built in
- Car wash example
- Have no memory
- Agent is LLM with access to tools
    - claude code
    - codex
    - opencode
    - amp
- Story on read-only access
- Go is great
- Code
    - concat.go
    - show codex solve it
- Plan mode vs edit mode
    - Plan: width, think: depth
- Objective, Technical Requirements, Data Models, Edge Cases
    - What, not how
    - Make it testable
- CLAUDE.md/AGENTS.md
    - project, local, system
    - Will not follow it at times
    - Show mine
- Skills
    - LGTM: Looks good to Miki
- Context engineering: Providing the right data to your agent
    - No memory
        - Manage context
        - Claude 200K, keep it at 75%
- /new, /compact, double ESC
- Random things


---

Take away: Plan & Review

cognetive offset vs surrender

resume

- XKCD waiting 
    - https://www.reddit.com/r/xkcd/comments/1l5q7en/thinking/
    - llm.png
- Why Go is good language? (typing, simple, error handing, fast feedback)
    - Wes and Armin moving to Go
    - Wes ergonomics article
- What's an agent (vs LLM)
- Predict next token
    - https://bjornaustraat.substack.com/p/llms-are-amazing-but-theres-still
    - Randomness, hallucination
- You still need to understand code
    - Write a Go function that concatenates two string slices, show only the code.
        - concat.go
- CLAUDE.md/AGENTS.md
    - project, local
- Skills 
- Plan mode vs edit mode
    - Plan: width, think: depth
- Objective, Technical Requirements, Data Models, Edge Cases
- What, not how
- Make it testable
- Context engineering: Providing the right data to your agent
    - No memory
        - Manage context
        - Claude 200K, keep it at 75%
    - PreToolUse hook to prevent reading .env
    - /new, /compact, double ESC
- Go LSP server
- worktree
- /simplify
- Using agents to review code (a look at roborev)
    - roborev

        You are a code reviewer. Review the code for:
        - bugs
        - security
        - testing gaps
        - regressions
        - code quality

        Print only final review, wait until all tools are done before emitting review.
        For every issue found, print the following information:
        - Location
        - Severity
        - Problem
        - Suggested fix

- /plugin marketplace add JetBrains/go-modern-guidelines
/plugin install modern-go-guidelines


- claude -r to resume, then /export to save the prompt
- commit a lot

