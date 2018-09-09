package pmg

import (
	"bufio"
	"fmt"
	"gopkg.in/jdkato/prose.v2"
	"os"
	"strings"
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
)

func evaluateInput(input string, george *George) {
	fmt.Println("Starting eval")
	doc, _ := prose.NewDocument(input)

	var (
		adjs    []string
		nouns   []string
		actions []action
		bling   Bling
		ok/* let's */ bool
	)

	for _, el := range doc.Tokens() {
		fmt.Println(el)
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
		//TODO Call Stephen's code
	} else if actions[0] == quit {
		fmt.Println("Exiting!")
		os.Exit(0)
	} else if actions[0] == list {
		fmt.Println("George is wearing:")
		for _, bling := range george.Accessories {
			fmt.Println(bling)
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
