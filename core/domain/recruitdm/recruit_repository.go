package recruitdm

type RecruitRepository interface {
	Create(recruit *Recruit) error
	FetchByID(recruitID RecruitID) (*Recruit, error)
	// UpdateContractStatus(recruitID RecruitID, contractStatus RecruitStatus) error
}
