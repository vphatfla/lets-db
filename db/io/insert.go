package io

import (
	"fmt"

	"github.com/vphatfla/lets-db/db/engine"
	"github.com/vphatfla/lets-db/db/output"
)

type IOInsertHandler struct {
}

func (h *IOInsertHandler) Execute(key string, value string) error {
	if engine.Table == nil {
		e := fmt.Errorf("Error table db engine is nil")
		output.PrintOutputWithError(e)
		return 	e
	}

	if value, exists := engine.Table[key]; exists {
		// key, value existed, override now
		output.PrintOutput(fmt.Sprintf("Key existed, override, wrote %d ", len([]byte(value))))
	} else {
		engine.Table[key] = value
		output.PrintOutput(fmt.Sprintf("Key NOT exist, wrote %d ", len([]byte(value))))
	}

	return nil

}
// func Insert(key string, value string) error {
// }
