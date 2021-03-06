package users

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var DB map[string]User

func init() {
	DB = make(map[string]User)
	if !Exists("db.txt") {
		_, e := os.Create("db.txt")
		if e != nil {
			panic(e)
		}
	}
}

func GetUser(number string) (User, error) {
	if v, ok := DB[number]; ok {
		return v, nil
	}
	f, e := os.OpenFile("db.txt", os.O_RDONLY, os.ModeAppend)
	if e != nil {
		return User{}, e
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		spl := strings.Split(l, " ")
		nbr := spl[0]
		u := User{
			Number: number,
			Key:    spl[1],
		}
		if nbr == number {
			return u, nil
		}
	}
	return User{}, errors.New("This user could not be logged in")
}
