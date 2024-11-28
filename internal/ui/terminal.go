package ui

import (
    "bookmarks/internal/core"
    "fmt"
    "strings"
)

const (
    Reset  = "\033[0m"
    Bold   = "\033[1m"
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Blue   = "\033[34m"
    Purple = "\033[35m"
    Cyan   = "\033[36m"
)

type TerminalUI struct {
    manager *core.BookmarkManager
}

func NewTerminalUI(manager *core.BookmarkManager) *TerminalUI {
    return &TerminalUI{
        manager: manager,
    }
}

func (ui *TerminalUI) Run() error {
    for {
        ui.showMenu()
        choice := ui.readInput("Choice: ")

        switch choice {
        case "1":
            ui.addBookmark()
        case "2":
            ui.viewBookmarks()
        case "3":
            ui.searchBookmarks()
        case "4":
            ui.showStats()
        case "5":
            return nil
        default:
            fmt.Println(Red + "Invalid choice" + Reset)
        }
    }
}

func (ui *TerminalUI) showMenu() {
    fmt.Print("\033[H\033[2J") // Clear screen
    fmt.Println(Bold + Purple + "🔖 Bookmark Manager" + Reset)
    fmt.Println(Cyan + "\n1. Add Bookmark")
    fmt.Println("2. View Bookmarks")
    fmt.Println("3. Search")
    fmt.Println("4. Statistics")
    fmt.Println("5. Exit" + Reset)
}

func (ui *TerminalUI) readInput(prompt string) string {
    fmt.Print(Yellow + prompt + Reset)
    var input string
    fmt.Scanln(&input)
    return strings.TrimSpace(input)
}

func (ui *TerminalUI) addBookmark() {
    // Implementation
}

func (ui *TerminalUI) viewBookmarks() {
    // Implementation
}

func (ui *TerminalUI) searchBookmarks() {
    // Implementation
}

func (ui *TerminalUI) showStats() {
    // Implementation
}
