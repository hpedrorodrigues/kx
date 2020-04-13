package cmd

type CommandOptions interface {
	Validate() error
	Run() error
}

func Run(o CommandOptions) error {
	if err := o.Validate(); err != nil {
		return err
	}

	if err := o.Run(); err != nil {
		return err
	}

	return nil
}
