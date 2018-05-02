package domain

type MessageRepository interface {
	NewMessage(Message) Message
	GetMessages(id uint32) []Message
	InitMessage(request FriendRequest)
}

type Message struct {
	Id        uint32
	UserId    uint32
	DialogId  uint32
	Text      string
	Timestamp int64
}
