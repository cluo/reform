package models

// Generated with gopkg.in/reform.v1. Do not edit by hand.

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type extraTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *extraTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("extra").
func (v *extraTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *extraTableType) Columns() []string {
	return []string{"id", "name", "byte", "uint8", "bytep", "uint8p", "bytes", "uint8s", "bytesa", "uint8sa", "bytest", "uint8st"}
}

// NewStruct makes a new struct for that view or table.
func (v *extraTableType) NewStruct() reform.Struct {
	return new(Extra)
}

// NewRecord makes a new record for that table.
func (v *extraTableType) NewRecord() reform.Record {
	return new(Extra)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *extraTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ExtraTable represents extra view or table in SQL database.
var ExtraTable = &extraTableType{
	s: parse.StructInfo{Type: "Extra", SQLSchema: "", SQLName: "extra", Fields: []parse.FieldInfo{{Name: "ID", Type: "Integer", Column: "id"}, {Name: "Name", Type: "*String", Column: "name"}, {Name: "Byte", Type: "uint8", Column: "byte"}, {Name: "Uint8", Type: "uint8", Column: "uint8"}, {Name: "ByteP", Type: "*uint8", Column: "bytep"}, {Name: "Uint8P", Type: "*uint8", Column: "uint8p"}, {Name: "Bytes", Type: "[]uint8", Column: "bytes"}, {Name: "Uint8s", Type: "[]uint8", Column: "uint8s"}, {Name: "BytesA", Type: "[512]uint8", Column: "bytesa"}, {Name: "Uint8sA", Type: "[512]uint8", Column: "uint8sa"}, {Name: "BytesT", Type: "Bytes", Column: "bytest"}, {Name: "Uint8sT", Type: "Uint8s", Column: "uint8st"}}, PKFieldIndex: 0},
	z: new(Extra).Values(),
}

// String returns a string representation of this struct or record.
func (s Extra) String() string {
	res := make([]string, 12)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Name: " + reform.Inspect(s.Name, true)
	res[2] = "Byte: " + reform.Inspect(s.Byte, true)
	res[3] = "Uint8: " + reform.Inspect(s.Uint8, true)
	res[4] = "ByteP: " + reform.Inspect(s.ByteP, true)
	res[5] = "Uint8P: " + reform.Inspect(s.Uint8P, true)
	res[6] = "Bytes: " + reform.Inspect(s.Bytes, true)
	res[7] = "Uint8s: " + reform.Inspect(s.Uint8s, true)
	res[8] = "BytesA: " + reform.Inspect(s.BytesA, true)
	res[9] = "Uint8sA: " + reform.Inspect(s.Uint8sA, true)
	res[10] = "BytesT: " + reform.Inspect(s.BytesT, true)
	res[11] = "Uint8sT: " + reform.Inspect(s.Uint8sT, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Extra) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Name,
		s.Byte,
		s.Uint8,
		s.ByteP,
		s.Uint8P,
		s.Bytes,
		s.Uint8s,
		s.BytesA,
		s.Uint8sA,
		s.BytesT,
		s.Uint8sT,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Extra) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Name,
		&s.Byte,
		&s.Uint8,
		&s.ByteP,
		&s.Uint8P,
		&s.Bytes,
		&s.Uint8s,
		&s.BytesA,
		&s.Uint8sA,
		&s.BytesT,
		&s.Uint8sT,
	}
}

// View returns View object for that struct.
func (s *Extra) View() reform.View {
	return ExtraTable
}

// Table returns Table object for that record.
func (s *Extra) Table() reform.Table {
	return ExtraTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Extra) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Extra) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Extra) HasPK() bool {
	return s.ID != ExtraTable.z[ExtraTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Extra) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = Integer(i64)
	} else {
		s.ID = pk.(Integer)
	}
}

// check interfaces
var (
	_ reform.View   = ExtraTable
	_ reform.Struct = (*Extra)(nil)
	_ reform.Table  = ExtraTable
	_ reform.Record = (*Extra)(nil)
	_ fmt.Stringer  = (*Extra)(nil)
)

type extra2TableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *extra2TableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("extra2").
func (v *extra2TableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *extra2TableType) Columns() []string {
	return []string{"id"}
}

// NewStruct makes a new struct for that view or table.
func (v *extra2TableType) NewStruct() reform.Struct {
	return new(Extra2)
}

// NewRecord makes a new record for that table.
func (v *extra2TableType) NewRecord() reform.Record {
	return new(Extra2)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *extra2TableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// Extra2Table represents extra2 view or table in SQL database.
var Extra2Table = &extra2TableType{
	s: parse.StructInfo{Type: "Extra2", SQLSchema: "", SQLName: "extra2", Fields: []parse.FieldInfo{{Name: "ID", Type: "[]uint8", Column: "id"}}, PKFieldIndex: 0},
	z: new(Extra2).Values(),
}

// String returns a string representation of this struct or record.
func (s Extra2) String() string {
	res := make([]string, 1)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Extra2) Values() []interface{} {
	return []interface{}{
		s.ID,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Extra2) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
	}
}

// View returns View object for that struct.
func (s *Extra2) View() reform.View {
	return Extra2Table
}

// Table returns Table object for that record.
func (s *Extra2) Table() reform.Table {
	return Extra2Table
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Extra2) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Extra2) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Extra2) HasPK() bool {
	return s.ID != Extra2Table.z[Extra2Table.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Extra2) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = []uint8(i64)
	} else {
		s.ID = pk.([]uint8)
	}
}

// check interfaces
var (
	_ reform.View   = Extra2Table
	_ reform.Struct = (*Extra2)(nil)
	_ reform.Table  = Extra2Table
	_ reform.Record = (*Extra2)(nil)
	_ fmt.Stringer  = (*Extra2)(nil)
)

type extra3TableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *extra3TableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("extra3").
func (v *extra3TableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *extra3TableType) Columns() []string {
	return []string{"id"}
}

// NewStruct makes a new struct for that view or table.
func (v *extra3TableType) NewStruct() reform.Struct {
	return new(Extra3)
}

// NewRecord makes a new record for that table.
func (v *extra3TableType) NewRecord() reform.Record {
	return new(Extra3)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *extra3TableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// Extra3Table represents extra3 view or table in SQL database.
var Extra3Table = &extra3TableType{
	s: parse.StructInfo{Type: "Extra3", SQLSchema: "", SQLName: "extra3", Fields: []parse.FieldInfo{{Name: "ID", Type: "[16]uint8", Column: "id"}}, PKFieldIndex: 0},
	z: new(Extra3).Values(),
}

// String returns a string representation of this struct or record.
func (s Extra3) String() string {
	res := make([]string, 1)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Extra3) Values() []interface{} {
	return []interface{}{
		s.ID,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Extra3) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
	}
}

// View returns View object for that struct.
func (s *Extra3) View() reform.View {
	return Extra3Table
}

// Table returns Table object for that record.
func (s *Extra3) Table() reform.Table {
	return Extra3Table
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Extra3) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Extra3) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Extra3) HasPK() bool {
	return s.ID != Extra3Table.z[Extra3Table.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Extra3) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = [16]uint8(i64)
	} else {
		s.ID = pk.([16]uint8)
	}
}

// check interfaces
var (
	_ reform.View   = Extra3Table
	_ reform.Struct = (*Extra3)(nil)
	_ reform.Table  = Extra3Table
	_ reform.Record = (*Extra3)(nil)
	_ fmt.Stringer  = (*Extra3)(nil)
)

type notExportedTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *notExportedTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("not_exported").
func (v *notExportedTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *notExportedTableType) Columns() []string {
	return []string{"id"}
}

// NewStruct makes a new struct for that view or table.
func (v *notExportedTableType) NewStruct() reform.Struct {
	return new(notExported)
}

// NewRecord makes a new record for that table.
func (v *notExportedTableType) NewRecord() reform.Record {
	return new(notExported)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *notExportedTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// notExportedTable represents not_exported view or table in SQL database.
var notExportedTable = &notExportedTableType{
	s: parse.StructInfo{Type: "notExported", SQLSchema: "", SQLName: "not_exported", Fields: []parse.FieldInfo{{Name: "ID", Type: "string", Column: "id"}}, PKFieldIndex: 0},
	z: new(notExported).Values(),
}

// String returns a string representation of this struct or record.
func (s notExported) String() string {
	res := make([]string, 1)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *notExported) Values() []interface{} {
	return []interface{}{
		s.ID,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *notExported) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
	}
}

// View returns View object for that struct.
func (s *notExported) View() reform.View {
	return notExportedTable
}

// Table returns Table object for that record.
func (s *notExported) Table() reform.Table {
	return notExportedTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *notExported) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *notExported) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *notExported) HasPK() bool {
	return s.ID != notExportedTable.z[notExportedTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *notExported) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = string(i64)
	} else {
		s.ID = pk.(string)
	}
}

// check interfaces
var (
	_ reform.View   = notExportedTable
	_ reform.Struct = (*notExported)(nil)
	_ reform.Table  = notExportedTable
	_ reform.Record = (*notExported)(nil)
	_ fmt.Stringer  = (*notExported)(nil)
)

func init() {
	parse.AssertUpToDate(&ExtraTable.s, new(Extra))
	parse.AssertUpToDate(&Extra2Table.s, new(Extra2))
	parse.AssertUpToDate(&Extra3Table.s, new(Extra3))
	parse.AssertUpToDate(&notExportedTable.s, new(notExported))
}
