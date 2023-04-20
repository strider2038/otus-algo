package goroh

import (
	"fmt"
	"strconv"
	"strings"
)

func Calculate(expression string) (string, error) {
	n0, n1, m0, m1, err := parse(expression)
	if err != nil {
		return "", err
	}

	d0 := n1
	d1 := m1
	d := d0 * d1

	n0 *= d1
	m0 *= d0

	r0 := n0 + m0
	r1 := d

	g := GCD(r0, r1)
	r0 /= g
	r1 /= g

	return fmt.Sprintf("%d/%d", r0, r1), nil
}

func parse(expression string) (int, int, int, int, error) {
	p := strings.Split(expression, "+")
	if len(p) != 2 {
		return 0, 0, 0, 0, fmt.Errorf("invalid expression")
	}

	n := strings.Split(p[0], "/")
	m := strings.Split(p[1], "/")
	if len(n) != 2 || len(m) != 2 {
		return 0, 0, 0, 0, fmt.Errorf("invalid expression")
	}

	n0, err := strconv.Atoi(strings.TrimSpace(n[0]))
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("not a number: %w", err)
	}
	n1, err := strconv.Atoi(strings.TrimSpace(n[1]))
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("not a number: %w", err)
	}

	m0, err := strconv.Atoi(strings.TrimSpace(m[0]))
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("not a number: %w", err)
	}
	m1, err := strconv.Atoi(strings.TrimSpace(m[1]))
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("not a number: %w", err)
	}

	return n0, n1, m0, m1, nil
}
