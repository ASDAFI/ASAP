package alerts

import "context"

type AlertQueryHandler struct {
	alertRepo IRepository
}

func NewQueryHandler(alertRepo IRepository) *AlertQueryHandler {
	return &AlertQueryHandler{alertRepo}
}

func (h *AlertQueryHandler) GetLastAlertByDeviceSerial(ctx context.Context, query FindLastAlertByDeviceSerialQuery) (*Alert, error) {

	alert, err := h.alertRepo.FindLastByDeviceSerial(ctx, query.DeviceSerial)
	if err != nil {
		return nil, err
	}

	return alert, nil
}

func (h *AlertQueryHandler) GetAlertsByDeviceSerial(ctx context.Context, query FindAlertsByDeviceSerialQuery) ([]*Alert, error) {

	alerts, err := h.alertRepo.FindByDeviceSerial(ctx, query.DeviceSerial)
	if err != nil {
		return nil, err
	}

	return alerts, nil
}

func (h *AlertQueryHandler) GetRecentAlertsByDeviceSerial(ctx context.Context, query FindRecentAlertsByDeviceSerialQuery) ([]*Alert, error) {

	alerts, err := h.alertRepo.FindRecentByDeviceSerial(ctx, query.DeviceSerial)
	if err != nil {
		return nil, err
	}

	return alerts, nil
}
