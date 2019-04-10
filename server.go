package buttplug

import "github.com/pidurentry/buttplug-go/server"

type Server interface {
	ServerName() server.Name
	MajorVersion() server.MajorVersion
	MinorVersion() server.MinorVersion
	BuildVersion() server.BuildVersion
	MaxPingTime() server.MaxPingTime
}

type buttplugServer struct {
	serverName   server.Name
	majorVersion server.MajorVersion
	minorVersion server.MinorVersion
	buildVersion server.BuildVersion
	maxPingTime  server.MaxPingTime
}

func newButtplugServer(serverName server.Name, majorVersion server.MajorVersion, minorVersion server.MinorVersion, buildVersion server.BuildVersion, maxPingTime server.MaxPingTime) Server {
	return &buttplugServer{
		serverName:   serverName,
		majorVersion: majorVersion,
		minorVersion: minorVersion,
		buildVersion: buildVersion,
		maxPingTime:  maxPingTime,
	}
}

func (buttplugServer *buttplugServer) ServerName() server.Name {
	return buttplugServer.serverName
}

func (buttplugServer *buttplugServer) MajorVersion() server.MajorVersion {
	return buttplugServer.majorVersion
}

func (buttplugServer *buttplugServer) MinorVersion() server.MinorVersion {
	return buttplugServer.minorVersion
}

func (buttplugServer *buttplugServer) BuildVersion() server.BuildVersion {
	return buttplugServer.buildVersion
}

func (buttplugServer *buttplugServer) MaxPingTime() server.MaxPingTime {
	return buttplugServer.maxPingTime
}
