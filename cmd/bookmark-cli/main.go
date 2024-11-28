package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Base bookmark type
type BaseBookmark struct {
	URL       string    `json:"url"`
	Title     string    `json:"title"`
	Notes     string    `json:"notes,omitempty"`
	Tags      []string  `json:"tags,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

// Design inspiration bookmark
type DesignBookmark struct {
	BaseBookmark
	Type        string `json:"type"`
	Category    string `json:"category"`
	ColorScheme string `json:"colors"`
	HasRepo     bool   `json:"hasRepo"`
	RepoURL     string `json:"repoUrl,omitempty"`
}

// Code bookmark
type CodeBookmark struct {
	BaseBookmark
	Language       string `json:"language"`
	Solution       string `json:"solution"`
	WorksInVersion string `json:"worksInVersion,omitempty"`
}

// Article bookmark
type ArticleBookmark struct {
	BaseBookmark
	Author      string `json:"author,omitempty"`
	ReadingTime int    `json:"readingTime,omitempty"`
	Category    string `json:"category"`
}

// Generic bookmark
type GenericBookmark struct {
	BaseBookmark
	Category string `json:"category"`
}

type BookmarkManager struct {
	Designs  []DesignBookmark  `json:"designs"`
	Code     []CodeBookmark    `json:"code"`
	Articles []ArticleBookmark `json:"articles"`
	Generic  []GenericBookmark `json:"generic"`
}

// Colors and styles
const (
	Reset    = "\033[0m"
	Bold     = "\033[1m"
	Red      = "\033[31m"
	Green    = "\033[32m"
	Yellow   = "\033[33m"
	Blue     = "\033[34m"
	Magenta  = "\033[35m"
	Cyan     = "\033[36m"
	White    = "\033[37m"
	BgRed    = "\033[41m"
	BgGreen  = "\033[42m"
	BgYellow = "\033[43m"
	BgBlue   = "\033[44m"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func readLine(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(Yellow + prompt + Reset)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func showMainMenu() {
	clearScreen()
	fmt.Println(Bold + Magenta + "\n📚 Bookmark Manager v1.0" + Reset)
	fmt.Println(Cyan + "\n1. 📝 Add New Bookmark")
	fmt.Println("2. 👀 View Bookmarks")
	fmt.Println("3. 🔍 Search")
	fmt.Println("4. 📤 Export to Markdown")
	fmt.Println("5. ❌ Exit" + Reset)
}

func addBookmark(manager *BookmarkManager) {
	clearScreen()
	fmt.Println(Bold + Blue + "\n=== Add New Bookmark ===" + Reset)
	fmt.Println(Cyan + "\n1. 🎨 Design Inspiration")
	fmt.Println("2. 💻 Code/Stack Overflow")
	fmt.Println("3. 📚 Article/Blog")
	fmt.Println("4. 🔖 Generic Bookmark" + Reset)

	choice := readLine("\nChoice: ")

	base := BaseBookmark{
		CreatedAt: time.Now(),
	}

	fmt.Println(Yellow + "\n--- Basic Information ---" + Reset)
	base.URL = readLine("URL: ")
	base.Title = readLine("Title: ")
	base.Notes = readLine("Notes (optional): ")
	tags := readLine("Tags (comma-separated): ")
	if tags != "" {
		base.Tags = strings.Split(tags, ",")
		for i := range base.Tags {
			base.Tags[i] = strings.TrimSpace(base.Tags[i])
		}
	}

	switch choice {
	case "1":
		design := DesignBookmark{BaseBookmark: base}
		fmt.Println(Yellow + "\n--- Design Details ---" + Reset)
		design.Type = readLine("Type (website/app/design): ")
		design.Category = readLine("Category (saas/dashboard/landing): ")
		design.ColorScheme = readLine("Color scheme (light/dark): ")
		hasRepo := readLine("Has repository? (y/n): ")
		if strings.ToLower(hasRepo) == "y" {
			design.HasRepo = true
			design.RepoURL = readLine("Repository URL: ")
		}
		manager.Designs = append(manager.Designs, design)

	case "2":
		code := CodeBookmark{BaseBookmark: base}
		fmt.Println(Yellow + "\n--- Code Details ---" + Reset)
		code.Language = readLine("Programming Language: ")
		code.Solution = readLine("Solution summary: ")
		code.WorksInVersion = readLine("Works in version: ")
		manager.Code = append(manager.Code, code)

	case "3":
		article := ArticleBookmark{BaseBookmark: base}
		fmt.Println(Yellow + "\n--- Article Details ---" + Reset)
		article.Author = readLine("Author: ")
		readingTime := readLine("Reading time (minutes): ")
		fmt.Sscan(readingTime, &article.ReadingTime)
		article.Category = readLine("Category (tutorial/blog/documentation): ")
		manager.Articles = append(manager.Articles, article)

	case "4":
		generic := GenericBookmark{BaseBookmark: base}
		fmt.Println(Yellow + "\n--- Category Details ---" + Reset)
		generic.Category = readLine("Category: ")
		manager.Generic = append(manager.Generic, generic)
	}

	saveBookmarks(*manager)
	fmt.Println(Green + "\n✅ Bookmark saved successfully!" + Reset)
	fmt.Print(Cyan + "\nPress Enter to continue..." + Reset)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func viewBookmarks(manager BookmarkManager) {
	clearScreen()
	fmt.Println(Bold + Blue + "\n=== View Bookmarks ===" + Reset)
	fmt.Println(Cyan + "\n1. 🎨 Design Inspiration")
	fmt.Println("2. 💻 Code/Stack Overflow")
	fmt.Println("3. 📚 Articles/Blogs")
	fmt.Println("4. 🔖 Generic Bookmarks")
	fmt.Println("5. 📋 All Bookmarks" + Reset)

	choice := readLine("\nChoice: ")

	switch choice {
	case "1":
		printDesignBookmarks(manager.Designs)
	case "2":
		printCodeBookmarks(manager.Code)
	case "3":
		printArticleBookmarks(manager.Articles)
	case "4":
		printGenericBookmarks(manager.Generic)
	case "5":
		printAllBookmarks(manager)
	}

	fmt.Print(Cyan + "\nPress Enter to continue..." + Reset)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func printDesignBookmark(b DesignBookmark) {
	fmt.Printf(Bold+"\n%s%s%s\n"+Reset, Blue, b.Title, Reset)
	fmt.Printf(Yellow+"🔗 URL: %s%s\n", Reset, b.URL)
	fmt.Printf(Yellow+"📱 Type: %s%s\n", Reset, b.Type)
	fmt.Printf(Yellow+"📂 Category: %s%s\n", Reset, b.Category)
	fmt.Printf(Yellow+"🎨 Color Scheme: %s%s\n", Reset, b.ColorScheme)
	if b.HasRepo {
		fmt.Printf(Yellow+"📦 Repository: %s%s\n", Reset, b.RepoURL)
	}
	printBaseInfo(b.BaseBookmark)
}

func printCodeBookmark(b CodeBookmark) {
	fmt.Printf(Bold+"\n%s%s%s\n"+Reset, Magenta, b.Title, Reset)
	fmt.Printf(Yellow+"🔗 URL: %s%s\n", Reset, b.URL)
	fmt.Printf(Yellow+"💻 Language: %s%s\n", Reset, b.Language)
	fmt.Printf(Yellow+"✨ Solution: %s%s\n", Reset, b.Solution)
	if b.WorksInVersion != "" {
		fmt.Printf(Yellow+"📌 Version: %s%s\n", Reset, b.WorksInVersion)
	}
	printBaseInfo(b.BaseBookmark)
}

func printArticleBookmark(b ArticleBookmark) {
	fmt.Printf(Bold+"\n%s%s%s\n"+Reset, Cyan, b.Title, Reset)
	fmt.Printf(Yellow+"🔗 URL: %s%s\n", Reset, b.URL)
	fmt.Printf(Yellow+"✍️  Author: %s%s\n", Reset, b.Author)
	fmt.Printf(Yellow+"⏱️  Reading Time: %d minutes%s\n", b.ReadingTime, Reset)
	fmt.Printf(Yellow+"📂 Category: %s%s\n", Reset, b.Category)
	printBaseInfo(b.BaseBookmark)
}

func printGenericBookmark(b GenericBookmark) {
	fmt.Printf(Bold+"\n%s%s%s\n"+Reset, Green, b.Title, Reset)
	fmt.Printf(Yellow+"🔗 URL: %s%s\n", Reset, b.URL)
	fmt.Printf(Yellow+"📂 Category: %s%s\n", Reset, b.Category)
	printBaseInfo(b.BaseBookmark)
}

func printBaseInfo(b BaseBookmark) {
	if b.Notes != "" {
		fmt.Printf(Yellow+"📝 Notes: %s%s\n", Reset, b.Notes)
	}
	if len(b.Tags) > 0 {
		fmt.Printf(Yellow+"🏷️  Tags: %s%s\n", Reset, strings.Join(b.Tags, ", "))
	}
	fmt.Printf(Yellow+"📅 Added: %s%s\n", b.CreatedAt.Format("2006-01-02 15:04:05"), Reset)
	fmt.Println(Yellow + "---------------" + Reset)
}

func searchBookmarks(manager BookmarkManager) {
	clearScreen()
	fmt.Println(Bold + Blue + "\n=== Search Bookmarks ===" + Reset)
	term := readLine("\nSearch term: ")
	term = strings.ToLower(term)

	fmt.Println(Yellow + "\nSearch Results:" + Reset)
	fmt.Println(Yellow + "---------------" + Reset)

	found := false
	for _, b := range manager.Designs {
		if containsSearchTerm(b.BaseBookmark, term) {
			printDesignBookmark(b)
			found = true
		}
	}
	for _, b := range manager.Code {
		if containsSearchTerm(b.BaseBookmark, term) {
			printCodeBookmark(b)
			found = true
		}
	}
	for _, b := range manager.Articles {
		if containsSearchTerm(b.BaseBookmark, term) {
			printArticleBookmark(b)
			found = true
		}
	}
	for _, b := range manager.Generic {
		if containsSearchTerm(b.BaseBookmark, term) {
			printGenericBookmark(b)
			found = true
		}
	}

	if !found {
		fmt.Println(Red + "No matches found" + Reset)
	}

	fmt.Print(Cyan + "\nPress Enter to continue..." + Reset)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func containsSearchTerm(b BaseBookmark, term string) bool {
	return strings.Contains(strings.ToLower(b.URL), term) ||
		strings.Contains(strings.ToLower(b.Title), term) ||
		strings.Contains(strings.ToLower(b.Notes), term) ||
		containsTag(b.Tags, term)
}

func containsTag(tags []string, term string) bool {
	for _, tag := range tags {
		if strings.Contains(strings.ToLower(tag), term) {
			return true
		}
	}
	return false
}

func printDesignBookmarks(bookmarks []DesignBookmark) {
	for _, b := range bookmarks {
		printDesignBookmark(b)
	}
}

func printCodeBookmarks(bookmarks []CodeBookmark) {
	for _, b := range bookmarks {
		printCodeBookmark(b)
	}
}

func printArticleBookmarks(bookmarks []ArticleBookmark) {
	for _, b := range bookmarks {
		printArticleBookmark(b)
	}
}

func printGenericBookmarks(bookmarks []GenericBookmark) {
	for _, b := range bookmarks {
		printGenericBookmark(b)
	}
}

func printAllBookmarks(manager BookmarkManager) {
	printDesignBookmarks(manager.Designs)
	printCodeBookmarks(manager.Code)
	printArticleBookmarks(manager.Articles)
	printGenericBookmarks(manager.Generic)
}

func loadBookmarks() BookmarkManager {
	data, err := os.ReadFile("bookmarks.json")
	if err != nil {
		return BookmarkManager{}
	}

	var manager BookmarkManager
	json.Unmarshal(data, &manager)
	return manager
}

func saveBookmarks(manager BookmarkManager) {
	data, _ := json.MarshalIndent(manager, "", "  ")
	os.WriteFile("bookmarks.json", data, 0644)
}

func exportToMarkdown(manager BookmarkManager) {
	var md strings.Builder

	md.WriteString("# 📚 Bookmark Collection\n\n")

	if len(manager.Designs) > 0 {
		md.WriteString("## 🎨 Design Inspirations\n\n")
		for _, b := range manager.Designs {
			md.WriteString(fmt.Sprintf("### [%s](%s)\n", b.Title, b.URL))
			md.WriteString(fmt.Sprintf("- **Type:** %s\n", b.Type))
			md.WriteString(fmt.Sprintf("- **Category:** %s\n", b.Category))
			md.WriteString(fmt.Sprintf("- **Color Scheme:** %s\n", b.ColorScheme))
			if b.HasRepo {
				md.WriteString(fmt.Sprintf("- **Repository:** [Link](%s)\n", b.RepoURL))
			}
			writeCommonMarkdown(&md, b.BaseBookmark)
		}
	}

	if len(manager.Code) > 0 {
		md.WriteString("\n## 💻 Code Solutions\n\n")
		for _, b := range manager.Code {
			md.WriteString(fmt.Sprintf("### [%s](%s)\n", b.Title, b.URL))
			md.WriteString(fmt.Sprintf("- **Language:** %s\n", b.Language))
			md.WriteString(fmt.Sprintf("- **Solution:** %s\n", b.Solution))
			if b.WorksInVersion != "" {
				md.WriteString(fmt.Sprintf("- **Version:** %s\n", b.WorksInVersion))
			}
			writeCommonMarkdown(&md, b.BaseBookmark)
		}
	}

	if len(manager.Articles) > 0 {
		md.WriteString("\n## 📚 Articles\n\n")
		for _, b := range manager.Articles {
			md.WriteString(fmt.Sprintf("### [%s](%s)\n", b.Title, b.URL))
			md.WriteString(fmt.Sprintf("- **Author:** %s\n", b.Author))
			md.WriteString(fmt.Sprintf("- **Reading Time:** %d minutes\n", b.ReadingTime))
			md.WriteString(fmt.Sprintf("- **Category:** %s\n", b.Category))
			writeCommonMarkdown(&md, b.BaseBookmark)
		}
	}

	if len(manager.Generic) > 0 {
		md.WriteString("\n## 🔖 Other Bookmarks\n\n")
		for _, b := range manager.Generic {
			md.WriteString(fmt.Sprintf("### [%s](%s)\n", b.Title, b.URL))
			md.WriteString(fmt.Sprintf("- **Category:** %s\n", b.Category))
			writeCommonMarkdown(&md, b.BaseBookmark)
		}
	}

	filename := fmt.Sprintf("bookmarks_%s.md", time.Now().Format("2006-01-02"))
	os.WriteFile(filename, []byte(md.String()), 0644)
	fmt.Printf(Green+"\n✅ Successfully exported to %s\n"+Reset, filename)
	fmt.Print(Cyan + "\nPress Enter to continue..." + Reset)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func writeCommonMarkdown(md *strings.Builder, b BaseBookmark) {
	if b.Notes != "" {
		md.WriteString(fmt.Sprintf("- **Notes:** %s\n", b.Notes))
	}
	if len(b.Tags) > 0 {
		md.WriteString(fmt.Sprintf("- **Tags:** %s\n", strings.Join(b.Tags, ", ")))
	}
	md.WriteString(fmt.Sprintf("- **Added:** %s\n\n", b.CreatedAt.Format("2006-01-02 15:04:05")))
}

func main() {
	manager := loadBookmarks()

	for {
		showMainMenu()
		choice := readLine("\nChoice: ")

		switch choice {
		case "1":
			addBookmark(&manager)
		case "2":
			viewBookmarks(manager)
		case "3":
			searchBookmarks(manager)
		case "4":
			exportToMarkdown(manager)
		case "5":
			fmt.Println(Green + "\nGoodbye! 👋" + Reset)
			return
		default:
			fmt.Println(Red + "\n❌ Invalid choice, please try again" + Reset)
			fmt.Print(Cyan + "\nPress Enter to continue..." + Reset)
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}
}