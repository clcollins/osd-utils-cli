package cluster

import (
	"fmt"

	k8spkg "github.com/openshift/osdctl/pkg/k8s"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

type mustGatherOptions struct {
	k8sclusterresourcefactory k8spkg.ClusterResourceFactoryOptions
	genericclioptions.IOStreams
}

// newCmdMustGather implements the mustGather command
// to retrieve a must-gather from the must-gather-operator on a cluster
func newCmdMustGather(streams genericclioptions.IOStreams, flags *genericclioptions.ConfigFlags) *cobra.Command {
	ops := newMustGatherOptions(streams, flags)
	mustGatherCmd := &cobra.Command{
		Use:               "must-gather",
		Short:             "retrieves a must-gather from the currently logged-in cluster",
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			// cmdutil.CheckErr(ops.complete(cmd, args))
			cmdutil.CheckErr(ops.run())
		},
	}
	ops.k8sclusterresourcefactory.AttachCobraCliFlags(mustGatherCmd)

	// mustGatherCmd.Flags().StringVarP(&ops.output, "out", "o", "default", "Output format [default | json | env]")
	// mustGatherCmd.Flags().BoolVarP(&ops.verbose, "verbose", "", false, "Verbose output")

	fmt.Printf("%+v", ops)

	return mustGatherCmd
}

func newMustGatherOptions(streams genericclioptions.IOStreams, flags *genericclioptions.ConfigFlags) *mustGatherOptions {
	return &mustGatherOptions{
		k8sclusterresourcefactory: k8spkg.ClusterResourceFactoryOptions{
			Flags: flags,
		},
		IOStreams: streams,
	}
}

func (o *mustGatherOptions) run() error {
	fmt.Printf("+%v", o)

	return nil
}
