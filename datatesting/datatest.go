package datatesting

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

var (
	ErrEmptyInput = errors.New("empty input")
)

type Solver interface {
	Solve(input []string) (string, error)
}

type SolverFunc func(input []string) (string, error)

func (f SolverFunc) Solve(input []string) (string, error) {
	return f(input)
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
	if err != nil {
		t.Fatal("open testdata dir:", err)
	}
	t.Log("ЗАДАЧА.\n", string(problem))

	for i := 0; ; i++ {
		inputFilename := fmt.Sprintf("%stest.%d.in", r.workdir, i)
		outputFilename := fmt.Sprintf("%stest.%d.out", r.workdir, i)

		input, err := os.ReadFile(inputFilename)
		if errors.Is(err, os.ErrNotExist) {
			break
		}
		if err != nil {
			t.Fatalf(`open input file "%s": %s`, inputFilename, err)
		}
		output, err := os.ReadFile(outputFilename)
		if err != nil {
			t.Fatalf(`open output file "%s": %s`, outputFilename, err)
		}

		if r.limit > 0 && i >= r.limit {
			t.Logf("test limit %d reached", i)
			break
		}

		t.Run(fmt.Sprintf("test %d: %s", i, string(input)), func(t *testing.T) {
			start := time.Now()
			defer func() {
				t.Log("elapsed time:", time.Since(start).String())
			}()

			actual, err := solver.Solve(r.parseInput(input))
			if err != nil {
				t.Error(err)
			} else if actual != strings.TrimSpace(string(output)) {
				t.Errorf(`fail: want "%s", got "%s"`, output, actual)
			}
		})
	}
}

func (r *Runner) parseInput(input []byte) []string {
	args := strings.Split(string(input), r.separator)
	if len(args) > 0 && args[len(args)-1] == "" {
		return args[:len(args)-1]
	}

	return args
}
