package storage

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type file struct {
	Filename string
}

func NewFile(filename string) *file {
	if filename == "" {
		rand.Seed(time.Now().UTC().UnixNano())
		filename = "file" + strconv.Itoa(rand.Int()) + ".txt"
	}

	return &file{
		Filename: filename,
	}
}

func (f *file) Write(data []string) error {
	fd, err := os.Create(f.Filename)
	if err != nil {
		return err
	}
	defer func() {
		err = fd.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	for _, v := range data {
		_, err = fd.Write([]byte(v + "\r\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *file) Read() ([]byte, error) {
	return []byte{}, nil
}
