// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type CandyMachineDataV2 struct {
	Price                uint64
	ItemsAvailable       uint64
	GoLiveDate           int64
	Symbol               string
	SellerFeeBasisPoints uint16
	Creators             []Creator
	IsMutable            bool
	RetainAuthority      bool
	BaseUrl              string
	MintsPerUser         *uint32      `bin:"optional"`
	Whitelist            *WhitelistV2 `bin:"optional"`
}

func (obj CandyMachineDataV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `ItemsAvailable` param:
	err = encoder.Encode(obj.ItemsAvailable)
	if err != nil {
		return err
	}
	// Serialize `GoLiveDate` param:
	err = encoder.Encode(obj.GoLiveDate)
	if err != nil {
		return err
	}
	// Serialize `Symbol` param:
	err = encoder.Encode(obj.Symbol)
	if err != nil {
		return err
	}
	// Serialize `SellerFeeBasisPoints` param:
	err = encoder.Encode(obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Serialize `Creators` param:
	err = encoder.Encode(obj.Creators)
	if err != nil {
		return err
	}
	// Serialize `IsMutable` param:
	err = encoder.Encode(obj.IsMutable)
	if err != nil {
		return err
	}
	// Serialize `RetainAuthority` param:
	err = encoder.Encode(obj.RetainAuthority)
	if err != nil {
		return err
	}
	// Serialize `BaseUrl` param:
	err = encoder.Encode(obj.BaseUrl)
	if err != nil {
		return err
	}
	// Serialize `MintsPerUser` param (optional):
	{
		if obj.MintsPerUser == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MintsPerUser)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Whitelist` param (optional):
	{
		if obj.Whitelist == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Whitelist)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *CandyMachineDataV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `ItemsAvailable`:
	err = decoder.Decode(&obj.ItemsAvailable)
	if err != nil {
		return err
	}
	// Deserialize `GoLiveDate`:
	err = decoder.Decode(&obj.GoLiveDate)
	if err != nil {
		return err
	}
	// Deserialize `Symbol`:
	err = decoder.Decode(&obj.Symbol)
	if err != nil {
		return err
	}
	// Deserialize `SellerFeeBasisPoints`:
	err = decoder.Decode(&obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Deserialize `Creators`:
	err = decoder.Decode(&obj.Creators)
	if err != nil {
		return err
	}
	// Deserialize `IsMutable`:
	err = decoder.Decode(&obj.IsMutable)
	if err != nil {
		return err
	}
	// Deserialize `RetainAuthority`:
	err = decoder.Decode(&obj.RetainAuthority)
	if err != nil {
		return err
	}
	// Deserialize `BaseUrl`:
	err = decoder.Decode(&obj.BaseUrl)
	if err != nil {
		return err
	}
	// Deserialize `MintsPerUser` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MintsPerUser)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Whitelist` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Whitelist)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type WhitelistV2 struct {
	MerkleRoot   [32]uint8
	MintsPerUser *uint32 `bin:"optional"`
	GoLiveDate   int64
	Price        *uint64 `bin:"optional"`
}

func (obj WhitelistV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MerkleRoot` param:
	err = encoder.Encode(obj.MerkleRoot)
	if err != nil {
		return err
	}
	// Serialize `MintsPerUser` param (optional):
	{
		if obj.MintsPerUser == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MintsPerUser)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `GoLiveDate` param:
	err = encoder.Encode(obj.GoLiveDate)
	if err != nil {
		return err
	}
	// Serialize `Price` param (optional):
	{
		if obj.Price == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Price)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *WhitelistV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MerkleRoot`:
	err = decoder.Decode(&obj.MerkleRoot)
	if err != nil {
		return err
	}
	// Deserialize `MintsPerUser` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MintsPerUser)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `GoLiveDate`:
	err = decoder.Decode(&obj.GoLiveDate)
	if err != nil {
		return err
	}
	// Deserialize `Price` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Price)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type CandyMachineData struct {
	Uuid           string
	Price          uint64
	ItemsAvailable uint64
	GoLiveDate     *int64 `bin:"optional"`
}

func (obj CandyMachineData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Uuid` param:
	err = encoder.Encode(obj.Uuid)
	if err != nil {
		return err
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `ItemsAvailable` param:
	err = encoder.Encode(obj.ItemsAvailable)
	if err != nil {
		return err
	}
	// Serialize `GoLiveDate` param (optional):
	{
		if obj.GoLiveDate == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.GoLiveDate)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *CandyMachineData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Uuid`:
	err = decoder.Decode(&obj.Uuid)
	if err != nil {
		return err
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `ItemsAvailable`:
	err = decoder.Decode(&obj.ItemsAvailable)
	if err != nil {
		return err
	}
	// Deserialize `GoLiveDate` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.GoLiveDate)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type ConfigData struct {
	Uuid                 string
	Symbol               string
	SellerFeeBasisPoints uint16
	Creators             []Creator
	MaxSupply            uint64
	IsMutable            bool
	RetainAuthority      bool
	MaxNumberOfLines     uint32
}

func (obj ConfigData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Uuid` param:
	err = encoder.Encode(obj.Uuid)
	if err != nil {
		return err
	}
	// Serialize `Symbol` param:
	err = encoder.Encode(obj.Symbol)
	if err != nil {
		return err
	}
	// Serialize `SellerFeeBasisPoints` param:
	err = encoder.Encode(obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Serialize `Creators` param:
	err = encoder.Encode(obj.Creators)
	if err != nil {
		return err
	}
	// Serialize `MaxSupply` param:
	err = encoder.Encode(obj.MaxSupply)
	if err != nil {
		return err
	}
	// Serialize `IsMutable` param:
	err = encoder.Encode(obj.IsMutable)
	if err != nil {
		return err
	}
	// Serialize `RetainAuthority` param:
	err = encoder.Encode(obj.RetainAuthority)
	if err != nil {
		return err
	}
	// Serialize `MaxNumberOfLines` param:
	err = encoder.Encode(obj.MaxNumberOfLines)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ConfigData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Uuid`:
	err = decoder.Decode(&obj.Uuid)
	if err != nil {
		return err
	}
	// Deserialize `Symbol`:
	err = decoder.Decode(&obj.Symbol)
	if err != nil {
		return err
	}
	// Deserialize `SellerFeeBasisPoints`:
	err = decoder.Decode(&obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Deserialize `Creators`:
	err = decoder.Decode(&obj.Creators)
	if err != nil {
		return err
	}
	// Deserialize `MaxSupply`:
	err = decoder.Decode(&obj.MaxSupply)
	if err != nil {
		return err
	}
	// Deserialize `IsMutable`:
	err = decoder.Decode(&obj.IsMutable)
	if err != nil {
		return err
	}
	// Deserialize `RetainAuthority`:
	err = decoder.Decode(&obj.RetainAuthority)
	if err != nil {
		return err
	}
	// Deserialize `MaxNumberOfLines`:
	err = decoder.Decode(&obj.MaxNumberOfLines)
	if err != nil {
		return err
	}
	return nil
}

type ConfigLine struct {
	Name string
	Uri  string
}

func (obj ConfigLine) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Uri` param:
	err = encoder.Encode(obj.Uri)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ConfigLine) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Uri`:
	err = decoder.Decode(&obj.Uri)
	if err != nil {
		return err
	}
	return nil
}

type Creator struct {
	Address  ag_solanago.PublicKey
	Verified bool
	Share    uint8
}

func (obj Creator) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Address` param:
	err = encoder.Encode(obj.Address)
	if err != nil {
		return err
	}
	// Serialize `Verified` param:
	err = encoder.Encode(obj.Verified)
	if err != nil {
		return err
	}
	// Serialize `Share` param:
	err = encoder.Encode(obj.Share)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Creator) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Address`:
	err = decoder.Decode(&obj.Address)
	if err != nil {
		return err
	}
	// Deserialize `Verified`:
	err = decoder.Decode(&obj.Verified)
	if err != nil {
		return err
	}
	// Deserialize `Share`:
	err = decoder.Decode(&obj.Share)
	if err != nil {
		return err
	}
	return nil
}

type ErrorCode ag_binary.BorshEnum

const (
	ErrorCodeIncorrectOwner ErrorCode = iota
	ErrorCodeUninitialized
	ErrorCodeMintMismatch
	ErrorCodeIndexGreaterThanLength
	ErrorCodeConfigMustHaveAtleastOneEntry
	ErrorCodeNumericalOverflowError
	ErrorCodeTooManyCreators
	ErrorCodeUuidMustBeExactly6Length
	ErrorCodeNotEnoughTokens
	ErrorCodeNotEnoughSOL
	ErrorCodeTokenTransferFailed
	ErrorCodeCandyMachineEmpty
	ErrorCodeCandyMachineNotLiveYet
	ErrorCodeConfigLineMismatch
	ErrorCodeWhitelistExists
	ErrorCodeWhiteListMissing
	ErrorCodeWrongWhitelist
	ErrorCodeNotWhitelisted
	ErrorCodeInvalidFraction
	ErrorCodeBumpMissing
)

func (value ErrorCode) String() string {
	switch value {
	case ErrorCodeIncorrectOwner:
		return "IncorrectOwner"
	case ErrorCodeUninitialized:
		return "Uninitialized"
	case ErrorCodeMintMismatch:
		return "MintMismatch"
	case ErrorCodeIndexGreaterThanLength:
		return "IndexGreaterThanLength"
	case ErrorCodeConfigMustHaveAtleastOneEntry:
		return "ConfigMustHaveAtleastOneEntry"
	case ErrorCodeNumericalOverflowError:
		return "NumericalOverflowError"
	case ErrorCodeTooManyCreators:
		return "TooManyCreators"
	case ErrorCodeUuidMustBeExactly6Length:
		return "UuidMustBeExactly6Length"
	case ErrorCodeNotEnoughTokens:
		return "NotEnoughTokens"
	case ErrorCodeNotEnoughSOL:
		return "NotEnoughSOL"
	case ErrorCodeTokenTransferFailed:
		return "TokenTransferFailed"
	case ErrorCodeCandyMachineEmpty:
		return "CandyMachineEmpty"
	case ErrorCodeCandyMachineNotLiveYet:
		return "CandyMachineNotLiveYet"
	case ErrorCodeConfigLineMismatch:
		return "ConfigLineMismatch"
	case ErrorCodeWhitelistExists:
		return "WhitelistExists"
	case ErrorCodeWhiteListMissing:
		return "WhiteListMissing"
	case ErrorCodeWrongWhitelist:
		return "WrongWhitelist"
	case ErrorCodeNotWhitelisted:
		return "NotWhitelisted"
	case ErrorCodeInvalidFraction:
		return "InvalidFraction"
	case ErrorCodeBumpMissing:
		return "BumpMissing"
	default:
		return ""
	}
}
