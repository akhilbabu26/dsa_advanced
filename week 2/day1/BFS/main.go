package main

import "fmt"

func BFS(graph map[string][]string, start string){
	visited := make(map[string]bool)

	queue := []string{start}
	visited[start] = true

	for len(queue) > 0{
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node)

		for _, neighbor := range graph[node]{
			if !visited[neighbor]{
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func main(){
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"D", "E"},
		"C": {},
		"D": {},
		"E": {},
	}

	BFS(graph, "A")
}