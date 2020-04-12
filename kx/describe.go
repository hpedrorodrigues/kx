package kx

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/resource"
	cmddescribe "k8s.io/kubectl/pkg/cmd/describe"
	"k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/describe"
	"kx/kx/common"
	"kx/kx/fuzzyfinder"
	"strings"
)

func Describe(ctx context.Context, cmd *cobra.Command) error {
	resources, err := common.ListResources(ctx)
	if err != nil {
		return err
	}

	idx, err := fuzzyfinder.Find(resources, func(i int) string {
		r := resources[i]
		kind, name := strings.ToLower(r.GetKind()), r.GetName()
		return fmt.Sprintf("%s/%s", kind, name)
	})

	if err != nil {
		return err
	}

	r := resources[idx]

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
