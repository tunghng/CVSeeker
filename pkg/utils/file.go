package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// ReadJsonFile
func ReadJsonFile(file string) []byte {
	jsonFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	value, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	return value
}

func ReadFile(path string) ([]byte, bool, error) {
	if !FileExists(path) {
		return nil, false, nil
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, true, err
	}
	return content, true, nil
}

// Mkdir ...
func Mkdir(path string) error {
	if FileExists(path) {
		return nil
	}
	return os.Mkdir(path, 0700)
}

// FileExists ...
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func DeleteFile(path string) error {
	return os.Remove(path)
}

func GetFileContentType(out *os.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}

func GenLinkFileGCS(baseURL string, bucket string, dirToFile string) string {
	return fmt.Sprintf(`https://storage.cloud.google.com/%v/%v`, bucket, dirToFile)
}
