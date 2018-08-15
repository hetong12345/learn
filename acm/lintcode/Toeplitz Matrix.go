package lintcode

import "fmt"

func ToeplitzMatrix(matrix [][]int) bool {
	var M ,N ,cur int
	M= len(matrix)
	N=len(matrix[0])
	fmt.Println(M,N,cur)
	for i := 0; i<M; i++ {
		for j := 0; j<N; j++ {
			//cur=matrix[M][N]
			//fmt.Println(i,j)
		}
	}
	return true
}