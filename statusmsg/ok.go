package statusmsg

//go:generate go run ../message/generator.go -- $GOFILE

type Ok struct {
	ID int `json:"Id"`
}
