# 🔍 Go Escape Analysis

> Demonstrating stack vs heap allocation performance in Go

## Quick Start

```bash
# Run benchmarks
go test -bench=. -benchmem

# See escape analysis
go build -gcflags="-m -l" main.go

# Run demo
go run main.go
```

## The Code

```go
type User struct {
    ID      int
    Age     int
    Balance float64
    Active  bool
    Data    [8192]byte  // 8KB struct
}

// Stack allocation
func ByValue(id int) User {
    return User{ID: id, Age: 30, Balance: 1000.0, Active: true}
}

// Heap allocation
func ByPointer(id int) *User {
    return &User{ID: id, Age: 30, Balance: 1000.0, Active: true}
}
```

## Results

```
BenchmarkSingleByValue-12       6650589     158.8 ns/op        0 B/op    0 allocs/op
BenchmarkSingleByPointer-12      1793710     680.8 ns/op     9472 B/op    1 allocs/op
```

> **Note**: The `-12` suffix indicates this benchmark ran on a 12-core CPU. Your results may vary depending on your machine's specifications.

| Metric | Stack | Heap | Impact |
|--------|-------|------|--------|
| Speed | 158.8 ns | 680.8 ns | **4.3x faster** ⚡ |
| Memory | 0 B | 9,472 B | **Zero GC** 💾 |
| Allocs | 0 | 1 | **No pressure** 🧹 |

## Escape Analysis

```bash
$ go build -gcflags="-m -l" main.go

./main.go:25:9: &User{...} escapes to heap
```

Returning a pointer forces heap allocation. Returning by value uses stack.

## Real-World Impact

- **API Servers**: Significantly reduces CPU time per request
- **Batch Jobs**: Completes several times faster with large datasets
- **Stream Processing**: Eliminates heap allocations, reducing GC pressure

## When to Use What

**Use Values (Stack):**
- Small/medium structs
- High-frequency operations
- Performance critical paths

**Use Pointers (Heap):**
- Very large structs
- Need to share/mutate
- Interface requirements
