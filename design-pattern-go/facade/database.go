package facade

import (
	"bufio"
	"os"
	"strings"
)

func GetProperties(dbname string) map[string]string {
	filename := "./facade/" + dbname + ".txt"
	properties := map[string]string{}

	db, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	scanner := bufio.NewScanner(db)
	for scanner.Scan() {
		prop := strings.Split(scanner.Text(), "=")
		properties[prop[0]] = strings.TrimRight(prop[1], "\n")
	}

	return properties
}
