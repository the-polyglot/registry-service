package enums

type TermsOfService uint32

const (
	Temporary TermsOfService = iota + 1
	Permanent
	Pensionable
	Contract
	Casual
)
