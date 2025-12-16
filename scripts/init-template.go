package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func updateGoMod(newModule string) error {
	path := filepath.Join("go.mod")

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "module ") {
			lines[i] = "module " + newModule
			break
		}
	}

	err = os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
	if err == nil {
		fmt.Printf("✔ go.mod updated → module %s\n", newModule)
	}
	return err
}

func updateReadme(newName string) error {
	path := filepath.Join("README.md")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	content := string(data)
	content = strings.ReplaceAll(content, "go-api-rest-template", newName)
	content = strings.ReplaceAll(content, "Go REST API Template", newName)

	err = os.WriteFile(path, []byte(content), 0644)
	if err == nil {
		fmt.Println("✔ README.md updated")
	}
	return err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("❌ Missing module name.")
		fmt.Println("Usage:")
		fmt.Println("  go run scripts/init-template.go github.com/your-org/your-project")
		os.Exit(1)
	}

	newModule := os.Args[1]

	fmt.Printf("🔧 Initializing Go template for: %s\n", newModule)

	if err := updateGoMod(newModule); err != nil {
		fmt.Println("❌ Failed updating go.mod:", err)
		os.Exit(1)
	}

	if err := updateReadme(newModule); err != nil {
		fmt.Println("❌ Failed updating README:", err)
		os.Exit(1)
	}

	fmt.Println("🚀 Template initialized successfully!")
	fmt.Println("👉 Next steps:")
	fmt.Println("   go mod tidy")
	fmt.Println("   git init && git commit -m \"init project\"")
}
