package test

import (
  "github.com/mlambda-net/net/pkg/net"
  "github.com/mlambda-net/store/pkg/infrastructure/server"
  "os"
)

type ActorTest struct {
  client net.Client
  server server.Server
}

func (t ActorTest) LoadActor(actor string) net.Request {
  return t.client.Actor(actor)
}

func (t ActorTest) close()  {
  t.server.Stop()
}

func NewClient() ActorTest {
  _ = os.Setenv("Debug", os.Getenv("DEBUG"))
  s := server.NewServer()
  s.Start()
  c := net.NewClient("localhost", "9000")
  return ActorTest{client: c, server: s}
}
