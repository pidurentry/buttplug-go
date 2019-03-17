package handshakemsg

//go:generate go run ../message/generator.go -- $GOFILE

type ServerName string
type MajorVersion int
type MinorVersion int
type BuildVersion int
type MaxPingTime int

type ServerInfo struct {
	ID             int            `json:"Id"`
	ServerName     ServerName     `json:"ServerName"`
	MajorVersion   MajorVersion   `json:"MajorVersion"`
	MinorVersion   MinorVersion   `json:"MinorVersion"`
	BuildVersion   BuildVersion   `json:"BuildVersion"`
	MessageVersion MessageVersion `json:"MessageVersion"`
	MaxPingTime    MaxPingTime    `json:"MaxPingTime"`
}
