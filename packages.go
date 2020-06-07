package main
import "fmt"


func main(){

	var sli = []string{"hello world","hello tarun","halo ma boiii"}
	fmt.Println(sli)
	fmt.Println(slice(sli))
	fmt.Println(oddoreven(4))

	for i:=0;i<=10;i++ {

		fmt.Println("its a gucci banana my dude" )
		fmt.Println(i)
	}




}
func slice(mylist []string)[]string{

	var length = len(mylist)

	return mylist[:(length-1)]

}

func oddoreven(num int)bool{

	if (num%2 == 0){

		return true
	}else{

		return false
	}


}