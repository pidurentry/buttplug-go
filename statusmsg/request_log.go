package statusmsg

//go:generate go run ../message/generator.go -- $GOFILE

type RequestLog struct {
	ID       int      `json:"Id"`
	LogLevel LogLevel `json:"LogLevel"`
}
