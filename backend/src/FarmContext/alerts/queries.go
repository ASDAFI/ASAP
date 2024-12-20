package alerts

type FindLastAlertByDeviceSerialQuery struct {
	DeviceSerial string
}

type FindAlertsByDeviceSerialQuery struct {
	DeviceSerial string
}

type FindRecentAlertsByDeviceSerialQuery struct {
	DeviceSerial string
}
