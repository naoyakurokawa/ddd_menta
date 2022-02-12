package mentoruc

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/userdm"
	"github.com/naoyakurokawa/ddd_menta/core/util"
)

type MentorCreateUsecase interface {
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
		planTitle []string,
		planCategory []string,
		planTag []string,
		planDetial []string,
		planType []string,
		planPrice []string,
		planStatus []string,
	) (*mentordm.Mentor, error)
}

type MentorCreateUsecaseImpl struct {
	mentorRepo mentordm.MentorRepository
}

// user usecaseのコンストラクタ
func NewMentorCreateUsecase(mentorRepo mentordm.MentorRepository) MentorCreateUsecase {
	return &MentorCreateUsecaseImpl{mentorRepo: mentorRepo}
}

// Create userを保存するときのユースケース
func (mu *MentorCreateUsecaseImpl) Create(
	userID string,
	title string,
	mainImg string,
	subImg string,
	category string,
	detail string,
	mentorTag []string,
	mentorAssessment []string,
	mentorExperienceYears []string,
	planTitle []string,
	planCategory []string,
	planTag []string,
	planDetial []string,
	planType []string,
	planPrice []string,
	planStatus []string,
) (*mentordm.Mentor, error) {
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
		return nil, err
	}

	//メンタースキル作成
	if len(mentorTag) > 0 {
		for i := 0; i < len(mentorTag); i++ {
			uintMentorAssessment, err := util.CastUint(mentorAssessment[i])
			if err != nil {
				return nil, err
			}

			uintMentorExperienceYears, err := util.CastUint(mentorExperienceYears[i])
			if err != nil {
				return nil, err
			}

			mentorExperienceYears, err := mentordm.NewExperienceYears(uintMentorExperienceYears)
			if err != nil {
				return nil, err
			}

			mentor.AddMentorSkill(
				mentorTag[i],
				uintMentorAssessment,
				mentorExperienceYears,
			)
		}
	}

	//メンタープラン追加
	if len(planTitle) > 0 {
		for i := 0; i < len(planTitle); i++ {
			uintPlanType, err := util.CastUint(planType[i])
			if err != nil {
				return nil, err
			}

			planType, err := mentordm.NewPlanType(uintPlanType)
			if err != nil {
				return nil, err
			}

			price, err := util.CastUint(planPrice[i])
			if err != nil {
				return nil, err
			}

			uintPlanStatus, err := util.CastUint(planStatus[i])
			if err != nil {
				return nil, err
			}
			planStatus, err := mentordm.NewPlanStatus(uintPlanStatus)
			if err != nil {
				return nil, err
			}

			mentor.AddPlan(
				planTitle[i],
				planCategory[i],
				planTag[i],
				planDetial[i],
				planType,
				price,
				planStatus,
			)
		}
	}

	//最終的にinfraのCreateメソッドを実行することになる
	createdMentor, err := mu.mentorRepo.Create(mentor)
	if err != nil {
		return nil, err
	}

	return createdMentor, nil
}
