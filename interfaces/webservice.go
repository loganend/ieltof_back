package interfaces

import (
	"github.com/ieltof/domain"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"strconv"
)

//type UserInteractor interface {
//	GetUser(domain.User) (domain.User, error)
//	NewUser(domain.User) error
//	GetUsers() ([]domain.User, error)
//}

type Interator interface {
	GetUser(domain.User) (domain.User, error)
	NewUser(domain.User) error
	GetUsers() ([]domain.User, error)

	GetFriends(userId int) ([]domain.Friend, error)
	FriendRequest(domain.FriendRequest) error

	GetMessages(orderId uint32) ([]domain.Message, error)
	NewMessage(userId, orderId, itemId int) error
	InitMessage(request domain.FriendRequest) error
}

//
//
//type FriendInteractor interface {
//	GetFriends(userId int) (domain.User, error)
//	FriendRequest(domain.FriendRequest) error
//}
//
//type MessageInteractor interface {
//	GetMessages(orderId int) (domain.User, error)
//	NewMessage(userId, orderId, itemId int) error
//	InitMessage(request domain.FriendRequest) error
//}

type WebserviceHandler struct {
	Interator Interator
	//UserInteractor    UserInteractor
	//FriendInteractor  FriendInteractor
	//MessageInteractor MessageInteractor
}

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

	handler.Interator.FriendRequest(friendRequest);
	handler.Interator.InitMessage(friendRequest);
}

func (handler WebserviceHandler) GetFriends(res http.ResponseWriter, req *http.Request) {

	id := req.URL.Query().Get("user_id")

	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	friends, err := handler.Interator.GetFriends(userId)

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