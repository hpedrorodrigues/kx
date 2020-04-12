package kx

import (
	"context"
	"github.com/spf13/cobra"
	cmdexec "k8s.io/kubectl/pkg/cmd/exec"
	"k8s.io/kubectl/pkg/cmd/util"
	"kx/kx/common"
)

func Exec(ctx context.Context, cmd *cobra.Command) error {
	r, err := common.FindPod(ctx)
	if err != nil {
		return err
	}

	o := &cmdexec.ExecOptions{
		StreamOptions: cmdexec.StreamOptions{
			Stdin:     true,
			TTY:       true,
			IOStreams: common.GetIOStreams(),
		},
		Executor: &cmdexec.DefaultRemoteExecutor{},
	}

	cf := common.GetConfigFlags()
	f := util.NewFactory(util.NewMatchVersionFlags(cf))

	args := []string{r.GetName(), "sh"}

	argsLenAtDash := 1

	if err := o.Complete(f, cmd, args, argsLenAtDash); err != nil {
		return err
	}
	if err := o.Validate(); err != nil {
		return err
	}
	if err := o.Run(); err != nil {
		return err
	}

	return nil
}
