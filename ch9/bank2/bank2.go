package main

var (
	sema    = make(chan struct{})
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{}
	balance += amount
	<-sema
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}
