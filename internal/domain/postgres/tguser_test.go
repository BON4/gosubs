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

func testTgusers(t *testing.T) {
	t.Parallel()

	query := Tgusers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTgusersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tguser{}
	if err = randomize.Struct(seed, o, tguserDBTypes, true, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
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

	count, err := Tgusers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTgusersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tguser{}
	if err = randomize.Struct(seed, o, tguserDBTypes, true, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Tgusers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Tgusers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTgusersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tguser{}
	if err = randomize.Struct(seed, o, tguserDBTypes, true, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TguserSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Tgusers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTgusersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tguser{}
	if err = randomize.Struct(seed, o, tguserDBTypes, true, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TguserExists(ctx, tx, o.UserID)
	if err != nil {
		t.Errorf("Unable to check if Tguser exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TguserExists to return true, but got false.")
	}
}

func testTgusersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tguser{}
	if err = randomize.Struct(seed, o, tguserDBTypes, true, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	tguserFound, err := FindTguser(ctx, tx, o.UserID)
	if err != nil {
		t.Error(err)
	}

	if tguserFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTgusersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tguser{}
	if err = randomize.Struct(seed, o, tguserDBTypes, true, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Tgusers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTgusersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tguser{}
	if err = randomize.Struct(seed, o, tguserDBTypes, true, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Tgusers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTgusersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tguserOne := &Tguser{}
	tguserTwo := &Tguser{}
	if err = randomize.Struct(seed, tguserOne, tguserDBTypes, false, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}
	if err = randomize.Struct(seed, tguserTwo, tguserDBTypes, false, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = tguserOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = tguserTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Tgusers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTgusersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tguserOne := &Tguser{}
	tguserTwo := &Tguser{}
	if err = randomize.Struct(seed, tguserOne, tguserDBTypes, false, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}
	if err = randomize.Struct(seed, tguserTwo, tguserDBTypes, false, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = tguserOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = tguserTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tgusers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func tguserBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Tguser) error {
	*o = Tguser{}
	return nil
}

func tguserAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Tguser) error {
	*o = Tguser{}
	return nil
}

func tguserAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Tguser) error {
	*o = Tguser{}
	return nil
}

func tguserBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Tguser) error {
	*o = Tguser{}
	return nil
}

func tguserAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Tguser) error {
	*o = Tguser{}
	return nil
}

func tguserBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Tguser) error {
	*o = Tguser{}
	return nil
}

func tguserAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Tguser) error {
	*o = Tguser{}
	return nil
}

func tguserBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Tguser) error {
	*o = Tguser{}
	return nil
}

func tguserAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Tguser) error {
	*o = Tguser{}
	return nil
}

func testTgusersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Tguser{}
	o := &Tguser{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, tguserDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Tguser object: %s", err)
	}

	AddTguserHook(boil.BeforeInsertHook, tguserBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	tguserBeforeInsertHooks = []TguserHook{}

	AddTguserHook(boil.AfterInsertHook, tguserAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	tguserAfterInsertHooks = []TguserHook{}

	AddTguserHook(boil.AfterSelectHook, tguserAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	tguserAfterSelectHooks = []TguserHook{}

	AddTguserHook(boil.BeforeUpdateHook, tguserBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	tguserBeforeUpdateHooks = []TguserHook{}

	AddTguserHook(boil.AfterUpdateHook, tguserAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	tguserAfterUpdateHooks = []TguserHook{}

	AddTguserHook(boil.BeforeDeleteHook, tguserBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	tguserBeforeDeleteHooks = []TguserHook{}

	AddTguserHook(boil.AfterDeleteHook, tguserAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	tguserAfterDeleteHooks = []TguserHook{}

	AddTguserHook(boil.BeforeUpsertHook, tguserBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	tguserBeforeUpsertHooks = []TguserHook{}

	AddTguserHook(boil.AfterUpsertHook, tguserAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	tguserAfterUpsertHooks = []TguserHook{}
}

func testTgusersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tguser{}
	if err = randomize.Struct(seed, o, tguserDBTypes, true, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tgusers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTgusersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tguser{}
	if err = randomize.Struct(seed, o, tguserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(tguserColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Tgusers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTguserToManyUserSubs(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Tguser
	var b, c Sub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, tguserDBTypes, true, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, subDBTypes, false, subColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, subDBTypes, false, subColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.UserID = a.UserID
	c.UserID = a.UserID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UserSubs().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.UserID == b.UserID {
			bFound = true
		}
		if v.UserID == c.UserID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TguserSlice{&a}
	if err = a.L.LoadUserSubs(ctx, tx, false, (*[]*Tguser)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserSubs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserSubs = nil
	if err = a.L.LoadUserSubs(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserSubs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testTguserToManyUserSubHistories(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Tguser
	var b, c SubHistory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, tguserDBTypes, true, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, subHistoryDBTypes, false, subHistoryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, subHistoryDBTypes, false, subHistoryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.UserID = a.UserID
	c.UserID = a.UserID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UserSubHistories().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.UserID == b.UserID {
			bFound = true
		}
		if v.UserID == c.UserID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TguserSlice{&a}
	if err = a.L.LoadUserSubHistories(ctx, tx, false, (*[]*Tguser)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserSubHistories); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserSubHistories = nil
	if err = a.L.LoadUserSubHistories(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserSubHistories); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testTguserToManyAddOpUserSubs(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Tguser
	var b, c, d, e Sub

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, tguserDBTypes, false, strmangle.SetComplement(tguserPrimaryKeyColumns, tguserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Sub{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, subDBTypes, false, strmangle.SetComplement(subPrimaryKeyColumns, subColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Sub{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserSubs(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.UserID != first.UserID {
			t.Error("foreign key was wrong value", a.UserID, first.UserID)
		}
		if a.UserID != second.UserID {
			t.Error("foreign key was wrong value", a.UserID, second.UserID)
		}

		if first.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserSubs[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserSubs[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserSubs().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testTguserToManyAddOpUserSubHistories(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Tguser
	var b, c, d, e SubHistory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, tguserDBTypes, false, strmangle.SetComplement(tguserPrimaryKeyColumns, tguserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*SubHistory{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, subHistoryDBTypes, false, strmangle.SetComplement(subHistoryPrimaryKeyColumns, subHistoryColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*SubHistory{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserSubHistories(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.UserID != first.UserID {
			t.Error("foreign key was wrong value", a.UserID, first.UserID)
		}
		if a.UserID != second.UserID {
			t.Error("foreign key was wrong value", a.UserID, second.UserID)
		}

		if first.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserSubHistories[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserSubHistories[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserSubHistories().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testTgusersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tguser{}
	if err = randomize.Struct(seed, o, tguserDBTypes, true, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
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

func testTgusersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tguser{}
	if err = randomize.Struct(seed, o, tguserDBTypes, true, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TguserSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTgusersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Tguser{}
	if err = randomize.Struct(seed, o, tguserDBTypes, true, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Tgusers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	tguserDBTypes = map[string]string{`UserID`: `bigint`, `TelegramID`: `bigint`, `Username`: `text`, `Status`: `enum.user_status('creator','administrator','member','restricted','left','kicked')`}
	_             = bytes.MinRead
)

func testTgusersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(tguserPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(tguserAllColumns) == len(tguserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Tguser{}
	if err = randomize.Struct(seed, o, tguserDBTypes, true, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tgusers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, tguserDBTypes, true, tguserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTgusersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(tguserAllColumns) == len(tguserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Tguser{}
	if err = randomize.Struct(seed, o, tguserDBTypes, true, tguserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tgusers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, tguserDBTypes, true, tguserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(tguserAllColumns, tguserPrimaryKeyColumns) {
		fields = tguserAllColumns
	} else {
		fields = strmangle.SetComplement(
			tguserAllColumns,
			tguserPrimaryKeyColumns,
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

	slice := TguserSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTgusersUpsert(t *testing.T) {
	t.Parallel()

	if len(tguserAllColumns) == len(tguserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Tguser{}
	if err = randomize.Struct(seed, &o, tguserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Tguser: %s", err)
	}

	count, err := Tgusers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, tguserDBTypes, false, tguserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Tguser struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Tguser: %s", err)
	}

	count, err = Tgusers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
