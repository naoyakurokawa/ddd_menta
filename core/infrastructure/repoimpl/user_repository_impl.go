package repoimpl

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/datamodel"
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

// todo:gormのスライスによるinsertは以下のエラー発生 要調査
// エラー：reflect: call of reflect.Value.Interface on zero Value
func (ur *UserRepositoryImpl) Create(user *userdm.User) (*userdm.User, error) {
	var u datamodel.User

	u.UserID = user.UserID()
	u.Name = user.Name()
	u.Email = user.Email()
	u.Password = user.Password()
	u.Profile = user.Profile()
	u.CreatedAt = user.CreatedAt()
	u.UserCareers = user.UserCareers()

	tx := ur.Conn.Begin()
	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	for i := 0; i < len(u.UserCareers); i++ {
		userCareer := &datamodel.UserCareer{
			UserCareerID: userdm.UserCareerID.Value(u.UserCareers[i].UserCareerID),
			UserID:       userdm.UserID.Value(u.UserID),
			From:         u.UserCareers[i].From,
			To:           u.UserCareers[i].To,
			Detail:       u.UserCareers[i].Detail,
			CreatedAt:    u.UserCareers[i].CreatedAt,
		}
		if err := tx.Create(&userCareer).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	return user, nil
}

func (ur *UserRepositoryImpl) FindByID(userID userdm.UserID) (*userdm.User, error) {
	var dataModelCareer []userdm.UserCareer
	dataModelUser := &datamodel.User{
		UserID:      "",
		Name:        "",
		Email:       "",
		Password:    "",
		Profile:     "",
		CreatedAt:   time.Now(),
		UserCareers: dataModelCareer,
	}
	if err := ur.Conn.Where("user_id = ?", userID.Value()).Find(&dataModelUser).Error; err != nil {
		return nil, err
	}
	user, err := userdm.NewUser(
		dataModelUser.UserID,
		dataModelUser.Name,
		dataModelUser.Email,
		dataModelUser.Password,
		dataModelUser.Profile,
		dataModelUser.UserCareers)
	if err != nil {
		return nil, err
	}
	return user, nil
}
