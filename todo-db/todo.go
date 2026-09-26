package main

import (
    "context"
    "fmt"

    "go.mongodb.org/mongo-driver/v2/bson"
)

type Todo struct {
    ID   bson.ObjectID `bson:"_id,omitempty"`
    Task string        `bson:"task"`
}

func addTodo(task string) {
    todo := Todo{
        Task: task,
    }

    _, err := collection.InsertOne(context.Background(), todo)

    if err != nil {
        fmt.Println("Error adding todo:", err)
        return
    }

    fmt.Println("Added successfully")
}

func listTodo() {
    cursor, err := collection.Find(context.Background(), bson.M{})

    if err != nil {
        fmt.Println("Error getting todos:", err)
        return
    }

    defer cursor.Close(context.Background())

    var todos []Todo

    err = cursor.All(context.Background(), &todos)

    if err != nil {
        fmt.Println("Error reading todos:", err)
        return
    }

    if len(todos) == 0 {
        fmt.Println("No tasks found")
        return
    }

    for index, todo := range todos {
        fmt.Println(index+1, ".", todo.Task)
    }
}

func deleteTodo(task string) {
    var todo Todo

    // Find the task
    err := collection.FindOne(
        context.Background(),
        bson.M{"task": task},
    ).Decode(&todo)

    if err != nil {
        fmt.Println("Task not found")
        return
    }

    // Delete the task
    _, err = collection.DeleteOne(
        context.Background(),
        bson.M{"_id": todo.ID},
    )

    if err != nil {
        fmt.Println("Error deleting todo:", err)
        return
    }

    fmt.Println("Deleted successfully")
}