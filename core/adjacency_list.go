package core

import "fmt"

//	graph := map[string][]string{
//		"a": {"b", "c"},
//		"b": {"d"},
//		"c": {"e"},
//		"d": {"f"},
//		"e": {""},
//		"f": {""},
//	}
func AdjacencyList(graph map[string][]string) {
	// Depth First Traversal -> stack approach
	// a,b,d
	dfs := func(gp map[string][]string, source string) {
		stack := []string{source}
		for len(stack) > 0 {
			current := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]

			fmt.Println(current)

			for i := 0; i < len(gp[current]); i++ {
				if len(gp[current][i]) > 0 {
					stack = append(stack, gp[current][i])
				}
			}
		}
	}

	fmt.Println("starting Depth First Search")
	dfs(graph, "a")

	// Breadth First Traversal -> queue approach
	// a,b,c
	bfs := func(gp map[string][]string, source string) {
		queue := []string{source}
		for len(queue) > 0 {
			current := queue[len(queue)-1]
			queue = queue[0 : len(queue)-1]
			fmt.Println(current)
			for i := 0; i < len(gp[current]); i++ {
				neighbor := gp[current][i]
				if len(neighbor) > 0 {
					queue = append([]string{neighbor}, queue...)
				}
			}
		}
	}
	fmt.Println("starting Breadth First Search")
	bfs(graph, "a")
}
