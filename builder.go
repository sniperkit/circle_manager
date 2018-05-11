package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

var beegoAppProcess *os.Process

func genBeegoAppResource() error {
	fmt.Println("Starting app of beego")
	if err := beegoBuild(); err != nil {
		return err
	}
	fmt.Println("Started app of beego")
	commentRouterExistChan := make(chan bool)
	go waitCommentRouterFile(commentRouterExistChan)

	fmt.Println("Wait for generate of resources")
	commentRouterExist := <-commentRouterExistChan
	if commentRouterExist {
		fmt.Println("Genarated resources! Exit app of beego")
		beegoAppProcess.Kill()
		return nil
	}
	return errors.New("비정상 종료")
}

func beegoBuild() error {
	fmt.Println("Executing app build")
	var err error
	beegoAppProcess, err = executer("beego run", "bee", []string{"run", "-runmode=dev", "-gendoc=true", "-e=./"}, false)
	return err
}

func dockerBuild(url string) error {
	fmt.Println("Executing docker build")
	_, err := executer("docker build", "docker", []string{"build", "-t", url, "."}, true)
	return err
}

func dockerPush(url string) error {
	fmt.Println("Executing docker push")
	_, err := executer("docker push", "docker", []string{"push", url}, true)
	return err
}

func executeSys(cmdType string, cmdName string, cmdArgs []string) error {
	binary, lookErr := exec.LookPath(envs.RootPath)
	if lookErr != nil {
		panic(lookErr)
	}

	env := os.Environ()

	execErr := syscall.Exec(binary, cmdArgs, env)
	if execErr != nil {
		panic(execErr)
	}
	return nil
}

func executer(cmdType string, cmdName string, cmdArgs []string, cmdWait bool) (*os.Process, error) {
	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Dir = envs.RootPath
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return nil, err
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
		return nil, err
	}

	if cmdWait {
		err = cmd.Wait()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
			return nil, err
		}
	}

	return cmd.Process, nil
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
			fmt.Println("Finding resource of beego")
			if cnt > limit {
				fmt.Println("Failed finding resource of beego")
				commentRouterExist <- false
				return
			}
			subDirectoryFiles(filepath.Join(envs.RootPath, "routers"), func(info os.FileInfo) error {
				if strings.Index(info.Name(), "commentsRouter") == 0 {
					fmt.Println("Suceessed finding resource of beego")
					t.Stop()
					commentRouterExist <- true
				}
				return nil
			})
		}
	}

}
