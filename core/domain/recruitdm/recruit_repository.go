package recruitdm

type RecruitRepository interface {
	Create(recruit *Recruit) error
	FindByID(recruitID RecruitID) (*Recruit, error)
	// UpdateContractStatus(recruitID RecruitID, contractStatus RecruitStatus) error
}
