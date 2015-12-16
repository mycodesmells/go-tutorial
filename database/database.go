package database

import (
	"time"

	"github.com/mycodesmells/go-tutorial/people"
	"fmt"
)

func MakeQuery() string {
	time.Sleep(5 * time.Second)
	return "--> Database query response <--"
}

func SaveFinalGrade(t people.Teacher, s people.Student) string {
	gr := t.FinalGrade(s.Grades())
	fn := s.FullName()
	return fmt.Sprintf("Giving %d for %s", gr, fn)
}
