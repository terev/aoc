package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"strconv"
	"strings"
	"time"
)

var Root = &cobra.Command{
	Use: "aoc",
}

func init() {
	Root.AddCommand(run)
}

var run = &cobra.Command{
	Use: "run",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			year string
			day  string
		)

		if strings.HasPrefix(args[0], "day") {
			year = strconv.Itoa(time.Now().Year())
			day = strings.TrimPrefix(args[0], "day")
		} else {
			t, err := time.Parse("200602", args[0])
			if err != nil {
				return err
			}

			year = strconv.Itoa(t.Year())
			day = strconv.Itoa(t.Day())
		}

		extern := year

		path, err := filepath.Abs("./" + extern)
		if err != nil {
			return err
		}

		if _, err := os.Stat(path); err != nil {
			return err
		}

		pluginPath := filepath.Join(path, year+".so")

		buildCmd := exec.Command("go", "build", "-buildmode=plugin", "-o", pluginPath, path)
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr

		if err := buildCmd.Run(); err != nil {
			return err
		}

		yearSoln, err := plugin.Open(pluginPath)
		if err != nil {
			return err
		}

		soln, err := yearSoln.Lookup("Day" + day)
		if err != nil {
			return err
		}

		callable, ok := soln.(func() error)

		if !ok {
			return fmt.Errorf("Sorry solution function Day%s doesnt satisfy interface func()error", day)
		}

		return callable()
	},
}
