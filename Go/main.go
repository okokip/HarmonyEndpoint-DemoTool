package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type test_case struct {
	index_c int
	name    string
}

func main() {
	var list_test [6]test_case
	var cont_test bool = true
	var user_input int
	list_test[0] = test_case{1, "Anti-bot"}
	list_test[1] = test_case{2, "Zero-Phishing"}
	list_test[2] = test_case{3, "Web download Anti-virus"}
	list_test[3] = test_case{4, "on-access Anti-virus Scanning"}
	list_test[4] = test_case{5, "Behaviour Guard"}
	list_test[5] = test_case{6, "URL filtering"}

	for cont_test {
		fmt.Println("Welcome to the Check Point Harmony Endpoint testing tool!")
		for i := 0; i < len(list_test); i++ {
			fmt.Println(list_test[i].index_c, ": ", list_test[i].name)
		}
		fmt.Println("7 : Quit the program")
		fmt.Println("Please enter the test case number that you would like to run: ")
		fmt.Scanln(&user_input)
		if !(user_input <= 7 && user_input > 0) {
			fmt.Println("Invalid input, please enter again. e.g. If you would like to run test case 2 anti-bot, please enter 2.")
			cont_test = true
		} else if user_input == 7 {
			fmt.Println("Bye Bye!")
			cont_test = false
		} else {
			switch user_input {
			case 1:
				fmt.Println("test case " + strconv.Itoa(user_input) + " " + list_test[user_input-1].name + " has been executed!")
				antibot()
			case 2:
				fmt.Println("test case " + strconv.Itoa(user_input) + " " + list_test[user_input-1].name + " has been executed!")
				phishing()
			case 3:
				avtest()
				fmt.Println("test case " + strconv.Itoa(user_input) + " " + list_test[user_input-1].name + " has been executed!")
			case 4:
				on_access_av_base64()
				fmt.Println("test case " + strconv.Itoa(user_input) + " " + list_test[user_input-1].name + " has been executed!")
			case 5:
				cred_dump()
				fmt.Println("test case " + strconv.Itoa(user_input) + " " + list_test[user_input-1].name + " has been executed!")
			case 6:
				url_filtering()
				fmt.Println("test case " + strconv.Itoa(user_input) + " " + list_test[user_input-1].name + " has been executed!")
			}
			var new_test string
			var valid_input bool = false
			for !valid_input {
				fmt.Println("Do you want to continue on other test case? (Y/N)")
				fmt.Scanln(&new_test)
				if new_test != "Y" && new_test != "N" {
					fmt.Println("Invalid input, please enter Y/N: ")
					valid_input = false
				} else {
					valid_input = true
				}
			}
			if new_test == "N" {
				cont_test = false
			} else {
				//continue
				fmt.Println("Clear the terminal!")
				cmd := exec.Command("cmd", "/c", "cls")
				cmd.Stdout = os.Stdout
				cmd.Run()
			}
		}
	}
}
