package stategen

import (
	"context"
	"testing"

	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/prysmaticlabs/go-ssz"
	testDB "github.com/prysmaticlabs/prysm/beacon-chain/db/testing"
	pb "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	"github.com/prysmaticlabs/prysm/shared/params"
	"github.com/prysmaticlabs/prysm/shared/testutil"
	logTest "github.com/sirupsen/logrus/hooks/test"
)

func TestMigrateToCold_NoBlock(t *testing.T) {
	hook := logTest.NewGlobal()
	ctx := context.Background()
	db := testDB.SetupDB(t)
	defer testDB.TeardownDB(t, db)

	service := New(db)

	beaconState, _ := testutil.DeterministicGenesisState(t, 32)
	beaconState.SetSlot(params.BeaconConfig().SlotsPerEpoch)
	if err := service.MigrateToCold(ctx, beaconState, [32]byte{}); err != nil {
		t.Fatal(err)
	}

	testutil.AssertLogsContain(t, hook, "Set hot and cold state split point")
}

func TestMigrateToCold_HigherSplitSlot(t *testing.T) {
	hook := logTest.NewGlobal()
	ctx := context.Background()
	db := testDB.SetupDB(t)
	defer testDB.TeardownDB(t, db)

	service := New(db)
	service.splitInfo.slot = 2

	beaconState, _ := testutil.DeterministicGenesisState(t, 32)
	beaconState.SetSlot(1)
	if err := service.MigrateToCold(ctx, beaconState, [32]byte{}); err != nil {
		t.Fatal(err)
	}

	testutil.AssertLogsDoNotContain(t, hook, "Set hot and cold state split point")
}

func TestMigrateToCold_NotEpochStart(t *testing.T) {
	hook := logTest.NewGlobal()
	ctx := context.Background()
	db := testDB.SetupDB(t)
	defer testDB.TeardownDB(t, db)

	service := New(db)

	beaconState, _ := testutil.DeterministicGenesisState(t, 32)
	beaconState.SetSlot(params.BeaconConfig().SlotsPerEpoch + 1)
	if err := service.MigrateToCold(ctx, beaconState, [32]byte{}); err != nil {
		t.Fatal(err)
	}

	testutil.AssertLogsDoNotContain(t, hook, "Set hot and cold state split point")
}

func TestMigrateToCold_MigrationCompletes(t *testing.T) {
	hook := logTest.NewGlobal()
	ctx := context.Background()
	db := testDB.SetupDB(t)
	defer testDB.TeardownDB(t, db)

	service := New(db)

	beaconState, _ := testutil.DeterministicGenesisState(t, 32)
	beaconState.SetSlot(params.BeaconConfig().SlotsPerEpoch)
	b := &ethpb.SignedBeaconBlock{
		Block: &ethpb.BeaconBlock{Slot: 2},
	}
	if err := service.beaconDB.SaveBlock(ctx, b); err != nil {
		t.Fatal(err)
	}
	bRoot, _ := ssz.HashTreeRoot(b.Block)
	if err := service.beaconDB.SaveStateSummary(ctx, &pb.StateSummary{Root: bRoot[:], Slot: 2}); err != nil {
		t.Fatal(err)
	}
	if err := service.beaconDB.SaveState(ctx, beaconState, bRoot); err != nil {
		t.Fatal(err)
	}
	service.slotsPerArchivedPoint = 2 // Ensure we can land on archived point.

	if err := service.MigrateToCold(ctx, beaconState, [32]byte{}); err != nil {
		t.Fatal(err)
	}

	if !service.beaconDB.HasArchivedPoint(ctx, 1) {
		t.Error("Did not preserve archived point")
	}

	testutil.AssertLogsContain(t, hook, "Saved archived point during state migration")
	testutil.AssertLogsContain(t, hook, "Deleted state during migration")
	testutil.AssertLogsContain(t, hook, "Set hot and cold state split point")
}
