package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	const day = 8
	copies := [][2]string{
		{"dayxx/dayxx.go", fmt.Sprintf("day%02d/day%02d.go", day, day)},
		{"dayxx/dayxx_test.go", fmt.Sprintf("day%02d/day%02d_test.go", day, day)},
		{"inputs/dayxx.txt", fmt.Sprintf("inputs/day%02d.txt", day)},
		{"inputs/dayxx.txt", fmt.Sprintf("inputs/day%02d_example1.txt", day)},
	}

	for _, cp := range copies {
		src, dst := cp[0], cp[1]

		srcStat, err := os.Stat(src)
		if err != nil {
			panic(err)
		}

		if _, err = os.Stat(dst); errors.Is(err, os.ErrNotExist) {
			if err = copyFile(src, dst, srcStat); err != nil {
				panic(err)
			}
		}
	}
}

func copyFile(src, dst string, stat os.FileInfo) error {
	if _, err := os.Stat(dst); !errors.Is(err, os.ErrNotExist) {
		return nil
	}

	if err := createDirIfNotExists(filepath.Dir(dst)); err != nil {
		return err
	}

	srcF, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcF.Close()

	dstF, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, stat.Mode().Perm())
	if err != nil {
		return err
	}
	defer dstF.Close()

	if _, err = io.Copy(dstF, srcF); err != nil {
		return err
	}

	return nil
}

func createDirIfNotExists(dir string) error {
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		return os.Mkdir(dir, 0755)
	}

	return nil
}
