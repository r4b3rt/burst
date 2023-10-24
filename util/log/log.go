package log

import "log/slog"

func Reason(err error) slog.Attr {
	return slog.String("reason", err.Error())
}
