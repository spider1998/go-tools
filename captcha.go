package tools

import (
	"bytes"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/pkg/errors"
	"image/color"
	"image/jpeg"
	"math/rand"
	"os"
)

var Captcha = &CaptchaService{}

type CaptchaService struct{}

func file2Bytes(filename string) ([]byte, error) {

	// File
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// FileInfo:
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// []byte
	data := make([]byte, stats.Size())
	count, err := file.Read(data)
	if err != nil {
		return nil, err
	}
	fmt.Printf("read file %s len: %d \n", filename, count)
	return data, nil
}

func (s *CaptchaService) Generate() (token string, image []byte, err error) {
	token = randString(32)
	c := captcha.New()
	data, err := file2Bytes("comic.ttf")
	if err != nil {
		return
	}
	err = c.AddFontFromBytes(data)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	c.SetBkgColor(color.RGBA{0xc8, 0xe1, 0xff, 1})
	value := RandString(4)
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, c.CreateCustom(value), nil)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return token, buf.Bytes(), nil
}

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

const letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
