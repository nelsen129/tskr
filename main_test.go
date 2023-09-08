package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func createDir(t *testing.T, d string) {
	err := os.Mkdir(d, 0755)
	if err != nil {
		t.Fatalf("Error: could not create directory %s, %v", d, err)
	}

	t.Cleanup(func() {
		err := os.RemoveAll(d)
		if err != nil {
			t.Fatalf("Error: could not delete directory %s, %s", d, err)
		}
	})
}

func logToBuffer(t *testing.T, buf *bytes.Buffer) {
	log.SetOutput(buf)
	t.Cleanup(func() {
		log.SetOutput(os.Stderr)
	})
}

func TestInitSuccess(t *testing.T) {
	tmpdir := filepath.Join(os.TempDir(), "tskr_testdir_main")
	createDir(t, tmpdir)
	os.Chdir(tmpdir)

	args := os.Args[0:1] // Name of the program
	args = append(args, "init")
	os.Args = args
	main()
}

func TestInitDirectoryIsNotEmpty(t *testing.T) {
	tmpdir := filepath.Join(os.TempDir(), "tskr_testdir_main")
	tmpfile := "tmptestfile"
	createDir(t, tmpdir)

	_, err := os.Create(filepath.Join(tmpdir, tmpfile))
	if err != nil {
		t.Fatalf("Error: could not create file %s in directory %s, %v", tmpfile, tmpdir, err)
	}
	os.Chdir(tmpdir)

	// Custom buffer to check log output
	var buf bytes.Buffer
	logToBuffer(t, &buf)

	args := os.Args[0:1] // Name of the program
	args = append(args, "init")
	os.Args = args
	main()

	msg := "init .: directory is not empty"
	if !strings.Contains(buf.String(), msg) {
		t.Fatalf("Error: expected directory to be empty. Got %s, expected %s", buf.String(), msg)
	}
}
