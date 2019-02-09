package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	tag        = flag.String("t", "", "Tag to look for. Will return whole file by default.")
	closingTag = flag.String("c", "", "Closing tag to match on. This only needs to be specified if it is not the companion tag to -t")
	file       = flag.String("f", "", "Which file to scan. Scans stdin if no file provided (enables piping to this utility)")
	out        = flag.String("o", "", "Output file. If left blank, will output to stdout")
	help       = flag.Bool("-help", false, "Show usage info")
	h          = flag.Bool("h", false, "Show usage info")
)

func main() {
	flag.Parse()

	if *h || *help {
		flag.Usage()
		os.Exit(0)
	}

	var in io.Reader
	if *file != "" {
		// Specified a file
		f, err := os.OpenFile(*file, os.O_RDONLY, 0665)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open specified filed [%s]: %v\n", *file, err)
			os.Exit(2)
		}
		defer f.Close()
		in = f
	} else {
		// Using stdin
		in = os.Stdin
	}

	var o io.Writer
	if *out != "" {
		f, err := os.OpenFile(*out, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0665)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening output file [%s]: %v\n", *out, err)
			os.Exit(3)
		}
		defer f.Close()
		o = f
	} else {
		o = os.Stdout
	}

	terminateTag := formEndTag(*tag)

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanWords)
	found := false
	for s.Scan() {
		word := s.Text()

		if found || *tag == "" {
			writeToOutput(word, o)
		}

		if *tag != "" {
			if matches, extra := matchStartingTag(word, *tag); matches {
				found = true
				writeToOutput(extra, o)
			} else if matches, extra := matchEndingTag(word, terminateTag); matches {
				found = false
				writeToOutput(extra, o)
			}
		}
	}
}

func writeToOutput(content string, out io.Writer) {
	if content == "" {
		return
	}
	_, err := out.Write([]byte(content))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write to output: %v\n", err)
		os.Exit(4)
	}
}

func formEndTag(start string) string {
	if start == "" {
		return ""
	}

	if strings.HasPrefix(start, "<") {
		// Specifically for XML
		return "</" + start[1:]
	}

	// This is reserved for non-xml formats later
	return ""
}

func matchStartingTag(content, tag string) (bool, string) {
	if !strings.Contains(content, tag) {
		return false, ""
	}

	if content == tag {
		return true, ""
	}

	return true, content[strings.Index(content, tag)+len(tag):]
}

func matchEndingTag(content, tag string) (bool, string) {
	if !strings.Contains(content, tag) {
		return false, ""
	}

	if content == tag {
		return true, ""
	}

	return true, content[:strings.Index(content, tag)]
}
