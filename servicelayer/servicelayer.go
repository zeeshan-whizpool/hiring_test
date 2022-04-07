package servicelayer

import (
	"encoding/json"
	"fmt"
	"hiring_test/repolayer"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

/*
	The search string to be included in GQL variables.
*/
func GetProjetList(searchString string) string {

	// Getting values from env for base address and # of projects to fetch
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gqlAddress := os.Getenv("GQL_URL")
	projectsCount, _ := strconv.Atoi(os.Getenv("PROJECT_COUNT"))

	// Executing GQL.
	projects := repolayer.GetProjetList(gqlAddress, searchString, projectsCount)
	if projects != nil {
		var names []string
		forks := 0

		for _, project := range projects {
			forks += project.ForksCount

			// I know, not ideal copying again into the array, in ideal world i would have copied it only once
			names = append(names, project.Name)
		}

		// Convet into json now.
		output := map[string]interface{}{
			"names": strings.Join(names, ", "),
			"forks": forks,
		}

		bytes, err := json.Marshal(output)
		if err != nil {
			log.Print("Unable to convert into json")
			return ""
		}
		return string(bytes)

	}

	return ""
}

/*
	Simple HTTP handler functions.
*/
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Hiring test. Use /gql?search=<projects search> endpoint")
	fmt.Println("Endpoint Hit: test")
}

func gqlSearch(w http.ResponseWriter, r *http.Request) {
	searchString := r.URL.Query().Get("search")
	fmt.Fprintf(w, GetProjetList(searchString))
	fmt.Println("Endpoint Hit: gql")
}

func HandleRequests() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/gql", gqlSearch)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
