package datatesting

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

var ErrNotEnoughArguments = errors.New("not enough arguments")

type Solver interface {
	Solve(t *testing.T, input, output []string)
}

type SolverFunc func(t *testing.T, input, output []string)

func (f SolverFunc) Solve(t *testing.T, input, output []string) {
	t.Helper()
	f(t, input, output)
}

var defaultRunner = NewRunner()

func Run(t *testing.T, solver Solver) {
	t.Helper()
	defaultRunner.Run(t, solver)
}

type Runner struct {
	workdir   string
	separator string
	limit     int
}

func NewRunner(options ...Option) *Runner {
	runner := &Runner{
		workdir:   "./testdata/",
		separator: "\r\n",
	}

	for _, set := range options {
		set(runner)
	}

	return runner
}

type Option func(r *Runner)

func WithWorkdir(workdir string) Option {
	return func(r *Runner) {
		r.workdir = workdir
	}
}

func WithLimit(limit int) Option {
	return func(r *Runner) {
		r.limit = limit
	}
}

func (r *Runner) Run(t *testing.T, solver Solver) {
	t.Helper()
	problem, err := os.ReadFile(r.workdir + "problem.txt")
	if err == nil {
		t.Log("ЗАДАЧА.\n", string(problem))
	} else if !errors.Is(err, os.ErrNotExist) {
		t.Fatal("open testdata dir:", err)
	}

	for i := 0; ; i++ {
		if r.limit > 0 && i >= r.limit {
			t.Logf("test limit %d reached", i)
			break
		}

		inputFilename := fmt.Sprintf("%stest.%d.in", r.workdir, i)
		outputFilename := fmt.Sprintf("%stest.%d.out", r.workdir, i)

		inputData, err := os.ReadFile(inputFilename)
		if errors.Is(err, os.ErrNotExist) {
			break
		}
		if err != nil {
			t.Fatalf(`open input file "%s": %s`, inputFilename, err)
		}
		outputData, err := os.ReadFile(outputFilename)
		if err != nil {
			t.Fatalf(`open output file "%s": %s`, outputFilename, err)
		}

		input := r.parseStrings(inputData)
		if len(input) < 1 {
			t.Fatalf(`empty input`)
		}
		output := r.parseStrings(outputData)

		t.Run(fmt.Sprintf("test %d (%s)", i, strings.TrimSpace(input[0])), func(t *testing.T) {
			solver.Solve(t, input, output)
		})
	}
}

func (r *Runner) parseStrings(input []byte) []string {
	args := strings.Split(string(input), r.separator)
	if len(args) > 0 && args[len(args)-1] == "" {
		return args[:len(args)-1]
	}
	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}

	return args
}

func ParseIntArray(input string) ([]int, error) {
	var err error

	rawNumbers := strings.Split(input, " ")
	numbers := make([]int, len(rawNumbers))
	for i, rawNumber := range rawNumbers {
		numbers[i], err = strconv.Atoi(rawNumber)
		if err != nil {
			return nil, fmt.Errorf("parse number at %d: %w", i, err)
		}
	}

	return numbers, nil
}
