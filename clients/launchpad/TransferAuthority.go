// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package launchpad

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// TransferAuthority is the `transferAuthority` instruction.
type TransferAuthority struct {
	Nonce  *ag_solanago.PublicKey
	Name   *string
	Symbol *string
	Uri    *string

	// [0] = [WRITE, SIGNER] launchpadManager
	//
	// [1] = [WRITE] authority
	//
	// [2] = [WRITE] collection
	//
	// [3] = [WRITE] nftMint
	//
	// [4] = [WRITE] collectionNftTa
	//
	// [5] = [WRITE] nftMetadata
	//
	// [6] = [WRITE] nftMasterEdition
	//
	// [7] = [WRITE] collectionMint
	//
	// [8] = [WRITE] collectionMetadata
	//
	// [9] = [WRITE] collectionMasterEdition
	//
	// [10] = [] mplTokenMetadata
	//
	// [11] = [] tokenProgram
	//
	// [12] = [] associatedTokenProgram
	//
	// [13] = [] systemProgram
	//
	// [14] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewTransferAuthorityInstructionBuilder creates a new `TransferAuthority` instruction builder.
func NewTransferAuthorityInstructionBuilder() *TransferAuthority {
	nd := &TransferAuthority{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 15),
	}
	return nd
}

// SetNonce sets the "nonce" parameter.
func (inst *TransferAuthority) SetNonce(nonce ag_solanago.PublicKey) *TransferAuthority {
	inst.Nonce = &nonce
	return inst
}

// SetName sets the "name" parameter.
func (inst *TransferAuthority) SetName(name string) *TransferAuthority {
	inst.Name = &name
	return inst
}

// SetSymbol sets the "symbol" parameter.
func (inst *TransferAuthority) SetSymbol(symbol string) *TransferAuthority {
	inst.Symbol = &symbol
	return inst
}

// SetUri sets the "uri" parameter.
func (inst *TransferAuthority) SetUri(uri string) *TransferAuthority {
	inst.Uri = &uri
	return inst
}

// SetLaunchpadManagerAccount sets the "launchpadManager" account.
func (inst *TransferAuthority) SetLaunchpadManagerAccount(launchpadManager ag_solanago.PublicKey) *TransferAuthority {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(launchpadManager).WRITE().SIGNER()
	return inst
}

// GetLaunchpadManagerAccount gets the "launchpadManager" account.
func (inst *TransferAuthority) GetLaunchpadManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *TransferAuthority) SetAuthorityAccount(authority ag_solanago.PublicKey) *TransferAuthority {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).WRITE()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *TransferAuthority) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetCollectionAccount sets the "collection" account.
func (inst *TransferAuthority) SetCollectionAccount(collection ag_solanago.PublicKey) *TransferAuthority {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(collection).WRITE()
	return inst
}

// GetCollectionAccount gets the "collection" account.
func (inst *TransferAuthority) GetCollectionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetNftMintAccount sets the "nftMint" account.
func (inst *TransferAuthority) SetNftMintAccount(nftMint ag_solanago.PublicKey) *TransferAuthority {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(nftMint).WRITE()
	return inst
}

// GetNftMintAccount gets the "nftMint" account.
func (inst *TransferAuthority) GetNftMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetCollectionNftTaAccount sets the "collectionNftTa" account.
func (inst *TransferAuthority) SetCollectionNftTaAccount(collectionNftTa ag_solanago.PublicKey) *TransferAuthority {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(collectionNftTa).WRITE()
	return inst
}

// GetCollectionNftTaAccount gets the "collectionNftTa" account.
func (inst *TransferAuthority) GetCollectionNftTaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetNftMetadataAccount sets the "nftMetadata" account.
func (inst *TransferAuthority) SetNftMetadataAccount(nftMetadata ag_solanago.PublicKey) *TransferAuthority {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(nftMetadata).WRITE()
	return inst
}

// GetNftMetadataAccount gets the "nftMetadata" account.
func (inst *TransferAuthority) GetNftMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetNftMasterEditionAccount sets the "nftMasterEdition" account.
func (inst *TransferAuthority) SetNftMasterEditionAccount(nftMasterEdition ag_solanago.PublicKey) *TransferAuthority {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(nftMasterEdition).WRITE()
	return inst
}

// GetNftMasterEditionAccount gets the "nftMasterEdition" account.
func (inst *TransferAuthority) GetNftMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetCollectionMintAccount sets the "collectionMint" account.
func (inst *TransferAuthority) SetCollectionMintAccount(collectionMint ag_solanago.PublicKey) *TransferAuthority {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(collectionMint).WRITE()
	return inst
}

// GetCollectionMintAccount gets the "collectionMint" account.
func (inst *TransferAuthority) GetCollectionMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetCollectionMetadataAccount sets the "collectionMetadata" account.
func (inst *TransferAuthority) SetCollectionMetadataAccount(collectionMetadata ag_solanago.PublicKey) *TransferAuthority {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(collectionMetadata).WRITE()
	return inst
}

// GetCollectionMetadataAccount gets the "collectionMetadata" account.
func (inst *TransferAuthority) GetCollectionMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetCollectionMasterEditionAccount sets the "collectionMasterEdition" account.
func (inst *TransferAuthority) SetCollectionMasterEditionAccount(collectionMasterEdition ag_solanago.PublicKey) *TransferAuthority {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(collectionMasterEdition).WRITE()
	return inst
}

// GetCollectionMasterEditionAccount gets the "collectionMasterEdition" account.
func (inst *TransferAuthority) GetCollectionMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetMplTokenMetadataAccount sets the "mplTokenMetadata" account.
func (inst *TransferAuthority) SetMplTokenMetadataAccount(mplTokenMetadata ag_solanago.PublicKey) *TransferAuthority {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(mplTokenMetadata)
	return inst
}

// GetMplTokenMetadataAccount gets the "mplTokenMetadata" account.
func (inst *TransferAuthority) GetMplTokenMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *TransferAuthority) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *TransferAuthority {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *TransferAuthority) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *TransferAuthority) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *TransferAuthority {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *TransferAuthority) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *TransferAuthority) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *TransferAuthority {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *TransferAuthority) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetRentAccount sets the "rent" account.
func (inst *TransferAuthority) SetRentAccount(rent ag_solanago.PublicKey) *TransferAuthority {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *TransferAuthority) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

func (inst TransferAuthority) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_TransferAuthority,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst TransferAuthority) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *TransferAuthority) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Nonce == nil {
			return errors.New("Nonce parameter is not set")
		}
		if inst.Name == nil {
			return errors.New("Name parameter is not set")
		}
		if inst.Symbol == nil {
			return errors.New("Symbol parameter is not set")
		}
		if inst.Uri == nil {
			return errors.New("Uri parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.LaunchpadManager is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Collection is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.NftMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.CollectionNftTa is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.NftMetadata is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.NftMasterEdition is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.CollectionMint is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.CollectionMetadata is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.CollectionMasterEdition is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.MplTokenMetadata is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *TransferAuthority) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("TransferAuthority")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=4]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" Nonce", *inst.Nonce))
						paramsBranch.Child(ag_format.Param("  Name", *inst.Name))
						paramsBranch.Child(ag_format.Param("Symbol", *inst.Symbol))
						paramsBranch.Child(ag_format.Param("   Uri", *inst.Uri))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=15]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       launchpadManager", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("              authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("             collection", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                nftMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        collectionNftTa", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("            nftMetadata", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("       nftMasterEdition", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("         collectionMint", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("     collectionMetadata", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("collectionMasterEdition", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("       mplTokenMetadata", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("           tokenProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta(" associatedTokenProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("          systemProgram", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("                   rent", inst.AccountMetaSlice.Get(14)))
					})
				})
		})
}

func (obj TransferAuthority) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Nonce` param:
	err = encoder.Encode(obj.Nonce)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Symbol` param:
	err = encoder.Encode(obj.Symbol)
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
func (obj *TransferAuthority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Nonce`:
	err = decoder.Decode(&obj.Nonce)
	if err != nil {
		return err
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Symbol`:
	err = decoder.Decode(&obj.Symbol)
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

// NewTransferAuthorityInstruction declares a new TransferAuthority instruction with the provided parameters and accounts.
func NewTransferAuthorityInstruction(
	// Parameters:
	nonce ag_solanago.PublicKey,
	name string,
	symbol string,
	uri string,
	// Accounts:
	launchpadManager ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	collection ag_solanago.PublicKey,
	nftMint ag_solanago.PublicKey,
	collectionNftTa ag_solanago.PublicKey,
	nftMetadata ag_solanago.PublicKey,
	nftMasterEdition ag_solanago.PublicKey,
	collectionMint ag_solanago.PublicKey,
	collectionMetadata ag_solanago.PublicKey,
	collectionMasterEdition ag_solanago.PublicKey,
	mplTokenMetadata ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *TransferAuthority {
	return NewTransferAuthorityInstructionBuilder().
		SetNonce(nonce).
		SetName(name).
		SetSymbol(symbol).
		SetUri(uri).
		SetLaunchpadManagerAccount(launchpadManager).
		SetAuthorityAccount(authority).
		SetCollectionAccount(collection).
		SetNftMintAccount(nftMint).
		SetCollectionNftTaAccount(collectionNftTa).
		SetNftMetadataAccount(nftMetadata).
		SetNftMasterEditionAccount(nftMasterEdition).
		SetCollectionMintAccount(collectionMint).
		SetCollectionMetadataAccount(collectionMetadata).
		SetCollectionMasterEditionAccount(collectionMasterEdition).
		SetMplTokenMetadataAccount(mplTokenMetadata).
		SetTokenProgramAccount(tokenProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}