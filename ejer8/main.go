package main

import "fmt"

func main() {

	championsLeague := map[string]int{
		"Barcelona":         5,
		"Real Madrid":       13,
		"Bayern Munich":     5,
		"Juventus":          4,
		"Manchester United": 3,
		"Chelsea":           1,
	}

	championsLeague["Liverpool"] = 6
	championsLeague["Boca Juniors"] = 14

	delete(championsLeague, "Chelsea")

	fmt.Println(championsLeague)

	for team, titles := range championsLeague {
		fmt.Printf("Team: %s, Titles: %d \n", team, titles)
	}

	titles, ok := championsLeague["Boca Juniors"]
	fmt.Println(titles, ok)

}
