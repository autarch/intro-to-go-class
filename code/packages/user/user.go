package user

type User struct {
	username string
}

func NewUser(u string) *User {
	user := userFromCache(u)
	if user != nil {
		return user
	}
	
	user = &User{username: u}
	setUserInCache(user)

	return user
}
