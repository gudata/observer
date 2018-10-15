package main

import (
	"github.com/gudata/observer/animals"
	"github.com/gudata/observer/artists"
	"github.com/gudata/observer/circus"
)

func main() {
	dog := &animals.Dog{}
	clown := &artists.Artist{}

	circus.AllCircuses.AddMember(dog)
	circus.AllCircuses.AddMember(clown)
	circus.AllCircuses.Notify()
}
