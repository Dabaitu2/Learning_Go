package main

import (
	"fmt"
	"sync"
)

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var wg sync.WaitGroup //定义一个同步等待的组

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

func main()  {
	wg.Add(5)
	go func() {
		fmt.Println("=", Balance())
		wg.Done() //减去一个计数
	}()
	go func() {
		fmt.Println("=", Balance())
		wg.Done() //减去一个计数
	}()
	go func() {
		Deposit(100)
		wg.Done() //减去一个计数
	}()
	go func() {
		Deposit(100)
		wg.Done() //减去一个计数
	}()
	go func() {
		Deposit(100)
		Deposit(100)
		wg.Done() //减去一个计数
	}()
	wg.Wait() //阻塞直到所有任务完成
	fmt.Println("=", Balance())
}