package main

import (
	"errors"
	"strings"
	"sync"
)

type Messenger interface {
	AddUser(username string, isBot bool) (int, error)
	AddChat(chatname string, isGroup bool, creator int, admins []int) (int, error)
	SendMessage(userId, chatId int, text string) (int, error)
	SendLike(userId, messageId int) error
	GetNumberOfLikes(messageId int) (int, error)
	SetChatAdmin(chatId, userId int) error
	GetLastMessage(chatId int) (string, int, error)
	GetLastUserMessage(userId int) (string, int, error)
}

type LastID struct {
	sync.Mutex
	Value int
}

var LastUID LastID
var LastChatID LastID
var LastMessageID LastID

type User struct {
	ID       int
	Username string
	IsBot    bool
}

type Message struct {
	ID   int
	Text string
}

type Chat struct {
	ID        int
	Name      string
	IsGroup   bool
	CreatorID int
	AdminsID  []int
	Messages  []Message
}

type MessengerImpl struct {
	Users []User
	Chats []Chat
}

func NewMessengerImpl() *MessengerImpl {
	users := make([]User, 0)
	chats := make([]Chat, 0)
	return &MessengerImpl{users, chats}
}

func generateID(id *LastID) int {
	id.Lock()
	defer id.Unlock()
	id.Value++
	return id.Value
}

func (messenger *MessengerImpl) AddUser(username string, isBot bool) (int, error) {
	username = strings.ToLower(username)
	if !ValidateUsername(username, messenger.Users) {
		return 0, errors.New("invalid username")
	}
	user := User{generateID(&LastUID), username, isBot}
	messenger.Users = append(messenger.Users, user)
	return user.ID, nil
}

func (messenger *MessengerImpl) AddChat(chatname string, isGroup bool, creator int, admins []int) (int, error) {
	if CheckChatnameExists(chatname, messenger.Chats) || !ValidateChatCreatorToAdd(creator, messenger.Users) || !ValidateChatAdminsToAdd(admins, messenger.Users) {
		return 0, errors.New("could not create chat")
	}
	chat := Chat{generateID(&LastChatID), chatname, isGroup, creator, admins, []Message{}}
	messenger.Chats = append(messenger.Chats, chat)
	return chat.ID, nil
}

func (messenger *MessengerImpl) SendMessage(userId, chatId int, text string) (int, error) {
	if !FindUserByID(userId, messenger.Users) || !FindChatByID(chatId, messenger.Chats) || !CheckSendAccess(userId, chatId, messenger.Chats) {
		return 0, errors.New("user could not send message")
	}

	chatIndex := 0
	for idx, c := range messenger.Chats {
		if c.ID == chatId {
			chatIndex = idx
		}
	}
	id := generateID(&LastMessageID)
	messenger.Chats[chatIndex].Messages = append(messenger.Chats[chatIndex].Messages, Message{id, text})
	return id, nil
}
