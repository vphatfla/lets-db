package io

import "github.com/vphatfla/lets-db/db/formatter"

type IOHandler interface {
	Execute(key string, value string) (formatter.Output, error)
}
