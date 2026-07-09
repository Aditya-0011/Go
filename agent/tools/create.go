package tools

import (
	"google.golang.org/adk/v2/agent"
	"google.golang.org/adk/v2/tool"
	"google.golang.org/adk/v2/tool/functiontool"
)

func create[Args any, Result any](
	name string,
	description string,
	fn func(ctx agent.Context, args Args) (Result, error),
) (tool.Tool, error) {
	return functiontool.New(
		functiontool.Config{
			Name:        name,
			Description: description,
		},
		fn,
	)
}
