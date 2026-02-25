package logging
import(
	"log/slog"
	"os"
	"context"
	// "net/http"
	// "strings"
	// "github.com/gin-gonic/gin"
)

type AuditEvent struct {
	Namespace       string
	Action       string
	Result       string
	IP           string
	ErrorType 	string
	ErrorMessage string
}

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

// logger.Info("audit",
// 	slog.String("user_id", userID),
// 	slog.String("action", "instance_create"),
// 	slog.String("resource_type", "vm"),
// 	slog.String("resource_id", vmID),
// 	slog.String("result", "success"),
// )

func ErrorTypeFromStatus(status int) string {
	switch status {
	case 400:
		return "validation"
	case 401, 403:
		return "permission"
	case 404:
		return "not_found"
	case 409:
		return "conflict"
	case 429:
		return "quota"
	case 408, 504:
		return "timeout"
	default:
		if status >= 500 {
			return "internal"
		}
		return "unknown"
	}
}




func AuditLog(ctx context.Context, ev AuditEvent, logger *slog.Logger) {
	attrs := []slog.Attr{
		slog.String("action", ev.Namespace),
		slog.String("action", ev.Action),
		slog.String("result", ev.Result),
		slog.String("ip", ev.IP),
}
	if ev.ErrorType != "" {
		attrs = append(attrs, slog.String("error_type", ev.ErrorType))
	}
	if ev.ErrorMessage != "" {
		attrs = append(attrs, slog.String("error_type", ev.ErrorMessage))
	}
	logger.LogAttrs(ctx, slog.LevelInfo, "audit", attrs...)
}