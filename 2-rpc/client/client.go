package main

import (
	"log"
	"net/rpc"

	"github.com/nachogoca/go-grpc-talk/2-rpc/types"
)

func main() {

	// create TCP connection
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("connection error: ", err)
	}
	defer client.Close()

	reply := &types.Task{}
	client.Call("ToDoList.CreateTask", "milk", reply)
	log.Printf("created: %+v\n", reply)

	client.Call("ToDoList.CreateTask", "ham", reply)
	log.Printf("created: %+v\n", reply)

	client.Call("ToDoList.CreateTask", "cheese", &reply)
	log.Printf("created: %+v\n", reply)

	replies := make([]types.Task, 0)
	client.Call("ToDoList.GetTasks", "", &replies)
	log.Println("tasks: ", replies)

	client.Call("ToDoList.CompleteTask", "milk", &reply)
	log.Printf("completed: %+v\n", reply)

	client.Call("ToDoList.GetTask", "milk", &reply)
	log.Printf("task: %+v\n", reply)
}
