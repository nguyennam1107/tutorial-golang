# Bài Tập Thực Hành - Intermediate

## Bài 1: Goroutines Basics
1. Tạo 10 goroutines in ra số từ 1-10
2. Sử dụng time.Sleep để sync goroutines
3. So sánh thời gian chạy sequential vs concurrent

## Bài 2: Channels
1. Tạo channel để gửi/nhận messages giữa goroutines
2. Implement producer-consumer pattern
3. Sử dụng buffered channel với capacity 5

## Bài 3: Worker Pool
1. Tạo worker pool với 5 workers
2. Process 20 jobs concurrently
3. Collect và in kết quả

## Bài 4: Select Statement
1. Sử dụng select để đọc từ nhiều channels
2. Implement timeout với time.After
3. Non-blocking channel operations

## Bài 5: Sync Package
1. Sử dụng sync.WaitGroup để đợi goroutines
2. Sử dụng sync.Mutex để protect shared data
3. Fix race condition trong counter

## Bài 6: Pipeline Pattern
1. Tạo pipeline: generate -> square -> filter
2. Chain multiple stages với channels
3. Close channels properly

## Bài 7: Fan-Out Fan-In
1. Fan-out: Phân phối work cho nhiều workers
2. Fan-in: Merge results từ nhiều channels
3. Process large dataset concurrently

## Bài 8: Context
1. Sử dụng context.WithTimeout
2. Sử dụng context.WithCancel
3. Pass values qua context

## 🎯 Challenges

### Challenge 1: Concurrent Web Scraper
- Scrape 10 URLs concurrently
- Use worker pool pattern
- Handle errors và timeouts
- Collect results

### Challenge 2: Rate Limiter
- Implement rate limiter với channels
- Limit: 5 requests per second
- Queue requests khi vượt limit

### Challenge 3: Concurrent File Processor
- Đọc nhiều files concurrently
- Process nội dung (word count, line count)
- Merge results và in report

### Challenge 4: Real-time Data Aggregator
- Simulate data stream với goroutines
- Aggregate data mỗi 1 giây
- Display real-time statistics

## 📝 Testing

Chạy với race detector để tìm race conditions:
```bash
go run -race bai1.go
```

## 💡 Tips

- Luôn close channels khi không dùng nữa
- Sử dụng WaitGroup để sync goroutines
- Dùng select cho non-blocking operations
- Tránh goroutine leaks
- Test với `-race` flag
