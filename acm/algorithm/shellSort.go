package algorithm

func ShellSort(arr []int) []int {
	var op, h, j, t, temp int
	incs := []int{ /*a1=3,a2=7,a3=16,a4=41,a5=101*/
		1391376, /*a1*a2*a3*a4*a5*/
		463792,  /*a2*a3*a4*a5*/
		198768,  /*a1*a3*a4*a5*/
		86961,   /*a1*a2*a4*a5*/
		33936,   /*a1*a2*a3*a5*/
		13776,   /*a1*a2*a3*a4*/
		4592,    /*a2*a3*a4*/
		1968,    /*a1*a3*a4*/
		861,     /*a1*a2*a4*/
		336,     /*a1*a2*a3*/
		112,     /*a2*a3*/
		48,      /*a1*a3*/
		21,      /*a1*a2*/
		7,       /*a2*/
		3,       /*a1*/
		1,
	}
	for t = 0; t < 16; t++ {
		h = incs[t]
		for i := h; i < len(arr); i++ {
			temp = arr[i]
			for j = i - h; j >= 0 && arr[j] > temp; j -= h {
				arr[j+h] = arr[j]
				op++
			}
			arr[j+h] = temp
			op++
		}
	}
	return nil
}
