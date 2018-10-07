package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

type warning struct {
	message   string
	positions []string
}

func run() error {
	cmd := exec.Command("sbt", "-no-colors", "test:compile")
	r, w := io.Pipe()
	cmd.Stdout, cmd.Stderr = w, w
	tee := io.TeeReader(r, os.Stdout)
	s := bufio.NewScanner(tee)
	if err := cmd.Start(); err != nil {
		return err
	}

	var ws []*warning
	go func() {
	L:
		for s.Scan() {
			line := s.Text()
			if strings.HasPrefix(line, "[warn] ") {
				xs := strings.SplitN(line[7:], ":", 4)
				if len(xs) != 4 {
					continue
				}
				if _, err := strconv.Atoi(xs[1]); err != nil {
					continue
				}
				pos, mess := xs[0]+":"+xs[1]+":"+xs[2], strings.TrimSpace(xs[3])
				for _, w := range ws {
					if w.message == mess {
						w.positions = append(w.positions, pos)
						continue L
					}
				}
				ws = append(ws, &warning{mess, []string{pos}})
			}
		}
	}()
	if err := cmd.Wait(); err != nil {
		return err
	}

	sort.Slice(ws, func(i, j int) bool {
		return len(ws[i].positions) > len(ws[j].positions)
	})
	var total int
	for _, w := range ws {
		fmt.Printf("%d: %s\n", len(w.positions), w.message)
		total += len(w.positions)
		for _, pos := range w.positions {
			fmt.Printf("  %s\n", pos)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Total: %d warnings\n", total)
	return nil
}
