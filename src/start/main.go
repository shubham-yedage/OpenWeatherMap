// main
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	mapParam := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter \n \t City Name: ")
	mapParam["Name"], _ = reader.ReadString('\n')
	fmt.Println("\t City Code/ID: ")
	mapParam["ID"], _ = reader.ReadString('\n')
	fmt.Println("\t Geographic Co-ordinates(long, lat): ")
	mapParam["Geolat"], _ = reader.ReadString('\n')
	mapParam["Geo long"], _ = reader.ReadString('\n')
	fmt.Println("\t Zip Code: ")
	mapParam["Zip"], _ = reader.ReadString('\n')
	fmt.Println("\t Country: ")
	mapParam["Country"], _ = reader.ReadString('\n')
	fmt.Println("\t Zone(bbox[Left long, Right long, Up lat, Bottom lat]): ")
	mapParam["bbox"], _ = reader.ReadString('\n')
	fmt.Println("\t Temprature Unit:")
	mapParam["Unit"], _ = reader.ReadString('\n')

	fmt.Println("Enter choice for data, \n \t B. Last 5 Days \n \t B. Last 16 Days \n \t C. Historical")
	v, _ := reader.ReadString('\n')

	if strings.Compare(v, "A") == 0 || strings.Compare(v, "B") == 0 {
		fmt.Println("\t Enter number of days:")
		mapParam["cnt"], _ = reader.ReadString('\n')
		fmt.Println("Its A or B!")
	} else if strings.Compare(v, "C") == 0 {
		fmt.Println("Its C!")
	} else {
		fmt.Println("Its Default!")
	}
}
