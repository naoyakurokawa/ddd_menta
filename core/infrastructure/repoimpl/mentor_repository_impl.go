package repoimpl

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/datamodel"
)

type MentorRepositoryImpl struct {
	Conn *gorm.DB
}

func NewMentorRepositoryImpl(conn *gorm.DB) mentordm.MentorRepository {
	return &MentorRepositoryImpl{Conn: conn}
}

func (mr *MentorRepositoryImpl) Create(mentor *mentordm.Mentor) (*mentordm.Mentor, error) {
	var m datamodel.Mentor
	m.UserID = string(mentor.UserID())
	m.MentorID = string(mentor.MentorID())
	m.Title = mentor.Title()
	m.MainImg = mentor.MainImg()
	m.SubImg = mentor.SubImg()
	m.Category = mentor.Category()
	m.Detail = mentor.Detail()
	m.CreatedAt = time.Time(mentor.CreatedAt())
	m.Plans = mentor.Plans()
	m.MentorSkills = mentor.MentorSkills()
	// メンター概要登録
	if err := mr.Conn.Create(&m).Error; err != nil {
		return nil, err
	}
	// メンタープラン登録
	for i := 0; i < len(m.Plans); i++ {
		plans := &datamodel.Plan{
			PlanID:     string(m.Plans[i].PlanID()),
			MentorID:   m.MentorID,
			Title:      m.Plans[i].Title(),
			Category:   m.Plans[i].Category(),
			Tag:        m.Plans[i].Tag(),
			Detail:     m.Plans[i].Detial(),
			PlanType:   uint16(m.Plans[i].PlanType()),
			Price:      m.Plans[i].Price(),
			PlanStatus: uint16(m.Plans[i].PlanStatus()),
			CreatedAt:  time.Time(m.Plans[i].CreatedAt()),
		}
		if err := mr.Conn.Create(&plans).Error; err != nil {
			return nil, err
		}
	}
	// メンタースキル登録
	for i := 0; i < len(m.MentorSkills); i++ {
		mentorSkills := &datamodel.MentorSkill{
			MentorSkillID:   string(m.MentorSkills[i].MentorSkillID()),
			MentorID:        m.MentorID,
			Tag:             m.MentorSkills[i].Tag(),
			Assessment:      m.MentorSkills[i].Assessment(),
			ExperienceYears: mentordm.ExperienceYears.Uint16(m.MentorSkills[i].ExperienceYears()),
			CreatedAt:       time.Time(m.MentorSkills[i].CreatedAt()),
		}
		if err := mr.Conn.Create(&mentorSkills).Error; err != nil {
			return nil, err
		}
	}

	return mentor, nil
}

func (mr *MentorRepositoryImpl) FindByID(mentorID mentordm.MentorID) (*mentordm.Mentor, error) {
	var (
		dataModelMentorSkills []mentordm.MentorSkill
		dataModelPlan         []mentordm.Plan
	)
	dataModeMentor := &datamodel.Mentor{
		UserID:       "",
		MentorID:     "",
		Title:        "",
		MainImg:      "",
		SubImg:       "",
		Category:     "",
		Detail:       "",
		CreatedAt:    time.Now(),
		Plans:        dataModelPlan,
		MentorSkills: dataModelMentorSkills,
	}
	if err := mr.Conn.Where("mentor_id = ?", string(mentorID)).Find(&dataModeMentor).Error; err != nil {
		return nil, err
	}
	mentorID, err := mentordm.NewMentorIDByVal(dataModeMentor.MentorID)
	if err != nil {
		return nil, err
	}
	mentor, err := mentordm.NewMentor(
		mentorID,
		userdm.UserIDType(dataModeMentor.UserID),
		dataModeMentor.Title,
		dataModeMentor.MainImg,
		dataModeMentor.SubImg,
		dataModeMentor.Category,
		dataModeMentor.Detail,
		dataModelMentorSkills,
		dataModelPlan,
	)
	if err != nil {
		return nil, err
	}
	return mentor, nil
}
