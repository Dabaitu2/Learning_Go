package bank3
import "sync"

var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	defer mu.Unlock()
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	// Unlock会在return语句读取完balance的值之后执行，所以Balance函数是并发安全的
	// 整个函数段都被加锁了
	return balance
}