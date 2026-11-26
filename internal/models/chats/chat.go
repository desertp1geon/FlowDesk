package chats

import "github.com/desertp1geon/FlowDesk/internal/models/users"

type Chat struct {
	Id         uint `gorm:"primaryKey"`
	OwenrId    uint
	ReceiverID uint

	Owner    users.User `gorm:"foreignKey:OwnerID;references:ID"`
	Receiver users.User `gorm:"foreignKey:ReceiverID;references:ID"`
}
