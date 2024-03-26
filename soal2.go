package main

import "fmt"


type Person struct {
	Name string
	Berat, Tinggi float32
}

type Data struct {
	Mark Person
	John Person
}

func calcBmi (person Person) float32 {
	return person.Berat/(person.Tinggi*person.Tinggi)
}

func isMarkHigher (markBMI float32, johnBMI float32) bool {
	if markBMI > johnBMI {
		return true
	}
	return false
}

func exec (data Data) {
	var markBMI float32 = calcBmi(data.Mark)
	var johnBMI float32 = calcBmi(data.John)
	var markHigherBMI = isMarkHigher(markBMI, johnBMI)

	fmt.Println("Mark's BMI: ", markBMI)
	fmt.Println("John's BMI: ", johnBMI)
	fmt.Println("Is Mark's BMI Higher: ", markHigherBMI)
	fmt.Println()
}

func main () {
	Mark1 := Person {"Mark", 78, 1.69}
	John1 := Person {"John", 92, 1.95}
	Data1 := Data {Mark1, John1}
	fmt.Println("Data 1")
	exec(Data1)

	Mark2 := Person {"Mark", 95, 1.88}
	John2 := Person {"John", 85, 1.76}
	Data2 := Data {Mark2, John2}
	fmt.Println("Data 2")
	exec(Data2)
}