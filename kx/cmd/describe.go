package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/resource"
	cmddescribe "k8s.io/kubectl/pkg/cmd/describe"
	"k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/describe"
	"kx/kx/common"
)

type DescribeOptions struct {
	Context context.Context
	Command *cobra.Command

	genericclioptions.IOStreams

	util.Factory
}

func (o *DescribeOptions) Validate() error {
	return nil
}

func (o *DescribeOptions) Run() error {
	r, err := common.FindResource(o.Context)
	if err != nil {
		return err
	}

	options := &cmddescribe.DescribeOptions{
		FilenameOptions: &resource.FilenameOptions{},
		DescriberSettings: &describe.DescriberSettings{
			ShowEvents: true,
		},
		IOStreams: o.IOStreams,
	}

	args := []string{r.GetKind(), r.GetName()}

	if err := options.Complete(o.Factory, o.Command, args); err != nil {
		return err
	}

	if err := options.Run(); err != nil {
		return err
	}

	return nil
}

func NewDescribeOptions(ctx context.Context, cmd *cobra.Command) *DescribeOptions {
	o := &DescribeOptions{
		Context:   ctx,
		Command:   cmd,
		IOStreams: common.GetIOStreams(),
		Factory:   util.NewFactory(common.GetConfigFlags()),
	}

	return o
}
