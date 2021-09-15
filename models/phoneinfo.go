package models

type PhoneInfo struct {
	primaryLine string
	otherLine   string
}

func NewPhoneInfo(primaryLine, otherLine string) PhoneInfo {
	return PhoneInfo{primaryLine, otherLine}
}
