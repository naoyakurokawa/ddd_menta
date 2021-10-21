package userdm

import "golang.org/x/xerrors"

type UserCareer struct {
	Id        int
	UserId    string
	From      string
	To        string
	Detail    string
	CreatedAt string
}

// NewUserCareer user_careerのコンストラクタ
func NewUserCareer(userId string, from string, to string, detail string) (*UserCareer, error) {
	//入力データチェック
	if len(userId) == 0 {
		return nil, xerrors.New("userId must not be empty")
	}
	if len(from) == 0 {
		return nil, xerrors.New("career period must not be empty")
	}
	if len(to) == 0 {
		return nil, xerrors.New("career period must not be empty")
	}
	if len(detail) == 0 {
		return nil, xerrors.New("career detail must not be empty")
	}

	user_career := &UserCareer{
		UserId: userId,
		From:   from,
		To:     to,
		Detail: detail,
	}

	return user_career, nil
}
