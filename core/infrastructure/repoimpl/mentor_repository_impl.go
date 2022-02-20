package repoimpl

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/datamodel"
)

type MentorRepositoryImpl struct {
	conn *gorm.DB
}

func NewMentorRepositoryImpl(conn *gorm.DB) mentordm.MentorRepository {
	return &MentorRepositoryImpl{conn: conn}
}

func (mr *MentorRepositoryImpl) Create(mentor *mentordm.Mentor) error {
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
	if err := mr.conn.Create(&m).Error; err != nil {
		return err
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
		if err := mr.conn.Create(&plans).Error; err != nil {
			return err
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
		if err := mr.conn.Create(&mentorSkills).Error; err != nil {
			return err
		}
	}

	return nil
}

func (mr *MentorRepositoryImpl) FindByID(mentorID mentordm.MentorID) (*mentordm.Mentor, error) {
	dataModeMentor := &datamodel.Mentor{}
	dataModelPlans := []datamodel.Plan{}
	dataModelMentorSkills := []datamodel.MentorSkill{}
	if err := mr.conn.Where("mentor_id = ?", string(mentorID)).Find(&dataModeMentor).Error; err != nil {
		return nil, err
	}

	if err := mr.conn.Where("mentor_id = ?", string(mentorID)).Find(&dataModelPlans).Error; err != nil {
		return nil, err
	}
	mentorPlans := make([]mentordm.Plan, len(dataModelPlans))
	for _, p := range dataModelPlans {
		plan, err := mentordm.ReconstructPlan(
			p.PlanID,
			p.Title,
			p.Category,
			p.Tag,
			p.Detail,
			p.PlanType,
			p.Price,
			p.PlanStatus,
		)
		if err != nil {
			return nil, err
		}
		mentorPlans = append(mentorPlans, *plan)
	}

	if err := mr.conn.Where("mentor_id = ?", string(mentorID)).Find(&dataModelMentorSkills).Error; err != nil {
		return nil, err
	}
	mentorSkills := make([]mentordm.MentorSkill, len(dataModelMentorSkills))
	for _, ms := range dataModelMentorSkills {
		mentorSkill, err := mentordm.ReconstructMentorSkill(
			ms.MentorSkillID,
			ms.Tag,
			ms.Assessment,
			ms.ExperienceYears,
		)
		if err != nil {
			return nil, err
		}
		mentorSkills = append(mentorSkills, *mentorSkill)
	}

	mentor, err := mentordm.Reconstruct(
		dataModeMentor.MentorID,
		dataModeMentor.UserID,
		dataModeMentor.Title,
		dataModeMentor.MainImg,
		dataModeMentor.SubImg,
		dataModeMentor.Category,
		dataModeMentor.Detail,
		mentorSkills,
		mentorPlans,
	)
	if err != nil {
		return nil, err
	}

	return mentor, nil
}
