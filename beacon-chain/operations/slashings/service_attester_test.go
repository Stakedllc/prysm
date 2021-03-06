package slashings

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/gogo/protobuf/proto"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/helpers"
	"github.com/prysmaticlabs/prysm/shared/params"
	"github.com/prysmaticlabs/prysm/shared/testutil"
)

func attesterSlashingForValIdx(valIdx ...uint64) *ethpb.AttesterSlashing {
	return &ethpb.AttesterSlashing{
		Attestation_1: &ethpb.IndexedAttestation{
			AttestingIndices: valIdx,
		},
		Attestation_2: &ethpb.IndexedAttestation{
			AttestingIndices: valIdx,
		},
	}
}

func pendingSlashingForValIdx(valIdx ...uint64) *PendingAttesterSlashing {
	return &PendingAttesterSlashing{
		attesterSlashing: attesterSlashingForValIdx(valIdx...),
		validatorToSlash: valIdx[0],
	}
}

func TestPool_InsertAttesterSlashing(t *testing.T) {
	type fields struct {
		pending  []*PendingAttesterSlashing
		included map[uint64]bool
		wantErr  bool
		err      string
	}
	type args struct {
		slashings []*ethpb.AttesterSlashing
	}

	beaconState, privKeys := testutil.DeterministicGenesisState(t, 64)
	pendingSlashings := make([]*PendingAttesterSlashing, 20)
	slashings := make([]*ethpb.AttesterSlashing, 20)
	for i := 0; i < len(pendingSlashings); i++ {
		sl, err := testutil.GenerateAttesterSlashingForValidator(beaconState, privKeys[i], uint64(i))
		if err != nil {
			t.Fatal(err)
		}
		pendingSlashings[i] = &PendingAttesterSlashing{
			attesterSlashing: sl,
			validatorToSlash: uint64(i),
		}
		slashings[i] = sl
	}
	if err := beaconState.SetSlot(helpers.StartSlot(1)); err != nil {
		t.Fatal(err)
	}

	// We mark the following validators with some preconditions.
	exitedVal, _ := beaconState.ValidatorAtIndex(uint64(2))
	exitedVal.ExitEpoch = 0
	futureExitedVal, _ := beaconState.ValidatorAtIndex(uint64(4))
	futureExitedVal.ExitEpoch = 17
	slashedVal, _ := beaconState.ValidatorAtIndex(uint64(5))
	slashedVal.Slashed = true
	if err := beaconState.UpdateValidatorAtIndex(uint64(2), exitedVal); err != nil {
		t.Fatal(err)
	}
	if err := beaconState.UpdateValidatorAtIndex(uint64(4), futureExitedVal); err != nil {
		t.Fatal(err)
	}
	if err := beaconState.UpdateValidatorAtIndex(uint64(5), slashedVal); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*PendingAttesterSlashing
		err    string
	}{
		{
			name: "Empty list",
			fields: fields{
				pending:  make([]*PendingAttesterSlashing, 0),
				included: make(map[uint64]bool),
			},
			args: args{
				slashings: slashings[0:1],
			},
			want: []*PendingAttesterSlashing{
				{
					attesterSlashing: slashings[0],
					validatorToSlash: 0,
				},
			},
		},
		{
			name: "Empty list two validators slashed",
			fields: fields{
				pending:  make([]*PendingAttesterSlashing, 0),
				included: make(map[uint64]bool),
			},
			args: args{
				slashings: slashings[0:2],
			},
			want: pendingSlashings[0:2],
		},
		{
			name: "Duplicate identical slashing",
			fields: fields{
				pending: []*PendingAttesterSlashing{
					pendingSlashings[1],
				},
				included: make(map[uint64]bool),
			},
			args: args{
				slashings: slashings[1:2],
			},
			want: pendingSlashings[1:2],
		},
		{
			name: "Slashing for already slashed validator",
			fields: fields{
				pending:  []*PendingAttesterSlashing{},
				included: make(map[uint64]bool),
			},
			args: args{
				slashings: slashings[5:6],
			},
			want: []*PendingAttesterSlashing{},
		},
		{
			name: "Slashing for exited validator",
			fields: fields{
				pending:  []*PendingAttesterSlashing{},
				included: make(map[uint64]bool),
			},
			args: args{
				slashings: slashings[2:3],
			},
			want: []*PendingAttesterSlashing{},
		},
		{
			name: "Slashing for futuristic exited validator",
			fields: fields{
				pending:  []*PendingAttesterSlashing{},
				included: make(map[uint64]bool),
			},
			args: args{
				slashings: slashings[4:5],
			},
			want: pendingSlashings[4:5],
		},
		{
			name: "Already included",
			fields: fields{
				pending: []*PendingAttesterSlashing{},
				included: map[uint64]bool{
					1: true,
				},
			},
			args: args{
				slashings: slashings[1:2],
			},
			want: []*PendingAttesterSlashing{},
		},
		{
			name: "Maintains sorted order",
			fields: fields{
				pending: []*PendingAttesterSlashing{
					pendingSlashings[0],
					pendingSlashings[2],
				},
				included: make(map[uint64]bool),
			},
			args: args{
				slashings: slashings[1:2],
			},
			want: pendingSlashings[0:3],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pool{
				pendingAttesterSlashing: tt.fields.pending,
				included:                tt.fields.included,
			}
			var err error
			for i := 0; i < len(tt.args.slashings); i++ {
				err = p.InsertAttesterSlashing(context.Background(), beaconState, tt.args.slashings[i])
			}
			if err != nil && tt.fields.wantErr && !strings.Contains(err.Error(), tt.fields.err) {
				t.Fatalf("Wanted err: %v, received %v", tt.fields.err, err)
			}
			if len(p.pendingAttesterSlashing) != len(tt.want) {
				t.Fatalf(
					"Mismatched lengths of pending list. Got %d, wanted %d.",
					len(p.pendingAttesterSlashing),
					len(tt.want),
				)
			}
			for i := range p.pendingAttesterSlashing {
				if p.pendingAttesterSlashing[i].validatorToSlash != tt.want[i].validatorToSlash {
					t.Errorf(
						"Pending attester to slash at index %d does not match expected. Got=%v wanted=%v",
						i,
						p.pendingAttesterSlashing[i].validatorToSlash,
						tt.want[i].validatorToSlash,
					)
				}
				if !proto.Equal(p.pendingAttesterSlashing[i].attesterSlashing, tt.want[i].attesterSlashing) {
					t.Errorf(
						"Pending attester slashing at index %d does not match expected. Got=%v wanted=%v",
						i,
						p.pendingAttesterSlashing[i],
						tt.want[i],
					)
				}
			}
		})
	}
}

func TestPool_InsertAttesterSlashing_SigFailsVerify_ClearPool(t *testing.T) {
	conf := params.BeaconConfig()
	conf.MaxAttesterSlashings = 2
	params.OverrideBeaconConfig(conf)
	beaconState, privKeys := testutil.DeterministicGenesisState(t, 64)
	pendingSlashings := make([]*PendingAttesterSlashing, 2)
	slashings := make([]*ethpb.AttesterSlashing, 2)
	for i := 0; i < 2; i++ {
		sl, err := testutil.GenerateAttesterSlashingForValidator(beaconState, privKeys[i], uint64(i))
		if err != nil {
			t.Fatal(err)
		}
		pendingSlashings[i] = &PendingAttesterSlashing{
			attesterSlashing: sl,
			validatorToSlash: uint64(i),
		}
		slashings[i] = sl
	}
	// We mess up the signature of the second slashing.
	badSig := make([]byte, 96)
	copy(badSig, "muahaha")
	pendingSlashings[1].attesterSlashing.Attestation_1.Signature = badSig
	slashings[1].Attestation_1.Signature = badSig
	p := &Pool{
		pendingAttesterSlashing: make([]*PendingAttesterSlashing, 0),
	}
	if err := p.InsertAttesterSlashing(
		context.Background(),
		beaconState,
		slashings[0],
	); err != nil {
		t.Fatal(err)
	}
	if err := p.InsertAttesterSlashing(
		context.Background(),
		beaconState,
		slashings[1],
	); err == nil {
		t.Error("Expected error when inserting slashing with bad sig, got nil")
	}
	// We expect to only have 1 pending attester slashing in the pool.
	if len(p.pendingAttesterSlashing) != 1 {
		t.Error("Expected failed attester slashing to have been cleared from pool")
	}
}

func TestPool_MarkIncludedAttesterSlashing(t *testing.T) {
	type fields struct {
		pending  []*PendingAttesterSlashing
		included map[uint64]bool
	}
	type args struct {
		slashing *ethpb.AttesterSlashing
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{
			name: "Included, does not exist in pending",
			fields: fields{
				pending: []*PendingAttesterSlashing{
					{
						attesterSlashing: attesterSlashingForValIdx(1),
						validatorToSlash: 1,
					},
				},
				included: make(map[uint64]bool),
			},
			args: args{
				slashing: attesterSlashingForValIdx(3),
			},
			want: fields{
				pending: []*PendingAttesterSlashing{
					pendingSlashingForValIdx(1),
				},
				included: map[uint64]bool{
					3: true,
				},
			},
		},
		{
			name: "Removes from pending list",
			fields: fields{
				pending: []*PendingAttesterSlashing{
					pendingSlashingForValIdx(1),
					pendingSlashingForValIdx(2),
					pendingSlashingForValIdx(3),
				},
				included: map[uint64]bool{
					0: true,
				},
			},
			args: args{
				slashing: attesterSlashingForValIdx(2),
			},
			want: fields{
				pending: []*PendingAttesterSlashing{
					pendingSlashingForValIdx(1),
					pendingSlashingForValIdx(3),
				},
				included: map[uint64]bool{
					0: true,
					2: true,
				},
			},
		},
		{
			name: "Removes from long pending list",
			fields: fields{
				pending: []*PendingAttesterSlashing{
					pendingSlashingForValIdx(1),
					pendingSlashingForValIdx(2),
					pendingSlashingForValIdx(3),
					pendingSlashingForValIdx(4),
					pendingSlashingForValIdx(5),
					pendingSlashingForValIdx(6),
					pendingSlashingForValIdx(7),
					pendingSlashingForValIdx(8),
					pendingSlashingForValIdx(9),
					pendingSlashingForValIdx(10),
					pendingSlashingForValIdx(11),
				},
				included: map[uint64]bool{
					0: true,
				},
			},
			args: args{
				slashing: attesterSlashingForValIdx(6),
			},
			want: fields{
				pending: []*PendingAttesterSlashing{
					pendingSlashingForValIdx(1),
					pendingSlashingForValIdx(2),
					pendingSlashingForValIdx(3),
					pendingSlashingForValIdx(4),
					pendingSlashingForValIdx(5),
					pendingSlashingForValIdx(7),
					pendingSlashingForValIdx(8),
					pendingSlashingForValIdx(9),
					pendingSlashingForValIdx(10),
					pendingSlashingForValIdx(11),
				},
				included: map[uint64]bool{
					0: true,
					6: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pool{
				pendingAttesterSlashing: tt.fields.pending,
				included:                tt.fields.included,
			}
			p.MarkIncludedAttesterSlashing(tt.args.slashing)
			if len(p.pendingAttesterSlashing) != len(tt.want.pending) {
				t.Fatalf(
					"Mismatched lengths of pending list. Got %d, wanted %d.",
					len(p.pendingAttesterSlashing),
					len(tt.want.pending),
				)
			}
			for i := range p.pendingAttesterSlashing {
				if !reflect.DeepEqual(p.pendingAttesterSlashing[i], tt.want.pending[i]) {
					t.Errorf(
						"Pending attester slashing at index %d does not match expected. Got=%v wanted=%v",
						i,
						p.pendingAttesterSlashing[i],
						tt.want.pending[i],
					)
				}
			}
			if !reflect.DeepEqual(p.included, tt.want.included) {
				t.Errorf("Included map is not as expected. Got=%v wanted=%v", p.included, tt.want.included)
			}
		})
	}
}

func TestPool_PendingAttesterSlashings(t *testing.T) {
	type fields struct {
		pending []*PendingAttesterSlashing
	}
	conf := params.BeaconConfig()
	conf.MaxAttesterSlashings = 1
	params.OverrideBeaconConfig(conf)
	beaconState, privKeys := testutil.DeterministicGenesisState(t, 64)
	pendingSlashings := make([]*PendingAttesterSlashing, 20)
	slashings := make([]*ethpb.AttesterSlashing, 20)
	for i := 0; i < len(pendingSlashings); i++ {
		sl, err := testutil.GenerateAttesterSlashingForValidator(beaconState, privKeys[i], uint64(i))
		if err != nil {
			t.Fatal(err)
		}
		pendingSlashings[i] = &PendingAttesterSlashing{
			attesterSlashing: sl,
			validatorToSlash: uint64(i),
		}
		slashings[i] = sl
	}
	tests := []struct {
		name   string
		fields fields
		want   []*ethpb.AttesterSlashing
	}{
		{
			name: "Empty list",
			fields: fields{
				pending: []*PendingAttesterSlashing{},
			},
			want: []*ethpb.AttesterSlashing{},
		},
		{
			name: "All eligible",
			fields: fields{
				pending: pendingSlashings,
			},
			want: slashings[0:1],
		},
		{
			name: "Multiple indices",
			fields: fields{
				pending: pendingSlashings[3:6],
			},
			want: slashings[3:4],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pool{
				pendingAttesterSlashing: tt.fields.pending,
			}
			if got := p.PendingAttesterSlashings(
				context.Background(),
			); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Unexpected return from PendingAttesterSlashings, wanted %v, received %v", tt.want, got)
			}
		})
	}
}

func TestPool_PendingAttesterSlashings_NoDuplicates(t *testing.T) {
	conf := params.BeaconConfig()
	conf.MaxAttesterSlashings = 2
	params.OverrideBeaconConfig(conf)
	beaconState, privKeys := testutil.DeterministicGenesisState(t, 64)
	pendingSlashings := make([]*PendingAttesterSlashing, 3)
	slashings := make([]*ethpb.AttesterSlashing, 3)
	for i := 0; i < 2; i++ {
		sl, err := testutil.GenerateAttesterSlashingForValidator(beaconState, privKeys[i], uint64(i))
		if err != nil {
			t.Fatal(err)
		}
		pendingSlashings[i] = &PendingAttesterSlashing{
			attesterSlashing: sl,
			validatorToSlash: uint64(i),
		}
		slashings[i] = sl
	}
	// We duplicate the last slashing.
	pendingSlashings[2] = pendingSlashings[1]
	slashings[2] = slashings[1]
	p := &Pool{
		pendingAttesterSlashing: pendingSlashings,
	}
	want := slashings[0:2]
	if got := p.PendingAttesterSlashings(
		context.Background(),
	); !reflect.DeepEqual(want, got) {
		t.Errorf("Unexpected return from PendingAttesterSlashings, wanted %v, received %v", want, got)
	}
}
