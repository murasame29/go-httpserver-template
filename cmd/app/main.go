package main

import (
	"flag"

	"github.com/murasame29/go-httpserver-template/cmd/config"
	nr "github.com/murasame29/go-httpserver-template/internal/pkg/newrelic"
	"github.com/murasame29/go-httpserver-template/internal/router"
	"github.com/murasame29/go-httpserver-template/internal/server"
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrlogrus"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

func init() {
	envPath := flag.String("env", "", "path to .env file")
	flag.Parse()

	// 環境変数を読み込む
	if *envPath != "" {
		config.LoadEnv(*envPath)
	} else {
		config.LoadEnv()
	}
}

func setLogger() (l *logrus.Logger) {
	l = logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	l.SetReportCaller(true)

	return
}

func setNewRelic(l *logrus.Logger) (nrApp *newrelic.Application) {
	nrApp, err := nr.NewNrApp(
		config.Config.NewRelic.AppName,
		config.Config.NewRelic.License,
	)
	if err != nil {
		l.Fatal(err)
	}

	l.SetFormatter(nrlogrus.NewFormatter(nrApp, &logrus.JSONFormatter{}))

	return nrApp
}

func main() {
	// loggerの設定
	l := setLogger()

	// NewRelicの設定。　不要ならコメントアウトか削除してください。
	//nrApp := setNewRelic(l)

	// サーバーの起動
	server.New(
		router.NewEchoServer(
			router.WithLogger(l),
			//router.WithNewRelic(nrApp),
		),
		server.WithLogger(l),
		server.WithHost(config.Config.Server.Host),
		server.WithPort(config.Config.Server.Port),
		server.WithShutdownTimeout(config.Config.Server.ShutdownTimeout),
	).RunWithGracefulShutdown()
}
