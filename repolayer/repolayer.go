package repolayer

import (
	"context"
	"fmt"
	"log"

	"github.com/shurcooL/graphql"
)

// Struct type that exports the project settings.
type ProjectInfo struct {
	Name        string
	Description string
	ForksCount  int
}

/*
	Below is the implementation of follow query while passing $n as variable/param
	projects(last:$n) {
	 	nodes {
	 	  name
	 	  description
	 	  forksCount
	 	}
	}
*/
func GetProjetList(gqlAddress string, searchString string, projectsCount int) []ProjectInfo {

	client := graphql.NewClient(gqlAddress, nil)

	var QueryProjects struct {
		Projects struct {
			Nodes []struct {
				Name        graphql.String
				Description graphql.String
				ForksCount  graphql.Int
			}
		} `graphql:"projects(search: $searchterm, last:$last_projects)"`
	}
	// Setup variables
	variables := map[string]interface{}{
		"last_projects": graphql.Int(projectsCount),
		"searchterm":    graphql.String(searchString),
	}

	// Execute the query
	err := client.Query(context.Background(), &QueryProjects, variables)
	if err != nil {
		log.Println("Error return with : ", err.Error())
		return nil
	}

	var projects []ProjectInfo

	fmt.Println("Total Projects", len(QueryProjects.Projects.Nodes))
	for _, node := range QueryProjects.Projects.Nodes {

		// Return array of project info's.
		project := ProjectInfo{string(node.Name), string(node.Description), int(node.ForksCount)}
		projects = append(projects, project)
	}

	return projects
}
