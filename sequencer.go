package astria_sequencer

import (
	"context"

	"github.com/rollkit/go-sequencing"
)

type Sequencer struct{}

var _ sequencing.Sequencer = &Sequencer{}

func (s *Sequencer) SubmitRollupTransaction(ctx context.Context, rollupId sequencing.RollupId, tx sequencing.Tx) error {
	//TODO implement me
	panic("implement me")
}

func (s *Sequencer) GetNextBatch(ctx context.Context, lastBatch *sequencing.Batch) (*sequencing.Batch, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Sequencer) VerifyBatch(ctx context.Context, batch *sequencing.Batch) (bool, error) {
	//TODO implement me
	panic("implement me")
}
