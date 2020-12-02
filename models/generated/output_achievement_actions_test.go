// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package generated

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testOutputAchievementActions(t *testing.T) {
	t.Parallel()

	query := OutputAchievementActions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOutputAchievementActionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievementAction{}
	if err = randomize.Struct(seed, o, outputAchievementActionDBTypes, true, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
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

	count, err := OutputAchievementActions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOutputAchievementActionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievementAction{}
	if err = randomize.Struct(seed, o, outputAchievementActionDBTypes, true, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OutputAchievementActions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OutputAchievementActions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOutputAchievementActionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievementAction{}
	if err = randomize.Struct(seed, o, outputAchievementActionDBTypes, true, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OutputAchievementActionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OutputAchievementActions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOutputAchievementActionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievementAction{}
	if err = randomize.Struct(seed, o, outputAchievementActionDBTypes, true, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OutputAchievementActionExists(ctx, tx, o.OutputAchievementActionID)
	if err != nil {
		t.Errorf("Unable to check if OutputAchievementAction exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OutputAchievementActionExists to return true, but got false.")
	}
}

func testOutputAchievementActionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievementAction{}
	if err = randomize.Struct(seed, o, outputAchievementActionDBTypes, true, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	outputAchievementActionFound, err := FindOutputAchievementAction(ctx, tx, o.OutputAchievementActionID)
	if err != nil {
		t.Error(err)
	}

	if outputAchievementActionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOutputAchievementActionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievementAction{}
	if err = randomize.Struct(seed, o, outputAchievementActionDBTypes, true, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OutputAchievementActions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOutputAchievementActionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievementAction{}
	if err = randomize.Struct(seed, o, outputAchievementActionDBTypes, true, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OutputAchievementActions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOutputAchievementActionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	outputAchievementActionOne := &OutputAchievementAction{}
	outputAchievementActionTwo := &OutputAchievementAction{}
	if err = randomize.Struct(seed, outputAchievementActionOne, outputAchievementActionDBTypes, false, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}
	if err = randomize.Struct(seed, outputAchievementActionTwo, outputAchievementActionDBTypes, false, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = outputAchievementActionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = outputAchievementActionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OutputAchievementActions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOutputAchievementActionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	outputAchievementActionOne := &OutputAchievementAction{}
	outputAchievementActionTwo := &OutputAchievementAction{}
	if err = randomize.Struct(seed, outputAchievementActionOne, outputAchievementActionDBTypes, false, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}
	if err = randomize.Struct(seed, outputAchievementActionTwo, outputAchievementActionDBTypes, false, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = outputAchievementActionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = outputAchievementActionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OutputAchievementActions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func outputAchievementActionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievementAction) error {
	*o = OutputAchievementAction{}
	return nil
}

func outputAchievementActionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievementAction) error {
	*o = OutputAchievementAction{}
	return nil
}

func outputAchievementActionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievementAction) error {
	*o = OutputAchievementAction{}
	return nil
}

func outputAchievementActionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievementAction) error {
	*o = OutputAchievementAction{}
	return nil
}

func outputAchievementActionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievementAction) error {
	*o = OutputAchievementAction{}
	return nil
}

func outputAchievementActionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievementAction) error {
	*o = OutputAchievementAction{}
	return nil
}

func outputAchievementActionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievementAction) error {
	*o = OutputAchievementAction{}
	return nil
}

func outputAchievementActionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievementAction) error {
	*o = OutputAchievementAction{}
	return nil
}

func outputAchievementActionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievementAction) error {
	*o = OutputAchievementAction{}
	return nil
}

func testOutputAchievementActionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OutputAchievementAction{}
	o := &OutputAchievementAction{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, outputAchievementActionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction object: %s", err)
	}

	AddOutputAchievementActionHook(boil.BeforeInsertHook, outputAchievementActionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	outputAchievementActionBeforeInsertHooks = []OutputAchievementActionHook{}

	AddOutputAchievementActionHook(boil.AfterInsertHook, outputAchievementActionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	outputAchievementActionAfterInsertHooks = []OutputAchievementActionHook{}

	AddOutputAchievementActionHook(boil.AfterSelectHook, outputAchievementActionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	outputAchievementActionAfterSelectHooks = []OutputAchievementActionHook{}

	AddOutputAchievementActionHook(boil.BeforeUpdateHook, outputAchievementActionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	outputAchievementActionBeforeUpdateHooks = []OutputAchievementActionHook{}

	AddOutputAchievementActionHook(boil.AfterUpdateHook, outputAchievementActionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	outputAchievementActionAfterUpdateHooks = []OutputAchievementActionHook{}

	AddOutputAchievementActionHook(boil.BeforeDeleteHook, outputAchievementActionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	outputAchievementActionBeforeDeleteHooks = []OutputAchievementActionHook{}

	AddOutputAchievementActionHook(boil.AfterDeleteHook, outputAchievementActionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	outputAchievementActionAfterDeleteHooks = []OutputAchievementActionHook{}

	AddOutputAchievementActionHook(boil.BeforeUpsertHook, outputAchievementActionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	outputAchievementActionBeforeUpsertHooks = []OutputAchievementActionHook{}

	AddOutputAchievementActionHook(boil.AfterUpsertHook, outputAchievementActionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	outputAchievementActionAfterUpsertHooks = []OutputAchievementActionHook{}
}

func testOutputAchievementActionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievementAction{}
	if err = randomize.Struct(seed, o, outputAchievementActionDBTypes, true, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OutputAchievementActions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOutputAchievementActionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievementAction{}
	if err = randomize.Struct(seed, o, outputAchievementActionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(outputAchievementActionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OutputAchievementActions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOutputAchievementActionToOneOutputAchievementUsingOutputAchievement(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OutputAchievementAction
	var foreign OutputAchievement

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, outputAchievementActionDBTypes, false, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, outputAchievementDBTypes, false, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.OutputAchievementID = foreign.OutputAchievementID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.OutputAchievement().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.OutputAchievementID != foreign.OutputAchievementID {
		t.Errorf("want: %v, got %v", foreign.OutputAchievementID, check.OutputAchievementID)
	}

	slice := OutputAchievementActionSlice{&local}
	if err = local.L.LoadOutputAchievement(ctx, tx, false, (*[]*OutputAchievementAction)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.OutputAchievement == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.OutputAchievement = nil
	if err = local.L.LoadOutputAchievement(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.OutputAchievement == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOutputAchievementActionToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OutputAchievementAction
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, outputAchievementActionDBTypes, false, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
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

	slice := OutputAchievementActionSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*OutputAchievementAction)(&slice), nil); err != nil {
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

func testOutputAchievementActionToOneSetOpOutputAchievementUsingOutputAchievement(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OutputAchievementAction
	var b, c OutputAchievement

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, outputAchievementActionDBTypes, false, strmangle.SetComplement(outputAchievementActionPrimaryKeyColumns, outputAchievementActionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, outputAchievementDBTypes, false, strmangle.SetComplement(outputAchievementPrimaryKeyColumns, outputAchievementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, outputAchievementDBTypes, false, strmangle.SetComplement(outputAchievementPrimaryKeyColumns, outputAchievementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*OutputAchievement{&b, &c} {
		err = a.SetOutputAchievement(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.OutputAchievement != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OutputAchievementActions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OutputAchievementID != x.OutputAchievementID {
			t.Error("foreign key was wrong value", a.OutputAchievementID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OutputAchievementID))
		reflect.Indirect(reflect.ValueOf(&a.OutputAchievementID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OutputAchievementID != x.OutputAchievementID {
			t.Error("foreign key was wrong value", a.OutputAchievementID, x.OutputAchievementID)
		}
	}
}
func testOutputAchievementActionToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OutputAchievementAction
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, outputAchievementActionDBTypes, false, strmangle.SetComplement(outputAchievementActionPrimaryKeyColumns, outputAchievementActionColumnsWithoutDefault)...); err != nil {
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
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OutputAchievementActions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.UserID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.UserID {
			t.Error("foreign key was wrong value", a.UserID, x.UserID)
		}
	}
}

func testOutputAchievementActionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievementAction{}
	if err = randomize.Struct(seed, o, outputAchievementActionDBTypes, true, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
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

func testOutputAchievementActionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievementAction{}
	if err = randomize.Struct(seed, o, outputAchievementActionDBTypes, true, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OutputAchievementActionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOutputAchievementActionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievementAction{}
	if err = randomize.Struct(seed, o, outputAchievementActionDBTypes, true, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OutputAchievementActions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	outputAchievementActionDBTypes = map[string]string{`OutputAchievementActionID`: `int`, `OutputAchievementID`: `int`, `UserID`: `int`, `ActionType`: `varchar`, `CreatedBy`: `int`, `CreatedAt`: `timestamp`, `ModifiedBy`: `int`, `ModifiedAt`: `timestamp`}
	_                              = bytes.MinRead
)

func testOutputAchievementActionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(outputAchievementActionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(outputAchievementActionAllColumns) == len(outputAchievementActionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievementAction{}
	if err = randomize.Struct(seed, o, outputAchievementActionDBTypes, true, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OutputAchievementActions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, outputAchievementActionDBTypes, true, outputAchievementActionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOutputAchievementActionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(outputAchievementActionAllColumns) == len(outputAchievementActionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievementAction{}
	if err = randomize.Struct(seed, o, outputAchievementActionDBTypes, true, outputAchievementActionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OutputAchievementActions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, outputAchievementActionDBTypes, true, outputAchievementActionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(outputAchievementActionAllColumns, outputAchievementActionPrimaryKeyColumns) {
		fields = outputAchievementActionAllColumns
	} else {
		fields = strmangle.SetComplement(
			outputAchievementActionAllColumns,
			outputAchievementActionPrimaryKeyColumns,
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

	slice := OutputAchievementActionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOutputAchievementActionsUpsert(t *testing.T) {
	t.Parallel()

	if len(outputAchievementActionAllColumns) == len(outputAchievementActionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLOutputAchievementActionUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OutputAchievementAction{}
	if err = randomize.Struct(seed, &o, outputAchievementActionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OutputAchievementAction: %s", err)
	}

	count, err := OutputAchievementActions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, outputAchievementActionDBTypes, false, outputAchievementActionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OutputAchievementAction struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OutputAchievementAction: %s", err)
	}

	count, err = OutputAchievementActions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
