package database_test

import (
	"testing"

	"github.com/mycodesmells/go-tutorial/database"
)

type FakeTeacher struct{}

func (ft FakeTeacher) FinalGrade(grades []int) int {
	return 1
}

type FakeStudent struct{}

func (fs FakeStudent) FullName() string {
	return "Jerry Lemon"
}

func (fs FakeStudent) Grades() []int {
	return []int {5}
}


func TestShouldCreateFinalGradeQuery(t *testing.T) {
	ft := FakeTeacher{}
	fs := FakeStudent{}

	query := database.SaveFinalGrade(ft, fs)
	exp := "Giving 1 for Jerry Lemon"
	if query != exp {
		t.Errorf("Incorrect query. Expected %v but got %v", exp, query)
	}
}