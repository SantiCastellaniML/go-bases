package main

import (
	"bases/ejercicios/ej-panic/ejercicio1/client"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var errorAmount int

func main() {
	fileName := "customers.txt"

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}

		fmt.Println("ejecución finalizada")
	}()

	//clientData := map[int]client.Client{}

	clientData := ReadFile(fileName)

	//fmt.Println(str)

	fmt.Println(clientData)

	defer func() {
		println("End of execution")
		if err := recover(); err != nil {
			fmt.Println("Several errors were detected at runtime")
			fmt.Println(err)
		}
	}()

	//error por cliente repetido:
	//RegisterClient("Juan,1,123456789,Av. Siempreviva 742", &clientData)

	//error por campo vacío o zero value:
	//RegisterClient("Juan,0,123321,Av. Siempreviva 742", &clientData)

	RegisterClient("Juan,5,789789,Av. Siempreviva 742", &clientData)

	fmt.Println(clientData)
}

func validateFields(fields []string) (msg string, result bool) {
	if len(fields) != 4 {
		return "not enough fields", false
	}

	if fields[0] == "" {
		return "name is empty", false
	}

	if fields[1] == "" || fields[1] == "0" {
		return "ID is empty", false
	}

	if fields[2] == "" || fields[2] == "0" {
		return "phone is empty", false
	}

	if fields[3] == "" {
		return "home is empty", false
	}

	return "", true
}

func RegisterClient(data string, clientStorage *map[int]client.Client) (err error) {
	fields := strings.Split(data, ",")

	msg, ok := validateFields(fields)

	if !ok {
		panic(msg)
	}

	phone, err := strconv.Atoi(fields[2])
	if err != nil {
		return errors.New("Phone number is not valid")
	}

	id, err := strconv.Atoi(fields[1])

	if err != nil {
		return errors.New("ID is not valid")
	}

	if _, exists := (*clientStorage)[id]; exists {
		errorAmount++
		panic("Error: client already exists") //Tarea 1
	}

	client := client.Client{
		Name:  fields[0],
		Phone: phone,
		Home:  fields[3],
	}

	//Everything is ok, register client
	(*clientStorage)[id] = client
	return nil
}

func ReadFile(fileName string) map[int]client.Client {
	file, err := os.ReadFile(fileName)

	if err != nil {
		errorAmount++
		panic("The indicated file was not found or is damaged")
	}

	lines := strings.Split(string(file), "\n")

	clientData := map[int]client.Client{}

	for _, line := range lines {
		RegisterClient(line, &clientData)

		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
	}

	return clientData
}
