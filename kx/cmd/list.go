package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/client-go/deprecated/scheme"
	cmdget "k8s.io/kubectl/pkg/cmd/get"
	"k8s.io/kubectl/pkg/cmd/util"
	"kx/kx/common"
)

type ListOptions struct {
	Context       context.Context
	Command       *cobra.Command
	ChunkSize     int64
	ResourceNames string

	genericclioptions.IOStreams

	util.Factory
}

func (o *ListOptions) Validate() error {
	return nil
}

func (o *ListOptions) Run() error {
	options := &cmdget.GetOptions{
		PrintFlags:             cmdget.NewGetPrintFlags(),
		IOStreams:              o.IOStreams,
		Namespace:              common.GetNamespace(),
		ChunkSize:              o.ChunkSize,
		ServerPrint:            true,
		IgnoreNotFound:         true,
		IsHumanReadablePrinter: true,
		ExplicitNamespace:      true,
	}

	options.ToPrinter = func(mapping *meta.RESTMapping, outputObjects *bool, withNamespace bool, withKind bool) (printers.ResourcePrinterFunc, error) {
		printFlags := options.PrintFlags.Copy()

		if mapping != nil {
			printFlags.SetKind(mapping.GroupVersionKind.GroupKind())
		}

		if withNamespace {
			if err := printFlags.EnsureWithNamespace(); err != nil {
				return nil, err
			}
		}

		if withKind {
			if err := printFlags.EnsureWithKind(); err != nil {
				return nil, err
			}
		}

		printer, err := printFlags.ToPrinter()
		if err != nil {
			return nil, err
		}

		printer, err = printers.NewTypeSetter(scheme.Scheme).WrapToPrinter(printer, nil)
		if err != nil {
			return nil, err
		}

		printer = &cmdget.TablePrinter{Delegate: printer}

		return printer.PrintObj, nil
	}
	args := []string{o.ResourceNames}

	if err := options.Run(o.Factory, o.Command, args); err != nil {
		return err
	}

	return nil
}

func NewListOptions(ctx context.Context, cmd *cobra.Command) *ListOptions {
	o := &ListOptions{
		Context:       ctx,
		Command:       cmd,
		IOStreams:     common.GetIOStreams(),
		Factory:       util.NewFactory(common.GetConfigFlags()),
		ChunkSize:     500,
		ResourceNames: common.GetResourceNames(),
	}

	return o
}
