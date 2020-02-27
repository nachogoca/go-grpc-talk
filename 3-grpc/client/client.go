package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nachogoca/go-grpc-talk/3-grpc/todo"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}

	client := todo.NewTodoClient(conn)

	createTask(client, "milk")
	createTask(client, "ham")
	createTask(client, "cheese")

	list, err := client.GetTasks(context.Background(), &todo.Void{})
	if err != nil {
		log.Fatalf("could not list tasks: %v", err)
	}
	printTasks(list)

	list, err = client.CompleteTask(context.Background(), &todo.Text{Text: "cheese"})
	if err != nil {
		log.Fatalf("could not list tasks: %v", err)
	}
	printTasks(list)

	t, err := client.GetTask(context.Background(), &todo.Text{Text: "milk"})
	if err != nil {
		log.Fatalf("could not find task %v", err)
	} else {
		log.Println(t.GetTitle())
		log.Println(t.GetDone())
	}

}

func createTask(client todo.TodoClient, task string) {
	created, err := client.CreateTask(context.Background(), &todo.Text{Text: task})
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	log.Printf("task created: %v", created.GetTitle())
}

func printTasks(list *todo.TaskList) {
	log.Println("list:")
	for _, t := range list.GetTasks() {
		if t.GetDone() {
			fmt.Printf("[X]")
		} else {
			fmt.Printf("[ ]")
		}
		fmt.Println(t.GetTitle())
	}
}
