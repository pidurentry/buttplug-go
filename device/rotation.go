package device

type Direction bool

const (
	ANTICLOCKWISE Direction = false
	CLOCKWISE     Direction = true
)

type Rotation struct {
	Index     Index     `json:"Index"`
	Speed     float32   `json:"Speed"`
	Direction Direction `json:"Clockwise"`
}
