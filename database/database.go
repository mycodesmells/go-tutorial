package database
import "time"

func MakeQuery() string {
	time.Sleep(5 * time.Second)
	return "--> Database query response <--"
}
