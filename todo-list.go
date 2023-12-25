package main

import (
	"fmt"
)

type TodoItem struct {
	ID    int
	Title string
}

var todoList []TodoItem
var listID int

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
)

func logFunctionInfo() {
	fmt.Println(Green, "=======Todo App:=======", Reset)
	fmt.Print(Green, "1. Add Todo List", Reset, "\n")
	fmt.Print(Green, "2. View Todo List", Reset, "\n")
	fmt.Print(Green, "3. Delete Todo Item", Reset, "\n")
	fmt.Print(Green, "4. Exit Server", Reset, "\n")
	fmt.Print(Green, "=======================", Reset, "\n")
}

func addTodoItem(title string) {
	listID++
	todo := TodoItem{ID: listID, Title: title}
	todoList = append(todoList, todo)
	fmt.Print(Blue, "> Added:", todo, Reset, "\n")
	fmt.Println()
}

func viewTodoList() {
	fmt.Print(Blue, "> Todo List:", Reset, "\n")
	for _, todo := range todoList {
		fmt.Printf(Blue+"> %d: %s\n"+Reset, todo.ID, todo.Title)
	}
	fmt.Println()
}

func deleteTodoItem(id int) {
	for i, todo := range todoList {
		if todo.ID == id {
			todoList = append(todoList[:i], todoList[i+1:]...)
			fmt.Print(Blue, "> Deleted:", todo, Reset, "\n")
			fmt.Println()
			return
		}
	}
	fmt.Print(Red, "> Error: Item not found", Reset, "\n")
	fmt.Println()
}

func main() {
	listID = 0 // 初始化
	logFunctionInfo()

	for {
		var choice int

		fmt.Print("> Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var title string
			fmt.Print("> Enter todo title: ")
			fmt.Scanln(&title)
			addTodoItem(title)
		case 2:
			viewTodoList()
		case 3:
			var id int
			fmt.Print("> Enter todo ID to delete: ")
			fmt.Scanln(&id)
			deleteTodoItem(id)
		case 4:
			fmt.Print(Purple, "> Exiting...", Reset, "\n")
			return
		default:
			fmt.Print(Red, "> Error: Invalid choice "+fmt.Sprintf("%d", choice), Reset, "\n")
			fmt.Println()
		}
	}
}
