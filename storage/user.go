package storage

import (
	"fluffy-spoon/storage/mysql"
	"fmt"
)

//User simple user information
type User struct {
	Username     string
	Avatar       string
	Signature    string
	SignupAt     string
	LastActiveAt string
}

// UserSignup create a new user
func UserSignup(username string, password string) bool {
	stmt, err := mysql.DBConn().Prepare(
		"insert ignore into tbl_user (`user_name`, `user_pwd`) values (?,?)")
	if err != nil {
		fmt.Printf("Failed to insert, err: %s\n", err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, password)
	if err != nil {
		fmt.Printf("Failed to insert, err: %s\n", err.Error())
		return false
	}
	if rowsAffected, err := ret.RowsAffected(); err == nil && rowsAffected > 0 {
		return true
	}
	return false
}

// UserSignin user login
func UserSignin(username string, encodedPasswd string) bool {
	stmt, err := mysql.DBConn().Prepare("select * from tbl_user where user_name=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Printf("stmt exec failed %s\n", err.Error())
		return false
	} else if rows == nil {
		fmt.Printf("stmt exec success but username %s not found\n", username)
		return false
	}

	rowResult := mysql.ParseRow(rows)
	if len(rowResult) > 0 && string(rowResult[0]["user_pwd"].([]byte)) == encodedPasswd {
		return true
	}
	return false
}

// GetUserInfo query userinfo by username
func GetUserInfo(username string) (User, error) {
	user := User{}
	stmt, err := mysql.DBConn().Prepare("select user_name, signup_at from tbl_user where user_name=? limit 1")
	if err != nil {
		fmt.Printf("stmt prepare error %s\n", err.Error())
		return user, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(username).Scan(&user.Username, &user.SignupAt)
	if err != nil {
		return user, err
	}
	return user, nil
}
