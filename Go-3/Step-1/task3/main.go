package main

import "log/slog"

// Напишите функцию LogUserAction(logger *slog.Logger, user string, action string),
//  которая логирует действие пользователя с помощью slog в формате
// Info с полями user и action.

func LogUserAction(logger *slog.Logger, user string, action string) {
	logger.Info("user action","user",user,"action", action)
}
