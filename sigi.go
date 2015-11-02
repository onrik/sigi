package sigi

import (
	"log"
	"reflect"
	"runtime/debug"
	"sync"
)

var (
	handlers map[string][]Handler = map[string][]Handler{}

	mutex sync.Mutex
)

type Handler interface{}

// Connect signal with handler
func Connect(signal string, handler Handler) {
	if kind := reflect.TypeOf(handler).Kind(); kind != reflect.Func {
		log.Printf("Handler type is '%s' must be func", kind)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	if _, exists := handlers[signal]; exists {
		handlers[signal] = append(handlers[signal], handler)
	} else {
		handlers[signal] = []Handler{handler}
	}
}

// Disconnect signal from handler
func Disconnect(signal string, handler Handler) {
	mutex.Lock()
	defer mutex.Unlock()
	if signalHandlers, exists := handlers[signal]; exists {
		handlerValue := reflect.ValueOf(handler)
		for i := range signalHandlers {
			if handlerValue == reflect.ValueOf(signalHandlers[i]) {
				handlers[signal] = append(signalHandlers[:i], signalHandlers[i+1:]...)
				return
			}
		}
	}
}

// Emit signal
func Emit(signal string, args ...interface{}) {
	signalHandlers, exists := handlers[signal]
	if !exists {
		log.Printf("No handlers for signal '%s'\n", signal)
		return
	}

	for _, handler := range signalHandlers {
		callHandler(handler, args...)
	}
}

func callHandler(handler Handler, args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			log.Println(string(debug.Stack()))
		}
	}()

	in := make([]reflect.Value, len(args))
	for i := range args {
		in[i] = reflect.ValueOf(args[i])
	}

	reflect.ValueOf(handler).Call(in)
}
