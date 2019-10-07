package stubs

// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// LoggerStub provides a stub for the Logger client.
type LoggerStub struct {
}

//Info push info log to buffer
func (l *LoggerStub) Info(message string, a ...interface{}) {

}

//Warning push warning log to buffer
func (l *LoggerStub) Warning(message string, a ...interface{}) {

}

//Error push error log to buffer
func (l *LoggerStub) Error(message string, a ...interface{}) {

}

//Debug push debug log to buffer
func (l *LoggerStub) Debug(message string, a ...interface{}) {

}

//Close buffer and send messages to stackdriver
func (l *LoggerStub) Close() {

}
