package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdrollout "k8s.io/kubectl/pkg/cmd/rollout"
	"k8s.io/kubectl/pkg/cmd/util"
	"kx/kx/common"
)

type RolloutHistoryOptions struct {
	Context context.Context
	Command *cobra.Command

	genericclioptions.IOStreams

	util.Factory
}

func (o *RolloutHistoryOptions) Validate() error {
	return nil
}

func (o *RolloutHistoryOptions) Run() error {
	r, err := common.FindRolloutResources(o.Context)
	if err != nil {
		return err
	}

	options := cmdrollout.NewRolloutHistoryOptions(o.IOStreams)
	args := []string{r.GetKind(), r.GetName()}
	if err := options.Complete(o.Factory, o.Command, args); err != nil {
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

func NewRolloutHistoryOptions(ctx context.Context, cmd *cobra.Command) *RolloutHistoryOptions {
	o := &RolloutHistoryOptions{
		Context:   ctx,
		Command:   cmd,
		IOStreams: common.GetIOStreams(),
		Factory:   util.NewFactory(common.GetConfigFlags()),
	}

	return o
}
