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
		mentorTag []string,
		mentorAssessment []string,
		mentorExperienceYears []string,
		plans []Plan,
	) error
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
	mentorTag []string,
	mentorAssessment []string,
	mentorExperienceYears []string,
	mentorPlans []Plan,
) error {
	mentorID := mentordm.NewMentorID()
	userIDIns := userdm.UserIDType(userID)

	var (
		mentorSkills []mentordm.MentorSkill
		plans        []mentordm.Plan
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
		mentorSkills,
		plans,
	)
	if err != nil {
		return err
	}

	//メンタースキル作成
	if len(mentorTag) > 0 {
		for i := 0; i < len(mentorTag); i++ {
			uintMentorAssessment, err := mentordm.StrCastUint(mentorAssessment[i])
			if err != nil {
				return err
			}

			uintMentorExperienceYears, err := mentordm.StrCastUint(mentorExperienceYears[i])
			if err != nil {
				return err
			}

			mentorExperienceYears, err := mentordm.NewExperienceYears(uintMentorExperienceYears)
			if err != nil {
				return err
			}

			mentor.AddMentorSkill(
				mentorTag[i],
				uintMentorAssessment,
				mentorExperienceYears,
			)
		}
	}

	//メンタープラン追加
	if len(mentorPlans) > 0 {
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
	}

	//最終的にinfraのCreateメソッドを実行することになる
	err = mu.mentorRepo.Create(mentor)
	if err != nil {
		return err
	}

	return nil
}
