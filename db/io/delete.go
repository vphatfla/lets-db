package io

import (
	"fmt"

	"github.com/vphatfla/lets-db/db/engine"
	"github.com/vphatfla/lets-db/db/formatter"
)
type IODeleteHandler struct {
}

func (h *IODeleteHandler) Execute(key string, value string) (formatter.Output, error) {
	if engine.Table == nil {
		e := &formatter.Error{}
		err := fmt.Errorf("Error table db engine is nil")
		e.Format("Delete", key, value, err.Error())
		return  e, err
	}


	if v, exists := engine.Table[key]; exists {
		// key, value existed, override now
		delete(engine.Table, key)
		r := &formatter.Result{}
		comment := fmt.Sprintf("Key existed, deleted %d bytes", len([]byte(v)) + len([]byte(key)))
		r.Format("Delete", key, value, comment)
		return r, nil
	} else {
		e := &formatter.Error{}
		err := fmt.Errorf("Key %s does not exist", key)
		e.Format("Delete", key, value, err.Error())
		return e, err
	}
}

