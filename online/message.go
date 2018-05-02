package online

const (
	actionSendMessage = "sendMessage"
)

type Message struct {
	DialogId uint32 `json:"dialog_id"`
	FromId   uint32 `json:"from_id"`
	ToId     uint32 `json:"to_id"`
	Text     string `json:"text"`
}
