package funcs

import (
	"github.com/open-falcon/common/model"
)

func AgentMetrics() []*model.MetricValue {
	return []*model.MetricValue{GaugeValue("Tomcat.Monitor.alive", 1)}
}
