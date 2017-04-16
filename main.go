package main

import (
	fmt "github.com/jhunt/go-ansi"
	"net/http"
	"os"

	"github.com/jhunt/go-cli"
)

type Opt struct {
	Help bool `cli:"-h, --help"`
	Root string `cli:"-r, --root"`
	Port uint16 `cli:"-p, --port"`
}

func main() {
	var opt Opt
	opt.Root = "."
	opt.Port = 3001

	_, args, err := cli.Parse(&opt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "oops: @Y{%s}\n", err)
		os.Exit(1)
	}
	if len(args) != 0 {
		fmt.Fprintf(os.Stderr, "USAGE: @C{gow} [-r /path] [-p 3001]\n")
		os.Exit(1)
	}
	if opt.Help {
		fmt.Fprintf(os.Stderr, "USAGE: @C{gow} [-r /path] [-p 3001]\n\n")
		fmt.Fprintf(os.Stderr, "  -h, --help        Print this super-helpful message.\n")
		fmt.Fprintf(os.Stderr, "  -r, --root /path  Directory to use as the root (/) to serve.\n")
		fmt.Fprintf(os.Stderr, "  -p, --port 3001   TCP port to listen on.\n")
		os.Exit(0)
	}
	if opt.Port < 1 || opt.Port > 65535 {
		fmt.Fprintf(os.Stderr, "@R{invalid port %d (out of range)}\n", opt.Port)
		os.Exit(1)
	}
	if _, err := os.Stat(opt.Root); err != nil {
		fmt.Fprintf(os.Stderr, "@R{%s: %s}\n", opt.Root, err)
		os.Exit(1)
	}

	http.Handle("/", http.FileServer(http.Dir(opt.Root)))

	fmt.Printf("binding *:@G{%d} to serve @G{/} -> @G{%s}\n", opt.Port, opt.Root)
	http.ListenAndServe(fmt.Sprintf(":%d", opt.Port), nil)
}
