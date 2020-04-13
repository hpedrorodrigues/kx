package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdlogs "k8s.io/kubectl/pkg/cmd/logs"
	"k8s.io/kubectl/pkg/cmd/util"
	"kx/kx/common"
)

type LogsOptions struct {
	Context context.Context
	Command *cobra.Command

	genericclioptions.IOStreams

	util.Factory
}

func (o *LogsOptions) Validate() error {
	return nil
}

func (o *LogsOptions) Run() error {
	r, err := common.FindPod(o.Context)
	if err != nil {
		return err
	}
	args := []string{r.GetName()}
	options := cmdlogs.NewLogsOptions(o.IOStreams, false)
	options.Follow = true

	if err := options.Complete(o.Factory, o.Command, args); err != nil {
		return err
	}

	if err := options.Validate(); err != nil {
		return err
	}

	if err := options.RunLogs(); err != nil {
		return err
	}

	return nil
}

func NewLogsOptions(ctx context.Context, cmd *cobra.Command) *LogsOptions {
	o := &LogsOptions{
		Context:   ctx,
		Command:   cmd,
		IOStreams: common.GetIOStreams(),
		Factory:   util.NewFactory(common.GetConfigFlags()),
	}

	return o
}
