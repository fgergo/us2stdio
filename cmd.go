// Command us2stdio dials server serving on path and redirects it's io to stdio.
//
// To mount a server serving on my9pserver.socket in acme-sac:
// on windows:
// go get github.com/fgergo/us2stdio
// in acme-sac:
// Local mount { {os us2stdio.exe my9pserver.socket}>[1=0]} /n/my9pserver
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%v: %v path", os.Args[0], os.Args[0])
		os.Exit(1)
	}
	conn, err := net.Dial("unix", os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: error opening %v: %v", os.Args[0], os.Args[1], err)
		os.Exit(1)
	}
	go io.Copy(os.Stdout, conn)
	io.Copy(conn, os.Stdin)

	conn.Close()
}
