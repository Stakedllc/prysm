package kv

import (
	"context"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/shared/bytesutil"
	"github.com/prysmaticlabs/prysm/slasher/detection/attestations/types"
	log "github.com/sirupsen/logrus"
	"go.opencensus.io/trace"
)

var highestEpoch uint64

func saveToDB(db *Store) func(uint64, uint64, interface{}, int64) {
	// Returning the function here so we can access the DB properly from the OnEvict.
	return func(epoch uint64, _ uint64, value interface{}, cost int64) {
		log.Tracef("evicting span map for epoch: %d", epoch)
		err := db.update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket(validatorsMinMaxSpanBucket)
			epochBucket, err := bucket.CreateBucketIfNotExists(bytesutil.Bytes8(epoch))
			if err != nil {
				return err
			}
			spanMap := value.(map[uint64]types.Span)
			for k, v := range spanMap {
				err = epochBucket.Put(bytesutil.Bytes8(k), marshalSpan(v))
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			log.Errorf("failed to save span map to db on cache eviction: %v", err)
		}
	}
}

func unmarshalSpan(ctx context.Context, enc []byte) (types.Span, error) {
	ctx, span := trace.StartSpan(ctx, "SlasherDB.unmarshalSpan")
	defer span.End()
	r := types.Span{}
	if len(enc) != 7 {
		return r, errors.New("wrong data length for min max span")
	}
	r.MinSpan = bytesutil.FromBytes2(enc[:2])
	r.MaxSpan = bytesutil.FromBytes2(enc[2:4])
	sigB := [2]byte{}
	copy(sigB[:], enc[4:6])
	r.SigBytes = sigB
	r.HasAttested = bytesutil.ToBool(enc[6])
	return r, nil
}

func marshalSpan(span types.Span) []byte {
	return append(append(append(bytesutil.Bytes2(uint64(span.MinSpan)), bytesutil.Bytes2(uint64(span.MaxSpan))...), span.SigBytes[:]...), bytesutil.FromBool(span.HasAttested))
}

// EpochSpansMap accepts epoch and returns the corresponding spans map epoch=>spans
// for slashing detection.
// Returns nil if the span map for this validator index does not exist.
func (db *Store) EpochSpansMap(ctx context.Context, epoch uint64) (map[uint64]types.Span, error) {
	ctx, span := trace.StartSpan(ctx, "SlasherDB.EpochSpansMap")
	defer span.End()
	if db.spanCacheEnabled {
		v, ok := db.spanCache.Get(epoch)
		spanMap := make(map[uint64]types.Span)
		if ok {
			spanMap = v.(map[uint64]types.Span)
			return spanMap, nil
		}
	}
	var err error
	var spanMap map[uint64]types.Span
	err = db.view(func(tx *bolt.Tx) error {
		b := tx.Bucket(validatorsMinMaxSpanBucket)
		epochBucket := b.Bucket(bytesutil.Bytes8(epoch))
		if epochBucket == nil {
			return nil
		}
		keysLength := epochBucket.Stats().KeyN
		spanMap = make(map[uint64]types.Span, keysLength)
		return epochBucket.ForEach(func(k, v []byte) error {
			key := bytesutil.FromBytes8(k)
			value, err := unmarshalSpan(ctx, v)
			if err != nil {
				return err
			}
			spanMap[key] = value
			return nil
		})
	})
	if spanMap == nil {
		spanMap = make(map[uint64]types.Span)
	}
	return spanMap, err
}

// EpochSpanByValidatorIndex accepts validator index and epoch returns the corresponding spans
// for slashing detection.
// Returns error if the spans for this validator index and epoch does not exist.
func (db *Store) EpochSpanByValidatorIndex(ctx context.Context, validatorIdx uint64, epoch uint64) (types.Span, error) {
	ctx, span := trace.StartSpan(ctx, "SlasherDB.EpochSpanByValidatorIndex")
	defer span.End()
	var err error
	if db.spanCacheEnabled {
		v, ok := db.spanCache.Get(epoch)
		spanMap := make(map[uint64]types.Span)
		if ok {
			spanMap = v.(map[uint64]types.Span)
			spans, ok := spanMap[validatorIdx]
			if ok {
				return spans, nil
			}
		}
	}
	var spans types.Span
	err = db.view(func(tx *bolt.Tx) error {
		b := tx.Bucket(validatorsMinMaxSpanBucket)
		epochBucket := b.Bucket(bytesutil.Bytes8(epoch))
		if epochBucket == nil {
			return nil
		}
		key := bytesutil.Bytes8(validatorIdx)
		v := epochBucket.Get(key)
		if v == nil {
			return nil
		}
		value, err := unmarshalSpan(ctx, v)
		if err != nil {
			return err
		}
		spans = value
		return nil
	})
	return spans, err
}

// SaveValidatorEpochSpans accepts validator index epoch and spans returns.
// Returns error if the spans for this validator index and epoch does not exist.
func (db *Store) SaveValidatorEpochSpans(ctx context.Context, validatorIdx uint64, epoch uint64, spans types.Span) error {
	ctx, span := trace.StartSpan(ctx, "SlasherDB.SaveValidatorEpochSpans")
	defer span.End()
	defer span.End()
	if db.spanCacheEnabled {
		if epoch > highestEpoch {
			highestEpoch = epoch
		}
		v, ok := db.spanCache.Get(epoch)
		spanMap := make(map[uint64]types.Span)
		if ok {
			spanMap = v.(map[uint64]types.Span)
		}
		spanMap[validatorIdx] = spans
		saved := db.spanCache.Set(epoch, spanMap, 1)
		if !saved {
			return fmt.Errorf("failed to save span map to cache")
		}
		return nil
	}
	return db.update(func(tx *bolt.Tx) error {
		b := tx.Bucket(validatorsMinMaxSpanBucket)
		epochBucket, err := b.CreateBucketIfNotExists(bytesutil.Bytes8(epoch))
		if err != nil {
			return err
		}
		key := bytesutil.Bytes8(validatorIdx)
		value := marshalSpan(spans)
		return epochBucket.Put(key, value)
	})
}

// SaveEpochSpansMap accepts a epoch and span map epoch=>spans and writes it to disk.
func (db *Store) SaveEpochSpansMap(ctx context.Context, epoch uint64, spanMap map[uint64]types.Span) error {
	ctx, span := trace.StartSpan(ctx, "SlasherDB.SaveEpochSpansMap")
	defer span.End()
	if db.spanCacheEnabled {
		if epoch > highestEpoch {
			highestEpoch = epoch
		}
		saved := db.spanCache.Set(epoch, spanMap, 1)
		if !saved {
			return fmt.Errorf("failed to save span map to cache")
		}
		return nil
	}
	return db.update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(validatorsMinMaxSpanBucket)
		valBucket, err := bucket.CreateBucketIfNotExists(bytesutil.Bytes8(epoch))
		if err != nil {
			return err
		}
		for k, v := range spanMap {
			err = valBucket.Put(bytesutil.Bytes8(k), marshalSpan(v))
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (db *Store) enableSpanCache(enable bool) {
	db.spanCacheEnabled = enable
}

// SaveCachedSpansMaps saves all span map from cache to disk
// if no span maps are in db or cache is disabled it returns nil.
func (db *Store) SaveCachedSpansMaps(ctx context.Context) error {
	ctx, span := trace.StartSpan(ctx, "SlasherDB.SaveCachedSpansMaps")
	defer span.End()
	if db.spanCacheEnabled {
		db.enableSpanCache(false)
		defer db.enableSpanCache(true)
		for epoch := uint64(0); epoch <= highestEpoch; epoch++ {
			v, ok := db.spanCache.Get(epoch)
			if ok {
				spanMap := v.(map[uint64]types.Span)
				if err := db.SaveEpochSpansMap(ctx, epoch, spanMap); err != nil {
					return errors.Wrap(err, "failed to save span maps from cache")
				}

			}
		}
	}
	return nil
}

// DeleteEpochSpans deletes a epochs validators span map using a epoch index as bucket key.
func (db *Store) DeleteEpochSpans(ctx context.Context, epoch uint64) error {
	ctx, span := trace.StartSpan(ctx, "SlasherDB.DeleteEpochSpans")
	defer span.End()
	if db.spanCacheEnabled {

		_, ok := db.spanCache.Get(epoch)
		if ok {
			db.spanCache.Del(epoch)
			return nil
		}

	}
	return db.update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(validatorsMinMaxSpanBucket)
		key := bytesutil.Bytes8(epoch)
		return bucket.DeleteBucket(key)
	})
}

// DeleteValidatorSpanByEpoch deletes a validator span for a certain epoch
// using a validator index as bucket key.
func (db *Store) DeleteValidatorSpanByEpoch(ctx context.Context, validatorIdx uint64, epoch uint64) error {
	ctx, span := trace.StartSpan(ctx, "SlasherDB.DeleteValidatorSpanByEpoch")
	defer span.End()
	if db.spanCacheEnabled {
		v, ok := db.spanCache.Get(epoch)
		spanMap := make(map[uint64][2]uint16)
		if ok {
			spanMap = v.(map[uint64][2]uint16)
		}
		delete(spanMap, validatorIdx)
		saved := db.spanCache.Set(epoch, spanMap, 1)
		if !saved {
			return errors.New("failed to save span map to cache")
		}
		return nil
	}
	return db.update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(validatorsMinMaxSpanBucket)
		e := bytesutil.Bytes8(epoch)
		epochBucket := bucket.Bucket(e)
		v := bytesutil.Bytes8(validatorIdx)
		return epochBucket.Delete(v)
	})
}
