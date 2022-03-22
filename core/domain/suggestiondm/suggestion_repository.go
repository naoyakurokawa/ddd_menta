package suggestiondm

type SuggestionRepository interface {
	Create(suggestion *Suggestion) error
	FetchByID(suggestionID SuggestionID) (*Suggestion, error)
}
