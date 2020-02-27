package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/nachogoca/go-grpc-talk/3-grpc/todo"
	"google.golang.org/grpc"
)

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
func (t *ToDoList) GetTask(ctx context.Context, title *todo.Text) (*todo.Task, error) {
	for i := range t.Tasks {
		if t.Tasks[i].Title == title.GetText() {
			log.Printf("task found: %+v\n", t.Tasks[i])
			return &todo.Task{
				Done:  t.Tasks[i].Done,
				Title: t.Tasks[i].Title,
			}, nil
		}
	}
	return nil, fmt.Errorf("could not found task")
}

// CreateTask adds the task to the list
func (t *ToDoList) CreateTask(ctx context.Context, title *todo.Text) (*todo.Task, error) {
	task := Task{Title: title.GetText()}
	t.Tasks = append(t.Tasks, task)

	log.Printf("task created: %+v\n", task)
	return &todo.Task{
		Title: task.Title,
		Done:  task.Done,
	}, nil
}

// CompleteTask marks task as done
func (t *ToDoList) CompleteTask(ctx context.Context, title *todo.Text) (*todo.TaskList, error) {

	for i, task := range t.Tasks {
		if task.Title == title.GetText() {
			t.Tasks[i].Done = true
		}
	}
	return t.GetTasks(ctx, nil)
}

// GetTasks returns the ToDo list tasks
func (t *ToDoList) GetTasks(ctx context.Context, _ *todo.Void) (*todo.TaskList, error) {

	log.Printf("task list: %+v\n", t.Tasks)
	tasks := make([]*todo.Task, 0, len(t.Tasks))
	for _, task := range t.Tasks {
		tasks = append(tasks, &todo.Task{
			Title: task.Title,
			Done:  task.Done,
		})
	}
	log.Printf("task list: %+v\n", tasks)

	return &todo.TaskList{
		Tasks: tasks,
	}, nil
}

func main() {

	server := grpc.NewServer()

	list := NewToDoList()
	todo.RegisterTodoServer(server, list)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("could not listen to :8080 %v", err)
	}
	log.Fatal(server.Serve(listener))

}
