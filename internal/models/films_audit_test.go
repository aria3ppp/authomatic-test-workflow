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

func testFilmsAudits(t *testing.T) {
	t.Parallel()

	query := FilmsAudits()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFilmsAuditsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FilmsAudit{}
	if err = randomize.Struct(seed, o, filmsAuditDBTypes, true, filmsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
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

	count, err := FilmsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFilmsAuditsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FilmsAudit{}
	if err = randomize.Struct(seed, o, filmsAuditDBTypes, true, filmsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := FilmsAudits().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := FilmsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFilmsAuditsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FilmsAudit{}
	if err = randomize.Struct(seed, o, filmsAuditDBTypes, true, filmsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FilmsAuditSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := FilmsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFilmsAuditsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FilmsAudit{}
	if err = randomize.Struct(seed, o, filmsAuditDBTypes, true, filmsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := FilmsAuditExists(ctx, tx, o.ID, o.ContributedBy, o.ContributedAt)
	if err != nil {
		t.Errorf("Unable to check if FilmsAudit exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FilmsAuditExists to return true, but got false.")
	}
}

func testFilmsAuditsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FilmsAudit{}
	if err = randomize.Struct(seed, o, filmsAuditDBTypes, true, filmsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	filmsAuditFound, err := FindFilmsAudit(ctx, tx, o.ID, o.ContributedBy, o.ContributedAt)
	if err != nil {
		t.Error(err)
	}

	if filmsAuditFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFilmsAuditsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FilmsAudit{}
	if err = randomize.Struct(seed, o, filmsAuditDBTypes, true, filmsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = FilmsAudits().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testFilmsAuditsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FilmsAudit{}
	if err = randomize.Struct(seed, o, filmsAuditDBTypes, true, filmsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := FilmsAudits().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFilmsAuditsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	filmsAuditOne := &FilmsAudit{}
	filmsAuditTwo := &FilmsAudit{}
	if err = randomize.Struct(seed, filmsAuditOne, filmsAuditDBTypes, false, filmsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}
	if err = randomize.Struct(seed, filmsAuditTwo, filmsAuditDBTypes, false, filmsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = filmsAuditOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = filmsAuditTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := FilmsAudits().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFilmsAuditsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	filmsAuditOne := &FilmsAudit{}
	filmsAuditTwo := &FilmsAudit{}
	if err = randomize.Struct(seed, filmsAuditOne, filmsAuditDBTypes, false, filmsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}
	if err = randomize.Struct(seed, filmsAuditTwo, filmsAuditDBTypes, false, filmsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = filmsAuditOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = filmsAuditTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FilmsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func filmsAuditBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *FilmsAudit) error {
	*o = FilmsAudit{}
	return nil
}

func filmsAuditAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *FilmsAudit) error {
	*o = FilmsAudit{}
	return nil
}

func filmsAuditAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *FilmsAudit) error {
	*o = FilmsAudit{}
	return nil
}

func filmsAuditBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *FilmsAudit) error {
	*o = FilmsAudit{}
	return nil
}

func filmsAuditAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *FilmsAudit) error {
	*o = FilmsAudit{}
	return nil
}

func filmsAuditBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *FilmsAudit) error {
	*o = FilmsAudit{}
	return nil
}

func filmsAuditAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *FilmsAudit) error {
	*o = FilmsAudit{}
	return nil
}

func filmsAuditBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *FilmsAudit) error {
	*o = FilmsAudit{}
	return nil
}

func filmsAuditAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *FilmsAudit) error {
	*o = FilmsAudit{}
	return nil
}

func testFilmsAuditsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &FilmsAudit{}
	o := &FilmsAudit{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, filmsAuditDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FilmsAudit object: %s", err)
	}

	AddFilmsAuditHook(boil.BeforeInsertHook, filmsAuditBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	filmsAuditBeforeInsertHooks = []FilmsAuditHook{}

	AddFilmsAuditHook(boil.AfterInsertHook, filmsAuditAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	filmsAuditAfterInsertHooks = []FilmsAuditHook{}

	AddFilmsAuditHook(boil.AfterSelectHook, filmsAuditAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	filmsAuditAfterSelectHooks = []FilmsAuditHook{}

	AddFilmsAuditHook(boil.BeforeUpdateHook, filmsAuditBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	filmsAuditBeforeUpdateHooks = []FilmsAuditHook{}

	AddFilmsAuditHook(boil.AfterUpdateHook, filmsAuditAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	filmsAuditAfterUpdateHooks = []FilmsAuditHook{}

	AddFilmsAuditHook(boil.BeforeDeleteHook, filmsAuditBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	filmsAuditBeforeDeleteHooks = []FilmsAuditHook{}

	AddFilmsAuditHook(boil.AfterDeleteHook, filmsAuditAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	filmsAuditAfterDeleteHooks = []FilmsAuditHook{}

	AddFilmsAuditHook(boil.BeforeUpsertHook, filmsAuditBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	filmsAuditBeforeUpsertHooks = []FilmsAuditHook{}

	AddFilmsAuditHook(boil.AfterUpsertHook, filmsAuditAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	filmsAuditAfterUpsertHooks = []FilmsAuditHook{}
}

func testFilmsAuditsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FilmsAudit{}
	if err = randomize.Struct(seed, o, filmsAuditDBTypes, true, filmsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FilmsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFilmsAuditsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FilmsAudit{}
	if err = randomize.Struct(seed, o, filmsAuditDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(filmsAuditColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := FilmsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFilmsAuditsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FilmsAudit{}
	if err = randomize.Struct(seed, o, filmsAuditDBTypes, true, filmsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
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

func testFilmsAuditsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FilmsAudit{}
	if err = randomize.Struct(seed, o, filmsAuditDBTypes, true, filmsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FilmsAuditSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFilmsAuditsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FilmsAudit{}
	if err = randomize.Struct(seed, o, filmsAuditDBTypes, true, filmsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := FilmsAudits().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	filmsAuditDBTypes = map[string]string{`ID`: `integer`, `Title`: `character varying`, `Descriptions`: `character varying`, `DateReleased`: `date`, `Duration`: `integer`, `SeriesID`: `integer`, `SeasonNumber`: `integer`, `EpisodeNumber`: `integer`, `ContributedBy`: `integer`, `ContributedAt`: `timestamp with time zone`, `Invalidation`: `character varying`}
	_                 = bytes.MinRead
)

func testFilmsAuditsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(filmsAuditPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(filmsAuditAllColumns) == len(filmsAuditPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &FilmsAudit{}
	if err = randomize.Struct(seed, o, filmsAuditDBTypes, true, filmsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FilmsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, filmsAuditDBTypes, true, filmsAuditPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testFilmsAuditsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(filmsAuditAllColumns) == len(filmsAuditPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &FilmsAudit{}
	if err = randomize.Struct(seed, o, filmsAuditDBTypes, true, filmsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FilmsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, filmsAuditDBTypes, true, filmsAuditPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(filmsAuditAllColumns, filmsAuditPrimaryKeyColumns) {
		fields = filmsAuditAllColumns
	} else {
		fields = strmangle.SetComplement(
			filmsAuditAllColumns,
			filmsAuditPrimaryKeyColumns,
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

	slice := FilmsAuditSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testFilmsAuditsUpsert(t *testing.T) {
	t.Parallel()

	if len(filmsAuditAllColumns) == len(filmsAuditPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := FilmsAudit{}
	if err = randomize.Struct(seed, &o, filmsAuditDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert FilmsAudit: %s", err)
	}

	count, err := FilmsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, filmsAuditDBTypes, false, filmsAuditPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FilmsAudit struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert FilmsAudit: %s", err)
	}

	count, err = FilmsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
