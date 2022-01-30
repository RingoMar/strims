// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
// so force a package dependency in case they don't.
var _ = queries.Equal

func testNodes(t *testing.T) {
	t.Parallel()

	query := Nodes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNodesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNodesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Nodes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNodesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NodeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNodesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := NodeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Node exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NodeExists to return true, but got false.")
	}
}

func testNodesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	nodeFound, err := FindNode(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if nodeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNodesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Nodes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testNodesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Nodes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNodesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	nodeOne := &Node{}
	nodeTwo := &Node{}
	if err = randomize.Struct(seed, nodeOne, nodeDBTypes, false, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}
	if err = randomize.Struct(seed, nodeTwo, nodeDBTypes, false, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = nodeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = nodeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Nodes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNodesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	nodeOne := &Node{}
	nodeTwo := &Node{}
	if err = randomize.Struct(seed, nodeOne, nodeDBTypes, false, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}
	if err = randomize.Struct(seed, nodeTwo, nodeDBTypes, false, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = nodeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = nodeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testNodesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNodesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(nodeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNodesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNodesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NodeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNodesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Nodes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	nodeDBTypes = map[string]string{`ID`: `bigint`, `Active`: `boolean`, `StartedAt`: `bigint`, `StoppedAt`: `bigint`, `ProviderName`: `text`, `ProviderID`: `text`, `Name`: `text`, `Memory`: `integer`, `CPUs`: `integer`, `Disk`: `integer`, `IPV4`: `text`, `IPV6`: `text`, `RegionName`: `text`, `RegionLat`: `double precision`, `RegionLng`: `double precision`, `SKUName`: `text`, `SKUNetworkCap`: `integer`, `SKUNetworkSpeed`: `integer`, `SKUPriceMonthly`: `real`, `SKUPriceHourly`: `real`, `WireguardKey`: `text`, `WireguardIP`: `text`, `User`: `text`, `Type`: `enum.node_type('controller','worker')`}
	_           = bytes.MinRead
)

func testNodesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(nodePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(nodeAllColumns) == len(nodePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testNodesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(nodeAllColumns) == len(nodePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Node{}
	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, nodeDBTypes, true, nodePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(nodeAllColumns, nodePrimaryKeyColumns) {
		fields = nodeAllColumns
	} else {
		fields = strmangle.SetComplement(
			nodeAllColumns,
			nodePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := NodeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testNodesUpsert(t *testing.T) {
	t.Parallel()

	if len(nodeAllColumns) == len(nodePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Node{}
	if err = randomize.Struct(seed, &o, nodeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Node: %s", err)
	}

	count, err := Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, nodeDBTypes, false, nodePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Node struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Node: %s", err)
	}

	count, err = Nodes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
