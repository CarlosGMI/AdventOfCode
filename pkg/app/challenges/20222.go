package app

import "fmt"

type Exec20222 struct{}

func (e Exec20222) Exec(sub *chan struct{}) {
	fmt.Println("Challenge 2")
}
