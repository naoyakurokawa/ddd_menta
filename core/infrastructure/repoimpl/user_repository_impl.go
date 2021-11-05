package repoimpl

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"time"
)

type UserRepositoryImpl struct {
	Conn *gorm.DB
}

func NewUserRepositoryImpl(conn *gorm.DB) userdm.UserRepository {
	return &UserRepositoryImpl{Conn: conn}
}

// NewDB DBと接続する
func NewDB() *gorm.DB {
	// err := godotenv.Load()

	// if err != nil {
	// 	fmt.Printf("読み込み出来ませんでした: %v", err)
	// }

	// USER := os.Getenv("DB_USER")
	// PASS := os.Getenv("DB_PASS")
	// PROTOCOL := "tcp(" + os.Getenv("DB_ADDRESS") + ")"
	// DB_NAME := os.Getenv("DB_NAME")
	// CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DB_NAME

	// db, err := gorm.Open("mysql", CONNECT)
	// if err != nil {
	// 	panic(err)
	// }
	db, err := gorm.Open("mysql", "ddd_menta:ddd_menta@tcp(localhost)/ddd_menta?parseTime=true")
	if err != nil {
		panic(err)
	}

	return db
}

// Create userの保存
func (ur *UserRepositoryImpl) Create(user *userdm.User) (*userdm.User, error) {
	if err := ur.Conn.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) FindByID(userID userdm.UserID) (*userdm.User, error) {
	dataModelUser := &userdm.User{
		UserID:    "",
		Name:      "",
		Email:     "",
		Password:  "",
		Profile:   "",
		CreatedAt: time.Now(),
	}
	if err := ur.Conn.Where("user_id = ?", userID).Find(&dataModelUser).Error; err != nil {
		return nil, err
	}
	return dataModelUser, nil
}
