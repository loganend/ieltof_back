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

func (interactor *Interactor) GetFriends(userId int) ([]domain.Friend, error) {

	friends := interactor.FriendRepository.GetFriends(userId)

	for index, relation := range friends {
		friends[index].Messages = interactor.MessageRepository.GetMessages(relation.Id)
	}

	return friends, nil
}

func (interactor *Interactor) FriendRequest(friendRequest domain.FriendRequest) error {
	interactor.FriendRepository.FriendRequest(friendRequest)
	return nil
}

func (interactor *Interactor) InitMessage(friendRequest domain.FriendRequest) error {
	interactor.MessageRepository.InitMessage(friendRequest)
	return nil
}

func (interactor *Interactor) GetMessages(id uint32) ([]domain.Message, error) {
	return []domain.Message{}, nil
}

func (interactor *Interactor) NewMessage(userId, orderId, itemId int) error {
	return nil
}