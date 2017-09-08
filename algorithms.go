package main

import (
	"container/heap"
)

var (
	inf = 20000000
)

func dijkstra(start int, N int, M int, adj map[int]map[int]int) ([]int, []int) {
	dist := make([]int, N+1)
	father := make([]int, N+1)
	visited := make([]int, N+1)
	for i := 0; i < N+1; i++ {
		dist[i] = inf
		father[i] = -1
	}
	dist[start] = 0
	father[start] = start

	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		u:        start,
		priority: 0,
		index:    0,
	}
	heap.Init(&pq)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		u := item.u
		visited[u] = 1
		for v, s := range adj[u] {
			if visited[v] == 0 && dist[v] > dist[u]+s {
				dist[v] = dist[u] + s
				father[v] = u
				heap.Push(&pq, &Item{
					u:        v,
					priority: s,
				})
			}
		}
	}

	return dist, father
}

// func main() {
// 	var T, N, M, x, y, s int
// 	fmt.Scan(&T)
// 	fmt.Scanf("%d %d", &N, &M)
//
// 	for i := 0; i < T; i++ {
// 		adjacent := map[int]map[int]int{}
// 		var start int
// 		for i := 0; i < M; i++ {
// 			fmt.Scanf("%d %d %d", &x, &y, &s)
// 			mm, ok := adjacent[x]
// 			if !ok {
// 				mm = make(map[int]int)
// 				adjacent[x] = mm
// 			}
// 			mm[y] = s
// 		}
//
// 		fmt.Scan(&start)
//
// 		dist, father := dijkstra(start, N, M, adjacent)
// 		fmt.Println(dist, father)
// 	}
// }
