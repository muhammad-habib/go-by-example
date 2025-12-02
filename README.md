# Go by Example

> A collection of practical Go examples demonstrating key concepts, patterns, and performance considerations.

## Examples

### [01 - Escape Analysis](./01-escape-analysis)

Understanding when Go allocates on stack vs heap and the performance implications.

**Key Concepts:**
- Stack vs Heap allocation
- Escape analysis
- Performance benchmarking
- 4.3x speed difference demonstrated

**Topics Covered:**
- Return by value vs pointer
- Memory allocation patterns
- GC pressure considerations

---

## Structure

Each example follows this structure:

```
XX-topic-name/
├── README.md       # Detailed explanation and results
├── main.go         # Runnable demo code
└── main_test.go    # Benchmarks and tests
```

## Running Examples

```bash
# Run a specific example
cd 01-escape-analysis
go run main.go

# Run benchmarks
go test -bench=. -benchmem

# See detailed analysis (example-specific)
go build -gcflags="-m -l" main.go
```

## Contributing

Feel free to suggest new examples or improvements!

## License

MIT

