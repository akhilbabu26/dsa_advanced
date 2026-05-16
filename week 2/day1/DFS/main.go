package main

import "fmt"

func DFS(graph map[string][]string, node string, visited map[string]bool){
	if visited[node]{
		return
	}

	visited[node] = true
	fmt.Println(node)

	for _, neighbor := range graph[node]{
		DFS(graph, neighbor, visited)
	}
}

func main(){
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"D"},
		"C": {},
		"D": {},
		"E": {},
	}

	visited := make(map[string]bool)

	DFS(graph, "A", visited)
}