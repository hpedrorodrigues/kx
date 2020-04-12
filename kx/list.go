package kx

import (
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/cli-runtime/pkg/printers"
	cmdget "k8s.io/kubectl/pkg/cmd/get"
	"k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/scheme"
	"kx/kx/common"
)

func List(cmd *cobra.Command) error {
	o := &cmdget.GetOptions{
		PrintFlags:             cmdget.NewGetPrintFlags(),
		IOStreams:              common.GetIOStreams(),
		Namespace:              common.GetNamespace(),
		ChunkSize:              500,
		ServerPrint:            true,
		IgnoreNotFound:         true,
		IsHumanReadablePrinter: true,
		ExplicitNamespace:      true,
	}

	o.ToPrinter = func(mapping *meta.RESTMapping, outputObjects *bool, withNamespace bool, withKind bool) (printers.ResourcePrinterFunc, error) {
		printFlags := o.PrintFlags.Copy()

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

	f := util.NewFactory(common.GetConfigFlags())

	args := []string{common.GetResourceNames()}

	if err := o.Run(f, cmd, args); err != nil {
		return err
	}

	return nil
}
