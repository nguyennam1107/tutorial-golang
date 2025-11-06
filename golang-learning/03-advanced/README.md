# 03. Advanced - Nâng Cao

## 📋 Nội dung

1. **Reflection** - Reflection và Meta-programming
2. **Generics** - Type Parameters (Go 1.18+)
3. **Advanced Concurrency** - Patterns nâng cao
4. **Design Patterns** - Gang of Four patterns in Go
5. **Performance** - Profiling & Optimization
6. **Memory Management** - GC, sync.Pool, escape analysis
7. **Code Generation** - go generate, protobuf
8. **Unsafe Package** - Low-level programming
9. **CGo** - C interop
10. **Build & Deploy** - Cross-compilation, plugins

## 🎯 Mục tiêu học

Sau khi hoàn thành phần này, bạn sẽ:
- Hiểu sâu về reflection và khi nào sử dụng
- Áp dụng generics để viết code reusable
- Master advanced concurrency patterns
- Implement design patterns phổ biến
- Profile và optimize performance
- Hiểu Go memory model
- Build production-ready applications

## 📝 Kiến thức cần có

Trước khi học phần này, bạn cần:
- ✅ Hoàn thành Basics (01-basics)
- ✅ Hoàn thành Intermediate (02-intermediate)
- ✅ Thành thạo goroutines & channels
- ✅ Hiểu interfaces và methods

## 🚀 Roadmap

### Week 1: Reflection & Generics
- [ ] Reflection basics
- [ ] Struct tags & metadata
- [ ] Generics introduction
- [ ] Type constraints
- [ ] Generic data structures

**Bài tập:**
- JSON serializer với reflection
- Generic Stack/Queue/LinkedList
- Type-safe cache với generics

### Week 2: Design Patterns
- [ ] Creational patterns (Singleton, Factory, Builder)
- [ ] Structural patterns (Adapter, Decorator, Facade)
- [ ] Behavioral patterns (Observer, Strategy, Command)
- [ ] Go-specific patterns

**Bài tập:**
- Implement 5+ design patterns
- Plugin system với Factory
- Event system với Observer

### Week 3: Performance & Optimization
- [ ] Benchmarking
- [ ] CPU profiling
- [ ] Memory profiling
- [ ] Trace profiling
- [ ] Optimization techniques

**Bài tập:**
- Benchmark common operations
- Optimize slow code
- Memory leak detection
- Profile real application

### Week 4: Advanced Topics
- [ ] Memory management deep dive
- [ ] Code generation
- [ ] Unsafe package
- [ ] CGo basics
- [ ] Build optimization

**Projects:**
- Custom serializer
- High-performance cache
- Code generator tool

## 📚 Topics Chi Tiết

### 1. Reflection

Reflection cho phép inspect và modify structures tại runtime.

**Khi nào dùng:**
- JSON/XML encoding/decoding
- ORM frameworks
- Dependency injection
- Testing frameworks

**Best practices:**
- Tránh dùng khi không cần thiết (performance overhead)
- Cache reflection results
- Validate types trước khi modify

### 2. Generics

Type parameters giúp viết code reusable mà vẫn type-safe.

**Use cases:**
- Data structures (Stack, Queue, Tree)
- Algorithms (Sort, Search)
- Utilities (Map, Filter, Reduce)

**Guidelines:**
- Dùng khi logic giống nhau cho nhiều types
- Tránh over-engineering
- Prefer interfaces khi có thể

### 3. Advanced Concurrency

Patterns phức tạp cho concurrent programming.

**Patterns:**
- Rate limiting
- Circuit breaker
- Bulkhead
- Retry với backoff
- Distributed semaphore

### 4. Design Patterns

Classic patterns adapted cho Go.

**Important patterns:**
- **Singleton**: Shared instance
- **Factory**: Object creation
- **Builder**: Complex object construction
- **Observer**: Event notification
- **Strategy**: Algorithm selection

### 5. Performance

Measure, profile, optimize.

**Tools:**
- `go test -bench`
- `pprof`
- `trace`
- `gcvis`

**Optimization areas:**
- Algorithm efficiency
- Memory allocations
- Goroutine overhead
- Lock contention

## ⚡ Performance Tips

### Memory
```go
// Bad: Many allocations
func badConcat(strs []string) string {
    result := ""
    for _, s := range strs {
        result += s // New allocation each time
    }
    return result
}

// Good: Single allocation
func goodConcat(strs []string) string {
    var builder strings.Builder
    builder.Grow(len(strs) * 10) // Pre-allocate
    for _, s := range strs {
        builder.WriteString(s)
    }
    return builder.String()
}
```

### Concurrency
```go
// Bad: Too many goroutines
for i := 0; i < 1000000; i++ {
    go process(i)
}

// Good: Worker pool
jobs := make(chan int, 1000000)
for w := 0; w < 100; w++ {
    go worker(jobs)
}
for i := 0; i < 1000000; i++ {
    jobs <- i
}
```

### Slice pre-allocation
```go
// Bad
var result []int
for i := 0; i < 1000; i++ {
    result = append(result, i)
}

// Good
result := make([]int, 0, 1000)
for i := 0; i < 1000; i++ {
    result = append(result, i)
}
```

## 🔧 Tools

### Profiling
```bash
# CPU profiling
go test -cpuprofile=cpu.prof -bench=.
go tool pprof cpu.prof

# Memory profiling
go test -memprofile=mem.prof -bench=.
go tool pprof mem.prof

# Trace
go test -trace=trace.out
go tool trace trace.out
```

### Benchmarking
```bash
# Run benchmarks
go test -bench=. -benchmem

# Compare benchmarks
go test -bench=. -benchmem > old.txt
# Make changes
go test -bench=. -benchmem > new.txt
benchstat old.txt new.txt
```

## 📖 Reading List

### Books
- "The Go Programming Language" - Donovan & Kernighan
- "Concurrency in Go" - Katherine Cox-Buday
- "100 Go Mistakes and How to Avoid Them" - Teiva Harsanyi

### Articles
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Memory Model](https://go.dev/ref/mem)
- [Go Performance](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html)

### Videos
- [GopherCon talks](https://www.youtube.com/c/GopherAcademy)
- [JustForFunc](https://www.youtube.com/c/JustForFunc)

## 🎓 Projects

### Project 1: JSON Schema Validator
- Parse JSON schema
- Validate JSON documents
- Use reflection extensively
- Handle all JSON types

### Project 2: Generic Cache Library
- Thread-safe cache
- TTL support
- LRU eviction
- Type-safe với generics
- Benchmarks

### Project 3: Rate Limiter Service
- Multiple algorithms (Token bucket, Leaky bucket)
- Distributed support
- Redis backend
- High performance
- Complete tests

### Project 4: Performance Monitor
- CPU/Memory profiling
- Metrics collection
- Dashboard (simple HTTP)
- Alerts on thresholds

## ✅ Checklist

Hoàn thành được khi:
- [ ] Hiểu khi nào dùng reflection
- [ ] Viết được generic data structures
- [ ] Implement 5+ design patterns
- [ ] Profile và optimize code
- [ ] Hiểu Go memory model
- [ ] Build production-ready app
- [ ] Pass tất cả exercises
- [ ] Complete 2+ projects

## 🚨 Common Pitfalls

### 1. Over-using Reflection
```go
// Bad: Reflection cho simple task
func getValue(i interface{}) {
    v := reflect.ValueOf(i)
    // Complex reflection code
}

// Good: Type assertion
func getValue(i interface{}) {
    if s, ok := i.(string); ok {
        // Use s
    }
}
```

### 2. Premature Optimization
- Measure first!
- Profile to find bottlenecks
- Optimize hot paths only

### 3. Ignoring Escape Analysis
```go
// Bad: Escapes to heap
func bad() *int {
    x := 42
    return &x // x escapes to heap
}

// Good: Stack allocation
func good() int {
    return 42
}
```

### 4. Lock Contention
```go
// Bad: Single mutex for all operations
var mu sync.Mutex
var data map[string]string

// Good: Sharded locks
type ShardedMap struct {
    shards []*Shard
}

type Shard struct {
    mu   sync.RWMutex
    data map[string]string
}
```

## 🔥 Advanced Challenges

### Challenge 1: Build Your Own ORM
- Struct to SQL mapping
- CRUD operations
- Reflection-based
- Migrations support
- Transactions

### Challenge 2: Distributed Cache
- Consistent hashing
- Replication
- Failover
- Monitoring
- Client library

### Challenge 3: High-Performance JSON Parser
- Faster than encoding/json
- Zero allocation
- Streaming support
- Benchmarks vs standard lib

## 💡 Tips

1. **Learn by doing**: Build real projects
2. **Read source code**: Study stdlib và popular libraries
3. **Profile everything**: Data beats intuition
4. **Test thoroughly**: Unit, integration, benchmark tests
5. **Stay updated**: Follow Go blog và release notes

## 🎯 Next Steps

After completing advanced:
1. **Backend Development** (04-backend)
2. **Microservices**
3. **Cloud Native Go**
4. **Production Systems**

---

**Good luck! 🚀**

Phần Advanced khó hơn nhưng sẽ giúp bạn trở thành Go expert!
