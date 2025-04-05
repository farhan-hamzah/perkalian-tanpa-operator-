package main
import "fmt"

func perkalianCepat(a, b int)int{
	if b == 0{
		return 0
	}
	return a + perkalianCepat(a, b-1)
}
func main(){
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Print(perkalianCepat(a, b))
}