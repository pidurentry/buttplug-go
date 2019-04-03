package enumerationmsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go"

func init() {
    buttplug.MessageRepository["ScanningFinished"] = func() interface{} { return &ScanningFinished{} }
}

func (msg *ScanningFinished) Id() buttplug.MessageId {
    return msg.ID
}

func (msg *ScanningFinished) Serilize() interface{} {
    return map[string]*ScanningFinished{
        "ScanningFinished": msg,
    }
}

