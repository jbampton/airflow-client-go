// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

/*
 * Airflow API (Stable)
 *
 * Apache Airflow management API.
 *
 * API version: 1.0.0
 * Contact: dev@airflow.apache.org
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package airflow
// ScheduleInterval struct for ScheduleInterval
type ScheduleInterval struct {
	Type string `json:"__type"`
	Days int32 `json:"days"`
	Seconds int32 `json:"seconds"`
	Microseconds int32 `json:"microseconds"`
	Years int32 `json:"years"`
	Months int32 `json:"months"`
	Leapdays int32 `json:"leapdays"`
	Hours int32 `json:"hours"`
	Minutes int32 `json:"minutes"`
	Year int32 `json:"year"`
	Month int32 `json:"month"`
	Day int32 `json:"day"`
	Hour int32 `json:"hour"`
	Minute int32 `json:"minute"`
	Second int32 `json:"second"`
	Microsecond int32 `json:"microsecond"`
	Value string `json:"value"`
}
