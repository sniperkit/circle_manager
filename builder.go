package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func beegoBuild() error {
	return executer("beego run", "bee", []string{"run", "-runmode=dev", "-gendoc=true"})
}

func dockerBuild(url string) error {
	return executer("docker build", "docker", []string{"build", "-t", url, "."})
}

func dockerPush(url string) error {
	return executer("docker push", "docker", []string{"push", url})
}

func executer(cmdType string, cmdName string, cmdArgs []string) error {
	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Dir = envs.RootPath
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return err
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s out | %s\n", cmdType, scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return err
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		return err
	}
	return nil
}

func waitCommentRouterFile(commentRouterExist chan bool) {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	cnt := 0
	limit := 100
	for {
		select {
		case <-t.C:
			cnt++
			if cnt > limit {
				commentRouterExist <- false
				return
			}
			subDirectoryFiles(filepath.Join(envs.RootPath, "routers"), func(info os.FileInfo) error {
				if strings.Index(info.Name(), "commentsRouter") == 0 {
					commentRouterExist <- true
				}
				return nil
			})
		}
	}
}
