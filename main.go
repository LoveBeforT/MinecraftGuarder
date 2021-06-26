package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	server_file   = "forge-1.12.2-14.23.5.2847-universal.jar"
	server_script = "sh run.sh"
)

func main() {
	fmt.Println("Minecraft Guarder Start !")
	var result []byte
	var err error

	for {
		time.Sleep(15 * time.Second)
		log.Println("Scaning server ...")
		cmd := exec.Command("bash", "-c", "ps aux | grep java")
		if result, err = cmd.Output(); err != nil {
			log.Println(err)
			os.Exit(1)
		}

		if !strings.Contains(string(result), server_file) {
			log.Println("Server not found , try to restart server ...")
			cmd = exec.Command("bash", "-c", server_script)
			if _, err = cmd.Output(); err != nil {
				log.Println("restart server failed!", err)
			}
		}
	}
}
