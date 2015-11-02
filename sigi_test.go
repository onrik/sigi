package sigi

import (
	"reflect"
	"testing"
)

var testCount int

func testHandler1(count int) {
	testCount += count
}

func testHandler2() {
	testCount++
}

func isHandlerConnected(signal string, handler Handler) bool {
	signalHandlers, exists := handlers[signal]
	if !exists {
		return false
	}

	handlerValue := reflect.ValueOf(handler)
	for i := range signalHandlers {
		if handlerValue == reflect.ValueOf(signalHandlers[i]) {
			return true
		}
	}

	return false
}

func TestConnect(t *testing.T) {
	signal := "signal1"
	Connect(signal, testHandler1)

	if !isHandlerConnected(signal, testHandler1) {
		t.Error("Handler not connected")
	}
}

func TestDisconnect(t *testing.T) {
	signal := "signal2"
	Connect(signal, testHandler1)
	Connect(signal, testHandler2)

	Disconnect(signal, testHandler2)

	if isHandlerConnected(signal, testHandler2) {
		t.Error("Handler wasn't disconneted")
	}

	if !isHandlerConnected(signal, testHandler1) {
		t.Error("Invalid handler disconnected")
	}
}

func TestEmit(t *testing.T) {
	signal := "signal3"
	Connect(signal, testHandler1)
	Emit(signal, 2)

	if testCount != 2 {
		t.Error("Handler wasn't called")
	}
}
