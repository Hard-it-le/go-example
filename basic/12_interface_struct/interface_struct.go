package main

import "net/http"

type Server interface {
	Route(pattern string, headerFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
	age  int
}

type Header map[string][]string
