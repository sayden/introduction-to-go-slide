package main

import (
	"io"
	"os"
	"os/exec"
)

func main() {
	c1 := exec.Command("ls", "-lh", "/")
	c2 := exec.Command("awk", "{print $5, \"\011\" $9}")

	r1, w1 := io.Pipe()
	c1.Stdout = w1
	c2.Stdin = r1

	c2.Stdout = os.Stdout

	c1.Start()
	c2.Start()
	c1.Wait()
	w1.Close()
}
