package main

import (
	"logs/parser"

	"github.com/google/cel-go/cel"
)

var (
	env *cel.Env
)

func init() {
	var err error
	env, err = cel.NewEnv(
		cel.Variable("Host", cel.StringType),
		cel.Variable("Time", cel.TimestampType),
		cel.Variable("Request", cel.StringType),
		cel.Variable("Status", cel.IntType),
		cel.Variable("Bytes", cel.IntType),
	)

	if err != nil {
		panic(err)
	}
}

func logMap(log parser.Log) map[string]any {
	return map[string]any{
		"Host":    log.Host,
		"Time":    log.Time,
		"Request": log.Request,
		"Status":  log.Status,
		"Bytes":   log.Bytes,
	}
}

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
