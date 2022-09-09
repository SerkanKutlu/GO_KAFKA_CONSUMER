package handlerPkg

import "prConsumer/internal"

type Handler struct {
	LogHandler *LogHandler
}

var appHandler *Handler

func SetHandler(logService *internal.LogService) {
	appHandler = new(Handler)
	appHandler.LogHandler = new(LogHandler)
	appHandler.LogHandler.logService = logService
}

func GetHandler() *Handler {
	return appHandler
}
