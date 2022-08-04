package main

import (
	"regexp"
)

func ValidateUsername(username string, users []User) bool {
	match, _ := regexp.MatchString("^[a-z0-9]{4,}$", username)
	if !match {
		return false
	}

	for _, u := range users {
		if username == u.Username {
			return false
		}
	}
	return true
}

func CheckChatnameExists(chatname string, chats []Chat) bool {
	for _, chat := range chats {
		if chatname == chat.Name {
			return true
		}
	}
	return false
}

func ValidateChatCreatorToAdd(creator int, users []User) bool {
	for _, u := range users {
		if creator == u.ID {
			if u.IsBot {
				return false
			} else {
				return true
			}
		}
	}
	return false
}

func FindUserByID(id int, users []User) bool {
	for _, u := range users {
		if id == u.ID {
			return true
		}
	}
	return false
}

func FindChatByID(id int, chats []Chat) bool {
	for _, c := range chats {
		if id == c.ID {
			return true
		}
	}
	return false
}

func FindIntInSlice(id int, items []int) bool {
	for _, t := range items {
		if id == t {
			return true
		}
	}
	return false
}

func ValidateChatAdminsToAdd(admins []int, users []User) bool {
	for _, admin := range admins {
		if !FindUserByID(admin, users) {
			return false
		}
	}
	return true
}

func CheckSendAccess(userID int, chatID int, chats []Chat) bool {
	for _, c := range chats {
		if c.ID == chatID {
			if c.IsGroup {
				return true
			} else {
				if FindIntInSlice(userID, c.AdminsID) {
					return true
				}
			}
		}
	}
	return false
}
