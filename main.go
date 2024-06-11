package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var explain_verbose = true
var debug_mode = true
var desc_delim = "END\n"
var dry_run = false
var base_branch = "main"

func main() {
	var switch_branch = false
	var branch string
	current_branch := Check("git", "branch", "--show-current")
	if current_branch == base_branch {
		// give the branch a name
		output("Name the branch: spaces will be replaced with -")
		branch = input("\n")
		branch = strings.ReplaceAll(strings.Trim(branch, " "), " ", "-")
		switch_branch = true
	}

	bs := Check("git", "status", "-s")
	status := formatStatus(bs)
	var commit_once = false
	var cmessage string
	if status != "" {
		br()
		output("Modified files:")
		output(status)
		br()
		output("input commit message")
		cmessage = input("\n")
		commit_once = true
	}

	br()
	output("Input the title of PR")
	verbose("(leave empty to autofill from git commits)")
	title := input("\n")
	var autofill = false
	if title == "" {
		autofill = true
	}
	br()
	output("Add a description << END")
	desc := input(desc_delim)

	br()
	output("Skip check on browser?")
	verbose("input starting with [y] for Yes, anything else for No")
	var skip_browser_check = false
	if strings.HasPrefix(input("\n"), "y") {
		skip_browser_check = true
	}

	args := []string{"pr", "create", "--base", base_branch, "--head", branch, "--title", "\"" + title + "\"", "--body", "\"" + desc + "\""}
	if desc == "" {
		args = append(args, "\"\"")
	}
	if autofill {
		args = append(args, "--fill")
	}
	if !skip_browser_check {
		args = append(args, "--web")
	}

	// execute stuff
	if switch_branch {
		Run("git", "switch", "-c", branch)
	}
	if commit_once {
		Run("git", "add", "-A")
		Run("git", "commit", "-m", cmessage)
	}
	Run("git", "push", "--set-upstream", "origin", "HEAD")
	Run("gh", args...)
	Run("git", "switch", "-c", base_branch)
}

func verbose(s string) {
	if explain_verbose {
		fmt.Println(s)
	}
}

func todo(a ...any) {
	fmt.Print("TODO: use ")
	fmt.Print(a...)
}

func output(a ...any) {
	fmt.Println(a...)
}
func br() {
	fmt.Println()
}
func Run(command string, args ...string) {
	if dry_run {
		fmt.Println("executing", command, strings.Join(args, " "))
		return
	}
	err := exec.Command(command, args...).Run()
	if err != nil {
		// apparentally some gh commands fail on success
		debug("Command", command, strings.Join(args, " "), "Failed.")
		debug(err)
	}
}

// Check() also runs on dry_run. do NOT put commands that mutate the environment.
func Check(command string, args ...string) string {
	b, err := exec.Command(command, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Trim(string(b), "\n")
}

// prints prefix, and waits for the user to input.
// end_sig must end with \n or it will not work.
// - reason: it uses bufio.ReadLine() internally
// if user input \n on the first input, stops reading.
func input(delim string) string {
	r := bufio.NewReader(os.Stdin)
	total := ""
	is_first := true
	for {
		fmt.Print("> ")
		line, err := r.ReadString(byte('\n'))
		assertNil(err)
		if strings.Trim(line, " \n") == "q" {
			output("Detected 'q'. Exitting..")
			os.Exit(1)
		}
		total += line
		if strings.HasSuffix(total, delim) {
			break
		}
		if is_first && total == "\n" {
			return ""
		}
		is_first = false
	}
	return strings.TrimSuffix(total, delim)
}
func debug(a ...any) {
	if debug_mode {
		fmt.Println(a...)
	}
}
func assertNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func formatStatus(s string) string {
	return strings.Trim(s, "\n")
}
func oneLine(s string) string {
	return strings.ReplaceAll(s, "\n", " ")
}
func removeFloatingM(s string) string {
	return strings.ReplaceAll(s, " M ", " ")
}
func If[T any](b bool, onTrue, onFalse T) T {
	if b {
		return onTrue
	} else {
		return onFalse
	}
}
