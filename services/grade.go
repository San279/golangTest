package services

func CheckGrade(input int) string {

	switch {
	case input >= 80:
		return "A"
	case input >= 70:
		return "B"
	case input >= 60:
		return "C"
	case input >= 50:
		return "D"
	default:
		return "F"
	}
}
