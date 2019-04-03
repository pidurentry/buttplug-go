package enumerationmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["ScanningFinished"] = func() interface{} { return &ScanningFinished{} }
}

func (msg *ScanningFinished) Id() message.Id {
    return msg.ID
}

func (msg *ScanningFinished) Serilize() interface{} {
    return map[string]*ScanningFinished{
        "ScanningFinished": msg,
    }
}

