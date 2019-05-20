package main

import "fmt"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

func main() {
	done := make(chan struct{}, 5)
	go func() {
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()
	go func() {
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()
	go func() {
		Deposit(100)
		Deposit(100)
		done <- struct{}{}
	}()
	for i := 0; i < 5; i++ {
		<-done
	}
	fmt.Println("=", Balance())
	fmt.Println("=", Balance())
}
