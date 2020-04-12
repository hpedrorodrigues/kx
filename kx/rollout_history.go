package kx

import (
	"context"
	"github.com/spf13/cobra"
	cmdrollout "k8s.io/kubectl/pkg/cmd/rollout"
	"k8s.io/kubectl/pkg/cmd/util"
	"kx/kx/common"
)

func RolloutHistory(ctx context.Context, cmd *cobra.Command) error {
	r, err := common.FindDDS(ctx)
	if err != nil {
		return err
	}

	o := cmdrollout.NewRolloutHistoryOptions(common.GetIOStreams())
	f := util.NewFactory(common.GetConfigFlags())
	args := []string{r.GetKind(), r.GetName()}
	if err := o.Complete(f, cmd, args); err != nil {
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
