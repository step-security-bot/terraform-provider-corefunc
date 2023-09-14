package cmd

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"text/tabwriter"

	"github.com/gookit/color"
	"github.com/northwood-labs/golang-utils/archstring"
	"github.com/northwood-labs/golang-utils/exiterrorf"
	"github.com/spf13/cobra"
)

var (
	// Color text.
	colorHeader = color.New(color.FgWhite, color.BgBlue, color.OpBold)

	// Version represents the version of the software.
	Version = "dev"

	// Commit represents the git commit hash of the software.
	Commit = func() string {
		if info, ok := debug.ReadBuildInfo(); ok {
			for _, setting := range info.Settings {
				if setting.Key == "vcs.revision" {
					return setting.Value
				}
			}
		}

		return "unknown"
	}()

	// BuildDate represents the date the software was built.
	BuildDate = func() string {
		if info, ok := debug.ReadBuildInfo(); ok {
			for _, setting := range info.Settings {
				if setting.Key == "vcs.time" {
					return setting.Value
				}
			}
		}

		return "unknown"
	}()

	// Dirty represents whether or not the git repo was dirty when the software was built.
	Dirty = func() string {
		if info, ok := debug.ReadBuildInfo(); ok {
			for _, setting := range info.Settings {
				if setting.Key == "vcs.modified" {
					return setting.Value
				}
			}
		}

		return "unknown"
	}()

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Long-form version information",
		Long: `Long-form version information, including the build commit hash, build date, Go
version, and external dependencies.`,
		Run: func(cmd *cobra.Command, args []string) {
			colorHeader.Println(" BASIC ")

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

			fmt.Fprintf(w, " Version:\t%s\t\n", Version)
			fmt.Fprintf(w, " Go version:\t%s\t\n", runtime.Version())
			fmt.Fprintf(w, " Git commit:\t%s\t\n", Commit)
			if Dirty == "true" {
				fmt.Fprintf(w, " Dirty repo:\t%s\t\n", Dirty)
			}
			fmt.Fprintf(w, " Build date:\t%s\t\n", BuildDate)
			fmt.Fprintf(w, " OS/Arch:\t%s/%s\t\n", runtime.GOOS, runtime.GOARCH)
			fmt.Fprintf(w, " System:\t%s\t\n", archstring.GetFriendlyName(runtime.GOOS, runtime.GOARCH))
			fmt.Fprintf(w, " CPU Cores:\t%d\t\n", runtime.NumCPU())

			err := w.Flush()
			if err != nil {
				exiterrorf.ExitErrorf(err)
			}

			fmt.Println("")

			//----------------------------------------------------------------------

			if buildInfo, ok := debug.ReadBuildInfo(); ok {
				colorHeader.Println(" DEPENDENCIES ")

				w = tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

				for i := range buildInfo.Deps {
					dependency := buildInfo.Deps[i]
					fmt.Fprintf(w, " %s\t%s\t\n", dependency.Path, dependency.Version)
				}
			}

			err = w.Flush()
			if err != nil {
				exiterrorf.ExitErrorf(err)
			}

			fmt.Println("")
		},
	}
)

func init() { // lint:allow_init
	rootCmd.AddCommand(versionCmd)
}
