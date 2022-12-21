package main

import (
	"context"
	"io"
)

type processDef interface {
	Init(ctx context.Context) error
	Build(ctx context.Context, stdout, stderr io.Writer) error
	Run(ctx context.Context, stdout, stderr io.Writer) error
	Kill(ctx context.Context) error
}