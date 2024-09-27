package packages

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

type Student struct {
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	DateOfBirth string   `json:"date_of_birth"`
	CreditBook  []Credit `json:"credit book"`
}

type Credit struct {
	Subject  string `json:"subject"`
	ExamDate string `json:"exam_date"`
	Teacher  string `json:"teacher"`
	Grade    string `json:"grade"` // Changed from int to string
}

type Data struct {
	Students []Student `json:"students"`
}

// Find students with the highest average grade
func findBestStudents(students []Student) []Student {
	var bestStudents []Student
	bestAverage := 0.0

	for _, student := range students {
		totalGrades := 0
		for _, credit := range student.CreditBook {
			grade, err := strconv.Atoi(credit.Grade) // Convert grade from string to int
			if err != nil {
				log.Printf("Invalid grade for student %s %s in subject %s: %v", student.FirstName, student.LastName, credit.Subject, err)
				continue
			}
			totalGrades += grade
		}
		average := float64(totalGrades) / float64(len(student.CreditBook))

		if average > bestAverage {
			bestAverage = average
			bestStudents = []Student{student}
		} else if average == bestAverage {
			bestStudents = append(bestStudents, student)
		}
	}

	return bestStudents
}

// Print student data in table format
func printBestStudents(students []Student) {
	// Initialize a new tabwriter
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	// Print the table header
	fmt.Fprintln(writer, "Full Name\tSubject\tGrade\t")

	for _, student := range students {
		fullName := student.FirstName + " " + student.LastName
		for _, credit := range student.CreditBook {
			fmt.Fprintf(writer, "%s\t%s\t%s\t\n", fullName, credit.Subject, credit.Grade)
		}
	}

	// Flush the writer to output the table
	writer.Flush()
}

func GetBestStudents() {
	// Read the JSON file
	file, err := os.ReadFile("list_of_group.json")
	if err != nil {
		log.Fatalf("Failed to read the file: %v", err)
	}

	// Parse the JSON data
	var data Data
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatalf("Failed to parse JSON data: %v", err)
	}

	// Find the student(s) with the best average grade
	bestStudents := findBestStudents(data.Students)

	// Print the results in table format
	printBestStudents(bestStudents)
}
