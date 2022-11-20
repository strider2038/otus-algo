package datatesting

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

var ErrNotEnoughArguments = errors.New("not enough arguments")

type Solver interface {
	Solve(input, output []string) error
}

type SolverFunc func(input, output []string) error

func (f SolverFunc) Solve(input, output []string) error {
	return f(input, output)
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

		t.Run(fmt.Sprintf("test %d (%s)", i, strings.TrimSpace(string(input))), func(t *testing.T) {
			start := time.Now()
			defer func() {
				t.Log("elapsed time:", time.Since(start).String())
			}()

			err := solver.Solve(r.parseStrings(input), r.parseStrings(output))
			if err != nil {
				t.Error(err)
			}
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
