package main

import "fmt"

type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    lastIndex := len(s.items) - 1
    item := s.items[lastIndex]
    s.items = s.items[:lastIndex]
    return item, true
}

func main() {
    intStack := &Stack[int]{}
    intStack.Push(10)
    intStack.Push(20)
    val, _ := intStack.Pop()
    fmt.Println("Pop từ intStack:", val)

    stringStack := &Stack[string]{}
    stringStack.Push("A")
    stringStack.Push("B")
    str, _ := stringStack.Pop()
    fmt.Println("Pop từ stringStack:", str)
}