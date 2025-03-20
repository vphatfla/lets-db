package engine

var Table map[string]string

func InitTableDB() {
	Table = make(map[string]string)
}
