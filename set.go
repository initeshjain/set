package set

import "fmt"

type Set map[interface{}]struct{}

func NewSet() Set {
    return make(Set)
}

func (s Set) Add(item interface{}) {
    s[item] = struct{}{}
}

func (s Set) Remove(item interface{}) {
    delete(s, item)
}

func (s Set) Contains(item interface{}) bool {
    _, found := s[item]
    return found
}

func (s Set) Print() {
    for item := range s {
        fmt.Println(item)
    }
}
