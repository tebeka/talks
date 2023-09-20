package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

const (
	Version = "0.5.0"
)

func die(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func exists(path string, dir bool) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}

	if dir {
		return fi.IsDir()
	} else {
		return !fi.IsDir()
	}
}

func tmpZip() string {
	return fmt.Sprintf("%s/nrsc-%d.zip", os.TempDir(), time.Now().Unix())
}

func mkzip(root, zip string, zipArgs []string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	if err = os.Chdir(root); err != nil {
		return err
	}
	defer os.Chdir(pwd)

	args := []string{"-r", zip, "."}
	args = append(args, zipArgs...)

	cmd := exec.Command("zip", args...)
	if err = cmd.Run(); err != nil {
		return err
	}

	return nil
}

func appendFile(dest, src string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(dest, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	out.Seek(0, os.SEEK_END)
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}


func fixZipOffset(exe string) error {
	return exec.Command("zip", "-q", "-A", exe).Run()
}

func main() {
	flag.Usage = func() {
        fmt.Println("usage: nrsc EXECTABLE RESOURCE_DIR [ZIP OPTIONS]")
	}

	version := flag.Bool("version", false, "show version and exit")
	flag.Parse()

	if *version {
		fmt.Printf("nrsc version %s\n", Version)
	}

	if flag.NArg() < 2 {
		die("error: Wrong number of arguments\n")
	}

	exe := flag.Arg(0)
	if !exists(exe, false) {
		die("error: `%s` is not a file", exe)
	}

	root := flag.Arg(1)
	if !exists(root, true) {
		die("error: `%s` is not a directory", root)
	}

	zip := tmpZip()
	defer os.Remove(zip)

	if err := mkzip(root, zip, flag.Args()[2:]); err != nil {
		die("error: can't create zip - %s", err)
	}

	if err := appendFile(exe, zip); err != nil {
		die("error: can't append zip to %s - %s", exe, err)
	}

	if err := fixZipOffset(exe); err != nil {
		die("error: can't fix zip offset in %s - %s", exe, err)
	}
}
