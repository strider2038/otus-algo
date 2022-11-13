package prime

import (
	"math"
	"math/bits"
)

// CountByBruteForce - алгоритм поиска количества простых чисел через перебор делителей, O(N^2).
func CountByBruteForce(n int) int {
	count := 0

	for i := 2; i <= n; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			isPrime = isPrime && i%j != 0
		}
		if isPrime {
			count++
		}
	}

	return count
}

// CountByBruteForceOptimized - алгоритм поиска количества простых чисел через
// перебор делителей с применением оптимизаций.
func CountByBruteForceOptimized(n int) int {
	count := 0

	for i := 1; i <= n; i++ {
		if isPrime(i) {
			count++
		}
	}

	return count
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// CountByPrimes - алгоритм поиска простых чисел с оптимизациями поиска
// и делением только на простые числа, O(N * Sqrt(N) / Ln (N)).
func CountByPrimes(n int) int {
	if n <= 1 {
		return 0
	}

	count := 0
	primes := make([]int, n/2+1)
	primes[0] = 2
	count++

	for i := 3; i <= n; i += 2 {
		if isPrimeWithPrimes(i, primes) {
			primes[count] = i
			count++
		}
	}

	return count
}

func isPrimeWithPrimes(n int, primes []int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 0; primes[i] <= int(math.Sqrt(float64(n))); i++ {
		if n%primes[i] == 0 {
			return false
		}
	}

	return true
}

// CountBySieveOfEratosthenes - алгоритм "Решето Эратосфена" для быстрого поиска
// простых чисел O(N Log Log N).
func CountBySieveOfEratosthenes(n int) int {
	if n <= 1 {
		return 0
	}

	primes := NewBoolSieve(n)

	for i := 2; i*i <= n; i++ {
		if primes.IsSet(i) {
			for j := i * i; j <= n; j += i {
				primes.Unset(j)
			}
		}
	}

	return primes.Count()
}

// CountBySieveOfEratosthenesOptimized - алгоритм "Решето Эратосфена" с оптимизацией памяти,
// с использованием битовой матрицы, с сохранением по 32 значения в одном int,
// биты хранятся только для нечётных чисел.
func CountBySieveOfEratosthenesOptimized(n int) int {
	if n <= 1 {
		return 0
	}

	primes := NewBitSieve(n)

	for i := 2; i*i <= n; i++ {
		if primes.IsSet(i) {
			for j := i * i; j <= n; j += i {
				primes.Unset(j)
			}
		}
	}

	return primes.Count()
}

type BoolSieve []bool

func NewBoolSieve(n int) BoolSieve {
	sieve := make(BoolSieve, n-1)

	for i := 0; i < n-1; i++ {
		sieve[i] = true
	}

	return sieve
}

func (s BoolSieve) IsSet(i int) bool {
	return s[i-2]
}

func (s BoolSieve) Unset(i int) {
	s[i-2] = false
}

func (s BoolSieve) Count() int {
	count := 0

	for i := 0; i < len(s); i++ {
		if s[i] {
			count++
		}
	}

	return count
}

type BitSieve []uint32

func NewBitSieve(n int) BitSieve {
	sieve := make(BitSieve, n/64+1)

	for i := 0; i < len(sieve)-1; i++ {
		sieve[i] = 0xFFFFFFFF
	}

	// заполнение битами последнего блока
	m := n % 64
	b := uint32(1)
	last := uint32(0)
	for m > 0 {
		last = last | b
		b = b << 1
		m -= 2
	}

	sieve[len(sieve)-1] = last

	return sieve
}

func (s BitSieve) IsSet(i int) bool {
	if i&1 == 0 {
		return i == 2
	}

	index := i / 64
	bitIndex := (i % 64) / 2

	return s[index]&(1<<bitIndex) != 0
}

func (s BitSieve) Unset(i int) {
	if i&1 == 0 {
		return
	}

	index := i / 64
	bitIndex := (i % 64) / 2

	s[index] = s[index] & ^(1 << bitIndex)
}

func (s BitSieve) Count() int {
	count := 0

	for i := 0; i < len(s); i++ {
		// быстрый подсчет числа бит в 32-разрядном числе
		count += bits.OnesCount32(s[i])
	}

	return count
}
