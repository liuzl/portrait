package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

var cfmt = `curl -s localhost:10101/index/phone/query -X POST -d '%s' | jq -r '.results[].keys[]' > %s`

func pilosa(query, file string) {
	_, err := os.Stat(file)
	if err == nil || !os.IsNotExist(err) {
		return
	}
	c := fmt.Sprintf(cfmt, query, file)
	cmd := exec.Command("/bin/bash", "-c", c)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("cmd.StderrPipe()", err)
		return
	}
	if err = cmd.Start(); err != nil {
		fmt.Println("cmd.Start()", err)
		return
	}
	bytes, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = cmd.Wait(); err != nil {
		fmt.Println("cmd.Wait()", err, " stderr: ", string(bytes))
		os.Remove(file)
		return
	}
	okfile := file + ".ok"
	os.OpenFile(okfile, os.O_RDONLY|os.O_CREATE, 0666)
}

func setdiff(f1, f2, f3 string) error {
	_, err := os.Stat(f3)
	if err == nil {
		return fmt.Errorf("%s already exists", f3)
	}
	if !os.IsNotExist(err) {
		return err
	}
	if _, err = os.Stat(f1); err != nil {
		return err
	}
	if _, err = os.Stat(f2); err != nil {
		return err
	}
	c := fmt.Sprintf("grep -vxFf %s %s > %s", f2, f1, f3)
	cmd := exec.Command("/bin/bash", "-c", c)
	if err = cmd.Start(); err != nil {
		return err
	}
	if err = cmd.Wait(); err != nil {
		os.Remove(f3)
		return err
	}
	return nil
}
