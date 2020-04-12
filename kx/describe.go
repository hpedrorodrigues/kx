package kx

import (
	"context"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/resource"
	cmddescribe "k8s.io/kubectl/pkg/cmd/describe"
	"k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/describe"
	"kx/kx/common"
)

func Describe(ctx context.Context, cmd *cobra.Command) error {
	r, err := common.FindResource(ctx)
	if err != nil {
		return err
	}

	o := &cmddescribe.DescribeOptions{
		FilenameOptions: &resource.FilenameOptions{},
		DescriberSettings: &describe.DescriberSettings{
			ShowEvents: true,
		},
		IOStreams: common.GetIOStreams(),
	}
	f := util.NewFactory(common.GetConfigFlags())

	args := []string{r.GetKind(), r.GetName()}

	if err := o.Complete(f, cmd, args); err != nil {
		return err
	}

	if err := o.Run(); err != nil {
		return err
	}

	return nil
}
