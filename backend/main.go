// Code generated by hertz generator.

package main

import "github.com/cloudwego/hertz/pkg/app/server"

func main() {
	h := server.Default()

	register(h)
	h.Spin()
	//data, _ := dal.ContestList()
	//for i, contest := range data {
	//	fmt.Printf("Contest %d: %+v\n", i, *contest)
	//}
}
