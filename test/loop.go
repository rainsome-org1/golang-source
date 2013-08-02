package main
import "fmt"

func jump() (a int, err error){
	i := 0

	i++
	i++
	i++
	goto end
	end:
		fmt.Println("It is end")

	return  1, nil
}

func main() {

	a := []int{23, 45, 23, 34}

	for i , j := 0, len(a); i < j; i++ {
		fmt.Println(i, a[i])
	}

	fmt.Println("This is a loop until 100")
	sum := 0 
	for {
		sum++
		if sum > 100 {
			break
		}
	}

	p, err := jump()

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("No error: ", p)
	}
}
