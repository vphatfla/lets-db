package io

import (
	"fmt"

	"github.com/vphatfla/lets-db/db/engine"
	"github.com/vphatfla/lets-db/db/formatter"
)

type IOInsertHandler struct {
}

func (h *IOInsertHandler) Execute(key string, value string) (formatter.Output, error) {
	if engine.Table == nil {
		e := &formatter.Error{}
		err := fmt.Errorf("Error table db engine is nil")
		e.Format("Insert", key, value, err.Error())
		return  e, err
	}

	if _, exists := engine.Table[key]; exists {
		// key, value existed, override now
		engine.Table[key]= value
		r := &formatter.Result{}
		comment := fmt.Sprintf("Key existed, override, wrote %d bytes", len([]byte(value)))
		r.Format("Insert", key, value, comment)
		return r, nil
	} else {
		engine.Table[key] = value
		r := &formatter.Result{}
		comment := fmt.Sprintf("Key NOT exist, wrote %d bytes", len([]byte(value)))
		r.Format("Insert", key, value, comment)
		return r, nil
	}
}
