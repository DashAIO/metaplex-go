// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package gumdrop

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
)

type MerkleDistributor struct {
	Base     ag_solanago.PublicKey
	Bump     uint8
	Root     [32]uint8
	Temporal ag_solanago.PublicKey
}

var MerkleDistributorDiscriminator = [8]byte{77, 119, 139, 70, 84, 247, 12, 26}

func (obj MerkleDistributor) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(MerkleDistributorDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Base` param:
	err = encoder.Encode(obj.Base)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `Root` param:
	err = encoder.Encode(obj.Root)
	if err != nil {
		return err
	}
	// Serialize `Temporal` param:
	err = encoder.Encode(obj.Temporal)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MerkleDistributor) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(MerkleDistributorDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[77 119 139 70 84 247 12 26]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Base`:
	err = decoder.Decode(&obj.Base)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `Root`:
	err = decoder.Decode(&obj.Root)
	if err != nil {
		return err
	}
	// Deserialize `Temporal`:
	err = decoder.Decode(&obj.Temporal)
	if err != nil {
		return err
	}
	return nil
}

type ClaimStatus struct {
	IsClaimed bool
	Claimant  ag_solanago.PublicKey
	ClaimedAt int64
	Amount    uint64
}

var ClaimStatusDiscriminator = [8]byte{22, 183, 249, 157, 247, 95, 150, 96}

func (obj ClaimStatus) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ClaimStatusDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `IsClaimed` param:
	err = encoder.Encode(obj.IsClaimed)
	if err != nil {
		return err
	}
	// Serialize `Claimant` param:
	err = encoder.Encode(obj.Claimant)
	if err != nil {
		return err
	}
	// Serialize `ClaimedAt` param:
	err = encoder.Encode(obj.ClaimedAt)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ClaimStatus) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ClaimStatusDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[22 183 249 157 247 95 150 96]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `IsClaimed`:
	err = decoder.Decode(&obj.IsClaimed)
	if err != nil {
		return err
	}
	// Deserialize `Claimant`:
	err = decoder.Decode(&obj.Claimant)
	if err != nil {
		return err
	}
	// Deserialize `ClaimedAt`:
	err = decoder.Decode(&obj.ClaimedAt)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

type ClaimCount struct {
	Count    uint64
	Claimant ag_solanago.PublicKey
}

var ClaimCountDiscriminator = [8]byte{78, 134, 220, 213, 34, 152, 102, 167}

func (obj ClaimCount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ClaimCountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Count` param:
	err = encoder.Encode(obj.Count)
	if err != nil {
		return err
	}
	// Serialize `Claimant` param:
	err = encoder.Encode(obj.Claimant)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ClaimCount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ClaimCountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[78 134 220 213 34 152 102 167]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Count`:
	err = decoder.Decode(&obj.Count)
	if err != nil {
		return err
	}
	// Deserialize `Claimant`:
	err = decoder.Decode(&obj.Claimant)
	if err != nil {
		return err
	}
	return nil
}

type ClaimProof struct {
	Amount        uint64
	Count         uint64
	Claimant      ag_solanago.PublicKey
	Resource      ag_solanago.PublicKey
	ResourceNonce []byte
}

var ClaimProofDiscriminator = [8]byte{48, 173, 176, 137, 53, 116, 40, 112}

func (obj ClaimProof) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ClaimProofDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	// Serialize `Count` param:
	err = encoder.Encode(obj.Count)
	if err != nil {
		return err
	}
	// Serialize `Claimant` param:
	err = encoder.Encode(obj.Claimant)
	if err != nil {
		return err
	}
	// Serialize `Resource` param:
	err = encoder.Encode(obj.Resource)
	if err != nil {
		return err
	}
	// Serialize `ResourceNonce` param:
	err = encoder.Encode(obj.ResourceNonce)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ClaimProof) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ClaimProofDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[48 173 176 137 53 116 40 112]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	// Deserialize `Count`:
	err = decoder.Decode(&obj.Count)
	if err != nil {
		return err
	}
	// Deserialize `Claimant`:
	err = decoder.Decode(&obj.Claimant)
	if err != nil {
		return err
	}
	// Deserialize `Resource`:
	err = decoder.Decode(&obj.Resource)
	if err != nil {
		return err
	}
	// Deserialize `ResourceNonce`:
	err = decoder.Decode(&obj.ResourceNonce)
	if err != nil {
		return err
	}
	return nil
}

type CandyMachine struct {
	Authority     ag_solanago.PublicKey
	Wallet        ag_solanago.PublicKey
	TokenMint     *ag_solanago.PublicKey `bin:"optional"`
	Config        ag_solanago.PublicKey
	Data          CandyMachineData
	ItemsRedeemed uint64
	Bump          uint8
}

var CandyMachineDiscriminator = [8]byte{51, 173, 177, 113, 25, 241, 109, 189}

func (obj CandyMachine) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(CandyMachineDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Authority` param:
	err = encoder.Encode(obj.Authority)
	if err != nil {
		return err
	}
	// Serialize `Wallet` param:
	err = encoder.Encode(obj.Wallet)
	if err != nil {
		return err
	}
	// Serialize `TokenMint` param (optional):
	{
		if obj.TokenMint == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.TokenMint)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Config` param:
	err = encoder.Encode(obj.Config)
	if err != nil {
		return err
	}
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	// Serialize `ItemsRedeemed` param:
	err = encoder.Encode(obj.ItemsRedeemed)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CandyMachine) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(CandyMachineDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[51 173 177 113 25 241 109 189]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Authority`:
	err = decoder.Decode(&obj.Authority)
	if err != nil {
		return err
	}
	// Deserialize `Wallet`:
	err = decoder.Decode(&obj.Wallet)
	if err != nil {
		return err
	}
	// Deserialize `TokenMint` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.TokenMint)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Config`:
	err = decoder.Decode(&obj.Config)
	if err != nil {
		return err
	}
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	// Deserialize `ItemsRedeemed`:
	err = decoder.Decode(&obj.ItemsRedeemed)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	return nil
}

type Config struct {
	Authority ag_solanago.PublicKey
	Data      ConfigData
}

var ConfigDiscriminator = [8]byte{155, 12, 170, 224, 30, 250, 204, 130}

func (obj Config) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ConfigDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Authority` param:
	err = encoder.Encode(obj.Authority)
	if err != nil {
		return err
	}
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Config) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ConfigDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[155 12 170 224 30 250 204 130]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Authority`:
	err = decoder.Decode(&obj.Authority)
	if err != nil {
		return err
	}
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	return nil
}