package txs

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/utils/hashing"
)

func FuzzParserParseTx(f *testing.F) {
	parser, err := NewParser(nil)
	if err != nil {
		f.Fatalf("failed to create parser: %v", err)
	}

	f.Fuzz(func(t *testing.T, txBytes []byte) {
		tx, err := parser.ParseTx(txBytes)
		if err != nil {
			return
		}

		require := require.New(t)

		require.Equal(txBytes, tx.Bytes())
		require.Equal(len(txBytes), tx.Size())

		marshalled, err := parser.Codec().Marshal(CodecVersion, tx)
		require.NoError(err)
		require.Equal(txBytes, marshalled)

		size, err := parser.Codec().Size(CodecVersion, tx)
		require.NoError(err)
		require.Equal(len(txBytes), size)

		unsignedSize, err := parser.Codec().Size(CodecVersion, &tx.Unsigned)
		require.NoError(err)
		require.Equal(txBytes[:unsignedSize], tx.Unsigned.Bytes())

		expectedID := hashing.ComputeHash256Array(txBytes)
		require.Equal(expectedID, tx.ID())

		tx2, err := parser.ParseTx(marshalled)
		require.NoError(err)
		require.Equal(tx.ID(), tx2.ID())
		require.Equal(tx.Bytes(), tx2.Bytes())
		require.Equal(tx.Unsigned.Bytes(), tx2.Unsigned.Bytes())
	})
}
