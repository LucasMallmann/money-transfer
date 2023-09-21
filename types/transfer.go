package types

type TransferRequest struct {
	Receiver string  `json:"receiver" validate:"required"`
	Sender   string  `json:"sender" validate:"required"`
	Amount   float64 `json:"amount" validate:"required"`
}
