package repository

func AddUser(newuser User) {
	if GetDB().Create(&newuser).Error != nil {
		panic("failed to insert user record")
	}
}

func UsersLoginInfo(token string) (user User, exist bool) {
	if GetDB().Find(&user, "token=?", token).Error != nil {
		panic("err")
	}
	if user.Token == "" {
		return user, false
	}
	return user, true
}
