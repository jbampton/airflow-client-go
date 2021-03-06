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
import (
	"time"
)
// ListTaskInstanceForm struct for ListTaskInstanceForm
type ListTaskInstanceForm struct {
	// Return objects with specific DAG IDs. The value can be repeated to retrieve multiple matching values (OR condition).
	DagIds []string `json:"dag_ids,omitempty"`
	// Returns objects greater or equal to the specified date. This can be combined with execution_date_lte parameter to receive only the selected period. 
	ExecutionDateGte time.Time `json:"execution_date_gte,omitempty"`
	// Returns objects less than or equal to the specified date. This can be combined with execution_date_gte parameter to receive only the selected period. 
	ExecutionDateLte time.Time `json:"execution_date_lte,omitempty"`
	// Returns objects greater or equal the specified date. This can be combined with startd_ate_lte parameter to receive only the selected period. 
	StartDateGte time.Time `json:"start_date_gte,omitempty"`
	// Returns objects less or equal the specified date. This can be combined with start_date_gte parameter to receive only the selected period. 
	StartDateLte time.Time `json:"start_date_lte,omitempty"`
	// Returns objects greater or equal the specified date. This can be combined with start_date_lte parameter to receive only the selected period. 
	EndDateGte time.Time `json:"end_date_gte,omitempty"`
	// Returns objects less than or equal to the specified date. This can be combined with start_date_gte parameter to receive only the selected period. 
	EndDateLte time.Time `json:"end_date_lte,omitempty"`
	// Returns objects greater than or equal to the specified values. This can be combined with duration_lte parameter to receive only the selected period. 
	DurationGte float32 `json:"duration_gte,omitempty"`
	// Returns objects less than or equal to the specified values. This can be combined with duration_gte parameter to receive only the selected range. 
	DurationLte float32 `json:"duration_lte,omitempty"`
	// The value can be repeated to retrieve multiple matching values (OR condition).
	State []string `json:"State,omitempty"`
	// The value can be repeated to retrieve multiple matching values (OR condition).
	Pool []string `json:"Pool,omitempty"`
	// The value can be repeated to retrieve multiple matching values (OR condition).
	Queue []string `json:"Queue,omitempty"`
}
