# GitHub Copilot Go Demos 🧠🚀

A curated set of demos showcasing how to use [GitHub Copilot](https://github.com/features/copilot) effectively in Go projects.  
These examples are designed for internal team training and demonstrate real-world use cases across development, refactoring, testing, documentation, and debugging.

---

## 📚 Contents

| Demo | Description |
|------|-------------|
| **Demo 1** | Refactor function with error handling |
| **Demo 2** | Translate a Python class to idiomatic Go |
| **Demo 3** | Generate unit tests for a Go function |
| **Demo 4** | Identify and fix infinite loop (RCA) |
| **Demo 5** | Refactor monolithic function into helpers |
| **Demo 6** | Use Copilot for code review & mentoring |
| **Demo 7** | Use `/doc` & `@workspace` for documentation and project context |

---

## 🧪 Getting Started

```bash
git clone https://github.com/<your-username>/github-copilot-go-demos.git
cd github-copilot-go-demos
go mod init github.com/<your-username>/copilotdemos
```

You can run individual demos using:
```bash
cd demo1_refactor_function
go run main.go
```

---

## 🖥️ Training Material

The slide deck used for internal training is included:

- `github-copilot-go-demos.pptx`: Walkthrough of all 7 demos with speaker notes

---

## 🧑‍🏫 Tips for Using Copilot in Go

- Use clear and focused comments to guide suggestions
- Prefer `/doc` for function-level documentation
- Use `@workspace` to leverage project-wide context
- Apply the 4 S’s of prompting: **Single**, **Specific**, **Short**, **Surround**

---

## 🪪 License

[MIT License](LICENSE)

---

## 📣 Feedback

Feel free to open issues or contribute enhancements to these examples. This is a live learning space!

---

## 🔗 References

- [GitHub Copilot Docs](https://docs.github.com/en/copilot)
- [Effective Go](https://go.dev/doc/effective_go)
