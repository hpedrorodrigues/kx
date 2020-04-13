package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/cmd/util"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/cmd/util/editor"
	"kx/kx/common"
)

type EditOptions struct {
	Context context.Context
	Command *cobra.Command

	genericclioptions.IOStreams

	util.Factory
}

func (o *EditOptions) Validate() error {
	return nil
}

func (o *EditOptions) Run() error {
	r, err := common.FindResource(o.Context)
	if err != nil {
		return err
	}

	options := editor.NewEditOptions(editor.NormalEditMode, o.IOStreams)
	options.ValidateOptions = cmdutil.ValidateOptions{EnableValidation: true}

	args := []string{r.GetKind(), r.GetName()}

	if err := options.Complete(o.Factory, args, o.Command); err != nil {
		return err
	}
	if err := options.Run(); err != nil {
		return err
	}

	return nil
}

func NewEditOptions(ctx context.Context, cmd *cobra.Command) *EditOptions {
	o := &EditOptions{
		Context:   ctx,
		Command:   cmd,
		IOStreams: common.GetIOStreams(),
		Factory:   util.NewFactory(common.GetConfigFlags()),
	}

	return o
}
