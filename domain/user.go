package domain

type UserRepository interface {
	GetUser(User) User
	NewUser(User)
	GetUsers() []User
	GetOnlineUsers([]uint32) []User
}

type User struct {
	Id         uint32    `db:"id"`
	FacebookId string `db:"fid" json:"facebook_id"`
	Name       string `db:"name" json:"name"`
	Url 	   string `db:"url" json:"url"`

}
