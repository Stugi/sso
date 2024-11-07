package sl

import "log/slog"

// Err returns slog.Attr with error value
// sl.Err(err) вместо slog.Info("error", err.Error()).
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
