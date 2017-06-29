package koposcmd

import (
	"github.com/spf13/cobra"
)

var SapccKoposCmd = &cobra.Command{
	Use:   "openstack",
	Short: "Kubernetes operator Openstack",
	Run: func(cmd *cobra.Command, args []string) {
		main()
	},
}
