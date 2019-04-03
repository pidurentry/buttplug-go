package specificdevicemsg

// Auto generated file - do not directly edit!
import "github.com/pidurentry/buttplug-go/message"

func init() {
    message.Repository["FleshlightLaunchFW12Cmd"] = func() interface{} { return &FleshlightLaunchFW12Cmd{} }
}

func (msg *FleshlightLaunchFW12Cmd) Id() message.Id {
    return msg.ID
}

func (msg *FleshlightLaunchFW12Cmd) Serilize() interface{} {
    return map[string]*FleshlightLaunchFW12Cmd{
        "FleshlightLaunchFW12Cmd": msg,
    }
}

