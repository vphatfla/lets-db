package io

import (
	"fmt"

	"github.com/vphatfla/lets-db/db/engine"
	"github.com/vphatfla/lets-db/db/formatter"
)
type IOSelectHandler struct {
}

func (h *IOSelectHandler) Execute(key string, value string) (formatter.Output, error) {
	if engine.Table == nil {
		e := &formatter.Error{}
		err := fmt.Errorf("Error table db engine is nil")
		e.Format("Select", key, value, err.Error())
		return  e, err
	}


	if v, exists := engine.Table[key]; exists {
		// key, value existed, override now
		delete(engine.Table, key)
		r := &formatter.Result{}
		comment := fmt.Sprintf("Key existed, select %d bytes", len([]byte(v)))
		r.Format("Select", key, value, comment)
		return r, nil
	} else {
		r := &formatter.Result{}
		comment := fmt.Sprintf("Key does not existed, select %d bytes", len([]byte(v)))
		r.Format("Select", key, value, comment)
		return r, nil
	}
}
