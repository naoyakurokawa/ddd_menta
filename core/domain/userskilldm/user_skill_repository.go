package userskilldm

type UserSkillRepository interface {
	Create(userSkills []*UserSkill) ([]*UserSkill, error)
}
