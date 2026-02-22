package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	PORT = ":9000"
	codeTemplate = `
package main

import (
	"fmt"
	"unsafe"
)

type secret struct {
	flag  string
	dummy int
}

func (s *secret) String() string {
	return "Try again"
}

func main() {
	mySecret := &secret{flag: "%s"}

	var target interface{} = mySecret

	_ = target
	%s
}
`
)

var blacklist = []string{
	"os", "syscall", "ioutil", "exec", "cmd", "system", "log", "reflect", `THJCC{iT'%d%s%d%s%dSm%dY_gO%s!!!!!L%sAnG}`, 
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listening on %s\n", PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	_ = conn.SetDeadline(time.Now().Add(360 * time.Second))

	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	fmt.Fprint(writer, "=== Happy Cat Jail! ===\n")
	fmt.Fprint(writer, "This is a Go jail Challenge\n")
	fmt.Fprint(writer, "Pls type \"EOF\" at the Final line and Press Enter to Execute your Code\n")
	fmt.Fprint(writer, "=======================\n")
	writer.Flush()

	var userCode strings.Builder
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return
			}
			break
		}
		if strings.TrimSpace(line) == "EOF" {
			break
		}
		userCode.WriteString(line)
	}

	code := userCode.String()

	for _, bad := range blacklist {
		if strings.Contains(code, bad) {
			fmt.Fprint(writer, "Try syscall or sth else (?\n")
			writer.Flush()
			return
		}
	}

	lower := "abcdefghijklmnopqrstuvwxyz"
	mixed := lower + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	realFlag := fmt.Sprintf("THJCC{iT'%d%s%d%s%dSm%dY_gO%s!!!!!L%sAnG}",
		rand.Intn(9)+1,
		string(lower[rand.Intn(26)]),
		rand.Intn(9)+1,
		string(mixed[rand.Intn(52)]),
		rand.Intn(9)+1,
		rand.Intn(9)+1,
		string(mixed[rand.Intn(52)]),
		string(lower[rand.Intn(26)]),
	)

	fullSource := fmt.Sprintf(codeTemplate, realFlag, code)

	tmpFile, err := os.CreateTemp("", "jail_*.go")
	if err != nil {
		fmt.Fprint(writer, "error\n")
		writer.Flush()
		return
	}
	defer os.Remove(tmpFile.Name())
	_, _ = tmpFile.Write([]byte(fullSource))
	tmpFile.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 360*time.Second)
	defer cancel()

	blacklistHit := false
	for _, bad := range blacklist {
		if strings.Contains(code, bad) {
			blacklistHit = true
			break
		}
	}

	if blacklistHit {
		fmt.Fprintln(writer, "=======================")
		fmt.Fprintln(writer, "Try syscall or sth else (?")
		fmt.Fprintln(writer, "Connection closed.")
		writer.Flush()
		conn.Close()
		return
	}


	cmd := exec.CommandContext(ctx, "go", "run", tmpFile.Name())
	output, _ := cmd.CombinedOutput()
	outputStr := string(output)

	if strings.Contains(outputStr, "# command-line-arguments") ||
		strings.Contains(outputStr, "invalid use of unsafe") ||
		strings.Contains(outputStr, "cannot convert") ||
		strings.Contains(outputStr, "illegal") ||
		strings.Contains(outputStr, "invalid operation") {
		fmt.Fprintln(writer, "=======================")
		fmt.Fprintln(writer, "Try syscall or sth else (?")
	} else {
		fmt.Fprintln(writer, "=======================")
		fmt.Fprintln(writer, "Try syscall or sth else (?")
		fmt.Fprintln(writer, "=======================")
		fmt.Fprintln(writer, outputStr)
	}

	writer.Flush()


}
