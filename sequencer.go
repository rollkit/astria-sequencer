package astria_sequencer

import (
	"context"
	"fmt"

	astriaproto "buf.build/gen/go/astria/primitives/protocolbuffers/go/astria/primitive/v1"
	txproto "buf.build/gen/go/astria/protocol-apis/protocolbuffers/go/astria/protocol/transactions/v1alpha1"
	astria "github.com/astriaorg/astria-cli-go/modules/go-sequencer-client/client"

	"github.com/rollkit/go-sequencing"
)

type Sequencer struct {
	client *astria.Client
	signer *astria.Signer
}

var _ sequencing.Sequencer = &Sequencer{}

func New(url string) (*Sequencer, error) {
	signer, err := astria.GenerateSigner()
	if err != nil {
		return nil, fmt.Errorf("failed to create Astria signer: %w", err)
	}

	client, err := astria.NewClient(url)
	if err != nil {
		return nil, fmt.Errorf("failed to create Astria API client: %w", err)
	}

	return &Sequencer{
		client: client,
		signer: signer,
	}, nil
}

func (s *Sequencer) SubmitRollupTransaction(ctx context.Context, rollupId sequencing.RollupId, tx sequencing.Tx) error {
	unsignedTx := &txproto.UnsignedTransaction{
		Actions: []*txproto.Action{
			{
				Value: &txproto.Action_SequenceAction{
					SequenceAction: &txproto.SequenceAction{
						RollupId: &astriaproto.RollupId{Inner: rollupId[:]},
						Data:     []byte("test-data"),
						// TODO(tzdybal): what about `FeeAsset`?
					},
				},
			},
		},
	}

	signed, err := s.signer.SignTransaction(unsignedTx)
	if err != nil {
		panic(err)
	}

	_, err = s.client.BroadcastTxAsync(ctx, signed)

	return err
}

func (s *Sequencer) GetNextBatch(ctx context.Context, lastBatch *sequencing.Batch) (*sequencing.Batch, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Sequencer) VerifyBatch(ctx context.Context, batch *sequencing.Batch) (bool, error) {
	//TODO implement me
	panic("implement me")
}
