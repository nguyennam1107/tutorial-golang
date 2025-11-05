
// ============= BÀI 1: GOROUTINES BASICS =============

// TODO: Tạo function in ra số với goroutine
func printNumber(n int) {
	// Implement here
	// In ra số n và sleep 100ms
}

// TODO: So sánh sequential vs concurrent
func sequentialExecution() {
	// Implement here
	// Gọi printNumber tuần tự 10 lần
}

func concurrentExecution() {
	// Implement here
	// Gọi printNumber với goroutines 10 lần
	// Nhớ wait để goroutines hoàn thành
}

func exercise1() {
	fmt.Println("=== BÀI 1: GOROUTINES BASICS ===")

	// TODO: Đo thời gian sequential
	start := time.Now()
	sequentialExecution()
	seqTime := time.Since(start)
	fmt.Printf("Sequential time: %v\n", seqTime)

	// TODO: Đo thời gian concurrent
	start = time.Now()
	concurrentExecution()
	concTime := time.Since(start)
	fmt.Printf("Concurrent time: %v\n", concTime)
	fmt.Printf("Speedup: %.2fx\n", float64(seqTime)/float64(concTime))
}


// ============= BÀI 2: CHANNELS =============

// TODO: Producer - gửi numbers vào channel
func producer(ch chan<- int, count int) {
	// Implement here
	// Gửi số từ 1 đến count vào channel
	// Close channel khi done
}

// TODO: Consumer - nhận numbers từ channel
func consumer(ch <-chan int) {
	// Implement here
	// Nhận và in ra numbers từ channel
	// Sử dụng range để read until closed
}

// TODO: Buffered channel example
func bufferedChannelDemo() {
	// Implement here
	// Tạo buffered channel với capacity 5
	// Gửi 5 messages mà không cần goroutine
	// Đọc và in ra messages
}

func exercise2() {
	fmt.Println("\n=== BÀI 2: CHANNELS ===")

	// Test producer-consumer
	ch := make(chan int)
	go producer(ch, 10)
	consumer(ch)

	// Test buffered channel
	bufferedChannelDemo()
}


// ============= BÀI 3: WORKER POOL =============

// TODO: Worker function
func worker(id int, jobs <-chan int, results chan<- int) {
	// Implement here
	// Nhận jobs từ channel
	// Process: square the number
	// Gửi result vào results channel
	// In ra: "Worker {id} processing job {job}"
}

// TODO: Create worker pool
func createWorkerPool(numWorkers int, jobs <-chan int, results chan<- int) {
	// Implement here
	// Start numWorkers goroutines
	// Mỗi worker gọi worker()
}

func exercise3() {
	fmt.Println("\n=== BÀI 3: WORKER POOL ===")

	// TODO: Setup
	numJobs := 20
	numWorkers := 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// TODO: Start worker pool
	createWorkerPool(numWorkers, jobs, results)

	// TODO: Send jobs
	// Gửi 20 jobs vào channel
	
	// TODO: Close jobs channel

	// TODO: Collect results
	// Nhận và in ra 20 results
}


// ============= BÀI 4: SELECT STATEMENT =============

// TODO: Select từ nhiều channels
func multipleChannelsSelect() {
	// Implement here
	c1 := make(chan string)
	c2 := make(chan string)

	// Goroutine 1: gửi "message 1" sau 1 giây
	
	// Goroutine 2: gửi "message 2" sau 2 giây
	
	// Sử dụng select để nhận từ c1 hoặc c2
	// Nhận 2 messages
}

// TODO: Timeout với select
func selectWithTimeout() {
	// Implement here
	ch := make(chan string)

	// Goroutine gửi message sau 2 giây
	
	// Select với timeout 1 giây
	// In "Timeout!" nếu timeout
	// In message nếu nhận được
}

// TODO: Non-blocking select
func nonBlockingSelect() {
	// Implement here
	messages := make(chan string)
	signals := make(chan bool)

	// Try receive từ messages (không block)
	// Try send vào signals (không block)
	// Sử dụng default case
}

func exercise4() {
	fmt.Println("\n=== BÀI 4: SELECT STATEMENT ===")

	fmt.Println("Multiple channels:")
	multipleChannelsSelect()

	fmt.Println("\nTimeout:")
	selectWithTimeout()

	fmt.Println("\nNon-blocking:")
	nonBlockingSelect()
}


// ============= BÀI 5: SYNC PACKAGE =============

// TODO: WaitGroup example
func waitGroupExample() {
	// Implement here
	var wg sync.WaitGroup

	// Start 5 goroutines
	// Mỗi goroutine in ra số của nó
	// Sử dụng wg.Add(), wg.Done(), wg.Wait()
}

// TODO: Mutex để fix race condition
var (
	counter int
	mu      sync.Mutex
)

func incrementWithoutMutex(wg *sync.WaitGroup) {
	// Implement here
	// Increment counter 1000 lần
	// KHÔNG dùng mutex (sẽ có race condition)
}

func incrementWithMutex(wg *sync.WaitGroup) {
	// Implement here
	// Increment counter 1000 lần
	// SỬ DỤNG mutex để protect counter
}

func exercise5() {
	fmt.Println("\n=== BÀI 5: SYNC PACKAGE ===")

	fmt.Println("WaitGroup example:")
	waitGroupExample()

	// Test WITHOUT mutex (race condition)
	fmt.Println("\nWithout Mutex (có race condition):")
	counter = 0
	var wg1 sync.WaitGroup
	// TODO: Start 10 goroutines gọi incrementWithoutMutex
	// Mỗi goroutine increment 1000 lần
	// Expected: 10000, Actual: < 10000 (do race)

	// Test WITH mutex (safe)
	fmt.Println("\nWith Mutex (safe):")
	counter = 0
	var wg2 sync.WaitGroup
	// TODO: Start 10 goroutines gọi incrementWithMutex
	// Expected: 10000, Actual: 10000
}


// ============= BÀI 6: PIPELINE PATTERN =============

// TODO: Stage 1 - Generate numbers
func generate(nums ...int) <-chan int {
	// Implement here
	// Tạo channel và goroutine
	// Gửi tất cả nums vào channel
	// Close channel khi done
	// Return channel
	return nil
}

// TODO: Stage 2 - Square numbers
func square(in <-chan int) <-chan int {
	// Implement here
	// Nhận từ in channel
	// Square mỗi số
	// Gửi vào out channel
	// Close out channel khi in closed
	// Return out channel
	return nil
}

// TODO: Stage 3 - Filter even numbers
func filterEven(in <-chan int) <-chan int {
	// Implement here
	// Nhận từ in channel
	// Chỉ gửi số chẵn vào out channel
	// Return out channel
	return nil
}

func exercise6() {
	fmt.Println("\n=== BÀI 6: PIPELINE PATTERN ===")

	// TODO: Build pipeline
	// numbers := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// squared := square(numbers)
	// evens := filterEven(squared)

	// TODO: Consume results
	// for num := range evens {
	//     fmt.Println(num)
	// }
}


// ============= BÀI 7: FAN-OUT FAN-IN =============

// TODO: Generate numbers
func generate7(nums ...int) <-chan int {
	// Implement here
	return nil
}

// TODO: Square numbers (worker)
func square7(in <-chan int) <-chan int {
	// Implement here
	// Mỗi worker process từ in channel
	return nil
}

// TODO: Merge nhiều channels thành 1 (fan-in)
func merge(channels ...<-chan int) <-chan int {
	// Implement here
	var wg sync.WaitGroup
	out := make(chan int)

	// TODO: Start goroutine cho mỗi channel
	// Đọc từ channel và gửi vào out
	
	// TODO: Close out khi tất cả channels done
	
	return out
}

func exercise7() {
	fmt.Println("\n=== BÀI 7: FAN-OUT FAN-IN ===")

	// TODO: Setup pipeline
	// in := generate7(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// TODO: Fan-out - start 3 workers
	// c1 := square7(in)
	// c2 := square7(in)
	// c3 := square7(in)

	// TODO: Fan-in - merge results
	// results := merge(c1, c2, c3)

	// TODO: Consume results
	// for num := range results {
	//     fmt.Println(num)
	// }
}


// ============= BÀI 8: CONTEXT =============

// TODO: Context với timeout
func contextWithTimeout() {
	// Implement here
	// Tạo context với timeout 1 giây
	
	// Goroutine thực hiện task mất 2 giây
	
	// Select giữa done channel và context.Done()
}

// TODO: Context với cancel
func contextWithCancel() {
	// Implement here
	// Tạo context với cancel
	
	// Goroutine chạy vô hạn
	// Lắng nghe ctx.Done()
	
	// Cancel sau 2 giây
}

// TODO: Context với values
func contextWithValues() {
	// Implement here
	// Tạo context với values
	ctx := context.WithValue(context.Background(), "userID", 123)
	ctx = context.WithValue(ctx, "requestID", "abc-123")
	
	// Function nhận context và in values
	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	// TODO: Lấy values từ context
	// userID := ctx.Value("userID")
	// requestID := ctx.Value("requestID")
	// In ra values
}

func exercise8() {
	fmt.Println("\n=== BÀI 8: CONTEXT ===")

	fmt.Println("Context with timeout:")
	contextWithTimeout()

	fmt.Println("\nContext with cancel:")
	contextWithCancel()

	fmt.Println("\nContext with values:")
	contextWithValues()
}


// ============= CHALLENGE 1: CONCURRENT WEB SCRAPER =============

/*
YÊU CẦU:
1. Scrape 10 URLs concurrently
2. Sử dụng worker pool pattern (5 workers)
3. Handle timeout (3 seconds per URL)
4. Collect và in results
5. Handle errors gracefully
*/

type ScrapeResult struct {
	URL     string
	Title   string
	Content string
	Error   error
}

// TODO: Simulate fetch URL
func fetchURL(url string) (string, error) {
	// Implement here
	// Simulate HTTP request với time.Sleep
	// Return fake content hoặc error
	return "", nil
}

// TODO: Worker function
func scrapeWorker(id int, urls <-chan string, results chan<- ScrapeResult) {
	// Implement here
	// Nhận URL từ channel
	// Fetch với timeout
	// Gửi result vào results channel
}

func challenge1() {
	fmt.Println("\n=== CHALLENGE 1: WEB SCRAPER ===")

	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://stackoverflow.com",
		"https://reddit.com",
		"https://medium.com",
		"https://dev.to",
		"https://news.ycombinator.com",
		"https://lobste.rs",
		"https://twitter.com",
		"https://facebook.com",
	}

	// TODO: Setup channels
	// urlChan := make(chan string, len(urls))
	// results := make(chan ScrapeResult, len(urls))

	// TODO: Start worker pool (5 workers)

	// TODO: Send URLs

	// TODO: Collect results

	// TODO: Print summary
}


// ============= CHALLENGE 2: RATE LIMITER =============

/*
YÊU CẦU:
1. Implement rate limiter: 5 requests per second
2. Queue requests khi vượt limit
3. Process requests theo rate limit
4. Test với 20 requests
*/

// TODO: Rate limiter với channel
type RateLimiter struct {
	// Implement here
	// Sử dụng ticker channel
}

func NewRateLimiter(requestsPerSecond int) *RateLimiter {
	// Implement here
	return nil
}

func (rl *RateLimiter) Allow() {
	// Implement here
	// Block cho đến khi có slot available
}

// TODO: Test function
func makeRequest(id int) {
	fmt.Printf("Request %d at %v\n", id, time.Now().Format("15:04:05.000"))
	time.Sleep(100 * time.Millisecond) // Simulate work
}

func challenge2() {
	fmt.Println("\n=== CHALLENGE 2: RATE LIMITER ===")

	// TODO: Create rate limiter (5 req/sec)
	// limiter := NewRateLimiter(5)

	// TODO: Make 20 requests
	// Observe rate limiting in action
}
