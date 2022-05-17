package backup

import "example.com/dotfiles/pkg/domain"

type backupSvc struct {
	stdio domain.StdIOSvc
}

func NewBackupSvc(iosvc domain.StdIOSvc) domain.BackupSvc {
	return backupSvc{
		stdio: iosvc,
	}
}

func (bsvc backupSvc) ReadCfgFile(filePath string) (domain.DotFileCfg, error) {
	return bsvc.stdio.ReadCfgFile(filePath)
}

func (bsvc backupSvc) ProcessCfgFile(cfgFile domain.DotFileCfg) error {
	return bsvc.stdio.ProcessCfgFile(cfgFile)
}
