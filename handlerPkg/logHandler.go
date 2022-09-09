package handlerPkg

import (
	"prConsumer/internal"
	"prConsumer/model"
)

type LogHandler struct {
	logService *internal.LogService
}

func (lh *LogHandler) Log(log *model.Log) error {
	if err := lh.logService.LogRepository.Insert(log); err != nil {
		return err
	}
	return nil
}
