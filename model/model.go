package model

type OrderResponse struct {
	UserID   string `json:"userID"`
	Username string `json:"userName"`
	Password string `json:"password"`
	Broker   string `json:"broker"`
}
