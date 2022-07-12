package main

import (
	"archive/zip"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func phishing() {
	var err error
	//url := "http://google.com"
	url := "http://salesforce.sbm-demo.xyz/phishing"
	err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	fmt.Println("The browser has been opened.")
	if err != nil {
		log.Fatal(err)
	}
}

func antibot() {
	//url := "http://google.com"
	url := "http://www.threat-cloud.com/test/files/HighConfidenceBot.html"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("user-agent", "*<|>*")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func avtest() {
	//url := "https://go.dev/images/gophers/ladder.svg"
	//url := "https://admissions.hku.hk/sites/default/files/2021-10/HKU%20Non%20JUPAS%20Admissions%202022.pdf"
	url := "https://secure.eicar.org/eicar.com"
	r, _ := http.Get(url)
	//content, _ := io.ReadAll(r.Body)
	file, err := os.Create("eicar.com")
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()
	_, err_1 := io.Copy(file, r.Body)
	if err_1 != nil {
		fmt.Println(err_1)
	}
	defer file.Close()
	//fmt.Printf("%s", content)
}

func cred_dump() {
	var file_path string
	path, _ := os.Getwd()
	archive, err := zip.OpenReader("testtool.zip")
	if err != nil {
		fmt.Printf("Open archive failure, error=[%v]\n", err)
	}
	defer archive.Close()
	//fmt.Println(archive.File[0].Name)
	for _, f := range archive.File {
		if f.Name == "procdump.exe" {
			file_path := filepath.Join(path, f.Name)
			dstFile, e := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if e != nil {
				fmt.Printf("Create file failure, error=[%v]\n", e)
			}
			file_in_archive, e1 := f.Open()
			if e1 != nil {
				fmt.Printf("Open archive filefailure, error=[%v]\n", e1)
			}
			_, e2 := io.Copy(dstFile, file_in_archive)
			if e2 != nil {
				fmt.Printf("save archive file failure, error=[%v]\n", e2)
			}
			dstFile.Close()
			file_in_archive.Close()
			break
		}
	}
	//parameter := " -ma lsass.exe lsass.dmp"
	payload := file_path + " -ma lsass.exe lsass.dmp"
	bat_file, e := os.Create("dump.bat")
	if e != nil {
		fmt.Printf("Create bat file failure, error=[%v]\n", e)
	}
	payload_w := strings.NewReader(payload)
	_, e3 := io.Copy(bat_file, payload_w)
	if e3 != nil {
		fmt.Printf("Copy payload failure, error=[%v]\n", e3)
	}
	cmd := exec.Command("cmd", payload)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func on_access_av_base64() {
	file, e := os.Create("eicar_1.com")
	if e != nil {
		fmt.Printf("Create file failure, error=[%v]\n", e)
	}
	eicar := "WDVPIVAlQEFQWzRcUFpYNTQoUF4pN0NDKTd9JEVJQ0FSLVNUQU5EQVJELUFOVElWSVJVUy1URVNULUZJTEUhJEgrSCo="
	decode_e, err := base64.StdEncoding.DecodeString(eicar)
	if err != nil {
		fmt.Printf("Base64 decode failure, error=[%v]\n", err)
	} else {
		payload := strings.NewReader(string(decode_e))
		_, e1 := io.Copy(file, payload)
		if e1 != nil {
			fmt.Printf("Copy payload failure, error=[%v]\n", e1)
		} else {
			fmt.Println("Done!")
		}
	}

}

func url_filtering() {
	type url_list struct {
		index    int
		catagory string
		url      string
	}
	var url_cat [2]url_list
	url_cat[0] = url_list{0, "Gambling", "http://www.google.com"}
	url_cat[1] = url_list{1, "News", "http://bbc.com"}
	fmt.Println("ID : Catagory : URL")
	for i := 0; i < len(url_cat); i++ {
		fmt.Println("ID: ", url_cat[i].index, " Catagory:", url_cat[i].catagory, " URL:", url_cat[i].url)
	}
	var user_choice int
	var valid_input bool = false
	for !valid_input {
		fmt.Println("Please select the Id of the URL to start the test: ")
		fmt.Scanln(&user_choice)
		if !(user_choice == 0 || user_choice == 1) {
			fmt.Println("Invalid input, please enter again.")
			valid_input = false
		} else {
			valid_input = true
			err := exec.Command("rundll32", "url.dll,FileProtocolHandler", url_cat[user_choice].url).Start()
			fmt.Println("The browser has been opened.")
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
