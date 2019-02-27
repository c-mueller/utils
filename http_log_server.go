// Copyright 2019 Christian Müller <dev@c-mueller.xyz>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

var MaxLogLines = 250

var format = ""

var HttpLoggerInst = HttpLogger{}

func SetLoggerFormat(f string) {
	format = f
}

type LogRecord struct {
	Message   string `json:"message"`
	Level     string `json:"level"`
	Timestamp int64  `json:"timestamp"`
}

type HttpLogger struct {
	logs             []LogRecord
	FormattedBackend logging.Backend
}

func (a *HttpLogger) InitLogger() {
	formatter := logging.MustStringFormatter(format)
	backend := logging.AddModuleLevel(logging.NewBackendFormatter(a, formatter))
	backend.SetLevel(logging.DEBUG, "")
	a.FormattedBackend = backend
}

func (a *HttpLogger) GetLogs(ctx *gin.Context) {
	ctx.JSON(200, a.logs)
}

func (a *HttpLogger) Log(level logging.Level, i int, rec *logging.Record) error {
	a.appendLogMessage(level.String(), rec.Formatted(i+2), rec.Time.Unix())
	return nil
}

func (a *HttpLogger) appendLogMessage(level, message string, timestamp int64) {
	m := LogRecord{
		Level:     level,
		Message:   message,
		Timestamp: timestamp,
	}

	a.logs = append([]LogRecord{m}, a.logs...)

	if len(a.logs) > MaxLogLines {
		a.logs = a.logs[0:MaxLogLines]
	}
}
