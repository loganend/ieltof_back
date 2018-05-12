package domain

type FriendRepository interface {
	GetFriends(userId uint32) []Friend
	FriendRequest(friendRequest FriendRequest) bool
	AcceptFriendship(friend Friend) bool
	IgnoreFriendship(friend Friend) bool
}

type Friend struct {
	Id  uint32 `json:"dialog_id"`
	Uid uint32 `db:"uid" json:"user_id"`
	Fid uint32 `db:"fid" json:"friend_id"`
	Apt bool   `db:"apt" json:"accept"`
	Messages[] Message `json:"messages"`
	Name string `json:"name"`
	Url string `json:"url"`
}

type FriendRequest struct {
	FromId int `json:"from_id"`
	ToId   int `json:"to_id"`
}
