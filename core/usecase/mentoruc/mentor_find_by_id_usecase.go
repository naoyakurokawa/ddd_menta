package mentoruc

import "github.com/naoyakurokawa/ddd_menta/core/domain/mentordm"

type MentorFindByIDUsecase interface {
	FindByID(mentorID mentordm.MentorID) (*mentordm.Mentor, error)
}

type MentorFindByIDUsecaseImpl struct {
	mentorRepo mentordm.MentorRepository
}

func NewmentorFindByIDUsecase(mentorRepo mentordm.MentorRepository) MentorFindByIDUsecase {
	return &MentorFindByIDUsecaseImpl{mentorRepo: mentorRepo}
}

func (mu *MentorFindByIDUsecaseImpl) FindByID(mentorID mentordm.MentorID) (*mentordm.Mentor, error) {
	return mu.mentorRepo.FindByID(mentorID)
}
