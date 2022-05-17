package domain

type DotFileCfg struct {
	Dots map[string][]string `json:"dots"`
}

// IO Service wrapper
type BackupSvc interface {
	ReadCfgFile(filePath string) (DotFileCfg, error)
	ProcessCfgFile(cfgFile DotFileCfg) error
}

/*
Implment all IO based logic here.
This way we can extend this service
to read/write files in different 
ways (HTTP, SQL, FTP, etc...)
*/
type StdIOSvc interface {
	ReadCfgFile(filePath string) (DotFileCfg, error)
	ProcessCfgFile(cfgFile DotFileCfg) error
}
