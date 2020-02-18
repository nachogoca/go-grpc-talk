package main

import "fmt"

// Task is an independent ToDo task
type Task struct {
	Title string
	Done  bool
}

// ToDoList stores the tasks
type ToDoList struct {
	Tasks []Task
}

// NewToDoList creates a ToDoList
func NewToDoList() *ToDoList {
	return &ToDoList{
		Tasks: make([]Task, 0),
	}
}

// GetTask returns a task
func (todo *ToDoList) GetTask(title string) (Task, error) {
	for _, t := range todo.Tasks {
		if t.Title == title {
			return t, nil
		}
	}
	return Task{}, fmt.Errorf("could not found ToDo")
}

// CreateTask adds the task to the list
func (todo *ToDoList) CreateTask(title string) Task {
	t := Task{Title: title}
	todo.Tasks = append(todo.Tasks, t)
	return t
}

// CompleteTask marks task as done
func (todo *ToDoList) CompleteTask(title string) {
	for i, t := range todo.Tasks {
		if t.Title == title {
			todo.Tasks[i].Done = true
			return
		}
	}
}

// GetTasks returns the ToDo list tasks
func (todo *ToDoList) GetTasks() []Task {
	return todo.Tasks
}

func main() {

	toBuy := NewToDoList()
	toBuy.CreateTask("milk")
	toBuy.CreateTask("ham")
	toBuy.CreateTask("cheese")

	printTasks(toBuy.GetTasks())
	fmt.Println("")
	toBuy.CompleteTask("milk")
	printTasks(toBuy.GetTasks())

}

func printTasks(tasks []Task) {
	for _, t := range tasks {
		if t.Done {
			fmt.Printf("[X]")
		} else {
			fmt.Printf("[ ]")
		}
		fmt.Println(t.Title)
	}
}
