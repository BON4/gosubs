// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testSubs(t *testing.T) {
	t.Parallel()

	query := Subs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSubsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sub{}
	if err = randomize.Struct(seed, o, subDBTypes, true, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
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

	count, err := Subs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sub{}
	if err = randomize.Struct(seed, o, subDBTypes, true, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Subs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Subs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sub{}
	if err = randomize.Struct(seed, o, subDBTypes, true, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SubSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Subs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sub{}
	if err = randomize.Struct(seed, o, subDBTypes, true, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SubExists(ctx, tx, o.UserID, o.CreatorID)
	if err != nil {
		t.Errorf("Unable to check if Sub exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SubExists to return true, but got false.")
	}
}

func testSubsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sub{}
	if err = randomize.Struct(seed, o, subDBTypes, true, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	subFound, err := FindSub(ctx, tx, o.UserID, o.CreatorID)
	if err != nil {
		t.Error(err)
	}

	if subFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSubsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sub{}
	if err = randomize.Struct(seed, o, subDBTypes, true, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Subs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSubsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sub{}
	if err = randomize.Struct(seed, o, subDBTypes, true, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Subs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSubsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	subOne := &Sub{}
	subTwo := &Sub{}
	if err = randomize.Struct(seed, subOne, subDBTypes, false, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}
	if err = randomize.Struct(seed, subTwo, subDBTypes, false, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = subOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = subTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Subs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSubsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	subOne := &Sub{}
	subTwo := &Sub{}
	if err = randomize.Struct(seed, subOne, subDBTypes, false, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}
	if err = randomize.Struct(seed, subTwo, subDBTypes, false, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = subOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = subTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Subs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func subBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Sub) error {
	*o = Sub{}
	return nil
}

func subAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Sub) error {
	*o = Sub{}
	return nil
}

func subAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Sub) error {
	*o = Sub{}
	return nil
}

func subBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Sub) error {
	*o = Sub{}
	return nil
}

func subAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Sub) error {
	*o = Sub{}
	return nil
}

func subBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Sub) error {
	*o = Sub{}
	return nil
}

func subAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Sub) error {
	*o = Sub{}
	return nil
}

func subBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Sub) error {
	*o = Sub{}
	return nil
}

func subAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Sub) error {
	*o = Sub{}
	return nil
}

func testSubsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Sub{}
	o := &Sub{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, subDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Sub object: %s", err)
	}

	AddSubHook(boil.BeforeInsertHook, subBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	subBeforeInsertHooks = []SubHook{}

	AddSubHook(boil.AfterInsertHook, subAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	subAfterInsertHooks = []SubHook{}

	AddSubHook(boil.AfterSelectHook, subAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	subAfterSelectHooks = []SubHook{}

	AddSubHook(boil.BeforeUpdateHook, subBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	subBeforeUpdateHooks = []SubHook{}

	AddSubHook(boil.AfterUpdateHook, subAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	subAfterUpdateHooks = []SubHook{}

	AddSubHook(boil.BeforeDeleteHook, subBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	subBeforeDeleteHooks = []SubHook{}

	AddSubHook(boil.AfterDeleteHook, subAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	subAfterDeleteHooks = []SubHook{}

	AddSubHook(boil.BeforeUpsertHook, subBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	subBeforeUpsertHooks = []SubHook{}

	AddSubHook(boil.AfterUpsertHook, subAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	subAfterUpsertHooks = []SubHook{}
}

func testSubsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sub{}
	if err = randomize.Struct(seed, o, subDBTypes, true, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Subs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSubsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sub{}
	if err = randomize.Struct(seed, o, subDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(subColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Subs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSubToOneCreatorUsingCreator(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Sub
	var foreign Creator

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, subDBTypes, false, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, creatorDBTypes, false, creatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Creator struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.CreatorID = foreign.CreatorID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Creator().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.CreatorID != foreign.CreatorID {
		t.Errorf("want: %v, got %v", foreign.CreatorID, check.CreatorID)
	}

	slice := SubSlice{&local}
	if err = local.L.LoadCreator(ctx, tx, false, (*[]*Sub)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Creator == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Creator = nil
	if err = local.L.LoadCreator(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Creator == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSubToOneTguserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Sub
	var foreign Tguser

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, subDBTypes, false, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, tguserDBTypes, false, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.UserID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.UserID != foreign.UserID {
		t.Errorf("want: %v, got %v", foreign.UserID, check.UserID)
	}

	slice := SubSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Sub)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSubToOneSetOpCreatorUsingCreator(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Sub
	var b, c Creator

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, subDBTypes, false, strmangle.SetComplement(subPrimaryKeyColumns, subColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, creatorDBTypes, false, strmangle.SetComplement(creatorPrimaryKeyColumns, creatorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, creatorDBTypes, false, strmangle.SetComplement(creatorPrimaryKeyColumns, creatorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Creator{&b, &c} {
		err = a.SetCreator(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Creator != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Subs[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CreatorID != x.CreatorID {
			t.Error("foreign key was wrong value", a.CreatorID)
		}

		if exists, err := SubExists(ctx, tx, a.UserID, a.CreatorID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testSubToOneSetOpTguserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Sub
	var b, c Tguser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, subDBTypes, false, strmangle.SetComplement(subPrimaryKeyColumns, subColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, tguserDBTypes, false, strmangle.SetComplement(tguserPrimaryKeyColumns, tguserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, tguserDBTypes, false, strmangle.SetComplement(tguserPrimaryKeyColumns, tguserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Tguser{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserSubs[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.UserID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		if exists, err := SubExists(ctx, tx, a.UserID, a.CreatorID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testSubsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sub{}
	if err = randomize.Struct(seed, o, subDBTypes, true, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
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

func testSubsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sub{}
	if err = randomize.Struct(seed, o, subDBTypes, true, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SubSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSubsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sub{}
	if err = randomize.Struct(seed, o, subDBTypes, true, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Subs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	subDBTypes = map[string]string{`UserID`: `bigint`, `CreatorID`: `bigint`, `ActivatedAt`: `timestamp with time zone`, `ExpiresAt`: `timestamp with time zone`, `Status`: `enum.sub_status('expired','active','cancelled','inactive')`, `Price`: `integer`}
	_          = bytes.MinRead
)

func testSubsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(subPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(subAllColumns) == len(subPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Sub{}
	if err = randomize.Struct(seed, o, subDBTypes, true, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Subs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, subDBTypes, true, subPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSubsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(subAllColumns) == len(subPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Sub{}
	if err = randomize.Struct(seed, o, subDBTypes, true, subColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Subs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, subDBTypes, true, subPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(subAllColumns, subPrimaryKeyColumns) {
		fields = subAllColumns
	} else {
		fields = strmangle.SetComplement(
			subAllColumns,
			subPrimaryKeyColumns,
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

	slice := SubSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSubsUpsert(t *testing.T) {
	t.Parallel()

	if len(subAllColumns) == len(subPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Sub{}
	if err = randomize.Struct(seed, &o, subDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Sub: %s", err)
	}

	count, err := Subs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, subDBTypes, false, subPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Sub struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Sub: %s", err)
	}

	count, err = Subs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}