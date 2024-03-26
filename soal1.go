package main

import "fmt"


type Team struct {
	scores [3]float64
}

type Data struct {
	teamLumbaLumba Team
	teamKoala Team
}

func calcAverage(team Team) float64 {
	var sum float64 = 0
	for i := 0; i < 3; i++ {
		sum += team.scores[i]
	}
	return sum/3;
}

func determineWinnerDefault (lumbaLumba float64, koala float64) string {
	if lumbaLumba > koala {
		return "Lumba Lumba"
	} else if lumbaLumba == koala {
		return "Seri"
	}
	return "Koala"
}

func determineWinnerBonus1 (lumbaLumba float64, koala float64) string {
	if lumbaLumba == koala {
		return "Seri"
	}
	if lumbaLumba < 100 && koala < 100 {
		return "Tidak Ada"
	} 
	if lumbaLumba > koala {
		return "Lumba Lumba"
	}
	return "Koala"
}

func determineWinnerBonus2 (lumbaLumba float64, koala float64) string {
	if lumbaLumba < 100 && koala < 100 {
		return "Tidak Ada"
	} 
	if lumbaLumba > koala {
		return "Lumba Lumba"
	} else if lumbaLumba == koala {
		return "Seri"
	}
	return "Koala"
}

func exec (data Data, determineWinner func (float64, float64) string) {
	var teamLumbaLumbaAverage = calcAverage(data.teamLumbaLumba)
	var teamKoalaAverage = calcAverage(data.teamKoala)
	var pemenang = determineWinner(teamLumbaLumbaAverage, teamKoalaAverage)

	fmt.Println("Team Lumba Lumba Average:", teamLumbaLumbaAverage)
	fmt.Println("Team Koala Average:", teamKoalaAverage)
	fmt.Println("Pemenang:", pemenang)
	fmt.Println()
}

func main () {
	data1 := Data {Team {[3]float64{96,108,89}}, Team {[3]float64{88,91,110}}}
	fmt.Println("Data 1")
	exec(data1, determineWinnerDefault)

	data2 := Data {Team {[3]float64{97,112,101}}, Team {[3]float64{109,95,123}}}
	fmt.Println("Data Bonus 1")
	exec(data2, determineWinnerBonus1)

	data3 := Data {Team {[3]float64{97,112,101}}, Team {[3]float64{109,95,106}}}
	fmt.Println("Data Bonus 2")
	exec(data3, determineWinnerBonus2)
}