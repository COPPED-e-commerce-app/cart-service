package models

type Cart struct {
	Id            string   `json:"id"`
	ItemIds       []string `json:"item_ids"`
	ShippingId    string   `json:"shipping_id"`
	UserId        string   `json:"user_id"`
	TotalPrice    string   `json:"total_price"`
	Tax           string   `json:"tax"`
	TimeCreated   string   `json:"time_created"`
	TimeCompleted string   `json:"time_completed"`
	Status        string   `json:"status"`
}
