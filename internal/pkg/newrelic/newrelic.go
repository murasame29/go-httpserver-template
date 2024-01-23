package newrelic

import (
	"github.com/newrelic/go-agent/v3/newrelic"
)

// NewNrApp は、NewRelic の Application を生成する関数です。
func NewNrApp(
	appName string,
	license string,
) (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(license),
		newrelic.ConfigAppLogForwardingEnabled(true),
		newrelic.ConfigAppLogEnabled(true),
	)
}
