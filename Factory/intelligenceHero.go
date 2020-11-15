package Factory

import (
	"bufio"
	"fmt"
	"os"
)


type intelligenceHero struct {
	hero
}

func newIntelligenceHero() iHero {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter hero mana: ")
	heroName, _ := reader.ReadString('\n')
	return &intelligenceHero{
		hero{
			name:      heroName,
			attribute: "intelligence",
		},
	}
}
