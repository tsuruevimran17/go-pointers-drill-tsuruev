package main

import "fmt"

type StudentCard struct {
    Owner string
    Active bool
}
func issue(card *StudentCard, owner string) {
	card.Owner = "Алиев Али"
	card.Active = true

}  // заполняет Owner и включает Active

func block(card *StudentCard)  {
	card.Active = false
}               // ставит Active = false

func rename(card *StudentCard, newOwner string){
	card.Owner = "Интукод Студент"
}

func main(){
	card:=StudentCard{ 
	}
issue(&card,"omar hayam")
fmt.Println(card)
block(&card)
rename(&card,"")
fmt.Println(card)
}