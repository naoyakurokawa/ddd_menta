package suggestiondm

type SuggestionRepository interface {
	Create(suggestion *Suggestion) error
	FindByID(suggestionID SuggestionID) (*Suggestion, error)
	// UpdateContractStatus(recruitID RecruitID, contractStatus RecruitStatus) error
}
