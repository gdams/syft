package pkg

type Digest struct {
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
}

// PythonFileRecord represents a single entry within a RECORD file for a python wheel or egg package
type PythonFileRecord struct {
	Path   string `json:"path"`
	Digest Digest `json:"digest"`
	Size   string `json:"size"`
}

// PythonPackageMetadata represents all captured data for a python egg or wheel package.
type PythonPackageMetadata struct {
	Name                 string             `json:"name" mapstruct:"Name"`
	Version              string             `json:"version" mapstruct:"Version"`
	License              string             `json:"license" mapstruct:"License"`
	Author               string             `json:"author" mapstruct:"Author"`
	AuthorEmail          string             `json:"authorEmail" mapstruct:"Authoremail"`
	Platform             string             `json:"platform" mapstruct:"Platform"`
	Files                []PythonFileRecord `json:"files,omitempty"`
	SitePackagesRootPath string             `json:"sitePackagesRootPath"`
	TopLevelPackages     []string           `json:"topLevelPackages,omitempty"`
}
