package kx

import (
	"context"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/cmd/util"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/cmd/util/editor"
	"kx/kx/common"
)

func Edit(ctx context.Context, cmd *cobra.Command) error {
	r, err := common.FindResource(ctx)
	if err != nil {
		return err
	}

	o := editor.NewEditOptions(editor.NormalEditMode, common.GetIOStreams())
	o.ValidateOptions = cmdutil.ValidateOptions{EnableValidation: true}

	f := util.NewFactory(common.GetConfigFlags())

	args := []string{r.GetKind(), r.GetName()}

	if err := o.Complete(f, args, cmd); err != nil {
		return err
	}
	if err := o.Run(); err != nil {
		return err
	}

	return nil
}
