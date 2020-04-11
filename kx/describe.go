package kx

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/resource"
	cmddescribe "k8s.io/kubectl/pkg/cmd/describe"
	"k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/describe"
	"os"
)

func Describe(parent string, cmd *cobra.Command) error {
	streams := genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
	o := &cmddescribe.DescribeOptions{
		FilenameOptions: &resource.FilenameOptions{},
		DescriberSettings: &describe.DescriberSettings{
			ShowEvents: true,
		},
		CmdParent: parent,
		IOStreams: streams,
	}

	cf := genericclioptions.NewConfigFlags(true)
	f := util.NewFactory(cf)

	args := []string{"pod", "fluentd-2zjxq"}

	if err := o.Complete(f, cmd, args); err != nil {
		return err
	}

	if err := o.Run(); err != nil {
		return err
	}

	return nil
}
