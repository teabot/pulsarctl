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

package topic

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func RemoveDeduplicationStatusCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Remove the deduplication policy for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	msg := cmdutils.Example{
		Desc:    "Remove the deduplication policy for a topic",
		Command: "pulsarctl topics remove-deduplication topic",
	}
	examples = append(examples, msg)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Remove the deduplication policy successfully for [topic]",
	}
	out = append(out, successOut, ArgError)
	out = append(out, TopicNameErrors...)
	out = append(out, NamespaceErrors...)
	desc.CommandOutput = out

	vc.SetDescription(
		"remove-deduplication",
		"Remove the deduplication policy for a topic",
		desc.ToString(),
		desc.ExampleToString(),
		"remove-deduplication",
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doRemoveDeduplicationStatus(vc)
	}, "the topic name is not specified or the topic name is specified more than one")
}

func doRemoveDeduplicationStatus(vc *cmdutils.VerbCmd) error {
	// for testing
	if vc.NameError != nil {
		return vc.NameError
	}

	topic, err := utils.GetTopicName(vc.NameArg)
	if err != nil {
		return err
	}

	admin := cmdutils.NewPulsarClient()
	err = admin.Topics().RemoveDeduplicationStatus(*topic)
	if err == nil {
		vc.Command.Printf("Remove the deduplication policy successfully for [%s]\n", topic.String())
	}
	return err
}
