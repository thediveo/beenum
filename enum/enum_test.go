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
	"slices"

	"github.com/cilium/ebpf/btf"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/thediveo/success"
)

var _ = Describe("ebpf enums", Ordered, func() {

	var spec *btf.Spec

	BeforeAll(func() {
		spec = Successful(btf.LoadSpec("./test/test_bpfel.o"))
	})

	It("iterates over typedef'ed enumerations in a spec", func() {
		tenums := slices.Collect(AllTypedefedEnums(spec))
		Expect(tenums).To(HaveLen(2))
		Expect(tenums).To(ConsistOf(
			And(HaveField("Name()", "defines"),
				HaveField("Gotype()", "uint64"),
				HaveField("AllElements()", HaveKeyWithValue("DEFAULT_FOOBAR", "281474976710656"))),
			And(HaveField("Name()", "shortdefines"),
				HaveField("Gotype()", "int8"),
				HaveField("AllElements()", HaveKeyWithValue("NOT_MAX", "-127")))))
	})

	It("iterates over the enumerations in a spec", func() {
		Expect(AllEnums(spec)).To(ConsistOf(
			HaveField("Values", ContainElement(
				HaveField("Name", "MAX_BEES"))),
			HaveField("Values", ContainElement(
				HaveField("Name", "NOT_MAX")))))
	})

})
