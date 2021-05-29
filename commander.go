package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func sendCommand(commands ...string) (string, error) {
	var totalCommand string
	for i, c := range commands {
		if i == 0 {
			totalCommand = c
		} else {
			totalCommand = totalCommand + " " + c
		}
	}
	var outputText string
	cmd := exec.Command("bash", "-c", totalCommand)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Çıktı okunurken hata oluştu:", err)
	}
	//çıktı okuyucusunun tanımlanması
	output := bufio.NewScanner(cmdReader)
	go func() {
		for output.Scan() {
			outputText = output.Text()
		}
	}()
	//komutun başlatılması
	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "", err
	} else {
		err = cmd.Wait()
		if err != nil {
			return "", err
		}
	}
	//komutun çalışması

	return outputText, err
}
