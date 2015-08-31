package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
	"sync"
)

func PrettyUser(u *user.User) string {
	return fmt.Sprintf("USER: %s UID: %s GID:%s HOME: %s", u.Username, u.Uid, u.Gid, u.HomeDir)
}

var wg sync.WaitGroup

func main() {

	// NOTE: If you do not know the size of the slice
	// you need, set it to 0 and append takes care of it.
	users := make([]string, 0)
	fmt.Println("Reading /etc/passwd...")
	file, err := os.Open("/etc/passwd")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "#") {
			continue
		}
		// Get the username
		user := strings.Split(line, ":")[0]
		users = append(users, user)
	}
	for _, u := range users {
		wg.Add(1)
		go func(u string) {
			usr, err := user.Lookup(u)
			if usr == nil {
				fmt.Println("user:", u)
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(PrettyUser(usr))
			defer wg.Done()
		}(u)
	}
	wg.Wait()

}
