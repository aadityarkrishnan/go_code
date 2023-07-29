package main

import "fmt"

func sieveOfEratosthenes(limit int) []int {
	// Create a boolean array "prime[0..limit]" and initialize
	// all entries as true. A value in prime[i] will finally
	// be false if i is Not a prime, and true otherwise.
	prime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		prime[i] = true
	}

	// Run the Sieve of Eratosthenes algorithm
	for p := 2; p*p <= limit; p++ {
		// If prime[p] is not changed, then it is a prime
		if prime[p] == true {
			// Update all multiples of p
			for i := p * p; i <= limit; i += p {
				fmt.Println("Value of I", i)

				prime[i] = false
			}
		}
	}

	// Create a list of prime numbers
	primeNumbers := []int{}
	for p := 2; p <= limit; p++ {
		if prime[p] == true {
			primeNumbers = append(primeNumbers, p)
		}
	}

	return primeNumbers
}

func main() {
	// Define the upper limit
	limit := 10

	// Find prime numbers up to the given limit
	primes := sieveOfEratosthenes(limit)

	// Display the prime numbers
	fmt.Println("Prime numbers up to", limit, "are:")
	fmt.Println(primes)
}
