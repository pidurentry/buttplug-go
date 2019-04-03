package handshakemsg

import "github.com/pidurentry/buttplug-go/message"

//go:generate go run ../tools/message_generator.go -- $GOFILE

type ClientName string
type MessageVersion int

const MESSAGE_VERSION MessageVersion = 1

type RequestServerInfo struct {
	ID             message.Id     `json:"Id"`
	ClientName     ClientName     `json:"ClientName"`
	MessageVersion MessageVersion `json:"MessageVersion"`
}
