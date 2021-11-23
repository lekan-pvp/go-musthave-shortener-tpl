package shorter

import "fmt"

const endpoint = "http://localhost:8080/"

func Shorting(id int) string {
	return fmt.Sprintf("%s%d", endpoint, id)
}
