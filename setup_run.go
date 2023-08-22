package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func runTerraformCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	// Run terraform init
	err = runTerraformCommand("terraform", "init")
	if err != nil {
		fmt.Println("Error running terraform init:", err)
		return
	}

	// Run terraform plan
	err = runTerraformCommand("terraform", "plan")
	if err != nil {
		fmt.Println("Error running terraform plan:", err)
		return
	}

	// Run terraform apply
	err = runTerraformCommand("terraform", "apply", "-auto-approve")
	if err != nil {
		fmt.Println("Error running terraform apply:", err)
		return
	}

	// Run Terraform Destroy
	err = runTerraformCommand("terraform", "destroy", "-auto-approve")
	if err != nil {
		fmt.Println("Error running terraform destroy:", err)
		return
	}

}
