package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

// numTasks - число заданий которые можно выполнять параллельно;
// numError - максимальное число ошибок после которого нужно приостановить обработку.
func createTasks(tasks []func(chan error), numTasks int, numError int, errChan chan error) {
	sumNumError := 0
	for _, task := range tasks {
		i := 0
	L:
		i++
		if i < 5 {
			fmt.Printf("numTasks = %d \truntimeNumgoroutine = %d\n", numTasks, runtime.NumGoroutine())
		}
		if runtime.NumGoroutine() == numTasks {
			goto L
		}
		fmt.Println("new task")
		go task(errChan)
	}
	for err := range errChan {
		if err != nil {
			sumNumError++
		}
		if sumNumError > numError {
			fmt.Println("Совершено максимальное число ошибок.Выполнение задач прекращено.")
			return
		}
	}
}

func main() {
	errChan := make(chan error, 5)
	a := func(in chan error) {
		fmt.Println("вызывается ошибка")
		time.Sleep(time.Second * 5)
		in <- errors.New("error")

	}
	b := func(in chan error) {
		fmt.Println("ошибки нет")
		time.Sleep(time.Second * 5)
		in <- nil

	}
	c, d := a, a
	e := b
	tasks := make([]func(chan error), 0)
	tasks = append(tasks, a, b, c, d, e)
	createTasks(tasks, 3, 2, errChan)
}
