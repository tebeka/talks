---
paths:
    - "**/*.go"
    - "go.mod"
---

- When scanning code, ignore testdata directory and any file or directory starting with and underscore (_)
- Interfaces should not have more than 4 methods
- Avoid using pointers as much as possible
- Avoid generics as much as possible
- Use `go run` to check that the code compiles (*not* build and then run)
- Add empty line after closing } for if/switch/select .... For example:
    if !ok {
        return fmt.Errorf("failed")
    }

    saveUser(user)


- Default to latest Go version when starting a project
