// Copyright 2023 youwkey. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// main
package main

import (
	"flag"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

const (
	defaultMountDir          = "./"
	defaultListenHost        = "localhost"
	defaultListenPort        = 3000
	defaultReadHeaderTimeout = 3
)

type flags struct {
	rootDir string
	host    string
	port    uint
}

type options struct {
	flags
	addr string
}

//nolint:gochecknoglobals
var (
	fRootDir string
	fHost    string
	fPort    uint
)

//nolint:gochecknoinits
func init() {
	flag.StringVar(&fRootDir, "dir", defaultMountDir, "mount root directory")
	flag.StringVar(&fHost, "host", defaultListenHost, "listen host")
	flag.UintVar(&fPort, "port", defaultListenPort, "listen port")
}

func parseOptions() options {
	flag.Parse()

	addr := fHost + ":" + strconv.FormatUint(uint64(fPort), 10)

	return options{
		flags: flags{
			rootDir: fRootDir,
			host:    fHost,
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
		Addr:              "localhost:3000",
		Handler:           handler,
		ReadHeaderTimeout: defaultReadHeaderTimeout * time.Second,
	}

	slog.Info("root dir mounted.", "dir", opts.rootDir)
	slog.Info("server started.", "address", "http://"+opts.addr)

	if err := server.ListenAndServe(); err != nil {
		slog.Error(err.Error())
	}
}
