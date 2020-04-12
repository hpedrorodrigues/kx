package kx

import (
	"context"
	"github.com/spf13/cobra"
	cmdlogs "k8s.io/kubectl/pkg/cmd/logs"
	"k8s.io/kubectl/pkg/cmd/util"
	"kx/kx/common"
)

func Logs(ctx context.Context, cmd *cobra.Command) error {
	r, err := common.FindPod(ctx)
	if err != nil {
		return err
	}
	args := []string{r.GetName()}
	o := cmdlogs.NewLogsOptions(common.GetIOStreams(), false)
	o.Follow = true

	f := util.NewFactory(common.GetConfigFlags())
	if err := o.Complete(f, cmd, args); err != nil {
		return err
	}
	if err := o.Validate(); err != nil {
		return err
	}
	if err := o.RunLogs(); err != nil {
		return err
	}
	return nil
}
