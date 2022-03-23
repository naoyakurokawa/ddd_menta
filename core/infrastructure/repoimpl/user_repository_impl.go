package repoimpl

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/datamodel"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImpl struct {
	Conn *gorm.DB
}

func NewUserRepositoryImpl(conn *gorm.DB) userdm.UserRepository {
	return &UserRepositoryImpl{Conn: conn}
}

// todo:gormのスライスによるinsertは以下のエラー発生 要調査
// エラー：reflect: call of reflect.Value.Interface on zero Value
func (ur *UserRepositoryImpl) Create(user *userdm.User) error {
	var u datamodel.User

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password().Value()), 12)
	if err != nil {
		return err
	}
	// hex.EncodeToString([]byte("あ"))

	u.UserID = user.UserID().String()
	u.Name = user.Name()
	u.Email = user.Email().Value()
	u.Password = string(hash)
	u.Profile = user.Profile()
	u.CreatedAt = user.CreatedAt()
	u.UserCareers = user.UserCareers()
	u.UserSkills = user.UserSkills()

	// User登録
	if err := ur.Conn.Create(&u).Error; err != nil {
		return err
	}
	// UserCareer登録
	for _, uc := range u.UserCareers {
		userCareer := &datamodel.UserCareer{
			UserCareerID: uc.UserCareerID().Value(),
			UserID:       u.UserID,
			From:         uc.From(),
			To:           uc.To(),
			Detail:       uc.Detail(),
			CreatedAt:    uc.CreatedAt(),
		}
		if err := ur.Conn.Create(&userCareer).Error; err != nil {
			return err
		}
	}

	// UserSkill登録
	for _, us := range u.UserSkills {
		userSkill := &datamodel.UserSkill{
			UserSkillID:     us.UserSkillID().Value(),
			UserID:          u.UserID,
			Tag:             us.Tag(),
			Assessment:      us.Assessment(),
			ExperienceYears: us.Assessment(),
			CreatedAt:       us.CreatedAt(),
		}
		if err := ur.Conn.Create(&userSkill).Error; err != nil {
			return err
		}
	}
	return err
}

func (ur *UserRepositoryImpl) FetchById(userID userdm.UserID) (*userdm.User, error) {
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
	if err := ur.Conn.Where("user_id = ?", userID.String()).Find(&dataModelUser).Error; err != nil {
		return nil, err
	}
	user, err := userdm.NewUser(
		userdm.UserIDType(dataModelUser.UserID),
		dataModelUser.Name,
		userdm.EmailType(dataModelUser.Email),
		userdm.PasswordType(dataModelUser.Password),
		dataModelUser.Profile,
		dataModelUser.UserCareers,
		dataModelUser.UserSkills)
	if err != nil {
		return nil, err
	}
	return user, nil
}
