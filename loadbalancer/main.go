package main

import "fmt"

type Server struct {
	server map[int]Weight
}

type Weight struct {
	maxWeight     int
	currentWeight int
}

func main() {
	// 假设有服务器 1，2，3
	var server = []int{1, 2, 3}

	// roundrobin 负载均衡策略
	var i = 0
	for j := 0; j < 10; j++ {
		fmt.Println(server[i])
		i++

		if i%len(server) == 0 {
			i = 0
		}
	}

	// roundrobind with weight

	s := Server{server: make(map[int]Weight)}
	s.server[0] = Weight{3, 0}
	s.server[1] = Weight{2, 0}
	s.server[2] = Weight{1, 0}

	i = 0
	for j := 0; j < 10; j++ {
		fmt.Println(s.server[i])

		if s.server[i].currentWeight+1 == s.server[i].maxWeight {
			s.server[i] = Weight{s.server[i].maxWeight, 0}
			i++
		} else {
			s.server[i] = Weight{s.server[i].maxWeight, s.server[i].currentWeight + 1}
		}

		if i%len(s.server) == 0 {
			i = 0
		}
	}
}
