package main

import "testing"

func Test_sendInt(t *testing.T) {
	type args struct {
		ch chan int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendInt(tt.args.ch)
		})
	}
}

func Test_sendBool(t *testing.T) {
	type args struct {
		ch chan bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendBool(tt.args.ch)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_sendMessage(t *testing.T) {
	type args struct {
		ch      chan<- string
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendMessage(tt.args.ch, tt.args.message)
		})
	}
}

func Test_receiveMessage(t *testing.T) {
	type args struct {
		helloCh   <-chan string
		goodByeCh <-chan string
		quitCh    chan<- bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiveMessage(tt.args.helloCh, tt.args.goodByeCh, tt.args.quitCh)
		})
	}
}
