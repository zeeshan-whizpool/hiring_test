1 - Repolayer just executes the Graphql query over gitlab

2 - Service layer takes values for env, calls RepoLayer and then formats the results and presents it on endpoint localhost:10000/gql

3 - Tests of servicelayer is also written.

4. server.go starts the endpoint and client/client.go consumes it.
