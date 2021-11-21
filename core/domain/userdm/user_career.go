package userdm

import (
	"regexp"
	"time"

	"golang.org/x/xerrors"
)

type UserCareer struct {
	userCareerID UserCareerID
	userID       UserID
	from         time.Time
	to           time.Time
	detail       string
	createdAt    time.Time
}

const detailMaxLength = 1000

var oldestCareerYear = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)

// NewUserCareer user_careerのコンストラクタ
func NewUserCareer(userCareerID UserCareerID, userID UserID, from string, to string, detail string) (*UserCareer, error) {
	//入力データチェック
	if len(userCareerID) == 0 {
		return nil, xerrors.New("userCareerID must not be empty")
	}
	if len(userID) == 0 {
		return nil, xerrors.New("userID must not be empty")
	}
	if len(from) == 0 {
		return nil, xerrors.New("career period must not be empty")
	}

	if !dateCheck(from) {
		return nil, xerrors.New("career from must be in date format")
	}
	if len(to) == 0 {
		return nil, xerrors.New("career period must not be empty")
	}
	if !dateCheck(to) {
		return nil, xerrors.New("career to must be in date format")
	}
	fromTime := stringToTime(from)
	toTime := stringToTime(to)
	if fromTime.After(toTime) {
		return nil, xerrors.New("career from must smaller than career to")
	}
	//1970年以上となっているか確認
	if fromTime.Before(oldestCareerYear) || toTime.Before(oldestCareerYear) {
		return nil, xerrors.Errorf("career year must larger than %d", oldestCareerYear.Year())
	}
	if len(detail) == 0 {
		return nil, xerrors.New("career detail must not be empty")
	}
	if len(detail) > detailMaxLength {
		return nil, xerrors.Errorf("detail must less than %d: %s", detailMaxLength, detail)
	}
	now := time.Now()
	userCareer := &UserCareer{
		userCareerID: userCareerID,
		userID:       userID,
		from:         fromTime,
		to:           toTime,
		detail:       detail,
		createdAt:    now,
	}

	return userCareer, nil
}

// 日付チェック処理
func dateCheck(dateStr string) bool {
	// 削除する文字列を定義
	reg := regexp.MustCompile(`[-|/|:| |　]`)
	// 指定文字を削除
	str := reg.ReplaceAllString(dateStr, "")
	// 数値の値に対してフォーマットを定義
	format := string([]rune("20060102150405")[:len(str)])
	// パース処理 → 日付ではない場合はエラー
	_, err := time.Parse(format, str)
	return err == nil
}

const layout = "2006-01-02 15:04:05"

func stringToTime(str string) time.Time {
	t, _ := time.Parse(layout, str)
	return t
}

func (u *UserCareer) UserCareerID() UserCareerID {
	return u.userCareerID
}

func (u *UserCareer) UserID() UserID {
	return u.userID
}

func (u *UserCareer) From() time.Time {
	return u.from
}

func (u *UserCareer) To() time.Time {
	return u.to
}

func (u *UserCareer) Detail() string {
	return u.detail
}

func (u *UserCareer) CreatedAt() time.Time {
	return u.createdAt
}
