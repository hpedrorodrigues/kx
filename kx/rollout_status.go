package kx

import (
	"context"
	cmdrollout "k8s.io/kubectl/pkg/cmd/rollout"
	"k8s.io/kubectl/pkg/cmd/util"
	"kx/kx/common"
)

func RolloutStatus(ctx context.Context) error {
	r, err := common.FindDDS(ctx)
	if err != nil {
		return err
	}

	o := cmdrollout.NewRolloutStatusOptions(common.GetIOStreams())
	f := util.NewFactory(common.GetConfigFlags())
	args := []string{r.GetKind(), r.GetName()}
	if err := o.Complete(f, args); err != nil {
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
