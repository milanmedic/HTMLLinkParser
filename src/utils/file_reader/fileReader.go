package filereader

import "io/ioutil"

type FileReader struct{}

func CreateFileReader() *FileReader {
	return &FileReader{}
}

func (fr *FileReader) ReadFile(filename string) ([]byte, error) {
	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return contents, err
}

func (fr *FileReader) ReadFileAsString(filename string) (string, error) {
	contents, err := fr.ReadFile(filename)

	if err != nil {
		return "", err
	}

	return string(contents), nil
}
