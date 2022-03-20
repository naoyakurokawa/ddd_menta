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
	m.UserID = mentor.UserID().String()
	m.MentorID = mentor.MentorID().String()
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
	for _, p := range m.Plans {
		plans := &datamodel.Plan{
			PlanID:     p.PlanID().String(),
			MentorID:   m.MentorID,
			Title:      p.Title(),
			Category:   p.Category(),
			Tag:        p.Tag(),
			Detail:     p.Detial(),
			PlanType:   p.PlanType().Uint16(),
			Price:      p.Price(),
			PlanStatus: p.PlanStatus().Uint16(),
			CreatedAt:  p.CreatedAt().Time(),
		}
		if err := mr.conn.Create(&plans).Error; err != nil {
			return err
		}
	}
	// メンタースキル登録
	for _, s := range m.MentorSkills {
		mentorSkills := &datamodel.MentorSkill{
			MentorSkillID:   s.MentorSkillID().String(),
			MentorID:        m.MentorID,
			Tag:             s.Tag(),
			Assessment:      s.Assessment(),
			ExperienceYears: mentordm.ExperienceYears.Uint16(s.ExperienceYears()),
			CreatedAt:       s.CreatedAt().Time(),
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
	if err := mr.conn.Where("mentor_id = ?", mentorID.String()).Find(&dataModeMentor).Error; err != nil {
		return nil, err
	}

	if err := mr.conn.Where("mentor_id = ?", mentorID.String()).Find(&dataModelPlans).Error; err != nil {
		return nil, err
	}
	mentorPlans := make([]mentordm.Plan, len(dataModelPlans))
	for i, p := range dataModelPlans {
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
		mentorPlans[i] = *plan
	}

	if err := mr.conn.Where("mentor_id = ?", string(mentorID)).Find(&dataModelMentorSkills).Error; err != nil {
		return nil, err
	}
	mentorSkills := make([]mentordm.MentorSkill, len(dataModelMentorSkills))
	for i, ms := range dataModelMentorSkills {
		mentorSkill, err := mentordm.ReconstructMentorSkill(
			ms.MentorSkillID,
			ms.Tag,
			ms.Assessment,
			ms.ExperienceYears,
		)
		if err != nil {
			return nil, err
		}
		mentorSkills[i] = *mentorSkill
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
