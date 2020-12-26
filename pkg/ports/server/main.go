package main

import "github.com/mlambda-net/store/pkg/infrastructure/server"

func main() {
	s := server.NewServer()
	s.Start()
	s.Wait()
}
