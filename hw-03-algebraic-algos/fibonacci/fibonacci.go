package fibonacci

import (
	"fmt"
	"math/big"
)

func Recursive(n int) *big.Int {
	if n <= 0 {
		return big.NewInt(0)
	}
	if n <= 2 {
		return big.NewInt(1)
	}

	f := new(big.Int)
	f.Add(Recursive(n-1), Recursive(n-2))

	return f
}

func Iterative(n int) *big.Int {
	if n <= 0 {
		return big.NewInt(0)
	}
	if n <= 2 {
		return big.NewInt(1)
	}
	a := big.NewInt(1)
	b := big.NewInt(1)

	for i := 2; i < n; i++ {
		next := new(big.Int)
		next = next.Add(a, b)
		b = a
		a = next
	}

	return a
}

func ByMatrix(n int) *big.Int {
	if n <= 0 {
		return big.NewInt(0)
	}
	if n <= 2 {
		return big.NewInt(1)
	}

	m := Matrix{
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(0)},
	}

	m = m.Pow(n - 1)

	return m[0][0]
}

type Matrix [2][2]*big.Int

func (m Matrix) String() string {
	return fmt.Sprintf(
		"((%s, %s), (%s, %s))",
		m[0][0].String(),
		m[0][1].String(),
		m[1][0].String(),
		m[1][1].String(),
	)
}

func (m Matrix) Mul(m2 Matrix) Matrix {
	return Matrix{
		{
			new(big.Int).Add(new(big.Int).Mul(m[0][0], m2[0][0]), new(big.Int).Mul(m[0][1], m2[1][0])),
			new(big.Int).Add(new(big.Int).Mul(m[0][0], m2[0][1]), new(big.Int).Mul(m[0][1], m2[1][1])),
		},
		{
			new(big.Int).Add(new(big.Int).Mul(m[1][0], m2[0][0]), new(big.Int).Mul(m[1][1], m2[1][0])),
			new(big.Int).Add(new(big.Int).Mul(m[1][0], m2[0][1]), new(big.Int).Mul(m[1][1], m2[1][1])),
		},
	}
}

func (m Matrix) Pow(n int) Matrix {
	d := m
	pow := Matrix{
		{big.NewInt(1), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(1)},
	}

	for n > 0 {
		if n&1 == 1 {
			pow = pow.Mul(d)
		}
		n = n >> 1
		d = d.Mul(d)
	}

	return pow
}
