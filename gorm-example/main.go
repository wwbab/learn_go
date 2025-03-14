package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 定义用户模型
type User struct {
    ID   uint
    Name string
    Age  int
}

func main() {
    // 数据库连接信息
    dsn := "user:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
    // 替换为你自己的数据库用户名、密码、主机、端口和数据库名

    // 连接数据库
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

    // 自动迁移模型
    db.AutoMigrate(&User{})

    // 创建用户
    user := User{Name: "John Doe", Age: 30}
    result := db.Create(&user)
    if result.Error != nil {
        log.Fatalf("failed to create user: %v", result.Error)
    }
    log.Printf("created user with ID: %d", user.ID)

    // 查询用户
    var foundUser User
    db.First(&foundUser, user.ID)
    log.Printf("found user: %+v", foundUser)
}