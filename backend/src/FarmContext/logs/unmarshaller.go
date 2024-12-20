package logs

import (
	"context"
	"encoding/json"
	"errors"
	"farm/src/FarmContext/alerts"
	"farm/src/FarmContext/devices"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type Unmarshaller struct {
	logRepository    IDeviceLogRepository
	deviceRepository devices.IRepository
}
type PureDeviceLog struct {
	Data []float64 `json:"data"`
}

type DeviceLogMessage struct {
	PayLoad []byte
	Topic   string
}

func NewUnmarshaller(logRepository IDeviceLogRepository, deviceRepository devices.IRepository) *Unmarshaller {
	return &Unmarshaller{logRepository: logRepository, deviceRepository: deviceRepository}
}

func (u *Unmarshaller) Unmarshal(ctx context.Context, message DeviceLogMessage) error {

	// Get the device id from the topic
	splitedTopic := strings.Split(message.Topic, "/")
	if len(splitedTopic) != 3 {
		log.Error("Topic is not valid")
		return errors.New("Topic is not valid")
	}

	deviceSerial := splitedTopic[2]
	_, err := u.deviceRepository.FindBySerial(ctx, deviceSerial)

	if err != nil {
		log.WithFields(log.Fields{"deviceSerial": deviceSerial}).Errorln("Device not found.")
		return errors.New("Device not found")
	}
	logFromPast := false
	// remove {ee} from the payload
	stringMessage := string(message.PayLoad)

	logLenght := len(stringMessage)
	if stringMessage[logLenght-4:] == "{ee}" {
		logFromPast = true
		stringMessage = stringMessage[:logLenght-4]
	}

	// Unmarshal the payload
	message.PayLoad = []byte(stringMessage)
	var pureDeviceLog PureDeviceLog
	err = json.Unmarshal(message.PayLoad, &pureDeviceLog)

	if err != nil {
		log.WithFields(log.Fields{
			"deviceSerial": deviceSerial,
			"logFromPast":  logFromPast,
			"error":        err,
		}).Error("Error unmarshalling the payload")
		return err
	}

	// fix device time
	stringTime := strconv.Itoa(int(pureDeviceLog.Data[4]))

	year, err := strconv.ParseInt(stringTime[:4], 10, 64)
	if err != nil {
		log.WithFields(log.Fields{
			"deviceSerial": deviceSerial,
			"logFromPast":  logFromPast,
			"error":        err,
		}).Error("Error fix the deviceTime")
	}

	month, err := strconv.ParseInt(stringTime[4:6], 10, 64)
	if err != nil {
		log.WithFields(log.Fields{
			"deviceSerial": deviceSerial,
			"logFromPast":  logFromPast,
			"error":        err,
		}).Error("Error fix the deviceTime")
	}

	day, err := strconv.ParseInt(stringTime[6:8], 10, 64)
	if err != nil {
		log.WithFields(log.Fields{
			"deviceSerial": deviceSerial,
			"logFromPast":  logFromPast,
			"error":        err,
		}).Error("Error fix the deviceTime")
	}

	hour, err := strconv.ParseInt(stringTime[8:10], 10, 64)
	if err != nil {
		log.WithFields(log.Fields{
			"deviceSerial": deviceSerial,
			"logFromPast":  logFromPast,
			"error":        err,
		}).Error("Error fix the deviceTime")
	}

	minute, err := strconv.ParseInt(stringTime[10:12], 10, 64)
	if err != nil {
		log.WithFields(log.Fields{
			"deviceSerial": deviceSerial,
			"logFromPast":  logFromPast,
			"error":        err,
		}).Error("Error fix the deviceTime")
	}

	second, err := strconv.ParseInt(stringTime[12:14], 10, 64)
	if err != nil {
		log.WithFields(log.Fields{
			"deviceSerial": deviceSerial,
			"logFromPast":  logFromPast,
			"error":        err,
		}).Error("Error fix the deviceTime")
	}

	deviceTime := time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, time.UTC)

	// Save the device log
	humidity := uint(float32(1024-uint(pureDeviceLog.Data[19])) / 1024 * 100)
	serverTIme := time.Now().UTC()

	log.WithFields(log.Fields{
		"deviceSerial": deviceSerial,
		"logFromPast":  logFromPast,
		"deviceTime":   deviceTime,
		"serverTime":   serverTIme,
		"humidity":     humidity,
	}).Info("Device Log")
	deviceLog := &DeviceLog{
		DeviceSerial: deviceSerial,
		DeviceTime:   deviceTime,
		Humidity:     humidity,
		ServerTime:   serverTIme,
	}
	if logFromPast == false {
		alerts.CheckAlert(ctx, deviceSerial, humidity)
	}
	return u.logRepository.Save(ctx, deviceLog)
}
