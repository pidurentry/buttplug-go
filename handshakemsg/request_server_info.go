package handshakemsg

import "github.com/pidurentry/buttplug-go"

//go:generate go run ../message/generator.go -- $GOFILE

type ClientName string
type MessageVersion int

const MESSAGE_VERSION MessageVersion = 1

type RequestServerInfo struct {
	ID             buttplug.MessageId `json:"Id"`
	ClientName     ClientName         `json:"ClientName"`
	MessageVersion MessageVersion     `json:"MessageVersion"`
}
