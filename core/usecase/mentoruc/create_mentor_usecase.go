package mentoruc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
)

type CreateMentorUsecase interface {
	Create(
		userID string,
		title string,
		mainImg string,
		subImg string,
		category string,
		detial string,
		mentorSkills []MentorSkill,
		plans []Plan,
	) error
}

type MentorSkill struct {
	Tag             string
	Assessment      string
	ExperienceYears string
}

type Plan struct {
	PlanTitle    string
	PlanCategory string
	PlanTag      string
	PlanDetial   string
	PlanType     string
	PlanPrice    string
	PlanStatus   string
}

type CreateMentorUsecaseImpl struct {
	mentorRepo mentordm.MentorRepository
}

// user usecaseのコンストラクタ
func NewCreateMentorUsecase(mentorRepo mentordm.MentorRepository) CreateMentorUsecase {
	return &CreateMentorUsecaseImpl{mentorRepo: mentorRepo}
}

// Create userを保存するときのユースケース
func (mu *CreateMentorUsecaseImpl) Create(
	userID string,
	title string,
	mainImg string,
	subImg string,
	category string,
	detail string,
	mentorSkills []MentorSkill,
	mentorPlans []Plan,
) error {
	mentorID := mentordm.NewMentorID()
	userIDIns := userdm.UserIDType(userID)

	var (
		initMentorSkills []mentordm.MentorSkill
		initPlans        []mentordm.Plan
	)

	// メンター作成
	mentor, err := mentordm.NewMentor(
		mentorID,
		userIDIns,
		title,
		mainImg,
		subImg,
		category,
		detail,
		initMentorSkills,
		initPlans,
	)
	if err != nil {
		return err
	}

	//メンタースキル作成
	for _, m := range mentorSkills {
		uintMentorAssessment, err := mentordm.StrCastUint(m.Assessment)
		if err != nil {
			return err
		}

		uintMentorExperienceYears, err := mentordm.StrCastUint(m.ExperienceYears)
		if err != nil {
			return err
		}

		mentorExperienceYears, err := mentordm.NewExperienceYears(uintMentorExperienceYears)
		if err != nil {
			return err
		}

		mentor.AddMentorSkill(
			m.Tag,
			uintMentorAssessment,
			mentorExperienceYears,
		)
	}

	//メンタープラン追加
	for _, p := range mentorPlans {
		uintPlanType, err := mentordm.StrCastUint(p.PlanType)
		if err != nil {
			return err
		}

		planType, err := mentordm.NewPlanType(uintPlanType)
		if err != nil {
			return err
		}

		price, err := mentordm.StrCastUint(p.PlanPrice)
		if err != nil {
			return err
		}

		uintPlanStatus, err := mentordm.StrCastUint(p.PlanStatus)
		if err != nil {
			return err
		}
		planStatus, err := mentordm.NewPlanStatus(uintPlanStatus)
		if err != nil {
			return err
		}

		mentor.AddPlan(
			p.PlanTitle,
			p.PlanCategory,
			p.PlanTag,
			p.PlanDetial,
			planType,
			price,
			planStatus,
		)
	}

	//最終的にinfraのCreateメソッドを実行することになる
	err = mu.mentorRepo.Create(mentor)
	if err != nil {
		return err
	}

	return nil
}
