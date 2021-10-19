package userdm

type UserCareer struct {
	Id        int
	UserId    string
	From      string
	To        string
	Detail    string
	CreatedAt string
}

// NewUserCareer user_careerのコンストラクタ
func NewUserCareer(user_id string, from string, to string, detail string) (*UserCareer, error) {
	user_career := &UserCareer{
		UserId: user_id,
		From:   from,
		To:     to,
		Detail: detail,
	}

	return user_career, nil
}
