package interfaces

import (
	"github.com/ieltof/domain"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"strconv"
)


type Interator interface {
	GetUser(domain.User) (domain.User, error)
	NewUser(domain.User) error
	GetUsers() ([]domain.User, error)
	GetOnlineUsers(ids []uint32) ([]domain.User, error)

	GetFriends(userId uint32) ([]domain.Friend, error)
	FriendRequest(domain.FriendRequest) (bool, error)
	AcceptFriendship(friend domain.Friend) (bool, error)
	IgnoreFriendship(friend domain.Friend) (bool, error)

	GetMessages(orderId uint32) ([]domain.Message, error)
	NewMessage(message domain.Message) (domain.Message, error)
	InitMessage(request domain.FriendRequest) error
}


type WebserviceHandler struct {
	Interator Interator
}

var HubInstance Hub;
var InteratorInstance Interator;

func (handler WebserviceHandler) GetUser(res http.ResponseWriter, req *http.Request) {

	var err error
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		res.WriteHeader(http.StatusPreconditionFailed)
	}

	var user domain.User
	err = json.Unmarshal(b, &user)
	if err != nil {
		res.WriteHeader(http.StatusPreconditionFailed)
	}
	u, err := handler.Interator.GetUser(user);

	resp, err := json.Marshal(u)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Write(resp)
}

func (handler WebserviceHandler) NewUser(res http.ResponseWriter, req *http.Request) {

	var err error
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		res.WriteHeader(http.StatusPreconditionFailed)
	}

	var user domain.User
	err = json.Unmarshal(b, &user)
	if err != nil {
		res.WriteHeader(http.StatusPreconditionFailed)
	}

	handler.Interator.NewUser(user);
}

func (handler WebserviceHandler) GetUsers(res http.ResponseWriter, req *http.Request) {

	users, err := handler.Interator.GetUsers();
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}

	resp, err := json.Marshal(users)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Write(resp)
}

func (handler WebserviceHandler) GetOnlineUsers(res http.ResponseWriter, req *http.Request) {
	onlineUsers := make(map[uint32]*User)
	onlineUsers = HubInstance.GetUsers()

	var ids [] uint32
	for k := range onlineUsers {
		ids = append(ids, k)
	}

	users, err := handler.Interator.GetOnlineUsers(ids);
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}

	resp, err := json.Marshal(users)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Write(resp)

}

func (handler WebserviceHandler) FriendRequest(res http.ResponseWriter, req *http.Request) {

	var err error
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		res.WriteHeader(http.StatusPreconditionFailed)
	}

	var friendRequest domain.FriendRequest
	err = json.Unmarshal(b, &friendRequest)
	if err != nil {
		res.WriteHeader(http.StatusPreconditionFailed)
	}

	isOk, _ := handler.Interator.FriendRequest(friendRequest);

	if(isOk) {
		handler.Interator.InitMessage(friendRequest);
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
}

func (handler WebserviceHandler) GetFriends(res http.ResponseWriter, req *http.Request) {

	id := req.URL.Query().Get("user_id")

	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	friends, err := handler.Interator.GetFriends(uint32(userId))

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}

	resp, err := json.Marshal(friends)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Write(resp)
}

func (handler WebserviceHandler) AcceptFriendship(res http.ResponseWriter, req *http.Request) {
	var err error
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		res.WriteHeader(http.StatusPreconditionFailed)
	}

	var friend domain.Friend
	err = json.Unmarshal(b, &friend)
	if err != nil {
		res.WriteHeader(http.StatusPreconditionFailed)
	}

	isOk, _ := handler.Interator.AcceptFriendship(friend);

	if(isOk) {

	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
}

func (handler WebserviceHandler) IgnoreFriendship(res http.ResponseWriter, req *http.Request) {
	var err error
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		res.WriteHeader(http.StatusPreconditionFailed)
	}

	var friend domain.Friend
	err = json.Unmarshal(b, &friend)
	if err != nil {
		res.WriteHeader(http.StatusPreconditionFailed)
	}

	isOk, _ := handler.Interator.IgnoreFriendship(friend);

	if(isOk) {

	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
}



func (handler WebserviceHandler) OptionRequest(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "*")
}