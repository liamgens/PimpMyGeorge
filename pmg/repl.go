package pmg

import (
	"bufio"
	"fmt"
	"gopkg.in/jdkato/prose.v2"
	"os"
	"strings"
	"time"
	"math/rand"
)

type action int

const (
	add action = 0 + iota
	export
	list
	quit
)

type verbAction struct {
	verb   string
	action action
}

var (
	georges  []George
	blang    []Bling
	blangMap map[string]Bling

	actionMap map[string]action

	verbList []verbAction = []verbAction{
		verbAction{"give", add},
		verbAction{"pimp", add},
		verbAction{"wear", add},

		verbAction{"export", export},
		verbAction{"render", export},
		verbAction{"gimme", export},

		verbAction{"list", list},
		verbAction{"explain", list},
		verbAction{"describe", list},

		verbAction{"quit", quit},
		verbAction{"exit", quit},
		verbAction{"leave", quit},
	}

	flavortown []string
)

func flavaflav(bling Bling) string{
	flav := flavortown[rand.Intn(len(flavortown))]
	if strings.Contains(flav, "%s") {
		return fmt.Sprintf(flav, bling.Adj, bling.Noun)
	}
	return flav
}

func evaluateInput(input string, george *George) {
	doc, _ := prose.NewDocument(input)

	var (
		adjs    []string
		nouns   []string
		actions []action
		bling   Bling
		ok/* let's */ bool
	)

	for _, el := range doc.Tokens() {
		switch el.Tag {
		case "JJ":
			adjs = append(adjs, el.Text)
		case "NN":
			nouns = append(nouns, el.Text)
		case "VB":
			if act, oke := actionMap[el.Text]; oke {
				actions = append(actions, act)
			}
		}
	}

	if len(actions) == 0 {
		return
	}

	if actions[0] == export {
		fmt.Println("Exporting that stuff")
		err := CreateBlingImage(*george, "output.png")
		if err != nil {
			fmt.Println("Error exporting")
		}
	} else if actions[0] == quit {
		fmt.Println("Exiting!")
		os.Exit(0)
	} else if actions[0] == list {
		if len(george.Accessories) == 0 {
			fmt.Println("George is bare ass naked and he looks great ;)")
			return
		}
		fmt.Println("George is wearing:")
		for ind, bling := range george.Accessories {
			fmt.Printf("%d: A %s %s\n", ind+1, bling.Adj, bling.Noun)
		}
	} else {
		for _, adj := range adjs {
			for _, noun := range nouns {
				if bling, ok = blangMap[adj+" "+noun]; ok {
					break
				}
			}
		}
		george.Accessories = append(george.Accessories, bling)
		fmt.Println(flavaflav(bling))
	}
}

func initializeVars() {
	// Grab a slice of George
	georges, blang = FetchGeorgeBlingData("./pmg/georges.json",
		"./pmg/blang.json")

	blangMap = make(map[string]Bling)
	for _, bling := range blang {
		blangMap[strings.ToLower(bling.Adj+" "+bling.Noun)] = bling
	}

	actionMap = make(map[string]action)
	for _, verb := range verbList {
		actionMap[verb.verb] = verb.action
	}

	file,_ := os.Open("./pmg/flavor.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		flavortown = append(flavortown, scanner.Text())
	}
	rand.Seed(time.Now().Unix())
}

func MainLoop() {
	initializeVars()

	george := georges[0]
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("--> ")

		// read
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("REPL caught error: %s\n", err.Error())
			continue
		}

		// eval, print
		evaluateInput(input, &george)
	} // loop
}
