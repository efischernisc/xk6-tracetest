package output

import (
	"fmt"
	"math/rand"

	"github.com/efischernisc/xk6-tracetest/models"
	"github.com/efischernisc/xk6-tracetest/modules/tracetest"
	"github.com/sirupsen/logrus"

	"go.k6.io/k6/metrics"
	"go.k6.io/k6/output"
)

type Output struct {
	config    models.OutputConfig
	testRunID int64
	logger    logrus.FieldLogger
	tracetest *tracetest.Tracetest
}

var _ output.Output = new(Output)

func New(params output.Params, tracetest *tracetest.Tracetest) (*Output, error) {
	if tracetest == nil {
		return nil, fmt.Errorf("tracetest must not be nil")
	}

	config, err := models.NewConfig(params)
	if err != nil {
		return nil, err
	}

	tracetest.UpdateFromConfig(config)

	return &Output{
		config:    config,
		tracetest: tracetest,
		logger:    params.Logger.WithField("component", "xk6-tracetest-output"),
	}, nil
}

func (o *Output) Description() string {
	return fmt.Sprintf("xk6-tracetest-output (TestRunID: %d)", o.testRunID)
}

func (o *Output) AddMetricSamples(samples []metrics.SampleContainer) {
	if len(samples) == 0 {
		return
	}

	for _, s := range samples {
		o.handleSample(s)
	}
}

func (o *Output) Stop() error {
	o.logger.Debug("Stopping...")
	defer o.logger.Debug("Stopped!")

	return nil
}

func (o *Output) Start() error {
	o.logger.Debug("Starting...")
	o.testRunID = 10000 + rand.Int63n(99999-10000)
	o.logger.Debug("Started!")

	return nil
}
