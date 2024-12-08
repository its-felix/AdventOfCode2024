package inputs

import (
	"bufio"
	"embed"
	"io"
	"io/fs"
)

//go:embed *.txt
var inputs embed.FS

func GetInput(name string) fs.File {
	f, err := inputs.Open(name)
	if err != nil {
		panic(err)
	}

	return f
}

func GetInputLines(name string) <-chan string {
	ch := make(chan string)
	go func() {
		f := GetInput(name)
		defer f.Close()
		defer close(ch)

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}()

	return ch
}

func GetInputText(name string) string {
	f := GetInput(name)
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return string(b)
}
