# Learn Go with Tests 🚀

My personal journey through the ["Learn Go with Tests"](https://quii.gitbook.io/learn-go-with-tests) course. This repo is my playground for practicing Go fundamentals using test-driven development.

## What's Inside

Each folder represents a different concept I'm learning:

- **helloworld/** - Starting with the basics: functions, constants, and testing
- **integers/** - Working with numbers and documentation
- **arrays/** - Collections, slices, and iteration
- **iteration/** - Loops and benchmarking
- More to come as I progress...

## Running the Code

```bash
# Run all tests
go test ./...

# Run tests for a specific package
go test ./arrays

# Run tests with coverage
go test -cover ./...

# Run benchmarks
go test -bench=. ./...
```

## My Learning Notes

This is purely for my own learning - expect messy commits, experimental code, and lots of "aha!" moments captured in git history. The beauty of TDD is writing the test first, watching it fail, then making it pass. It's oddly satisfying!

## Progress Tracker

- [x] Hello World
- [x] Integers
- [x] Arrays and Slices
- [x] Iteration
- [ ] Structs, methods & interfaces
- [ ] Pointers & errors
- [ ] Maps
- [ ] Dependency injection
- [ ] And much more...

---

_"The only way to learn a new programming language is by writing programs in it." - Dennis Ritchie_
