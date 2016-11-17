package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

const (
    GOOLE_HOST = "https://raw.githubusercontent.com/racaljk/hosts/master/hosts"
	LOCAL_FILE = "C:/Windows/System32/drivers/etc/hosts"
)

func main() {
	resp, _ := http.Get(GOOLE_HOST)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	err := ioutil.WriteFile(LOCAL_FILE, body, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Update host success.")
	cmd := exec.Command("cmd.exe", "/c", "ipconfig /flushdns")
	err = cmd.Run()
	if err != nil {
		fmt.Println("flush DNS error. ", err)
		return
	}
	fmt.Println("flush DNS success.")
}
