// Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
// This program and the accompanying materials are made available under the terms of the under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

package integration_new_tests

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-cf/on-demand-service-broker/broker"
)

func matchingUpdateOperationWith(updateTaskId int) types.GomegaMatcher {
	return &updateOperationMatcher{
		expected: MatchJSON(fmt.Sprintf(`{"BoshTaskID": %d, "OperationType": "update"}`, updateTaskId)),
	}
}

type updateOperationMatcher struct {
	expected types.GomegaMatcher
}

func (uo *updateOperationMatcher) Match(actual interface{}) (bool, error) {
	return uo.expected.Match(asOperationData(actualBytes(actual)))
}

func (uo *updateOperationMatcher) FailureMessage(actual interface{}) (message string) {
	return uo.expected.FailureMessage(asOperationData(actualBytes(actual)))
}

func (uo *updateOperationMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return uo.expected.NegatedFailureMessage(asOperationData(actualBytes(actual)))
}

func asOperationData(source []byte) []byte {
	var updateResponse brokerapi.UpdateResponse
	err := json.Unmarshal(source, &updateResponse)
	Expect(err).NotTo(HaveOccurred())

	var operationData broker.OperationData
	err = json.Unmarshal([]byte(updateResponse.OperationData), &operationData)
	Expect(err).NotTo(HaveOccurred())
	return serialized(&operationData)
}

func actualBytes(actual interface{}) []byte {
	bytes, isBytes := actual.([]byte)
	Expect(isBytes).To(BeTrue(), fmt.Sprintf("converting actual to string: %s", actual))
	return bytes
}

func serialized(source *broker.OperationData) []byte {
	bytes, err := json.Marshal(source)
	Expect(err).NotTo(HaveOccurred())
	return bytes
}
