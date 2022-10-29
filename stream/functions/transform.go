package functions

import "fmt"

func IntToInt64(a int) int64 {
	return int64(a)
}

func IntToString(a int) string {
	return fmt.Sprintf("%d", a)
}
