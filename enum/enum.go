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
	"strconv"

	"github.com/cilium/ebpf/btf"
	"github.com/thediveo/opt"
)

// TypedefEnum represents an “typedef enum { ... } NAME;” including the
// enumerated named values.
type TypedefEnum struct {
	name string
	enum *btf.Enum
}

// Name returns the type name of the enumeration.
//
// Important: do not confuse the type name with the name of the enumeration, if
// any.
func (te *TypedefEnum) Name() string { return te.name }

// Gotype returns the name of the Go type that can represent the values of this
// typedef'ed enumeration, such as “uint32”, “int8”, et cetera.
func (te *TypedefEnum) Gotype() string {
	vt := opt.If[string](!te.enum.Signed).Then("u").Else("")
	switch te.enum.Size {
	case 1:
		vt += "int8"
	case 2:
		vt += "int16"
	case 4:
		vt += "int32"
	case 8:
		vt += "int64"
	}
	return vt
}

// AllElements iterates over all named values of this typedef'ed enum, producing
// (name, formatted-value) pairs, taking signed-ness correctly into consideration.
func (te *TypedefEnum) AllElements() iter.Seq2[string, string] {
	return func(yield func(string, string) bool) {
		formatFn := opt.If[func(uint64, int) string](te.enum.Signed).
			Then(formatInt).Else(strconv.FormatUint)
		for _, enumval := range te.enum.Values {
			if !yield(enumval.Name, formatFn(enumval.Value, 10)) {
				return
			}
		}
	}
}

func formatInt(i uint64, base int) string { return strconv.FormatInt(int64(i), base) }

// AllTypedefedEnums iterates over all typedefs for an BTF enumeration type,
// return the name of the typedef as well the enumeration type. On purpose,
// AllTypedefedEnums ignores nested typedef's of enums.
func AllTypedefedEnums(spec *btf.Spec) iter.Seq[*TypedefEnum] {
	return func(yield func(*TypedefEnum) bool) {
		for typ, err := range spec.All() {
			if err != nil {
				return // first encountered error aborts spec iterator
			}
			typedefT, ok := typ.(*btf.Typedef)
			if !ok {
				continue
			}
			// no btf.As here on purpose, as we deal only with flat definitions.
			enumT, ok := typedefT.Type.(*btf.Enum)
			if !ok {
				continue
			}
			if !yield(&TypedefEnum{name: typedefT.Name, enum: enumT}) {
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
