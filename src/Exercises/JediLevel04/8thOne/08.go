package main
import "fmt"

func main(){
	m := map[string][]string{
		"bond_james": {"shaken but not stirred", "Martini"},
		"Money Penny":{"James bond", "literature", "computer science"},
		"no_dr":{"being evil", "ice creams", "sunsets"},
	}
	fmt.Println(m)
	for k,v := range m{
		fmt.Println(k)
		for i, val := range v{
			fmt.Println(i, val)
		}
	}
}