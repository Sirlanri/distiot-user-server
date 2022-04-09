package token

//通过token获取userID，如果redis中没有，则从mysql中获取并写入redis
func GetUserIDByToken(token string) (userID int, err error) {
	userID, err = GetUserIDByTokenRedis(token)
	if err != nil {
		userID, err = GetUserIDByTokenMysql(token)
		if err != nil {
			return 0, err
		}
		//在MySQL发现数据，存入Redis
		InsertTokenUserIDRedis(token, userID)
		return userID, nil
	}
	return userID, nil
}
