package alerts

import "context"

type CommandHandler struct {
	AlertRepository IRepository
}

func NewCommandHandler(alertRepo IRepository) *CommandHandler {
	return &CommandHandler{alertRepo}
}

func (h *CommandHandler) CreateAlert(ctx context.Context, command CreateAlertCommand) error {
	alert, err := NewAlert(CreateAlertParameters{command.DeviceSerial,
		command.Text,
		command.Humidity})

	if err != nil {
		return err
	}

	err = h.AlertRepository.Save(ctx, alert)

	return err

}
