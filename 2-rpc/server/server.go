package main

/*
Go RPC Requirements

The net/rpc package stipulates that only methods that satisfy the following
criteria will be made available for remote access; other methods will be ignored.

    The method’s type is exported.
    The method is exported
    The method has two arguments, both exported (or builtin types).
    The method’s second argument is a pointer
	The method has return type error
*/
import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/nachogoca/go-grpc-talk/2-rpc/types"
)

// ToDoList stores the tasks
type ToDoList struct {
	Tasks []types.Task
}

// NewToDoList creates a ToDoList
func NewToDoList() *ToDoList {
	return &ToDoList{
		Tasks: make([]types.Task, 0),
	}
}

// GetTask returns a task
func (todo *ToDoList) GetTask(title string, reply *types.Task) error {
	for i := range todo.Tasks {
		if todo.Tasks[i].Title == title {
			*reply = todo.Tasks[i]
			log.Printf("task found: %+v\n", reply)
			return nil
		}
	}
	return fmt.Errorf("could not found task")
}

// CreateTask adds the task to the list
func (todo *ToDoList) CreateTask(title string, reply *types.Task) error {
	t := types.Task{Title: title}
	todo.Tasks = append(todo.Tasks, t)

	*reply = t
	log.Printf("task created: %+v\n", reply)
	return nil
}

// CompleteTask marks task as done
func (todo *ToDoList) CompleteTask(title string, reply *types.Task) error {
	for i, t := range todo.Tasks {
		if t.Title == title {
			todo.Tasks[i].Done = true
			*reply = todo.Tasks[i]
			log.Printf("task completed: %+v\n", reply)
			return nil
		}
	}
	return fmt.Errorf("could not found task")
}

// GetTasks returns the ToDo list tasks
func (todo *ToDoList) GetTasks(dummy string, reply *[]types.Task) error {
	*reply = todo.Tasks
	log.Printf("tasks found: %#v\n", reply)
	return nil
}

func main() {

	todo := NewToDoList()
	// Publish the receiver methods
	err := rpc.Register(todo)
	if err != nil {
		log.Fatal("Format of todo is not correct", err)
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("could not create listener on port 8080: ", err)
	}
	defer listener.Close()

	log.Println("listening on port 8080")
	rpc.Accept(listener)

}
