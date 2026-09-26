package main

import "fmt"

func main() {

    connectDB()

    for {
        var action string

        fmt.Println("\nEnter action (add/list/delete/end):")
        fmt.Scanln(&action)

        if action == "add" {

            var task string

            fmt.Println("Enter task:")
            fmt.Scanln(&task)

            addTodo(task)

        } else if action == "list" {

            listTodo()

        } else if action == "delete" {

            var task string

            fmt.Println("Enter task to delete:")
            fmt.Scanln(&task)

            deleteTodo(task)

        } else if action == "end" {

            fmt.Println("Goodbye!")
            break

        } else {

            fmt.Println("Invalid action")
        }
    }
}