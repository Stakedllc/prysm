package kv

import (
	"context"
	"flag"
	"reflect"
	"testing"
	"time"

	"github.com/prysmaticlabs/prysm/slasher/detection/attestations/types"
	"github.com/prysmaticlabs/prysm/slasher/flags"
	"github.com/urfave/cli"
)

type spanMapTestStruct struct {
	epoch   uint64
	spanMap map[uint64]types.Span
}

var spanTests []spanMapTestStruct

func init() {
	spanTests = []spanMapTestStruct{
		{
			epoch: 1,
			spanMap: map[uint64]types.Span{
				1: {MinSpan: 10, MaxSpan: 20, HasAttested: false, SigBytes: [2]byte{1, 1}},
				2: {MinSpan: 11, MaxSpan: 21, HasAttested: true, SigBytes: [2]byte{1, 1}},
				3: {MinSpan: 12, MaxSpan: 22, HasAttested: false, SigBytes: [2]byte{1, 1}},
			},
		},
		{
			epoch: 2,
			spanMap: map[uint64]types.Span{
				1: {MinSpan: 10, MaxSpan: 20, HasAttested: false, SigBytes: [2]byte{1, 1}},
				2: {MinSpan: 11, MaxSpan: 21, HasAttested: true, SigBytes: [2]byte{1, 1}},
				3: {MinSpan: 12, MaxSpan: 22, HasAttested: true, SigBytes: [2]byte{1, 1}},
			},
		},
		{
			epoch: 3,
			spanMap: map[uint64]types.Span{
				1: {MinSpan: 10, MaxSpan: 20, HasAttested: true, SigBytes: [2]byte{1, 1}},
				2: {MinSpan: 11, MaxSpan: 21, SigBytes: [2]byte{1, 1}},
				3: {MinSpan: 12, MaxSpan: 22, SigBytes: [2]byte{1, 1}},
			},
		},
	}
}

func TestValidatorSpanMap_NilDB(t *testing.T) {
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	db := setupDB(t, cli.NewContext(app, set, nil))
	defer teardownDB(t, db)
	ctx := context.Background()

	validatorIdx := uint64(1)
	vsm, err := db.EpochSpansMap(ctx, validatorIdx)
	if err != nil {
		t.Fatalf("Nil EpochSpansMap should not return error: %v", err)
	}
	if !reflect.DeepEqual(vsm, map[uint64]types.Span{}) {
		t.Fatal("EpochSpansMap should return nil")
	}
}

func TestStore_SaveSpans(t *testing.T) {
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	db := setupDB(t, cli.NewContext(app, set, nil))
	defer teardownDB(t, db)
	ctx := context.Background()

	for _, tt := range spanTests {
		err := db.SaveEpochSpansMap(ctx, tt.epoch, tt.spanMap)
		if err != nil {
			t.Fatalf("Save validator span map failed: %v", err)
		}
		sm, err := db.EpochSpansMap(ctx, tt.epoch)
		if err != nil {
			t.Fatalf("Failed to get validator span map: %v", err)
		}

		if sm == nil || !reflect.DeepEqual(sm, tt.spanMap) {
			t.Fatalf("Get should return validator span map: %v got: %v", tt.spanMap, sm)
		}
		s, err := db.EpochSpanByValidatorIndex(ctx, 1, tt.epoch)
		if err != nil {
			t.Fatalf("Failed to get validator span for epoch 1: %v", err)
		}
		if !reflect.DeepEqual(s, tt.spanMap[1]) {
			t.Fatalf("Get should return validator spans for epoch 1: %v got: %v", tt.spanMap[1], s)
		}
	}
}

func TestStore_SaveCachedSpans(t *testing.T) {
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	set.Bool(flags.UseSpanCacheFlag.Name, true, "enable span map cache")
	db := setupDB(t, cli.NewContext(app, set, nil))
	defer teardownDB(t, db)
	ctx := context.Background()

	for _, tt := range spanTests {
		err := db.SaveEpochSpansMap(ctx, tt.epoch, tt.spanMap)
		if err != nil {
			t.Fatalf("Save validator span map failed: %v", err)
		}
		// wait for value to pass through cache buffers
		time.Sleep(time.Millisecond * 10)
		sm, err := db.EpochSpansMap(ctx, tt.epoch)
		if err != nil {
			t.Fatalf("Failed to get validator span map: %v", err)
		}

		if sm == nil || !reflect.DeepEqual(sm, tt.spanMap) {
			t.Fatalf("Get should return validator span map: %v got: %v", tt.spanMap, sm)
		}
		s, err := db.EpochSpanByValidatorIndex(ctx, 1, tt.epoch)
		if err != nil {
			t.Fatalf("Failed to get validator span for epoch 1: %v", err)
		}
		if !reflect.DeepEqual(s, tt.spanMap[1]) {
			t.Fatalf("Get should return validator spans for epoch 1: %v got: %v", tt.spanMap[1], s)
		}
	}
}

func TestStore_DeleteEpochSpans(t *testing.T) {
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	db := setupDB(t, cli.NewContext(app, set, nil))
	defer teardownDB(t, db)
	ctx := context.Background()

	for _, tt := range spanTests {
		err := db.SaveEpochSpansMap(ctx, tt.epoch, tt.spanMap)
		if err != nil {
			t.Fatalf("Save validator span map failed: %v", err)
		}
	}

	for _, tt := range spanTests {
		sm, err := db.EpochSpansMap(ctx, tt.epoch)
		if err != nil {
			t.Fatalf("Failed to get validator span map: %v", err)
		}
		if sm == nil || !reflect.DeepEqual(sm, tt.spanMap) {
			t.Fatalf("Get should return validator span map: %v got: %v", tt.spanMap, sm)
		}
		err = db.DeleteEpochSpans(ctx, tt.epoch)
		if err != nil {
			t.Fatalf("Delete validator span map error: %v", err)
		}
		sm, err = db.EpochSpansMap(ctx, tt.epoch)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(sm, map[uint64]types.Span{}) {
			t.Errorf("Expected validator span map to be deleted, received: %v", sm)
		}
	}
}

func TestValidatorSpanMap_DeletesOnCacheSavesToDB(t *testing.T) {
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	set.Bool(flags.UseSpanCacheFlag.Name, true, "enable span map cache")
	db := setupDB(t, cli.NewContext(app, set, nil))
	defer teardownDB(t, db)
	ctx := context.Background()

	for _, tt := range spanTests {
		err := db.SaveEpochSpansMap(ctx, tt.epoch, tt.spanMap)
		if err != nil {
			t.Fatalf("Save validator span map failed: %v", err)
		}
	}
	// Wait for value to pass through cache buffers.
	time.Sleep(time.Millisecond * 10)
	for _, tt := range spanTests {
		spanMap, err := db.EpochSpansMap(ctx, tt.epoch)
		if err != nil {
			t.Fatalf("Failed to get validator span map: %v", err)
		}
		if spanMap == nil || !reflect.DeepEqual(spanMap, tt.spanMap) {
			t.Fatalf("Get should return validator span map: %v got: %v", tt.spanMap, spanMap)
		}

		if err = db.DeleteEpochSpans(ctx, tt.epoch); err != nil {
			t.Fatalf("Delete validator span map error: %v", err)
		}
		// Wait for value to pass through cache buffers.
		db.enableSpanCache(false)
		time.Sleep(time.Millisecond * 10)
		spanMap, err = db.EpochSpansMap(ctx, tt.epoch)
		if err != nil {
			t.Fatal(err)
		}
		db.enableSpanCache(true)
		if !reflect.DeepEqual(spanMap, tt.spanMap) {
			t.Errorf("Expected validator span map to be deleted, received: %v", spanMap)
		}
	}
}

func TestValidatorSpanMap_SaveOnEvict(t *testing.T) {
	db := setupDBDiffCacheSize(t, 5)
	defer teardownDB(t, db)
	ctx := context.Background()

	tsm := &spanMapTestStruct{
		epoch: 1,
		spanMap: map[uint64]types.Span{
			1: {MinSpan: 10, MaxSpan: 20, SigBytes: [2]byte{0, 1}},
			2: {MinSpan: 11, MaxSpan: 21, HasAttested: true},
			3: {MinSpan: 12, MaxSpan: 22},
		},
	}
	for i := uint64(0); i < 6; i++ {
		err := db.SaveEpochSpansMap(ctx, i, tsm.spanMap)
		if err != nil {
			t.Fatalf("Save validator span map failed: %v", err)
		}
	}

	// Wait for value to pass through cache buffers.
	time.Sleep(time.Millisecond * 1000)
	for i := uint64(0); i < 6; i++ {
		sm, err := db.EpochSpansMap(ctx, i)
		if err != nil {
			t.Fatalf("Failed to get validator span map: %v", err)
		}
		if sm == nil || !reflect.DeepEqual(sm, tsm.spanMap) {
			t.Fatalf("Get should return validator: %d span map: %v got: %v", i, tsm.spanMap, sm)
		}
	}
}

func TestValidatorSpanMap_SaveCachedSpansMaps(t *testing.T) {
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	set.Bool(flags.UseSpanCacheFlag.Name, true, "enable span map cache")
	db := setupDB(t, cli.NewContext(app, set, nil))
	defer teardownDB(t, db)
	ctx := context.Background()

	for _, tt := range spanTests {
		err := db.SaveEpochSpansMap(ctx, tt.epoch, tt.spanMap)
		if err != nil {
			t.Fatalf("Save validator span map failed: %v", err)
		}
	}
	// wait for value to pass through cache buffers
	time.Sleep(time.Millisecond * 10)
	if err := db.SaveCachedSpansMaps(ctx); err != nil {
		t.Errorf("Failed to save cached span maps to db: %v", err)
	}
	db.spanCache.Clear()
	for _, tt := range spanTests {
		sm, err := db.EpochSpansMap(ctx, tt.epoch)
		if err != nil {
			t.Fatalf("Failed to get validator span map: %v", err)
		}
		if !reflect.DeepEqual(sm, tt.spanMap) {
			t.Fatalf("Get should return validator span map: %v got: %v", tt.spanMap, sm)
		}
	}
}
