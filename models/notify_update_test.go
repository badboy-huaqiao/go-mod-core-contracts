/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package models

import "testing"

var testNotifyUpdate = NotifyUpdate{Name: "For Testing", Operation: NotifyUpdateAdd}

func TestNotifyUpdateValidation(t *testing.T) {
	valid := testNotifyUpdate

	invalidName := testNotifyUpdate
	invalidName.Name = ""

	invalidOp := testNotifyUpdate
	invalidOp.Operation = ""

	invalidOpNotBlank := testNotifyUpdate
	invalidOpNotBlank.Operation = "blah"

	tests := []struct {
		name        string
		nu          NotifyUpdate
		expectError bool
	}{
		{"valid notify update", valid, false},
		{"invalid name", invalidName, true},
		{"invalid operation", invalidOp, true},
		{"invalid operation value", invalidOpNotBlank, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.nu.Validate()
			checkValidationError(err, tt.expectError, tt.name, t)
		})
	}
}
