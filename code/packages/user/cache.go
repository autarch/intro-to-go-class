package user

var cache map[string]*User

func userFromCache(username string) *User {
	cache := getCache()
	u, _ := cache[username]
	return u
}

func setUserInCache(user *User) {
	cache := getCache()
	cache[user.username] = user
}

func getCache() map[string]*User {
	if cache == nil {
		cache = make(map[string]*User)
	}
	return cache
}
