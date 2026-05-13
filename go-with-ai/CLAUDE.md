# General Guidelines

- Assume you are an expert developer
- Write the simplest code to solve the problem at hand with minimal code changes
- Prefer code clarity above all
- Use the standard library as much as possible
- Respect common idioms in the language you're producing code to and in the current code base
- Avoid mocking in tests
- Use primitive data types as must as possible
- Layer code in a way it's built as a library and possible APIs (such as web, command line ...)
- Use the "typos" tool to check for spelling mistakes
- Only comment non-obvious parts of the code. Prefer self explaining code.
- Never commit secrets or large binary files
- Write code with less indentation as possible, prefer early exit
- Don't add yourself to commit messages
- Do not commit or push unless explicitly asked to. If the task involves code changes, stop after the code is working.
- When creating Dockerfiles pin base image versions (never use `latest`)
- Follow https://www.conventionalcommits.org/ in commit messages
- Add a comment with example input next to every regular expression you write
