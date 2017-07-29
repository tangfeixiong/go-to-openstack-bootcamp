/*
   Inspired from
     https://github.com/openshift/origin/blob/master/pkg/cmd/openshift/openshift.go
*/
package cmd

import (
	// "os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"

	oldcmd "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/cmd"
	kopincmd "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/kopit/cmd/kopin"
)

func CommandFor(basename string) *cobra.Command {
	var cmd *cobra.Command

	// in, out, errout := os.Stdin, os.Stdout, os.Stderr

	// Make case-insensitive and strip executable suffix if present
	if runtime.GOOS == "windows" {
		basename = strings.ToLower(basename)
		basename = strings.TrimSuffix(basename, ".exe")
	}

	switch basename {
	case "kopos":
		return oldcmd.RootCmd
	case "kopin":
		fallthrough
	default:
		cmd = kopincmd.NewRootCommand(basename)
	}

	//	if cmd.UsageFunc() == nil {
	//		templates.ActsAsRootCommand(cmd, []string{"options"})
	//	}
	/*
	   https://stackoverflow.com/questions/34053881/golang-how-can-i-use-pflag-with-other-packages-that-use-flag
	*/
	//  flagtypes.GLog(cmd.PersistentFlags())

	return cmd
}
