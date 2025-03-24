package formatter

import (
	"fmt"
	"strings"
)

type Output interface {
	Format(t string, key string, value string, comment string)
	GetOutput() string
}

type Result struct {
	Output string
}

type Error struct {
	Output string
}
// Output generic simple format for database output
func (r *Result) Format(t string, key string, value string, comment string) {
	r.Output = fmt.Sprintf("Executed %s for key = %s and value = %s | Comment = %s", strings.ToUpper(t), key, value, comment)
}

func (r *Result) GetOutput() string {
	return r.Output
}

// Output generic simple error
func (e *Error) Format(t string, key string, value string, comment string) {
	e.Output = fmt.Sprintf("Error executing %s for key = %s and value = %s | Comment =%s", strings.ToUpper(t), key, value, comment)
}

func (e *Error) GetOutput() string {
	return e.Output
}
