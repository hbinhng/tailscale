// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Code generated by tailscale/cmd/viewer; DO NOT EDIT.

package prefs

import (
	"encoding/json"
	"errors"
	"net/netip"

	jsonexpv2 "github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

//go:generate go run tailscale.com/cmd/cloner  -clonefunc=false -type=TestPrefs,TestBundle,TestValueStruct,TestGenericStruct,TestPrefsGroup -tags=test

// View returns a read-only view of TestPrefs.
func (p *TestPrefs) View() TestPrefsView {
	return TestPrefsView{ж: p}
}

// TestPrefsView provides a read-only view over TestPrefs.
//
// Its methods should only be called if `Valid()` returns true.
type TestPrefsView struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *TestPrefs
}

// Valid reports whether v's underlying value is non-nil.
func (v TestPrefsView) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v TestPrefsView) AsStruct() *TestPrefs {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v TestPrefsView) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v TestPrefsView) MarshalJSONV2(e *jsontext.Encoder, opt jsonexpv2.Options) error {
	return jsonexpv2.MarshalEncode(e, v.ж, opt)
}

func (v *TestPrefsView) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x TestPrefs
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v TestPrefsView) Int32Item() Item[int32]                   { return v.ж.Int32Item }
func (v TestPrefsView) UInt64Item() Item[uint64]                 { return v.ж.UInt64Item }
func (v TestPrefsView) StringItem1() Item[string]                { return v.ж.StringItem1 }
func (v TestPrefsView) StringItem2() Item[string]                { return v.ж.StringItem2 }
func (v TestPrefsView) BoolItem1() Item[bool]                    { return v.ж.BoolItem1 }
func (v TestPrefsView) BoolItem2() Item[bool]                    { return v.ж.BoolItem2 }
func (v TestPrefsView) StringSlice() ListView[string]            { return v.ж.StringSlice.View() }
func (v TestPrefsView) IntSlice() ListView[int]                  { return v.ж.IntSlice.View() }
func (v TestPrefsView) AddrItem() Item[netip.Addr]               { return v.ж.AddrItem }
func (v TestPrefsView) StringStringMap() MapView[string, string] { return v.ж.StringStringMap.View() }
func (v TestPrefsView) IntStringMap() MapView[int, string]       { return v.ж.IntStringMap.View() }
func (v TestPrefsView) AddrIntMap() MapView[netip.Addr, int]     { return v.ж.AddrIntMap.View() }
func (v TestPrefsView) Bundle1() ItemView[*TestBundle, TestBundleView] {
	return ItemViewOf(&v.ж.Bundle1)
}
func (v TestPrefsView) Bundle2() ItemView[*TestBundle, TestBundleView] {
	return ItemViewOf(&v.ж.Bundle2)
}
func (v TestPrefsView) Generic() ItemView[*TestGenericStruct[int], TestGenericStructView[int]] {
	return ItemViewOf(&v.ж.Generic)
}
func (v TestPrefsView) BundleList() StructListView[*TestBundle, TestBundleView] {
	return StructListViewOf(&v.ж.BundleList)
}
func (v TestPrefsView) StringBundleMap() StructMapView[string, *TestBundle, TestBundleView] {
	return StructMapViewOf(&v.ж.StringBundleMap)
}
func (v TestPrefsView) IntBundleMap() StructMapView[int, *TestBundle, TestBundleView] {
	return StructMapViewOf(&v.ж.IntBundleMap)
}
func (v TestPrefsView) AddrBundleMap() StructMapView[netip.Addr, *TestBundle, TestBundleView] {
	return StructMapViewOf(&v.ж.AddrBundleMap)
}
func (v TestPrefsView) Group() TestPrefsGroup { return v.ж.Group }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _TestPrefsViewNeedsRegeneration = TestPrefs(struct {
	Int32Item       Item[int32]
	UInt64Item      Item[uint64]
	StringItem1     Item[string]
	StringItem2     Item[string]
	BoolItem1       Item[bool]
	BoolItem2       Item[bool]
	StringSlice     List[string]
	IntSlice        List[int]
	AddrItem        Item[netip.Addr]
	StringStringMap Map[string, string]
	IntStringMap    Map[int, string]
	AddrIntMap      Map[netip.Addr, int]
	Bundle1         Item[*TestBundle]
	Bundle2         Item[*TestBundle]
	Generic         Item[*TestGenericStruct[int]]
	BundleList      StructList[*TestBundle]
	StringBundleMap StructMap[string, *TestBundle]
	IntBundleMap    StructMap[int, *TestBundle]
	AddrBundleMap   StructMap[netip.Addr, *TestBundle]
	Group           TestPrefsGroup
}{})

// View returns a read-only view of TestBundle.
func (p *TestBundle) View() TestBundleView {
	return TestBundleView{ж: p}
}

// TestBundleView provides a read-only view over TestBundle.
//
// Its methods should only be called if `Valid()` returns true.
type TestBundleView struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *TestBundle
}

// Valid reports whether v's underlying value is non-nil.
func (v TestBundleView) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v TestBundleView) AsStruct() *TestBundle {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v TestBundleView) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v TestBundleView) MarshalJSONV2(e *jsontext.Encoder, opt jsonexpv2.Options) error {
	return jsonexpv2.MarshalEncode(e, v.ж, opt)
}

func (v *TestBundleView) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x TestBundle
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v TestBundleView) Name() string                 { return v.ж.Name }
func (v TestBundleView) Nested() TestValueStructView  { return v.ж.Nested.View() }
func (v TestBundleView) Equal(v2 TestBundleView) bool { return v.ж.Equal(v2.ж) }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _TestBundleViewNeedsRegeneration = TestBundle(struct {
	Name   string
	Nested *TestValueStruct
}{})

// View returns a read-only view of TestValueStruct.
func (p *TestValueStruct) View() TestValueStructView {
	return TestValueStructView{ж: p}
}

// TestValueStructView provides a read-only view over TestValueStruct.
//
// Its methods should only be called if `Valid()` returns true.
type TestValueStructView struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *TestValueStruct
}

// Valid reports whether v's underlying value is non-nil.
func (v TestValueStructView) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v TestValueStructView) AsStruct() *TestValueStruct {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v TestValueStructView) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v TestValueStructView) MarshalJSONV2(e *jsontext.Encoder, opt jsonexpv2.Options) error {
	return jsonexpv2.MarshalEncode(e, v.ж, opt)
}

func (v *TestValueStructView) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x TestValueStruct
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v TestValueStructView) Value() int                        { return v.ж.Value }
func (v TestValueStructView) Equal(v2 TestValueStructView) bool { return v.ж.Equal(v2.ж) }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _TestValueStructViewNeedsRegeneration = TestValueStruct(struct {
	Value int
}{})

// View returns a read-only view of TestGenericStruct.
func (p *TestGenericStruct[T]) View() TestGenericStructView[T] {
	return TestGenericStructView[T]{ж: p}
}

// TestGenericStructView[T] provides a read-only view over TestGenericStruct[T].
//
// Its methods should only be called if `Valid()` returns true.
type TestGenericStructView[T ImmutableType] struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *TestGenericStruct[T]
}

// Valid reports whether v's underlying value is non-nil.
func (v TestGenericStructView[T]) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v TestGenericStructView[T]) AsStruct() *TestGenericStruct[T] {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v TestGenericStructView[T]) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v TestGenericStructView[T]) MarshalJSONV2(e *jsontext.Encoder, opt jsonexpv2.Options) error {
	return jsonexpv2.MarshalEncode(e, v.ж, opt)
}

func (v *TestGenericStructView[T]) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x TestGenericStruct[T]
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v TestGenericStructView[T]) Value() T                               { return v.ж.Value }
func (v TestGenericStructView[T]) Equal(v2 TestGenericStructView[T]) bool { return v.ж.Equal(v2.ж) }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
func _TestGenericStructViewNeedsRegeneration[T ImmutableType](TestGenericStruct[T]) {
	_TestGenericStructViewNeedsRegeneration(struct {
		Value T
	}{})
}

// View returns a read-only view of TestPrefsGroup.
func (p *TestPrefsGroup) View() TestPrefsGroupView {
	return TestPrefsGroupView{ж: p}
}

// TestPrefsGroupView provides a read-only view over TestPrefsGroup.
//
// Its methods should only be called if `Valid()` returns true.
type TestPrefsGroupView struct {
	// ж is the underlying mutable value, named with a hard-to-type
	// character that looks pointy like a pointer.
	// It is named distinctively to make you think of how dangerous it is to escape
	// to callers. You must not let callers be able to mutate it.
	ж *TestPrefsGroup
}

// Valid reports whether v's underlying value is non-nil.
func (v TestPrefsGroupView) Valid() bool { return v.ж != nil }

// AsStruct returns a clone of the underlying value which aliases no memory with
// the original.
func (v TestPrefsGroupView) AsStruct() *TestPrefsGroup {
	if v.ж == nil {
		return nil
	}
	return v.ж.Clone()
}

func (v TestPrefsGroupView) MarshalJSON() ([]byte, error) { return json.Marshal(v.ж) }

func (v TestPrefsGroupView) MarshalJSONV2(e *jsontext.Encoder, opt jsonexpv2.Options) error {
	return jsonexpv2.MarshalEncode(e, v.ж, opt)
}

func (v *TestPrefsGroupView) UnmarshalJSON(b []byte) error {
	if v.ж != nil {
		return errors.New("already initialized")
	}
	if len(b) == 0 {
		return nil
	}
	var x TestPrefsGroup
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	v.ж = &x
	return nil
}

func (v TestPrefsGroupView) FloatItem() Item[float64]             { return v.ж.FloatItem }
func (v TestPrefsGroupView) TestStringItem() Item[TestStringType] { return v.ж.TestStringItem }

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _TestPrefsGroupViewNeedsRegeneration = TestPrefsGroup(struct {
	FloatItem      Item[float64]
	TestStringItem Item[TestStringType]
}{})
