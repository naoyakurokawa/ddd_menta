package userskilldm

import (
	"golang.org/x/xerrors"
)

type ExperienceYears int

const (
	HalfYear ExperienceYears = iota
	LessThanYear
	LessThanThreeYears
	LessThanFiveYears
	MoreFiveYears
)

func (e ExperienceYears) Names() []string {
	return []string{
		"半年未満",
		"1年未満",
		"3年未満",
		"5年未満",
		"5年以上",
	}
}

func (e ExperienceYears) String() string {
	return e.Names()[e]
}

const (
	experienceYearsMinNum = 0
	experienceYearsMaxNum = 5
)

// コンストラクタ
func NewExperienceYears(experienceYears int) (ExperienceYears, error) {
	if experienceYears < experienceYearsMinNum || experienceYearsMaxNum < experienceYears {
		return -1, xerrors.Errorf("experienceYears must between %d and %d", experienceYearsMinNum, experienceYearsMaxNum)
	}

	return ExperienceYears(experienceYears), nil
}
