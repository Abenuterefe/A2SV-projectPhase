package controllers

import (
	"fmt"

	"task3/models"
	"task3/services"
)

func RunConsoleInterface(library services.LibraryManager, lib *services.Library) {
	for {
		fmt.Println("\nLibrary Menu:")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Member's Borrowed Books")
		fmt.Println("7. Add Member")
		fmt.Println("0. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id int
			var title, author string
			fmt.Print("Enter book ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter book title: ")
			fmt.Scan(&title)
			fmt.Print("Enter book author: ")
			fmt.Scan(&author)
			library.AddBook(models.Book{ID: id, Title: title, Author: author})
			fmt.Println("Book added.")
		case 2:
			var id int
			fmt.Print("Enter book ID to remove: ")
			fmt.Scan(&id)
			library.RemoveBook(id)
			fmt.Println("Book removed.")
		case 3:
			var bookID, memberID int
			fmt.Print("Enter book ID: ")
			fmt.Scan(&bookID)
			fmt.Print("Enter member ID: ")
			fmt.Scan(&memberID)
			err := library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed.")
			}
		case 4:
			var bookID, memberID int
			fmt.Print("Enter book ID: ")
			fmt.Scan(&bookID)
			fmt.Print("Enter member ID: ")
			fmt.Scan(&memberID)
			err := library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned.")
			}
		case 5:
			books := library.ListAvailableBooks()
			fmt.Println("Available Books:")
			for _, b := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", b.ID, b.Title, b.Author)
			}
		case 6:
			var memberID int
			fmt.Print("Enter member ID: ")
			fmt.Scan(&memberID)
			books := library.ListBorrowedBooks(memberID)
			fmt.Printf("Borrowed Books for Member %d:\n", memberID)
			for _, b := range books {
				fmt.Printf("ID: %d, Title: %s\n", b.ID, b.Title)
			}
		case 7:
			var id int
			var name string
			fmt.Print("Enter member ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter member name: ")
			fmt.Scan(&name)
			lib.Members[id] = &models.Member{ID: id, Name: name}
			fmt.Println("Member added.")
		case 0:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
