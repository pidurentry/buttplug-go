package buttplug

type Timeout struct{}
type UnknownMessageType struct{}
type MessageIdReused struct{}
type UnexpectedWebsocketMessageType struct{}

func (*Timeout) Error() string {
	return "timeout reached"
}

func (*UnknownMessageType) Error() string {
	return "unknown message type"
}

func (*MessageIdReused) Error() string {
	return "message id already used"
}

func (*UnexpectedWebsocketMessageType) Error() string {
	return "unexpected websocket message type"
}
