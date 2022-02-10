package mentordm

import (
	"golang.org/x/xerrors"
)

type ExperienceYears uint16

const (
	HalfYear ExperienceYears = iota + 1
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
	return e.Names()[e-1]
}

func (e ExperienceYears) Uint16() uint16 {
	return uint16(e)
}

const (
	experienceYearsMinNum = 0
	experienceYearsMaxNum = 5
)

func NewExperienceYears(experienceYears uint16) (ExperienceYears, error) {
	if experienceYears < experienceYearsMinNum || experienceYearsMaxNum < experienceYears {
		return 0, xerrors.Errorf("experienceYears must be between %d and %d", experienceYearsMinNum, experienceYearsMaxNum)
	}

	return ExperienceYears(experienceYears), nil
}
