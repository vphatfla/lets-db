package io

type IOHandler interface {
	Execute(key string, value string) error
}
