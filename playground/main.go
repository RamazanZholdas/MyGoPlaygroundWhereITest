package main

import "fmt"

type color struct {
	a, b, c, d int
}

func main() {
	mainMem := [16][16]color{}
	cache := [8]color{} //block size = 4
	g := 0
	counter := 0
	counter2 := 0
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			mainMem[j][i].a = 0
			mainMem[j][i].b = 0
			mainMem[j][i].c = 1
			mainMem[j][i].d = 0
			counter2++
		}
		for k := range cache {
			cache[k] = mainMem[k][g]
			g++
			counter++
		}
		g = 0
	}
	fmt.Println("number of accesses to cache: ", counter)
	fmt.Println("number of accesses to main memory: ", counter2)
	fmt.Println("total:", counter2+counter)

}
