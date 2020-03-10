# Going Open Source
[AppsFlyer meetup][meetup], March 18 2020

Talk on open sourcing a Go project.

## Abstract

You wrote some cool code and decided to open source it so others can enjoy it. But... how?

In this talk, I'll start by trying to convince you not to open source your code. After failing to do so, we'll cover several topics that will help you succeed in open sourcing your code.

We'll cover code structure, picking a licence, testing & CI, versioning, packaging, documentation and more. Some bots might make an appearance as well.

## Topics
- Why?
    - Fame (personal brand)
    - Help the community
    - Fortune (help company - nuclio)
- Why Not?
    - Time
    - Maintenance
	- fs story, Brad story
	- Quote from bolt
    - Mean people out there
- Pick a project (selenium? )
- Code structure
    - lib or exe?
- Licence
    - IANAL
    - https://github.blog/2015-03-09-open-source-license-usage-on-github-com/
- Documentation
    - README
	- shields
    - Documenting
	- Example test
- Dependencies
    - How to pick
    - License
- Versioning 
    - tag
    - semver
- Tests & CI
    - golangci-lint/staticcheck
    - gosec
- Issues & PRs
- Upload artifacts (CI?)
    - Docker 
- Bots
    - dependabot
    - rivi
    - security (gosec?)

[meetup]: https://www.meetup.com/AppsFlyer/events/268055610/
