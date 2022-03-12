package funcTesting

import "testing"

type color struct {
	a, b, c, d int
}

func BenchmarkDirectMapping1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainMem := [16][16]color{}
		cache := [8]color{} //block size = 4
		g := 0
		counter := 0
		counter2 := 0
		for i := 0; i < 16; i++ {
			for j := 0; j < 16; j++ {
				mainMem[i][j].a = 0
				mainMem[i][j].b = 0
				mainMem[i][j].c = 1
				mainMem[i][j].d = 0
				counter2++
			}
			for k := range cache {
				cache[k] = mainMem[k][g]
				g++
				counter++
			}
			g = 0
		}
	}
}

/*
number of accesses to cache:  128
number of accesses to main memory:  256
total: 384
*/

func BenchmarkDirectMapping2(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
	}
}

/*
number of accesses to cache:  128
number of accesses to main memory:  256
total: 384
*/
func BenchmarkDirectMapping3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainMem := [16][16]color{}
		cache := [8]color{} //block size = 4
		g := 0
		counter := 0
		counter2 := 0

		for i := 0; i < 16; i++ {
			for j := 0; j < 16; j++ {
				mainMem[j][i].c = 1
				counter2++
			}
		}

		for i := 0; i < 16; i++ {
			for j := 0; j < 16; j++ {
				mainMem[j][i].a = 0
				mainMem[j][i].b = 0
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
	}
}

/*
number of accesses to cache:  128
number of accesses to main memory:  512
total: 640
*/
