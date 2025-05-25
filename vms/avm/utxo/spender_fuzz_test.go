package utxo

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/vms/avm/fxs"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/nftfx"
	"github.com/ava-labs/avalanchego/vms/propertyfx"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

func FuzzSpenderSpendMint(f *testing.F) {
	f.Fuzz(func(t *testing.T, randSeed int64) {
		require := require.New(t)
		r := rand.New(rand.NewSource(randSeed)) // #nosec G404

		parser, err := txs.NewParser([]fxs.Fx{
			&secp256k1fx.Fx{},
			&nftfx.Fx{},
			&propertyfx.Fx{},
		})
		require.NoError(err)

		clock := &mockable.Clock{}
		s := NewSpender(clock, parser.Codec())

		numKeys := r.Intn(5)
		keys := make([]*secp256k1.PrivateKey, numKeys)
		for i := 0; i < numKeys; i++ {
			key, err := secp256k1.NewPrivateKey()
			require.NoError(err)
			keys[i] = key
		}
		kc := secp256k1fx.NewKeychain(keys...)

		numAssets := r.Intn(3) + 1
		assetIDs := make([]ids.ID, numAssets)
		for i := 0; i < numAssets; i++ {
			assetIDs[i] = ids.GenerateTestID()
		}

		amounts := make(map[ids.ID]uint64, numAssets)
		for _, assetID := range assetIDs {
			amounts[assetID] = r.Uint64()
		}

		numUTXOs := r.Intn(10)
		utxos := make([]*avax.UTXO, numUTXOs)
		for i := 0; i < numUTXOs; i++ {
			assetID := assetIDs[r.Intn(numAssets)]
			utxoID := avax.UTXOID{TxID: ids.GenerateTestID(), OutputIndex: uint32(r.Intn(10))}
			if r.Intn(2) == 0 {
				addrsLen := r.Intn(numKeys + 1)
				addrs := make([]ids.ShortID, addrsLen)
				for j := 0; j < addrsLen; j++ {
					if j < len(keys) {
						addrs[j] = keys[j].PublicKey().Address()
					} else {
						addrs[j] = ids.GenerateTestShortID()
					}
				}
				threshold := uint32(0)
				if addrsLen > 0 {
					threshold = uint32(r.Intn(addrsLen + 1))
				}
				out := &secp256k1fx.TransferOutput{
					Amt: r.Uint64(),
					OutputOwners: secp256k1fx.OutputOwners{
						Threshold: threshold,
						Addrs:     addrs,
					},
				}
				utxos[i] = &avax.UTXO{UTXOID: utxoID, Asset: avax.Asset{ID: assetID}, Out: out}
			} else {
				addrsLen := r.Intn(numKeys + 1)
				addrs := make([]ids.ShortID, addrsLen)
				for j := 0; j < addrsLen; j++ {
					if j < len(keys) {
						addrs[j] = keys[j].PublicKey().Address()
					} else {
						addrs[j] = ids.GenerateTestShortID()
					}
				}
				threshold := uint32(0)
				if addrsLen > 0 {
					threshold = uint32(r.Intn(addrsLen + 1))
				}
				out := &secp256k1fx.MintOutput{
					OutputOwners: secp256k1fx.OutputOwners{
						Threshold: threshold,
						Addrs:     addrs,
					},
				}
				utxos[i] = &avax.UTXO{UTXOID: utxoID, Asset: avax.Asset{ID: assetID}, Out: out}
			}
		}

		_, _, _, _ = s.Spend(utxos, kc, amounts)

		mintAmounts := make(map[ids.ID]uint64, len(amounts))
		for id, amt := range amounts {
			mintAmounts[id] = amt
		}
		_, _, _ = s.Mint(utxos, kc, mintAmounts, ids.GenerateTestShortID())
	})
}
