package init

import (
	"os"
	"path/filepath"
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
			t.Fatalf("Error: could not delete directory %s, %v", d, err)
		}
	})
}

func TestDirEmptyDoesNotExist(t *testing.T) {
	tmpdir := filepath.Join(os.TempDir(), "tskr_testdir")

	res, err := IsDirEmpty(tmpdir)
	if res || err == nil {
		t.Fatalf("Error: tmpdir was supposed to not exist. Got %t, %v, want false, open %s: no such file or directory", res, err, tmpdir)
	}
}

func TestDirEmptyIsEmpty(t *testing.T) {
	tmpdir := filepath.Join(os.TempDir(), "tskr_testdir")
	createDir(t, tmpdir)

	res, err := IsDirEmpty(tmpdir)
	if !res || err != nil {
		t.Fatalf("Error: tmpdir was supposed to be empty. Got %t, %v, want true, nil", res, err)
	}
}

func TestDirEmptyIsNotEmpty(t *testing.T) {
	tmpdir := filepath.Join(os.TempDir(), "tskr_testdir")
	tmpfile := "tmptestfile"
	createDir(t, tmpdir)

	_, err := os.Create(filepath.Join(tmpdir, tmpfile))
	if err != nil {
		t.Fatalf("Error: could not create file %s in directory %s, %v", tmpfile, tmpdir, err)
	}

	res, err := IsDirEmpty(tmpdir)
	if res || err != nil {
		t.Fatalf("Error: tmpdir was supposed to be empty. Got %t, %v, want true, nil", res, err)
	}
}

func TestCreateDirsSuccess(t *testing.T) {
	tmpdir := filepath.Join(os.TempDir(), "tskr_testdir")
	createDir(t, tmpdir)

	err := CreateDirs(tmpdir)
	if err != nil {
		t.Fatalf("Error: could not create dirs in directory %s. Got %v, expected nil", tmpdir, err)
	}
}

func TestCreateDirsFailure(t *testing.T) {
	tmpdir := filepath.Join(os.TempDir(), "tskr_testdir")
	tmpsubdir := "inbox"
	createDir(t, tmpdir)

	err := os.Mkdir(filepath.Join(tmpdir, tmpsubdir), 0755)
	if err != nil {
		t.Fatalf("Error: could not create directory %s, %v", filepath.Join(tmpdir, tmpsubdir), err)
	}

	err = CreateDirs(tmpdir)
	if err == nil {
		t.Fatalf("Error: didn't fail to create dirs in directory %s. Got %v, expected mkdir %s: file exists", tmpdir, err, filepath.Join(tmpdir, tmpsubdir))
	}
}

func TestInitSuccess(t *testing.T) {
	tmpdir := filepath.Join(os.TempDir(), "tskr_testdir")
	createDir(t, tmpdir)

	err := Init(tmpdir)
	if err != nil {
		t.Fatalf("Error: expected to succeed to init directory %s. Got %v, expected nil", tmpdir, err)
	}
}

func TestInitNoDirectory(t *testing.T) {
	tmpdir := filepath.Join(os.TempDir(), "tskr_testdir")

	err := Init(tmpdir)
	if err == nil {
		t.Fatalf("Error: expected to fail to init directory %s. Got %v, expected mkdir %s: no such file or directory", tmpdir, err, tmpdir)
	}
}

func TestInitDirectoryNotEmpty(t *testing.T) {
	tmpdir := filepath.Join(os.TempDir(), "tskr_testdir")
	tmpfile := "tmptestfile"
	createDir(t, tmpdir)

	_, err := os.Create(filepath.Join(tmpdir, tmpfile))
	if err != nil {
		t.Fatalf("Error: could not create file %s in directory %s, %v", tmpfile, tmpdir, err)
	}

	err = Init(tmpdir)
	if err == nil {
		t.Fatalf("Error: expected to not error in init directory %s. Got %v, expected init %s: directory is not empty", tmpdir, err, tmpdir)
	}
}
