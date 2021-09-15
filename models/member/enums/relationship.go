package enums

type Relationship int

const (
	Father Relationship = iota + 1
	Mother
	Brother
	Sister
	Wife
	Husband
	Son
	Daughter
	Friend
)
