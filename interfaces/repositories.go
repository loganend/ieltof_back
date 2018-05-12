package interfaces

import (
	"github.com/ieltof/domain"
	"fmt"
	"time"
	"strconv"
	"strings"
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

		row = repo.dbHandler.Query(fmt.Sprintf("SELECT * FROM users WHERE fid = '%s' LIMIT 1", user.FacebookId))
		row.Next()
		row.Scan(&id, &fid, &name, &url)
		u = domain.User{Id: id, FacebookId: fid, Name: name, Url: url}
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

func (repo *DbUserRepo) GetOnlineUsers(ids []uint32) []domain.User {
	var values string = "("
	for k := range ids {
		values = values +  strconv.FormatInt(int64(ids[k]), 10) + ", "
	}
	values = values + "100)"

	row := repo.dbHandler.Query(fmt.Sprintf("SELECT * FROM users WHERE id in " + values))
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

//Friends

func NewDbFriendRepo(dbHandlers map[string]DbHandler) *DbFriendRepo {
	dbFriendRepo := new(DbFriendRepo)
	dbFriendRepo.dbHandlers = dbHandlers
	dbFriendRepo.dbHandler = dbHandlers["DbFriendRepo"]
	return dbFriendRepo
}

func (repo *DbFriendRepo) GetFriends(userId uint32) []domain.Friend {
	//row := repo.dbHandler.Query(fmt.Sprintf("SELECT * FROM friends WHERE uid = '%d' or fid = '%d'", userId, userId))
	row := repo.dbHandler.Query(fmt.Sprintf("select t1.id, t1.uid, t1.fid, t1.apt, t2.uid as uid2, t2.text, t2.tmp, t3.id as id1, t3.name, t3.url, t4.id as id2, t4.name as name2, t4.url as url2 from friends as t1 "+
		"join messages as t2 on t1.id =t2.did "+
		"join users as t3 on t1.uid = t3.id "+
		"join users as t4 on t1.fid = t4.id "+
		"where t1.uid = '%d' or t1.fid = '%d'", userId, userId))

	var id uint32
	var uid uint32
	var fid uint32
	var apt bool

	var mUid uint32
	var text string
	var timestamp int64
	var id1 uint32
	var id2 uint32
	var name string
	var url string
	var name2 string
	var url2 string
	//var friends []domain.Friend;
	friends := make(map[uint32]*domain.Friend)

	for row.Next() {
		row.Scan(&id, &uid, &fid, &apt, &mUid, &text, &timestamp, &id1, &name, &url, &id2, &name2, &url2)
		if friends[id] != nil {
			m := domain.Message{Id: 0, UserId: mUid, DialogId: id, Text: text, Timestamp: timestamp}
			friends[id].Messages = append(friends[id].Messages, m)
		} else {
			f := domain.Friend{Id: id, Uid: uid, Fid: fid, Apt: apt, Messages: nil, Name: name, Url: url}
			friends[id] = &f
			if (userId != id2) {
				friends[id].Name = name2
				friends[id].Url = url2
			} else {
				friends[id].Name = name
				friends[id].Url = url
			}
		}
	}

	var fds []domain.Friend;
	for k := range friends {
		fds = append(fds, *friends[k])
	}

	return fds;
}

func (repo *DbFriendRepo) FindById(id int) domain.Friend {

	var f = domain.Friend{}
	return f;
}

func (repo *DbFriendRepo) FriendRequest(friendRequest domain.FriendRequest) bool {

	row := repo.dbHandler.Query(fmt.Sprintf("SELECT * FROM friends WHERE uid = '%d' AND fid = '%d'", friendRequest.FromId, friendRequest.ToId))
	if row.Next() {
		return false;
	}

	repo.dbHandler.Execute(fmt.Sprintf("INSERT INTO friends (uid, fid, apt) VALUES ('%d', '%d', '%t')",
		friendRequest.FromId, friendRequest.ToId, false))
	return true
}

func (repo *DbFriendRepo) AcceptFriendship(friend domain.Friend) bool {

	repo.dbHandler.Execute(fmt.Sprintf("UPDATE friends SET apt = '%t' WHERE id = '%d'", true, friend.Id))

	return true;
}

func (repo *DbFriendRepo) IgnoreFriendship(friend domain.Friend) bool {

	repo.dbHandler.Execute(fmt.Sprintf("DELETE FROM friends WHERE id = '%d'", friend.Id))

	return true;
}

//Messages

func NewDbMessageRepo(dbHandlers map[string]DbHandler) *DbMessageRepo {
	dbMessageRepo := new(DbMessageRepo)
	dbMessageRepo.dbHandlers = dbHandlers
	dbMessageRepo.dbHandler = dbHandlers["DbMessageRepo"]
	return dbMessageRepo
}

func (repo *DbMessageRepo) InitMessage(friendRequest domain.FriendRequest) {

	row := repo.dbHandler.Query(fmt.Sprintf("SELECT * FROM friends WHERE uid = '%d' AND fid = '%d' LIMIT 1", friendRequest.FromId, friendRequest.ToId))

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

func (repo *DbMessageRepo) NewMessage(message domain.Message) domain.Message {
	message.Timestamp = time.Now().Unix()
	message.Text = strings.Replace(message.Text, "'", "''", -1)
	repo.dbHandler.Execute(fmt.Sprintf("INSERT INTO messages (uid, did, text, tmp) VALUES ('%d', '%d', '%s', '%d')",
		message.UserId, message.DialogId, message.Text, message.Timestamp))
	return message
}

func (repo *DbMessageRepo) GetMessages(dialogId uint32) []domain.Message {
	row := repo.dbHandler.Query(fmt.Sprintf("SELECT * FROM messages WHERE did = '%d'", dialogId))

	var id uint32
	var uid uint32
	var did uint32
	var text string
	var timestamp int64
	var messages []domain.Message;
	for row.Next() {
		row.Scan(&id, &uid, &did, &text, &timestamp)
		f := domain.Message{Id: id, UserId: uid, DialogId: did, Text: text, Timestamp: timestamp}
		messages = append(messages, f)
	}

	return messages;
}
