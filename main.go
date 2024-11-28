package main

import (
	"bookmarks/internal/handlers"
	"bookmarks/internal/storage"
	"bookmarks/internal/ui"
	"fmt"
)

func main() {
	store := storage.NewStorage("bookmarks.json")
	bookmarkHandler := handlers.NewBookmarkHandler(store)

	for {
		ui.ShowMainMenu()
		choice := ui.ReadLine("\nChoice: ")

		switch choice {
		case "1":
			bookmarkHandler.AddBookmark()
		case "2":
			bookmarkHandler.ViewBookmarks()
		case "3":
			bookmarkHandler.SearchBookmarks()
		case "4":
			bookmarkHandler.ExportToMarkdown()
		case "5":
			fmt.Println(ui.Green + "\nGoodbye! 👋" + ui.Reset)
			return
		default:
			fmt.Println(ui.Red + "\n❌ Invalid choice, please try again" + ui.Reset)
			ui.WaitForEnter()
		}
	}
}
