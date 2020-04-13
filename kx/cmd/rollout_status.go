package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdrollout "k8s.io/kubectl/pkg/cmd/rollout"
	"k8s.io/kubectl/pkg/cmd/util"
	"kx/kx/common"
)

type RolloutStatusOptions struct {
	Context context.Context
	Command *cobra.Command

	genericclioptions.IOStreams

	util.Factory
}

func (o *RolloutStatusOptions) Validate() error {
	return nil
}

func (o *RolloutStatusOptions) Run() error {
	r, err := common.FindRolloutResources(o.Context)
	if err != nil {
		return err
	}

	options := cmdrollout.NewRolloutStatusOptions(o.IOStreams)
	args := []string{r.GetKind(), r.GetName()}

	if err := options.Complete(o.Factory, args); err != nil {
		return err
	}

	if err := options.Validate(); err != nil {
		return err
	}

	if err := options.Run(); err != nil {
		return err
	}

	return nil
}

func NewRolloutStatusOptions(ctx context.Context, cmd *cobra.Command) *RolloutStatusOptions {
	o := &RolloutStatusOptions{
		Context:   ctx,
		Command:   cmd,
		IOStreams: common.GetIOStreams(),
		Factory:   util.NewFactory(common.GetConfigFlags()),
	}

	return o
}
