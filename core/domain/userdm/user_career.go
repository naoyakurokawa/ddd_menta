package userdm

import (
	"regexp"
	"strconv"
	"time"

	"golang.org/x/xerrors"
)

type UserCareer struct {
	Id        int
	UserId    UserId
	From      string
	To        string
	Detail    string
	CreatedAt string
}

const detailMaxLength = 1000

// NewUserCareer user_careerのコンストラクタ
func NewUserCareer(userId UserId, from string, to string, detail string) (*UserCareer, error) {
	//入力データチェック
	if len(userId) == 0 {
		return nil, xerrors.New("userId must not be empty")
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
	fromInt := dateConvert(from)
	toInt := dateConvert(to)
	if fromInt >= toInt {
		return nil, xerrors.New("career from must smaller than career to")
	}

	if len(detail) == 0 {
		return nil, xerrors.New("career detail must not be empty")
	}
	if len(detail) > detailMaxLength {
		return nil, xerrors.Errorf("detail must less than %d: %s", detailMaxLength, detail)
	}

	userCareer := &UserCareer{
		UserId: userId,
		From:   from,
		To:     to,
		Detail: detail,
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

func dateConvert(dateStr string) int {
	reg := regexp.MustCompile(`[-|/|:| |　]`)
	str := reg.ReplaceAllString(dateStr, "")
	dateInt, _ := strconv.Atoi(str)
	return dateInt
}
