package set

import (
	"errors"
	"fmt"
)

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

func (s Set) GetByIndex(index int) (interface{}, error) {
    if index < 0 || index >= len(s) {
        return nil, errors.New("index out of range")
    }

    var count int
    for item := range s {
        if count == index {
            return item, nil
        }
        count++
    }

    return nil, errors.New("index out of range")
}