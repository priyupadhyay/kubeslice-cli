package internal

const (
	ProjectObject             = "projects.controller.kubeslice.io"
	ClusterObject             = "clusters.controller.kubeslice.io"
	SliceConfigObject         = "sliceconfigs.controller.kubeslice.io"
	ServiceExportConfigObject = "serviceexportconfigs.controller.kubeslice.io"

	// skipStep suffix is for the set of installation steps to skip.
	Kind_skipStep                = "kind"
	Calico_skipStep              = "calico"
	Controller_skipStep          = "controller"
	Worker_registration_skipStep = "worker-registration"
	UI_install_skipStep          = "enterprise"
	Worker_skipStep              = "worker"
	Demo_skipStep                = "demo"
	SecretObject                 = "secrets"
	OutputFormatYaml             = "yaml"
	OutputFormatJson             = "json"
)