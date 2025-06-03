package main

import (
	"fmt"
	"math/big"
	"sync"
	"time"
)

func calculateFactorial(n int, result chan<- *big.Int, wg *sync.WaitGroup) {

	fmt.Printf("Starting factorial calculation for %d...\n", n)
	res := big.NewInt(1)
	for i := 2; i <= n; i++ {
		res.Mul(res, big.NewInt(int64(i)))
	}
	fmt.Printf("Factorial of %d is %s\n", n, res.String())
	time.Sleep(2 * time.Second)
	fmt.Printf("Finished factorial calculation for %d\n", n)
	result <- res
}

func findPrimes(count int, resultChan chan<- []int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Starting search for the first %d prime numbers...\n", count)
	var primes []int
	num := 2
	for len(primes) < count {
		if isPrime(num) {
			primes = append(primes, num)
		}
		num++
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("Finished searching for the first %d prime numbers.\n", count)

	resultChan <- primes
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {

}
