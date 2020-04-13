package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdexec "k8s.io/kubectl/pkg/cmd/exec"
	"k8s.io/kubectl/pkg/cmd/util"
	"kx/kx/common"
)

const defaultShell = "/bin/sh"

type ExecOptions struct {
	Context context.Context
	Command *cobra.Command

	genericclioptions.IOStreams

	util.Factory
}

func (o *ExecOptions) Validate() error {
	return nil
}

func (o *ExecOptions) Run() error {
	r, err := common.FindPod(o.Context)
	if err != nil {
		return err
	}

	options := &cmdexec.ExecOptions{
		StreamOptions: cmdexec.StreamOptions{
			Stdin:     true,
			TTY:       true,
			IOStreams: o.IOStreams,
		},
		Executor: &cmdexec.DefaultRemoteExecutor{},
	}

	args := []string{r.GetName(), defaultShell}

	argsLenAtDash := 1

	if err := options.Complete(o.Factory, o.Command, args, argsLenAtDash); err != nil {
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

func NewExecOptions(ctx context.Context, cmd *cobra.Command) *ExecOptions {
	o := &ExecOptions{
		Context:   ctx,
		Command:   cmd,
		IOStreams: common.GetIOStreams(),
		Factory:   util.NewFactory(util.NewMatchVersionFlags(common.GetConfigFlags())),
	}

	return o
}
