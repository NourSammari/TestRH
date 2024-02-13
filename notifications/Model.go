package notifications

type PriorityType int

const (
	High PriorityType = iota
	Medium
	Low
)

type Notification struct {
	IdNotification   int          `json:"idNotif" gorm:"primaryKey;autoIncrement"`
	TypeNotification string       `json:"typeNotif"`
	Subject          string       `json:"subject"`
	Content          string       `json:"content"`
	Priority         PriorityType `json:"priority"`
	DeliveryStatus   string       `json:"deliveryStatus"`
}
