package circus

import (
	"fmt"
)

var AllCircuses = Members{speakers: make([]Speaker, 0)}

type Speaker interface {
	Speaks() string
}

type Members struct {
	speakers []Speaker
}

func (c *Members) AddMember(s Speaker) {
	c.speakers = append(c.speakers, s)
}

func (c *Members) Notify() {
	for _, s := range c.speakers {
		fmt.Println(s.Speaks())
	}
}

// func Perform(a Speaker) string { return a.Speaks() }
