package main

import "fmt"

type Item struct {
	value       interface{}
	prevElement *Item
	nextElement *Item
}

func (item *Item) Prev() *Item {
	return item.prevElement
}

func (item *Item) Next() *Item {
	return item.nextElement
}

func (item *Item) Value() interface{} {
	return item.value
}

type List struct {
	Container []Item
}

func (list *List) Len() interface{} {
	return len(list.Container)
}

func (list *List) Last() (*Item, error) {
	length := len(list.Container)
	var err error
	if length == 0 {
		return &Item{}, fmt.Errorf("Last: Zero length container")
	}
	return &list.Container[len(list.Container)-1], err
}

func (list *List) First() (*Item, error) {
	length := len(list.Container)
	var err error
	if length == 0 {
		return &Item{}, fmt.Errorf("First: Zero length container")
	}
	return &list.Container[0], err
}

func (list *List) PushFront(value interface{}) {
	item := value.(Item)
	if len(list.Container) != 0 {
		item.nextElement, _ = list.First()
	}
	list.Container = append([]Item{item}, list.Container...)
}
func (list *List) PushBack(value interface{}) {
	item := value.(Item)
	if len(list.Container) != 0 {
		item.prevElement, _ = list.Last()
		item.prevElement.nextElement = &item
	}
	list.Container = append(list.Container, item)
}
func (list *List) Remove(value interface{}) {
	for i, val := range list.Container {
		if val.value == value {
			copy(list.Container[i:], list.Container[i+1:])
			list.Container[len(list.Container)-1] = Item{}
			list.Container = list.Container[:len(list.Container)-1]
		}
	}
}

func (list *List) dataPrint() {
	fmt.Printf("list = %v\tlen = %d\n", list, list.Len())
	first, err := list.First()
	if err == nil {
		fmt.Printf("first  = %v\t", first.value)
	} else {
		fmt.Println(err)
	}
	last, err := list.Last()
	if err == nil {
		fmt.Printf("last  = %v\n", last.value)
	} else {
		fmt.Println(err)
	}

}
func main() {
	list := List{
		Container: []Item{},
	}
	list.dataPrint()
	list.PushFront(Item{
		value:       9,
		prevElement: &Item{},
		nextElement: &Item{},
	})
	list.dataPrint()
	list.PushBack(Item{
		value:       20,
		prevElement: &Item{},
		nextElement: &Item{},
	})
	list.dataPrint()
	list.PushBack(Item{
		value:       10,
		prevElement: &Item{},
		nextElement: &Item{},
	})
	list.dataPrint()
	list.PushFront(Item{
		value:       "asd",
		prevElement: &Item{},
		nextElement: &Item{},
	})
	list.dataPrint()
	list.Remove(20)
	list.dataPrint()
}
