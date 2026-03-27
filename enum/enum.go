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
	"iter"

	"github.com/cilium/ebpf/btf"
)

// AllTypedefedEnums iterates over all typedefs for an BTF enumeration type,
// return the name of the typedef as well the enumeration type. On purpose,
// AllTypedefedEnums ignores nested typedef's of enums.
func AllTypedefedEnums(spec *btf.Spec) iter.Seq2[string, *btf.Enum] {
	return func(yield func(string, *btf.Enum) bool) {
		for typ, err := range spec.All() {
			if err != nil {
				return // first encountered error aborts spec iterator
			}
			typedefT, ok := typ.(*btf.Typedef)
			if !ok {
				continue
			}
			// no btf.As on purpose, as we deal only with flat definitions.
			enumT, ok := typedefT.Type.(*btf.Enum)
			if !ok {
				continue
			}
			if !yield(typedefT.Name, enumT) {
				return
			}
		}
	}
}

// AllEnums iterates over all BTF enumeration types in a spec.
func AllEnums(spec *btf.Spec) iter.Seq[*btf.Enum] {
	return func(yield func(*btf.Enum) bool) {
		for typ, err := range spec.All() {
			if err != nil {
				return // first encountered error aborts spec iterator
			}
			enumT, ok := typ.(*btf.Enum)
			if !ok {
				continue
			}
			if !yield(enumT) {
				return
			}
		}
	}
}
