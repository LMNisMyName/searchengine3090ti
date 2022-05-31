package pack

import (
	"searchengine3090ti/cmd/user/dal/db"
	"searchengine3090ti/kitex_gen/userModel"
)

func User(u *db.User) *userModel.User {
	if u == nil {
		return nil
	}

	return &userModel.User{UserId: int64(u.ID), UserName: u.UserName, Avatar: "test"}
}

func Users(us []*db.User) []*userModel.User {
	users := make([]*userModel.User, 0)
	for _, u := range us {
		if userItr := User(u); userItr != nil {
			users = append(users, userItr)
		}
	}
	return users
}
