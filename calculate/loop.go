package calculate

import ("fmt"
"reflect")

// public and Private Variable
var PublicVariable = "public variable"
var privateVariable = " privateVariable"

func Looping(){
	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}



	for i:=0; ; i++{
		if i==10{
			break
		}
		if i == 3 {
			fmt.Println("i value", i)
			}else{
			continue
		}
		
	}

	for _,value:= range strDict{
		fmt.Println("map value", value)
	}
}

func Structure(){
	type rectangle struct {
		length  float64
		breadth float64
		color   string
		area float64
	}
	var rect rectangle 
	rect.length = 10
	rect.breadth = 20
	rect.color = "Green"
	rect.area = rect.length * rect.breadth
	fmt.Println(rect)
}

func ArrayAndSlice(){
	var theArray [3]string
	x := [5]int{10, 20, 30, 40, 50}
	
	var intSlice [] int
	fmt.Println("_+_+_+_+_+_+_+_+_+_+_+_+_+_+_+_+_+_+_",reflect.ValueOf(intSlice).Kind())
	theArray[0] = "India"  // Assign a value to the first element
	theArray[1] = "Canada" 
	theArray[2] = "Japan"
	fmt.Println(theArray)

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	//slice
	theSlice := make([]int, 2, 5)
	theSlice[0] = 11
	theSlice[1] = 10
	fmt.Println("Slice A:", theSlice)
	fmt.Printf("Length is %d Capacity is %d\n", len(theSlice), cap(theSlice))

	theSlice = append(theSlice, 30, 40, 50, 60, 70, 80, 90)
	fmt.Println("Slice A after appending data:", theSlice)
	fmt.Printf("Length is %d Capacity is %d\n", len(theSlice), cap(theSlice))
}