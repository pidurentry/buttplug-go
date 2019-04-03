package message

var Repository = make(map[string]func() interface{})
