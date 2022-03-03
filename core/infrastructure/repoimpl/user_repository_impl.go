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
func (ur *UserRepositoryImpl) Create(user *userdm.User) (*userdm.User, error) {
	var u datamodel.User

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password().Value()), 12)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	// UserCareer登録
	for i := 0; i < len(u.UserCareers); i++ {
		userCareer := &datamodel.UserCareer{
			UserCareerID: userdm.UserCareerID.Value(u.UserCareers[i].UserCareerID()),
			UserID:       u.UserID,
			From:         u.UserCareers[i].From(),
			To:           u.UserCareers[i].To(),
			Detail:       u.UserCareers[i].Detail(),
			CreatedAt:    u.UserCareers[i].CreatedAt(),
		}
		if err := ur.Conn.Create(&userCareer).Error; err != nil {
			return nil, err
		}
	}
	// UserSkill登録
	for i := 0; i < len(u.UserSkills); i++ {
		userSkill := &datamodel.UserSkill{
			UserSkillID:     userdm.UserSkillID.Value(u.UserSkills[i].UserSkillID()),
			UserID:          u.UserID,
			Tag:             u.UserSkills[i].Tag(),
			Assessment:      u.UserSkills[i].Assessment(),
			ExperienceYears: userdm.ExperienceYears.Uint16(u.UserSkills[i].ExperienceYears()),
			CreatedAt:       u.UserSkills[i].CreatedAt(),
		}
		if err := ur.Conn.Create(&userSkill).Error; err != nil {
			return nil, err
		}
	}
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
