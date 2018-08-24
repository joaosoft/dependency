package dependency

type CmdDependency string
type Imports map[string]*Import

type Memory struct {
	lockedImports    Imports
	internalImports  Imports
	externalImports  Imports
	loadedImports    map[string]bool
	installedImports Imports
	update           bool
}

type Import struct {
	Branch   string   `json:"branch,omitempty" yaml:"branch,omitempty"`
	Packages []string `json:"package,omitempty" yaml:"package,omitempty"`
	Revision string   `json:"revision,omitempty" yaml:"revision,omitempty"`
	Version  string   `json:"version,omitempty" yaml:"version,omitempty"`
	internal Internal
}

type Internal struct {
	repo Repo
}

type Repo struct {
	host    string
	user    string
	project string
	packag  string
	https   string
	ssh     string
	path    string
	vendor  string
	save    string
}

type Cache struct {
	imports Imports
	path    string
	config  string
}

type MovePackage struct {
	old string
	new string
}
