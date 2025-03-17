package main

import (
	"logs/parser"

	"github.com/google/cel-go/cel"
)

var (
	env *cel.Env
)

func init() {
	// Initialize log environment.
	var err error
	env, err = cel.NewEnv(
		cel.Variable("host", cel.StringType),
		cel.Variable("time", cel.TimestampType),
		cel.Variable("request", cel.StringType),
		cel.Variable("status", cel.IntType),
		cel.Variable("length", cel.IntType),
	)

	if err != nil {
		panic(err)
	}
}

// logMap converts a Log to a map.
func logMap(log parser.Log) map[string]any {
	return map[string]any{
		"host":    log.Host,
		"time":    log.Time,
		"request": log.Request,
		"status":  log.Status,
		"length":  log.Bytes,
	}
}

// createFilter creates a filter from logs from expr.
func createFilter(expr string) (func(parser.Log) bool, error) {
	ast, issues := env.Compile(expr)
	if issues != nil && issues.Err() != nil {
		return nil, issues.Err()
	}

	prog, err := env.Program(ast)
	if err != nil {
		return nil, err
	}

	fn := func(log parser.Log) bool {
		out, _, err := prog.Eval(logMap(log))
		if err != nil {
			return false
		}

		v, ok := out.Value().(bool)
		if !ok {
			return false
		}
		return v
	}

	return fn, nil
}
