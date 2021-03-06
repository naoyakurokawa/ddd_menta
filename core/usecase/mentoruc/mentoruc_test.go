package mentoruc

import (
	"time"

	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"golang.org/x/xerrors"
)

type userParams struct {
	userID    userdm.UserID
	name      string
	email     userdm.Email
	password  userdm.Password
	profile   string
	createdAt time.Time
}

type mentorParams struct {
	userID       userdm.UserID
	mentorID     mentordm.MentorID
	title        string
	mainImg      string
	subImg       string
	category     string
	detail       string
	mentorSkills []MentorSkill
	plans        []Plan
	createdAt    time.Time
}

var (
	up userParams
	mp mentorParams
)

func setupUser() error {
	//ユーザー
	userID := userdm.NewUserID()
	email, err := userdm.NewEmail("test@gmail.com")
	if err != nil {
		return xerrors.New("error NewEmail")
	}
	password, err := userdm.NewPassword("test12345678")
	if err != nil {
		return xerrors.New("error NewPassword")
	}
	up = userParams{
		userID,
		"テストユーザー",
		email,
		password,
		"テストユーザーです",
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}
	return nil
}

func setupMentor() error {
	var mentrSkills []MentorSkill
	var plans []Plan
	mentorID := mentordm.NewMentorID()

	m := MentorSkill{
		"Golang",
		"5",
		"5",
	}
	mentrSkills = append(mentrSkills, m)

	p := Plan{
		"DDDのメンタリング",
		"設計",
		"DDD",
		"DDDの設計手法を学べます",
		"2",
		"1000",
		"1",
	}
	plans = append(plans, p)

	mp = mentorParams{
		up.userID,
		mentorID,
		"プログラミング全般のメンタリング",
		"/main.jpg",
		"/sub.jpg",
		"プログライミング",
		"設計・開発・テストの一覧をサポートできます",
		mentrSkills,
		plans,
		time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}
	return nil
}
