// Copyright 2026 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy
// of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package enum

import (
	"maps"

	"github.com/cilium/ebpf/btf"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/thediveo/success"
)

var _ = Describe("ebpf enums", Ordered, func() {

	var spec *btf.Spec

	BeforeAll(func() {
		spec = Successful(btf.LoadSpec("../examples/config/config_bpfel.o"))
	})

	It("iterates over typedef'ed enumerations in a spec", func() {
		tenums := maps.Collect(AllTypedefedEnums(spec))
		Expect(tenums).To(HaveLen(1))
		Expect(tenums).To(HaveKeyWithValue(
			"defines",
			HaveField("Values", ContainElements(
				HaveField("Name", "MAX_BEES"),
				HaveField("Name", "DEFAULT_FOOBAR")))))
	})

	It("iterates over the enumerations in a spec", func() {
		Expect(AllEnums(spec)).To(ConsistOf(
			HaveField("Values", ContainElement(
				HaveField("Name", "MAX_BEES")))))
	})

})
