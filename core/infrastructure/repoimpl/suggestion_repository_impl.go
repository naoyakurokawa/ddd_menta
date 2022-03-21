package repoimpl

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/naoyakurokawa/ddd_menta/core/domain/suggestiondm"
	"github.com/naoyakurokawa/ddd_menta/core/infrastructure/datamodel"
)

type SuggestionRepositoryImpl struct {
	conn *gorm.DB
}

func NewSuggestionRepositoryImpl(conn *gorm.DB) suggestiondm.SuggestionRepository {
	return &SuggestionRepositoryImpl{conn: conn}
}

func (sr *SuggestionRepositoryImpl) Create(suggestion *suggestiondm.Suggestion) error {
	var s datamodel.Suggestion
	s.SuggestionID = suggestion.SuggestionID().String()
	s.MentorID = suggestion.MentorID().String()
	s.RecruitID = suggestion.RecruitID().String()
	s.Price = suggestion.Price()
	s.SuggestionType = suggestion.SuggestionType().Uint16()
	s.Detail = suggestion.Detail()
	s.SuggestionStatus = suggestion.SuggestionStatus().Uint16()

	if err := sr.conn.Create(&s).Error; err != nil {
		return err
	}

	return nil
}

func (sr *SuggestionRepositoryImpl) FindByID(suggestionID suggestiondm.SuggestionID) (*suggestiondm.Suggestion, error) {
	dataModelSuggestion := datamodel.Suggestion{}
	if err := sr.conn.Where("suggestion_id = ?", suggestionID.String()).Find(&dataModelSuggestion).Error; err != nil {
		return nil, err
	}

	return suggestiondm.Reconstruct(
		dataModelSuggestion.SuggestionID,
		dataModelSuggestion.MentorID,
		dataModelSuggestion.RecruitID,
		dataModelSuggestion.Price,
		dataModelSuggestion.SuggestionType,
		dataModelSuggestion.Detail,
		dataModelSuggestion.SuggestionStatus,
	)

}
