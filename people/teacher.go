package people

type Teacher interface {
	FinalGrade(grades []int) int
}