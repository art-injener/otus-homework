package mysql

const (
	QueryGetAllAccounts    = "select * from users"
	QueryAddAccount        = "insert into users (login_id, name, surname, birthday, sex, hobby, city, avatar) values (?, ?, ?, ?, ?, ?, ?, ?)"
	QueryGetAccountByParam = "select * from users where `%s`=?"
	QueryUpdateAccount     = "update users set " +
		"name = '%s', " +
		"surname = '%s', " +
		"birthday = '%s', " +
		"sex = '%s', " +
		"hobby = '%s', " +
		"city = '%s', " +
		"avatar = '%s' " +
		"where login_id = '%d'"

	QueryGetLoginsInfo = "select * from logins_info"
	QueryAddUser       = "insert into logins_info (email, password) values (?, ?)"

	QueryMakeFriends  = "insert into friends (user_id, friend_id, accept) values (?, ?, 1)"
	QueryGetFriends   = "select users.id, name, surname from users join friends on users.id = friends.friend_id where user_id = ?"
	QueryCheckFriends = "select COUNT(id) from friends where friends.user_id = ? and friends.friend_id = ?"
)
