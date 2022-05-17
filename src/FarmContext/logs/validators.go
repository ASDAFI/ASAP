package logs

import (
	"errors"
)

func (d DeviceLog) validateForCreateNewInstance() error {
	return nil
}

func (w WaterLog) validateForCreateNewInstance() error {
	if w.Volume <= 0 {
		return errors.New("water volume cant be less than 0")
	}
	return nil
}
