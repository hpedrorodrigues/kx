package common

func GetNamespace() string {
	if ns := *cf.Namespace; ns != "" {
		return ns
	}

	cc := cf.ToRawKubeConfigLoader()

	if ns, _, err := cc.Namespace(); err != nil {
		return "default"
	} else {
		return ns
	}
}
