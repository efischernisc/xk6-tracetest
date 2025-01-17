package output

import (
	"fmt"
	"strconv"

	"github.com/efischernisc/xk6-tracetest/models"
	"github.com/efischernisc/xk6-tracetest/utils"
	"go.k6.io/k6/lib/netext/httpext"
	"go.k6.io/k6/metrics"
)

func (o *Output) handleSample(sample metrics.SampleContainer) {
	if httpSample, ok := sample.(*httpext.Trail); ok {
		o.handleHttpSample(httpSample)
	}
}

func (o *Output) handleHttpSample(trail *httpext.Trail) {
	traceID, hasTrace := trail.Metadata["trace_id"]
	testID, hasTestID := trail.Metadata["test_id"]
	variableName, _ := trail.Metadata["variable_name"]
	_, hasShouldWait := trail.Metadata["should_wait"]

	if !hasTrace || !hasTestID {
		return
	}

	totalDuration := trail.Blocked + trail.ConnDuration + trail.Duration
	startTime := trail.EndTime.Add(-totalDuration)

	getTag := func(name string) string {
		val, _ := trail.Tags.Get(name)
		return val
	}

	strStatus := getTag("status")
	status, err := strconv.ParseInt(strStatus, 10, 64)
	if err != nil {
		o.logger.Warnf("unexpected error parsing status '%s': %w", strStatus, err)
		return
	}

	metadata := models.Metadata{
		"startTimeUnixNano": fmt.Sprint(startTime.UnixNano()),
		"endTimeUnixNano":   fmt.Sprint(trail.EndTime.UnixNano()),
		"group":             getTag("group"),
		"scenario":          getTag("scenario"),
		"traceID":           traceID,
		"hTTPUrl":           getTag("url"),
		"hTTPMethod":        getTag("method"),
		"hTTPStatus":        fmt.Sprint(status),
		"extension":         "K6",
	}

	request := models.Request{
		Method:   getTag("method"),
		URL:      getTag("url"),
		ID:       utils.RandHexStringRunes(8),
		Metadata: metadata,
	}

	options := models.TracetestOptions{
		VariableName: variableName,
		TestID:       testID,
		ShouldWait:   hasShouldWait,
	}

	if hasTestID {
		o.tracetest.RunTest(traceID, options, request)
	}
}
