/*
Assigment 1
Nama	: Aden Hidayatuloh
kelas	: GLNG-KS08
Kode Peserta : GLNG-KS08-011
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type StudentData struct {
	Id                 string `json:"id"`
	Student_code       string `json:"student_code"`
	Student_name       string `json:"student_name"`
	Student_address    string `json:"student_address"`
	Student_occupation string `json:"student_occupation"`
	Joining_reason     string `json:"joining_reason"`
}

type Students struct {
	Students []StudentData `json:"participants"`
}

// function to display student data
func (StudentData *StudentData) ShowBio() {

	fmt.Printf("ID : %s \n", StudentData.Id)
	fmt.Printf("Kode Peserta : %s \n", StudentData.Student_code)
	fmt.Printf("nama : %s \n", StudentData.Student_name)
	fmt.Printf("alamat : %s \n", StudentData.Student_address)
	fmt.Printf("pekerjaan : %s \n", StudentData.Student_occupation)
	fmt.Printf("alasan : %s \n", StudentData.Joining_reason)

}

// function to find and display student data
func FindStudentData(StudentCode string, MapStudentData map[string]*StudentData) {

	isFound := false
	for key, value := range MapStudentData {

		if strings.EqualFold(StudentCode, key) {
			isFound = true
			value.ShowBio()
			break
		}
	}

	if !isFound {
		fmt.Printf("Maaf, data yang anda cari %s tidak ada", StudentCode)
	}

}

// function to convert json to data type Students
func ConvertJsonToStructStudents(filePath string, objectStruct *Students) {
	// Read JSON file content
	data, err := os.ReadFile("data/participants.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal JSON data into variable students
	if err := json.Unmarshal(data, &objectStruct); err != nil {
		log.Fatal(err)
	}

}

func main() {

	var students Students
	var checkStudentCode string

	//handling if the user doesn't input argument
	if len(os.Args) > 1 {
		checkStudentCode = os.Args[1]

	} else {
		checkStudentCode = ""
	}

	ConvertJsonToStructStudents("data/participants.json", &students)

	//Create a student data map and give the value with data type StudentData
	StudentDataMap := make(map[string]*StudentData)
	for i := range students.Students {
		StudentDataMap[students.Students[i].Student_code] = &students.Students[i]
	}

	if checkStudentCode != "" {

		FindStudentData(checkStudentCode, StudentDataMap)
	}

}
