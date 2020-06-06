package main
import "fmt"
import "math"

func main(){

	fmt.Println("this is for packages")
	fmt.Println(math.Floor(2.7))
	fmt.Println("sublime is still the best")
	var yeet = "tea boii"
	fmt.Println(yeet)
	fmt.Println("hello worl")
	fmt.Println(oddoreven(55))




}

func oddoreven(a int) bool{

	if(a%2 == 0){

		return true
	}else{
		return false
	}
}