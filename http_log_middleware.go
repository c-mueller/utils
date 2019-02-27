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
	"time"
)

func LogMiddleware(moduleName string) gin.HandlerFunc {
	log := logging.MustGetLogger(moduleName)
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if raw != "" {
			path = path + "?" + raw
		}

		log.Info("[%s]: %s %s - %d: %.3f MS %s", clientIP, method, path, statusCode, float64(latency.Nanoseconds()/1000)/1000, comment)

	}
}
