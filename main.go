// Copyright 2023 youwkey. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// main
package main

import (
	"flag"
	"log/slog"
	"net/http"
	"time"
)

const (
	defaultMountDir          = "./"
	defaultAllHost           = false
	defaultListenPort        = "3333"
	defaultReadHeaderTimeout = 3
)

type flags struct {
	rootDir string
	allHost bool
	port    string
}

type options struct {
	flags
	addr string
}

//nolint:gochecknoglobals
var (
	fRootDir string
	fAllHost bool
	fPort    string
)

//nolint:gochecknoinits
func init() {
	flag.StringVar(&fRootDir, "dir", defaultMountDir, "mount root directory")
	flag.BoolVar(&fAllHost, "all", defaultAllHost, "if set, bind any host 0.0.0.0")
	flag.StringVar(&fPort, "port", defaultListenPort, "listen port")
}

func parseOptions() options {
	flag.Parse()

	host := "127.0.0.1"
	if fAllHost {
		host = "0.0.0.0"
	}

	addr := host + ":" + fPort

	return options{
		flags: flags{
			rootDir: fRootDir,
			allHost: fAllHost,
			port:    fPort,
		},
		addr: addr,
	}
}

func buildHandler(rootDir string) http.Handler {
	return http.FileServer(http.Dir(rootDir))
}

func main() {
	opts := parseOptions()
	handler := buildHandler(opts.rootDir)

	//nolint:exhaustruct
	server := &http.Server{
		Addr:              opts.addr,
		Handler:           handler,
		ReadHeaderTimeout: defaultReadHeaderTimeout * time.Second,
	}

	slog.Info("root dir mounted.", "dir", opts.rootDir)
	slog.Info("server started.", "address", "http://localhost:"+opts.port)

	if err := server.ListenAndServe(); err != nil {
		slog.Error(err.Error())
	}
}
