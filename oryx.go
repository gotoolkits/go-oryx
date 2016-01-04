// The MIT License (MIT)
//
// Copyright (c) 2013-2015 Oryx(ossrs)
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// The os defines:
//      bsd: darwin dragonfly freebsd nacl netbsd openbsd solaris
//      unix: bsd linux
//      server: unix plan9
//      posix: bsd linux windows
// All os by go:
//      server windows
//      posix plan9

package main

import (
	"flag"
	"fmt"
	"github.com/ossrs/go-oryx/app"
	"github.com/ossrs/go-oryx/core"
	"os"
)

func serve(svr *app.Server) int {
	if err := svr.PrepareLogger(); err != nil {
		core.Error.Println("prepare logger failed, err is", err)
		return -1
	}

	oryxMain(svr)

	core.Trace.Println("Copyright (c) 2013-2015 Oryx(ossrs)")
	core.Trace.Println(fmt.Sprintf("go-oryx/%v is advanced SRS, focus on realtime live streaming.", core.Version()))

	if err := svr.Initialize(); err != nil {
		core.Error.Println("initialize server failed, err is", err)
		return -1
	}

	if err := svr.Run(); err != nil {
		core.Error.Println("run server failed, err is", err)
		return -1
	}

	return 0
}

func main() {
	// the startup argv:
	//          -c conf/oryx.json
	//          --c conf/oryx.json
	//          -c=conf/oryx.json
	//          --c=conf/oryx.json
	//          --conf=conf/oryx.json
	var confFile string
	flag.StringVar(&confFile, "c", "", "the config file")
	flag.StringVar(&confFile, "conf", "", "the config file")

	flag.Usage = func(){
		fmt.Println(fmt.Sprintf("Usage: %v [-c|--conf <filename>] [-?|-h|--help]", os.Args[0]))
		fmt.Println(fmt.Sprintf("	-c, --conf filename     : the config file path"))
		fmt.Println(fmt.Sprintf("	-?, -h, --help          : show this help and exit"))
		fmt.Println(fmt.Sprintf("For example:"))
		fmt.Println(fmt.Sprintf("	%v -c conf/oryx.json", os.Args[0]))
	}
	flag.Parse()

	// show help without conf.
	if len(confFile) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	ret := func() int {
		svr := app.NewServer()
		defer svr.Close()

		if err := svr.ParseConfig(confFile); err != nil {
			core.Error.Println("parse config from", confFile, "failed, err is", err)
			return -1
		}

		return run(svr)
	}()

	os.Exit(ret)
}
