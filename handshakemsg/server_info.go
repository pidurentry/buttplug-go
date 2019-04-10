package handshakemsg

import (
	"github.com/pidurentry/buttplug-go/message"
	"github.com/pidurentry/buttplug-go/server"
)

//go:generate go run ../tools/message_generator.go -- $GOFILE
type ServerInfo struct {
	ID             message.Id          `json:"Id"`
	ServerName     server.Name         `json:"ServerName"`
	MajorVersion   server.MajorVersion `json:"MajorVersion"`
	MinorVersion   server.MinorVersion `json:"MinorVersion"`
	BuildVersion   server.BuildVersion `json:"BuildVersion"`
	MessageVersion MessageVersion      `json:"MessageVersion"`
	MaxPingTime    server.MaxPingTime  `json:"MaxPingTime"`
}
