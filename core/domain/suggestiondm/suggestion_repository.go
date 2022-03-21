package suggestiondm

type SuggestionRepository interface {
	Create(suggestion *Suggestion) error
	FindByID(suggestionID SuggestionID) (*Suggestion, error)
}
