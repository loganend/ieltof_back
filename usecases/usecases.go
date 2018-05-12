package usecases

import (
	"github.com/ieltof/domain"
)

//type UserInteractor struct {
//	UserRepository  domain.UserRepository
//}
//
//type FriendInteractor struct {
//	FriendRepository  domain.FriendRepository
//}
//
//type MessageInterator struct {
//	MessageRepository  domain.MessageRepository
//}


type Interactor struct {
	UserRepository  domain.UserRepository
	FriendRepository  domain.FriendRepository
	MessageRepository  domain.MessageRepository
}

func (interactor *Interactor) GetUser(user domain.User) (domain.User, error) {

	u := interactor.UserRepository.GetUser(user)
	return u, nil
}

func (interactor *Interactor) NewUser(user domain.User) error {
	interactor.UserRepository.NewUser(user)
	return nil
}

func (interator *Interactor) GetUsers() ([]domain.User, error) {
	return interator.UserRepository.GetUsers(), nil
}

func (interator *Interactor) GetOnlineUsers(ids []uint32) ([]domain.User, error) {
	users := interator.UserRepository.GetOnlineUsers(ids)

	return users, nil
}

func (interactor *Interactor) GetFriends(userId uint32) ([]domain.Friend, error) {

	friends := interactor.FriendRepository.GetFriends(userId)

	return friends, nil
}

func (interator *Interactor) AcceptFriendship(friend domain.Friend) (bool, error) {
	return interator.FriendRepository.AcceptFriendship(friend), nil
}

func (interator *Interactor) IgnoreFriendship(friend domain.Friend) (bool, error) {
	return interator.FriendRepository.IgnoreFriendship(friend), nil
}


func (interactor *Interactor) FriendRequest(friendRequest domain.FriendRequest) (bool, error) {
	res := interactor.FriendRepository.FriendRequest(friendRequest)
	return res, nil
}

func (interactor *Interactor) InitMessage(friendRequest domain.FriendRequest) error {
	interactor.MessageRepository.InitMessage(friendRequest)
	return nil
}

func (interactor *Interactor) GetMessages(id uint32) ([]domain.Message, error) {
	return []domain.Message{}, nil
}

func (interactor *Interactor) NewMessage(message domain.Message) (domain.Message, error) {
	msg := interactor.MessageRepository.NewMessage(message)
	return msg, nil
}