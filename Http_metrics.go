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
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var (
	requestExecTime = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:      "request_exectution_time",
		Namespace: "svcd_api",
		Help:      "Request Processing times of the HTTP Api",
	}, []string{"type", "path", "response_code"})
	Type = "server"
)

func init() {
	prometheus.MustRegister(requestExecTime)
}

type FilterFunc func(string) string

var DefaultFilterFunc = func(s string) string { return s }

func MetricsMiddleware() gin.HandlerFunc {
	return MetricsMiddlewareFiltered(DefaultFilterFunc)
}

func MetricsMiddlewareFiltered(filter FilterFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		context.Next()
		execTime := time.Now().Sub(t)

		status := fmt.Sprintf("%d", context.Writer.Status())

		path := filter(context.Request.URL.Path)

		requestExecTime.WithLabelValues(Type, path, status).Observe(float64(execTime.Nanoseconds()) / 1000000.0)
	}
}
