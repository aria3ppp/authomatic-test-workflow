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

func testSerieses(t *testing.T) {
	t.Parallel()

	query := Serieses()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSeriesesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Series{}
	if err = randomize.Struct(seed, o, seriesDBTypes, true, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
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

	count, err := Serieses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSeriesesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Series{}
	if err = randomize.Struct(seed, o, seriesDBTypes, true, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Serieses().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Serieses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSeriesesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Series{}
	if err = randomize.Struct(seed, o, seriesDBTypes, true, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SeriesSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Serieses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSeriesesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Series{}
	if err = randomize.Struct(seed, o, seriesDBTypes, true, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SeriesExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Series exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SeriesExists to return true, but got false.")
	}
}

func testSeriesesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Series{}
	if err = randomize.Struct(seed, o, seriesDBTypes, true, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	seriesFound, err := FindSeries(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if seriesFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSeriesesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Series{}
	if err = randomize.Struct(seed, o, seriesDBTypes, true, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Serieses().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSeriesesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Series{}
	if err = randomize.Struct(seed, o, seriesDBTypes, true, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Serieses().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSeriesesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	seriesOne := &Series{}
	seriesTwo := &Series{}
	if err = randomize.Struct(seed, seriesOne, seriesDBTypes, false, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}
	if err = randomize.Struct(seed, seriesTwo, seriesDBTypes, false, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = seriesOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = seriesTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Serieses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSeriesesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	seriesOne := &Series{}
	seriesTwo := &Series{}
	if err = randomize.Struct(seed, seriesOne, seriesDBTypes, false, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}
	if err = randomize.Struct(seed, seriesTwo, seriesDBTypes, false, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = seriesOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = seriesTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Serieses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func seriesBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Series) error {
	*o = Series{}
	return nil
}

func seriesAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Series) error {
	*o = Series{}
	return nil
}

func seriesAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Series) error {
	*o = Series{}
	return nil
}

func seriesBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Series) error {
	*o = Series{}
	return nil
}

func seriesAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Series) error {
	*o = Series{}
	return nil
}

func seriesBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Series) error {
	*o = Series{}
	return nil
}

func seriesAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Series) error {
	*o = Series{}
	return nil
}

func seriesBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Series) error {
	*o = Series{}
	return nil
}

func seriesAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Series) error {
	*o = Series{}
	return nil
}

func testSeriesesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Series{}
	o := &Series{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, seriesDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Series object: %s", err)
	}

	AddSeriesHook(boil.BeforeInsertHook, seriesBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	seriesBeforeInsertHooks = []SeriesHook{}

	AddSeriesHook(boil.AfterInsertHook, seriesAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	seriesAfterInsertHooks = []SeriesHook{}

	AddSeriesHook(boil.AfterSelectHook, seriesAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	seriesAfterSelectHooks = []SeriesHook{}

	AddSeriesHook(boil.BeforeUpdateHook, seriesBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	seriesBeforeUpdateHooks = []SeriesHook{}

	AddSeriesHook(boil.AfterUpdateHook, seriesAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	seriesAfterUpdateHooks = []SeriesHook{}

	AddSeriesHook(boil.BeforeDeleteHook, seriesBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	seriesBeforeDeleteHooks = []SeriesHook{}

	AddSeriesHook(boil.AfterDeleteHook, seriesAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	seriesAfterDeleteHooks = []SeriesHook{}

	AddSeriesHook(boil.BeforeUpsertHook, seriesBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	seriesBeforeUpsertHooks = []SeriesHook{}

	AddSeriesHook(boil.AfterUpsertHook, seriesAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	seriesAfterUpsertHooks = []SeriesHook{}
}

func testSeriesesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Series{}
	if err = randomize.Struct(seed, o, seriesDBTypes, true, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Serieses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSeriesesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Series{}
	if err = randomize.Struct(seed, o, seriesDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(seriesColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Serieses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSeriesToManySeriesFilms(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Series
	var b, c Film

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, seriesDBTypes, true, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, filmDBTypes, false, filmColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, filmDBTypes, false, filmColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.SeriesID, a.ID)
	queries.Assign(&c.SeriesID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.SeriesFilms().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.SeriesID, b.SeriesID) {
			bFound = true
		}
		if queries.Equal(v.SeriesID, c.SeriesID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := SeriesSlice{&a}
	if err = a.L.LoadSeriesFilms(ctx, tx, false, (*[]*Series)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SeriesFilms); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.SeriesFilms = nil
	if err = a.L.LoadSeriesFilms(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SeriesFilms); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testSeriesToManyAddOpSeriesFilms(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Series
	var b, c, d, e Film

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, seriesDBTypes, false, strmangle.SetComplement(seriesPrimaryKeyColumns, seriesColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Film{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, filmDBTypes, false, strmangle.SetComplement(filmPrimaryKeyColumns, filmColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Film{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSeriesFilms(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.SeriesID) {
			t.Error("foreign key was wrong value", a.ID, first.SeriesID)
		}
		if !queries.Equal(a.ID, second.SeriesID) {
			t.Error("foreign key was wrong value", a.ID, second.SeriesID)
		}

		if first.R.Series != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Series != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.SeriesFilms[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.SeriesFilms[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.SeriesFilms().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testSeriesToManySetOpSeriesFilms(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Series
	var b, c, d, e Film

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, seriesDBTypes, false, strmangle.SetComplement(seriesPrimaryKeyColumns, seriesColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Film{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, filmDBTypes, false, strmangle.SetComplement(filmPrimaryKeyColumns, filmColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetSeriesFilms(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.SeriesFilms().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetSeriesFilms(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.SeriesFilms().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.SeriesID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.SeriesID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.SeriesID) {
		t.Error("foreign key was wrong value", a.ID, d.SeriesID)
	}
	if !queries.Equal(a.ID, e.SeriesID) {
		t.Error("foreign key was wrong value", a.ID, e.SeriesID)
	}

	if b.R.Series != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Series != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Series != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Series != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.SeriesFilms[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.SeriesFilms[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testSeriesToManyRemoveOpSeriesFilms(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Series
	var b, c, d, e Film

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, seriesDBTypes, false, strmangle.SetComplement(seriesPrimaryKeyColumns, seriesColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Film{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, filmDBTypes, false, strmangle.SetComplement(filmPrimaryKeyColumns, filmColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddSeriesFilms(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.SeriesFilms().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveSeriesFilms(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.SeriesFilms().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.SeriesID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.SeriesID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Series != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Series != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Series != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Series != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.SeriesFilms) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.SeriesFilms[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.SeriesFilms[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testSeriesToOneUserUsingContributingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Series
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, seriesDBTypes, false, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ContributedBy = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ContributingUser().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SeriesSlice{&local}
	if err = local.L.LoadContributingUser(ctx, tx, false, (*[]*Series)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ContributingUser == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ContributingUser = nil
	if err = local.L.LoadContributingUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ContributingUser == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSeriesToOneSetOpUserUsingContributingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Series
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, seriesDBTypes, false, strmangle.SetComplement(seriesPrimaryKeyColumns, seriesColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetContributingUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ContributingUser != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ContributedSerieses[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ContributedBy != x.ID {
			t.Error("foreign key was wrong value", a.ContributedBy)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ContributedBy))
		reflect.Indirect(reflect.ValueOf(&a.ContributedBy)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ContributedBy != x.ID {
			t.Error("foreign key was wrong value", a.ContributedBy, x.ID)
		}
	}
}

func testSeriesesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Series{}
	if err = randomize.Struct(seed, o, seriesDBTypes, true, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
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

func testSeriesesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Series{}
	if err = randomize.Struct(seed, o, seriesDBTypes, true, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SeriesSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSeriesesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Series{}
	if err = randomize.Struct(seed, o, seriesDBTypes, true, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Serieses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	seriesDBTypes = map[string]string{`ID`: `integer`, `Title`: `character varying`, `Descriptions`: `character varying`, `DateStarted`: `date`, `DateEnded`: `date`, `ContributedBy`: `integer`, `ContributedAt`: `timestamp with time zone`, `Invalidation`: `character varying`}
	_             = bytes.MinRead
)

func testSeriesesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(seriesPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(seriesAllColumns) == len(seriesPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Series{}
	if err = randomize.Struct(seed, o, seriesDBTypes, true, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Serieses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, seriesDBTypes, true, seriesPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSeriesesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(seriesAllColumns) == len(seriesPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Series{}
	if err = randomize.Struct(seed, o, seriesDBTypes, true, seriesColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Serieses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, seriesDBTypes, true, seriesPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(seriesAllColumns, seriesPrimaryKeyColumns) {
		fields = seriesAllColumns
	} else {
		fields = strmangle.SetComplement(
			seriesAllColumns,
			seriesPrimaryKeyColumns,
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

	slice := SeriesSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSeriesesUpsert(t *testing.T) {
	t.Parallel()

	if len(seriesAllColumns) == len(seriesPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Series{}
	if err = randomize.Struct(seed, &o, seriesDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Series: %s", err)
	}

	count, err := Serieses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, seriesDBTypes, false, seriesPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Series struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Series: %s", err)
	}

	count, err = Serieses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
