package statusmsg

//go:generate go run ../message/generator.go -- $GOFILE

type TestString string

type Test struct {
	ID         int        `json:"Id"`
	TestString TestString `json:"TestString"`
}
