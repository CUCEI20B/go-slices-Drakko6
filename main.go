package main

import "fmt"

func main()  {
	var n uint64;

	fmt.Scan(&n);

	s:= make([]int64, 0, n);

	var e int64;
	var suma int64 = 0;
	for i:=0; i<int(n); i++ {
		
		fmt.Scan(&e)
		s = append(s, e)
		suma += e;
	}
	fmt.Println(suma);

}