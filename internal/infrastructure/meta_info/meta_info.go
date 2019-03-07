package meta_info

import "fmt"

type MetaInfo struct {
	Version   string
	Revision  string
	BuildDate string
	GoVersion string
	Mode      string
}

func NewMetaInfo(
	version string,
	revision string,
	buildDate string,
	goVersion string,
	mode string,
) *MetaInfo {
	return &MetaInfo{
		Version:   version,
		Revision:  revision,
		BuildDate: buildDate,
		GoVersion: goVersion,
		Mode:      mode,
	}
}

func (m *MetaInfo) GetMetaInfo() string {
	meta := fmt.Sprintf("version: %s (%s) \n", m.Version, m.Revision)
	meta += fmt.Sprintf("build at date: %s \n", m.BuildDate)
	meta += fmt.Sprintf("build at go version: %s \n", m.GoVersion)
	meta += fmt.Sprintf("build mode : %s \n", m.Mode)
	return meta
}
