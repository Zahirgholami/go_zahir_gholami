package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for project name
	fmt.Print("Enter the project name: ")
	projectName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading project name:", err)
		return
	}
	projectName = strings.TrimSpace(projectName)

	if strings.Contains(projectName, " ") {
		fmt.Println("Project name cannot contain spaces.")
		return
	}

	fmt.Print("Enter the project type (simple/api): ")
	projectType, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading project type:", err)
		return
	}
	projectType = strings.TrimSpace(projectType)

	err = os.Mkdir(projectName, 0755)
	if err != nil {
		fmt.Println("Error creating project directory:", err)
		return
	}

	err = os.Chdir(projectName)
	if err != nil {
		fmt.Println("Error changing to project directory:", err)
		return
	}

	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error initializing Go module:", err)
		return
	}

	switch projectType {
	case "simple":
		createSimpleProject()
	case "api":
		createAPIProject()
	default:
		fmt.Println("Invalid project type. Please enter 'simple' or 'api'.")
	}
}

func createSimpleProject() {
	mainContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
`
	err := os.WriteFile("main.go", []byte(mainContent), 0644)
	if err != nil {
		fmt.Println("Error creating main.go file:", err)
	}
}

func createAPIProject() {
	folders := []string{"routes", "models", "repositories", "services", "configs", "controllers"}
	for _, folder := range folders {
		err := os.Mkdir(folder, 0755)
		if err != nil {
			fmt.Println("Error creating folder:", folder, err)
			return
		}
	}

	mainContent := `package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})
	http.ListenAndServe(":8080", nil)
}
`
	err := os.WriteFile("main.go", []byte(mainContent), 0644)
	if err != nil {
		fmt.Println("Error creating main.go file:", err)
	}
}
