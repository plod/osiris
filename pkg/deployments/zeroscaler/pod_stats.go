package zeroscaler

import (
	"time"

	"github.com/plod/osiris/pkg/metrics"
)

type podStats struct {
	podDeletedTime *time.Time
	prevStatTime   *time.Time
	prevStats      *metrics.ProxyConnectionStats
	recentStatTime *time.Time
	recentStats    *metrics.ProxyConnectionStats
}
