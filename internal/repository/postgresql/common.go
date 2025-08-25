package postgresql

import (
	"os"
	path2 "path"
)

func getRequest(fileName string) (string, error) {

	path := path2.Join("./internal/sqlrequests/", fileName)
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(file), nil
}
