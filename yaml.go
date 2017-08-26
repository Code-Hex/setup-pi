package setup

import "os"

type Recipe struct {
	Items []*Item
}

type Item struct {
	Name    string
	Command string
	*Apt
	*Copy   `yaml:"copy"`
	*GetURL `yaml:"get_url"`
}

type Copy struct {
	Src  string
	Dest string
	Mode os.FileMode
}

type Apt struct {
	Packages     []string
	BeforeUpdate bool `yaml:"before_update"`
	IsUpgrade    bool `yaml:"is_upgrade"`
}

type GetURL struct {
	URL  string
	Dest string
}
