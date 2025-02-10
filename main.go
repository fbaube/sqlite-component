//go:generate go run go.bytecodealliance.org/cmd/wit-bindgen-go generate --world example --out gen ./wit

package main

import (
	"github.com/fbaube/scratch"
	"go.wasmcloud.dev/component/log/wasilog"
	"log/slog"
)

const InvokeResponse = "Hello from Invoker func #1!"
const InvokeResponse2 = "Hello from Invoker func #2!"

func init() {
	invoker.Exports.Call = invokerCall
	invoker.Exports.Func2 = func2call
}

// type LogValuer interface { LogValue() Value }

// You need a key for every value to be logged.
//  So after msg (and maybe ctx) arg(s), use
//  slog.Any(key,value) to pass a K-V pair.

func invokerCall() string {
     	// slog's func NewTextHandler is not listed at
	// https://pkg.go.dev/go.wasmcloud.dev/component/log/wasilog
     	// logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// So, instead: 
	logger := wasilog.ContextLogger("") // was: "Call"

	logger.Info("MyNewInfo", slog.Any("Loggable-nr", 42))
	
	logger.Info("Invoking function")
	return InvokeResponse
}

func func2call() string {
	logger := wasilog.ContextLogger("") // was: "Call"

	logger.Info("2-2-2-2", slog.Any("Loggable-nr", 2))
	
	logger.Info("Invoking func2")
	return InvokeResponse2
}

func main() {}
