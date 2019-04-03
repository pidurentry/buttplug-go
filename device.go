package buttplug

type DeviceName string
type DeviceIndex int

type DeviceMessages map[string]DeviceAttributes

type DeviceFeatureCount int

type DeviceAttributes struct {
	FeatureCount DeviceFeatureCount `json:"FeatureCount"`
}

type Device struct {
	Name     DeviceName     `json:"DeviceName"`
	Index    DeviceIndex    `json:"DeviceIndex"`
	Messages DeviceMessages `json:"DeviceMessages"`
}
