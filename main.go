package main

import (
	"fmt"
	"log"

	"main/HawkORM"
)

func main() {
	// データベース初期化
	// remember to add parseTime=true to parse time correctlt
	testDB, err := HawkORM.OpenMySQL("user:pass@localhost:3306/databasename?parseTime=true")
	if err != nil {
		log.Fatalf("connect to db failed!")
	}

	// equal to: SELECT id, name, email, created_at, updated_at FROM users WHERE (id = "userId" OR name = "testname") ORDER BY id ASC LIMIT 1
	// モデル定義
	type User struct {
		ID    string `hawkorm:"primarykey"`
		Name  string
		Email string
	}

	var users []User
	testDB.Select(&User{}).WhereOr(&User{ID: "userId", Name: "testname"}).First(&users)

	// レコードの作成
	insertTest := []User{
		{ID: "userid1", Name: "user1", Email: "test@gmail.com"},
		{ID: "userid2", Name: "user2", Email: "test2@gmail.com"},
	}
	// equal to: INSERT INTO users (id, name, email) VALUES ("userid1", "user1", "test@gmail.com"), ("userid2", "user2", "test2@gmail.com")
	res, err := testDB.Insert(&User{}).SetData(insertTest).Exec()
	fmt.Println(res)

}
