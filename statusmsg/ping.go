package statusmsg

//go:generate go run ../message/generator.go -- $GOFILE

type Ping struct {
	ID int `json:"Id"`
}
