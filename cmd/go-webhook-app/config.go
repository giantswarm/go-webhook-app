package main

import (
	"os"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

// CmdConfig represents the configuration of the command.
type CmdConfig struct {
	Debug               bool
	Development         bool
	WebhookListenAddr   string
	MetricsListenAddr   string
	MetricsPath         string
	TLSCertFilePath     string
	TLSKeyFilePath      string
	MinSMScrapeInterval time.Duration

	EnableHAUpdater   bool
	EnableHAUpdaterAZ string
}

// NewCmdConfig returns a new command configuration.
func NewCmdConfig() (*CmdConfig, error) {
	c := &CmdConfig{}

	app := kingpin.New("k8s-webhook-example", "A Kubernetes production-ready admission webhook example.")
	app.Version(Version)

	app.Flag("debug", "Enable debug mode.").BoolVar(&c.Debug)
	app.Flag("development", "Enable development mode.").BoolVar(&c.Development)
	app.Flag("webhook-listen-address", "the address where the HTTPS server will be listening to serve the webhooks.").Default(":8080").StringVar(&c.WebhookListenAddr)
	app.Flag("metrics-listen-address", "the address where the HTTP server will be listening to serve metrics, healthchecks, profiling...").Default(":8081").StringVar(&c.MetricsListenAddr)
	app.Flag("metrics-path", "the path where Prometheus metrics will be served.").Default("/metrics").StringVar(&c.MetricsPath)
	app.Flag("tls-cert-file-path", "the path for the webhook HTTPS server TLS cert file.").StringVar(&c.TLSCertFilePath)
	app.Flag("tls-key-file-path", "the path for the webhook HTTPS server TLS key file.").StringVar(&c.TLSKeyFilePath)
	app.Flag("webhook-haupdater", "enable haupdater webhook").Short('h').BoolVar(&c.EnableHAUpdater)
	app.Flag("webhook-haupdater-azs", "enable haupdater webhook").Short('a').StringVar(&c.EnableHAUpdaterAZ)
	app.Flag("webhook-sm-min-scrape-interval", "the minimum screate interval service monitors can have.").DurationVar(&c.MinSMScrapeInterval)

	_, err := app.Parse(os.Args[1:])
	if err != nil {
		return nil, err
	}

	return c, nil
}
