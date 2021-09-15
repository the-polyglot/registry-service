package member

type EmploymentInfo struct {
	employerId     string
	designation    string
	termsOfService uint8
}

func NewEmploymentInfo(employerId, designation string, termsOfService uint8) EmploymentInfo {
	return EmploymentInfo{employerId, designation, termsOfService}
}