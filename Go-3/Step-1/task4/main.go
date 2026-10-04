package main

import "log/slog"

func LogHTTPRequest(logger *slog.Logger, method, path string, status int, durationMs int64) {
	logger.Info("http request","method", method, "path", path, "status", status, "duration_ms", durationMs)
}
