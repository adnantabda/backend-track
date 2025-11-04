package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
)

func RunLibraryConsole() {
	library := services.NewLibrary()

	library.Members[1] = models.Member{ID: 1, Name: "first"}
	library.Members[2] = models.Member{ID: 2, Name: "second"}

	for {
		fmt.Println("\n===== Library Management System =====")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id int
			var title, author string
			fmt.Print("Enter Book ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Title: ")
			fmt.Scan(&title)
			fmt.Print("Enter Author: ")
			fmt.Scan(&author)

			library.AddBook(models.Book{ID: id, Title: title, Author: author})
			fmt.Println("Book added successfully.")

		case 2:
			var id int
			fmt.Print("Enter Book ID to remove: ")
			fmt.Scan(&id)
			library.RemoveBook(id)
			fmt.Println("Book removed successfully.")

		case 3:
			var bookID, memberID int
			fmt.Print("Enter Book ID: ")
			fmt.Scan(&bookID)
			fmt.Print("Enter Member ID: ")
			fmt.Scan(&memberID)

			err := library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed successfully.")
			}

		case 4:
			var bookID, memberID int
			fmt.Print("Enter Book ID: ")
			fmt.Scan(&bookID)
			fmt.Print("Enter Member ID: ")
			fmt.Scan(&memberID)

			err := library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned successfully.")
			}

		case 5:
			fmt.Println("Available Books:")
			for _, b := range library.ListAvailableBooks() {
				fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
			}

		case 6:
			var memberID int
			fmt.Print("Enter Member ID: ")
			fmt.Scan(&memberID)

			books := library.ListBorrowedBooks(memberID)
			if len(books) == 0 {
				fmt.Println("No borrowed books.")
			} else {
				fmt.Println("Borrowed Books:")
				for _, b := range books {
					fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
				}
			}

		case 7:
			fmt.Println("Exiting... Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
