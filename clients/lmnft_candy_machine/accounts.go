// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
)

type Whitelist struct {
	Authority    ag_solanago.PublicKey
	CandyMachine ag_solanago.PublicKey
	MerkleRoot   *[32]uint8 `bin:"optional"`
}

var WhitelistDiscriminator = [8]byte{204, 176, 52, 79, 146, 121, 54, 247}

func (obj Whitelist) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(WhitelistDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Authority` param:
	err = encoder.Encode(obj.Authority)
	if err != nil {
		return err
	}
	// Serialize `CandyMachine` param:
	err = encoder.Encode(obj.CandyMachine)
	if err != nil {
		return err
	}
	// Serialize `MerkleRoot` param (optional):
	{
		if obj.MerkleRoot == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MerkleRoot)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *Whitelist) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(WhitelistDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[204 176 52 79 146 121 54 247]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Authority`:
	err = decoder.Decode(&obj.Authority)
	if err != nil {
		return err
	}
	// Deserialize `CandyMachine`:
	err = decoder.Decode(&obj.CandyMachine)
	if err != nil {
		return err
	}
	// Deserialize `MerkleRoot` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MerkleRoot)
			if err != nil {
				return err
			}
		}
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

type CandyMachineV2 struct {
	Authority     ag_solanago.PublicKey
	Wallet        ag_solanago.PublicKey
	ItemsRedeemed uint64
	Data          CandyMachineDataV2
}

var CandyMachineV2Discriminator = [8]byte{50, 243, 71, 181, 164, 239, 110, 131}

func (obj CandyMachineV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(CandyMachineV2Discriminator[:], false)
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
	// Serialize `ItemsRedeemed` param:
	err = encoder.Encode(obj.ItemsRedeemed)
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

func (obj *CandyMachineV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(CandyMachineV2Discriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[50 243 71 181 164 239 110 131]",
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
	// Deserialize `ItemsRedeemed`:
	err = decoder.Decode(&obj.ItemsRedeemed)
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

type CandyMachineV3 struct {
	Seed          ag_solanago.PublicKey
	Bump          uint8
	Authority     ag_solanago.PublicKey
	Wallet        ag_solanago.PublicKey
	ItemsRedeemed uint64
	Data          CandyMachineDataV2
}

var CandyMachineV3Discriminator = [8]byte{221, 21, 200, 78, 142, 15, 223, 6}

func (obj CandyMachineV3) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(CandyMachineV3Discriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Seed` param:
	err = encoder.Encode(obj.Seed)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
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
	// Serialize `ItemsRedeemed` param:
	err = encoder.Encode(obj.ItemsRedeemed)
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

func (obj *CandyMachineV3) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(CandyMachineV3Discriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[221 21 200 78 142 15 223 6]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Seed`:
	err = decoder.Decode(&obj.Seed)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
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
	// Deserialize `ItemsRedeemed`:
	err = decoder.Decode(&obj.ItemsRedeemed)
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

type TotalMints struct {
	Total uint32
}

var TotalMintsDiscriminator = [8]byte{5, 252, 73, 108, 50, 30, 212, 224}

func (obj TotalMints) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(TotalMintsDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Total` param:
	err = encoder.Encode(obj.Total)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TotalMints) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(TotalMintsDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[5 252 73 108 50 30 212 224]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Total`:
	err = decoder.Decode(&obj.Total)
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
