package interfaces

import (
	"github.com/ieltof/domain"
	"fmt"
	"time"
)

type DbHandler interface {
	Execute(statement string)
	Query(statement string) Row
}

type Row interface {
	Scan(dest ...interface{})
	Next() bool
}

type DbRepo struct {
	dbHandlers map[string]DbHandler
	dbHandler  DbHandler
}

type DbUserRepo DbRepo
type DbFriendRepo DbRepo
type DbMessageRepo DbRepo

func NewDbUserRepo(dbHandlers map[string]DbHandler) *DbUserRepo {
	dbUserRepo := new(DbUserRepo)
	dbUserRepo.dbHandlers = dbHandlers
	dbUserRepo.dbHandler = dbHandlers["DbUserRepo"]
	return dbUserRepo
}

func (repo *DbUserRepo) NewUser(user domain.User) {
	repo.dbHandler.Execute(fmt.Sprintf("INSERT INTO users (fid, name) VALUES ('%s', '%s')", user.FacebookId, user.Name))
}

func (repo *DbUserRepo) GetUser(user domain.User) domain.User {

	row := repo.dbHandler.Query(fmt.Sprintf("SELECT * FROM users WHERE fid = '%s' LIMIT 1", user.FacebookId))

	var id uint32
	var fid string
	var name string
	var url string
	row.Next()
	row.Scan(&id, &fid, &name, &url)
	var u domain.User;
	if (id != 0) {
		u = domain.User{Id: id, FacebookId: fid, Name: name, Url: url}
		fmt.Println(user)
	} else {
		repo.dbHandler.Execute(fmt.Sprintf("INSERT INTO users (fid, name, url) VALUES ('%s', '%s', '%s')", user.FacebookId, user.Name, user.Url))
	}

	//var isAdmin string
	//var customerId int
	//row.Next()
	//row.Scan(&isAdmin, &customerId)
	//customerRepo := NewDbCustomerRepo(repo.dbHandlers)
	//u := domain.User{Id: id, Customer: customerRepo.FindById(customerId)}

	//return u
	return u;
}

func (repo *DbUserRepo) GetUsers() []domain.User {
	row := repo.dbHandler.Query(fmt.Sprintf("SELECT * FROM users"))
	var id uint32
	var fid string
	var name string
	var url string
	var users []domain.User;
	for row.Next() {
		row.Scan(&id, &fid, &name, &url)
		u := domain.User{Id: id, FacebookId: fid, Name: name, Url: url}
		users = append(users, u)
	}

	return users;
}

func NewDbFriendRepo(dbHandlers map[string]DbHandler) *DbFriendRepo {
	dbFriendRepo := new(DbFriendRepo)
	dbFriendRepo.dbHandlers = dbHandlers
	dbFriendRepo.dbHandler = dbHandlers["DbFriendRepo"]
	return dbFriendRepo
}

func (repo *DbFriendRepo) GetFriends(userId int) []domain.Friend{
	row := repo.dbHandler.Query(fmt.Sprintf("SELECT * FROM friends WHERE uid = '%d'", userId))
	var id uint32
	var uid uint32
	var fid uint32
	var apt bool
	var friends []domain.Friend;
	for row.Next() {
		row.Scan(&id, &uid, &fid, &apt)
		f := domain.Friend{Id: id, Uid: uid, Fid: fid, Apt: apt}
		friends = append(friends, f)
	}

	return friends;
}

func (repo *DbFriendRepo) FindById(id int) domain.Friend {

	var f = domain.Friend{}
	return f;
}

func (repo *DbFriendRepo) FriendRequest(friendRequest domain.FriendRequest) {
	repo.dbHandler.Execute(fmt.Sprintf("INSERT INTO friends (uid, fid, apt) VALUES ('%d', '%d', '%t')",
		friendRequest.FromId, friendRequest.ToId, false))
}

func NewDbMessageRepo(dbHandlers map[string]DbHandler) *DbMessageRepo {
	dbMessageRepo := new(DbMessageRepo)
	dbMessageRepo.dbHandlers = dbHandlers
	dbMessageRepo.dbHandler = dbHandlers["DbMessageRepo"]
	return dbMessageRepo
}

func (repo *DbMessageRepo) InitMessage(friendRequest domain.FriendRequest) {

	row := repo.dbHandler.Query(fmt.Sprintf("SELECT * FROM friends WHERE uid = '%d' LIMIT 1", friendRequest.FromId))

	var initMessage string = "Lets be friends"

	var id uint32
	var uid uint32
	var fid uint32
	var apt bool
	row.Next()
	row.Scan(&id, &uid, &fid, &apt)
	if (id != 0) {
		repo.dbHandler.Execute(fmt.Sprintf("INSERT INTO messages (uid, did, text, tmp) VALUES ('%d', '%d', '%s', '%d')",
			friendRequest.FromId, id, initMessage, time.Now().Unix()))
	}
}

func (repo *DbMessageRepo) NewMessage(user domain.Message) {

}

func (repo *DbMessageRepo) GetMessages(dialogId uint32) []domain.Message {
	row := repo.dbHandler.Query(fmt.Sprintf("SELECT * FROM messages WHERE did = '%d'", dialogId))

	var id uint32
	var uid uint32
	var did uint32
	var text string
	var timestamp uint32
	var messages []domain.Message;
	for row.Next() {
		row.Scan(&id, &uid, &did, &text, &timestamp)
		f := domain.Message{Id: id, UserId: uid, DialogId: did, Text: text, Timestamp: timestamp}
		messages = append(messages, f)
	}

	return messages;
}
