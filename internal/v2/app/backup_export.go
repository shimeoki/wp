package app

type ExportBackupQuery struct{}

type ExportBackupResult struct {
	Backup *BackupResult
}

func (s *BackupService) Export(
	Ctx,
	*ExportBackupQuery,
) (*ExportBackupResult, error) {
	return nil, nil // TODO: implement
}
