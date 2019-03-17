package handshakemsg

//go:generate go run ../message/generator.go -- $GOFILE

type ClientName string
type MessageVersion int

const MESSAGE_VERSION MessageVersion = 1

type RequestServerInfo struct {
	ID             int            `json:"Id"`
	ClientName     ClientName     `json:"ClientName"`
	MessageVersion MessageVersion `json:"MessageVersion"`
}
