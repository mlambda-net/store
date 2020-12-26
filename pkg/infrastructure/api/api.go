package api

import (
  "context"
  "errors"
  "fmt"
  "github.com/etherlabsio/healthcheck"
  "github.com/mlambda-net/net/pkg/metrics"
  "github.com/mlambda-net/net/pkg/net"
  "github.com/mlambda-net/store/pkg/infrastructure/api/command"
  "github.com/mlambda-net/store/pkg/infrastructure/api/conf"
  "github.com/mlambda-net/store/pkg/infrastructure/api/query"
  "github.com/sirupsen/logrus"
  "os"
  "time"
)

type Api interface {
  GetVersion() string
  GetHost() string
  Start()
  Path() string
  Docs() int
}

type setup struct {
  config  *conf.Configuration
  command command.Command
  query   query.Query
}

func (s *setup) Docs() int {
  return s.config.Docs.Port
}

func (s *setup) Path() string {
   if s.config.Docs.Path != "" {
     return fmt.Sprintf("/%s", s.config.Docs.Path)
   }
   return ""
}

func NewApi() Api  {
  return &setup{config: conf.LoadConfig()}
}

func (s *setup) GetVersion() string {
  version := os.Getenv("VERSION")
  if version == "" {
    version = "0.0.0"
  }
  return version
}

func (s *setup) GetHost() string {
  if s.config.Docs.Host == "localhost" {
    return fmt.Sprintf("%s:%d", s.config.Docs.Host, s.config.App.Port)
  }
  return s.config.Docs.Host
}

func (s *setup) Start() {
  client := net.NewClient(s.config.Remote.Server, s.config.Remote.Port)
  store := client.Actor("store")
  currency := client.Actor("currency")
  lookup := client.Actor("lookup")
  image := client.Actor("image")

  s.command = command.NewCommand(store, currency, lookup, image)
  s.query = query.NewQuery(store)

  local := net.NewApi(int32(s.config.App.Port), int32(s.config.Metric.Port))

  local.Metrics(func(mc *metrics.Configuration) {
    mc.App.Name = s.config.App.Name
    mc.App.Env = s.config.Env
    mc.App.Path = "/check/store"
    mc.App.Version = s.config.App.Version
    mc.Metric.Namespace = s.config.Metric.Namespace
    mc.Metric.SubSystem = "store"
  })

  local.Register(func(r net.Route) {
    s.command.Register(r)
    s.query.Register(r)
  })

  local.Checks(
    healthcheck.WithTimeout(5 * time.Second),
    healthcheck.WithChecker("server",healthcheck.CheckerFunc(func(ctx context.Context) error {
      status, err := client.Health()
      if err != nil {
        return err
      }
      if !status.Success {
        return errors.New(status.Message)
      }
      return nil
    })),
  )

  logrus.Info("listening API Host on:", s.config.App.Port)
  logrus.Info("listening API Metrics on:", s.config.Metric.Port)
  logrus.Info("listening API Docs on:", s.config.Docs.Port)
  local.Start()
  local.Wait()

}
