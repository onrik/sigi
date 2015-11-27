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
	_, err := Connect(signal, testHandler1)
	if err != nil {
		t.Error("Connect error", err.Error())
	}

	if !isHandlerConnected(signal, testHandler1) {
		t.Error("Handler not connected")
	}
}

func TestDisconnect(t *testing.T) {
	signal := "signal2"
	_, err := Connect(signal, testHandler1)
	if err != nil {
		t.Error("Connect error", err.Error())
	}

	_, err = Connect(signal, testHandler2)
	if err != nil {
		t.Error("Connect error", err.Error())
	}

	Disconnect(signal, testHandler2)

	if isHandlerConnected(signal, testHandler2) {
		t.Error("Handler wasn't disconneted")
	}

	if !isHandlerConnected(signal, testHandler1) {
		t.Error("Invalid handler disconnected")
	}
}

func TestDisconnectAnonymous(t *testing.T) {
	signal := "signal4"
	connector, err := Connect(signal, func() {})
	if err != nil {
		t.Error("Connect error", err.Error())
	}

	if !isHandlerConnected(signal, connector.handler) {
		t.Error("Handler is not conneted")
	}

	connector.Disconnect()

	if isHandlerConnected(signal, connector.handler) {
		t.Error("Handler wasn't disconneted")
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

func TestConnectError(t *testing.T) {
	connector, err := Connect("signal5", 22)
	if err == nil {
		t.Error("Check handler type error", err.Error())
	}

	if connector != nil {
		t.Error("Connector is not nil", connector)
	}
}
