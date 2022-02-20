package mentordm

import (
	"github.com/naoyakurokawa/ddd_menta/core/domain/sharedvo"
	"golang.org/x/xerrors"
)

type Plan struct {
	planID     PlanID
	title      string
	category   string
	tag        string
	detial     string
	planType   PlanType
	price      uint16
	planStatus PlanStatus
	createdAt  sharedvo.CreatedAt
}

const planTitleMaxLength = 255
const planDetialMaxLength = 2000

func NewPlan(
	planID PlanID,
	title string,
	category string,
	tag string,
	detial string,
	planType PlanType,
	price uint16,
	planStatus PlanStatus,
) (*Plan, error) {
	//入力データチェック
	if isEmpty(title) {
		return nil, xerrors.New("title must not be empty")
	}
	if isOver(title, planTitleMaxLength) {
		return nil, xerrors.Errorf("title must less than %d: %s", planTitleMaxLength, title)
	}
	if isEmpty(category) {
		return nil, xerrors.New("category must not be empty")
	}
	if isEmpty(tag) {
		return nil, xerrors.New("tag must not be empty")
	}
	if isEmpty(detial) {
		return nil, xerrors.New("detial must not be empty")
	}
	if isOver(detial, planDetialMaxLength) {
		return nil, xerrors.Errorf("detial must less than %d: %s", planDetialMaxLength, detial)
	}
	if isZero(price) {
		return nil, xerrors.New("price must more than 0")
	}

	plan := &Plan{
		planID:     planID,
		title:      title,
		category:   category,
		tag:        tag,
		detial:     detial,
		planType:   planType,
		price:      price,
		planStatus: planStatus,
		createdAt:  sharedvo.NewCreatedAt(),
	}

	return plan, nil
}

func ReconstructPlan(
	planID string,
	title string,
	category string,
	tag string,
	detial string,
	planType uint16,
	price uint16,
	planStatus uint16,
) (*Plan, error) {
	if isEmpty(title) {
		return nil, xerrors.New("title must not be empty")
	}
	if isOver(title, planTitleMaxLength) {
		return nil, xerrors.Errorf("title must less than %d: %s", planTitleMaxLength, title)
	}
	if isEmpty(category) {
		return nil, xerrors.New("category must not be empty")
	}
	if isEmpty(tag) {
		return nil, xerrors.New("tag must not be empty")
	}
	if isEmpty(detial) {
		return nil, xerrors.New("detial must not be empty")
	}
	if isOver(detial, planDetialMaxLength) {
		return nil, xerrors.Errorf("detial must less than %d: %s", planDetialMaxLength, detial)
	}
	if isZero(price) {
		return nil, xerrors.New("price must more than 0")
	}

	castedPlanID, err := NewPlanIDByVal(planID)
	if err != nil {
		return nil, xerrors.New("error NewMentorIDByVal")
	}
	planTypeIns, err := NewPlanType(planType)
	if err != nil {
		return nil, xerrors.New("error NewMentorIDByVal")
	}
	planStatusIns, err := NewPlanStatus(planStatus)
	if err != nil {
		return nil, xerrors.New("error NewMentorIDByVal")
	}
	plan := &Plan{
		planID:     castedPlanID,
		title:      title,
		category:   category,
		tag:        tag,
		detial:     detial,
		planType:   planTypeIns,
		price:      price,
		planStatus: planStatusIns,
		createdAt:  sharedvo.NewCreatedAt(),
	}

	return plan, nil
}

func (p *Plan) PlanID() PlanID {
	return p.planID
}
func (p *Plan) Title() string {
	return p.title
}

func (p *Plan) Category() string {
	return p.category
}

func (p *Plan) Tag() string {
	return p.tag
}

func (p *Plan) Detial() string {
	return p.detial
}

func (p *Plan) PlanType() PlanType {
	return p.planType
}

func (p *Plan) Price() uint16 {
	return p.price
}

func (p *Plan) PlanStatus() PlanStatus {
	return p.planStatus
}

func (p *Plan) CreatedAt() sharedvo.CreatedAt {
	return p.createdAt
}
