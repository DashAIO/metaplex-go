// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine_v2

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetCollectionDuringMint is the `setCollectionDuringMint` instruction.
type SetCollectionDuringMint struct {

	// [0] = [] candyMachine
	//
	// [1] = [] metadata
	//
	// [2] = [SIGNER] payer
	//
	// [3] = [WRITE] collectionPda
	//
	// [4] = [] tokenMetadataProgram
	//
	// [5] = [] instructions
	//
	// [6] = [] collectionMint
	//
	// [7] = [] collectionMetadata
	//
	// [8] = [] collectionMasterEdition
	//
	// [9] = [] authority
	//
	// [10] = [] collectionAuthorityRecord
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetCollectionDuringMintInstructionBuilder creates a new `SetCollectionDuringMint` instruction builder.
func NewSetCollectionDuringMintInstructionBuilder() *SetCollectionDuringMint {
	nd := &SetCollectionDuringMint{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *SetCollectionDuringMint) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *SetCollectionDuringMint {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(candyMachine)
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *SetCollectionDuringMint) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *SetCollectionDuringMint) SetMetadataAccount(metadata ag_solanago.PublicKey) *SetCollectionDuringMint {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(metadata)
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *SetCollectionDuringMint) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
func (inst *SetCollectionDuringMint) SetPayerAccount(payer ag_solanago.PublicKey) *SetCollectionDuringMint {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *SetCollectionDuringMint) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetCollectionPdaAccount sets the "collectionPda" account.
func (inst *SetCollectionDuringMint) SetCollectionPdaAccount(collectionPda ag_solanago.PublicKey) *SetCollectionDuringMint {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(collectionPda).WRITE()
	return inst
}

// GetCollectionPdaAccount gets the "collectionPda" account.
func (inst *SetCollectionDuringMint) GetCollectionPdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
func (inst *SetCollectionDuringMint) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *SetCollectionDuringMint {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
func (inst *SetCollectionDuringMint) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetInstructionsAccount sets the "instructions" account.
func (inst *SetCollectionDuringMint) SetInstructionsAccount(instructions ag_solanago.PublicKey) *SetCollectionDuringMint {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(instructions)
	return inst
}

// GetInstructionsAccount gets the "instructions" account.
func (inst *SetCollectionDuringMint) GetInstructionsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetCollectionMintAccount sets the "collectionMint" account.
func (inst *SetCollectionDuringMint) SetCollectionMintAccount(collectionMint ag_solanago.PublicKey) *SetCollectionDuringMint {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(collectionMint)
	return inst
}

// GetCollectionMintAccount gets the "collectionMint" account.
func (inst *SetCollectionDuringMint) GetCollectionMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetCollectionMetadataAccount sets the "collectionMetadata" account.
func (inst *SetCollectionDuringMint) SetCollectionMetadataAccount(collectionMetadata ag_solanago.PublicKey) *SetCollectionDuringMint {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(collectionMetadata)
	return inst
}

// GetCollectionMetadataAccount gets the "collectionMetadata" account.
func (inst *SetCollectionDuringMint) GetCollectionMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetCollectionMasterEditionAccount sets the "collectionMasterEdition" account.
func (inst *SetCollectionDuringMint) SetCollectionMasterEditionAccount(collectionMasterEdition ag_solanago.PublicKey) *SetCollectionDuringMint {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(collectionMasterEdition)
	return inst
}

// GetCollectionMasterEditionAccount gets the "collectionMasterEdition" account.
func (inst *SetCollectionDuringMint) GetCollectionMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SetCollectionDuringMint) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetCollectionDuringMint {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SetCollectionDuringMint) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetCollectionAuthorityRecordAccount sets the "collectionAuthorityRecord" account.
func (inst *SetCollectionDuringMint) SetCollectionAuthorityRecordAccount(collectionAuthorityRecord ag_solanago.PublicKey) *SetCollectionDuringMint {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(collectionAuthorityRecord)
	return inst
}

// GetCollectionAuthorityRecordAccount gets the "collectionAuthorityRecord" account.
func (inst *SetCollectionDuringMint) GetCollectionAuthorityRecordAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst SetCollectionDuringMint) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetCollectionDuringMint,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetCollectionDuringMint) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetCollectionDuringMint) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.CollectionPda is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Instructions is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.CollectionMint is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.CollectionMetadata is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.CollectionMasterEdition is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.CollectionAuthorityRecord is not set")
		}
	}
	return nil
}

func (inst *SetCollectionDuringMint) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetCollectionDuringMint")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             candyMachine", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                 metadata", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                    payer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("            collectionPda", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("     tokenMetadataProgram", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("             instructions", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           collectionMint", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("       collectionMetadata", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("  collectionMasterEdition", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                authority", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("collectionAuthorityRecord", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj SetCollectionDuringMint) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *SetCollectionDuringMint) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewSetCollectionDuringMintInstruction declares a new SetCollectionDuringMint instruction with the provided parameters and accounts.
func NewSetCollectionDuringMintInstruction(
	// Accounts:
	candyMachine ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	collectionPda ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	instructions ag_solanago.PublicKey,
	collectionMint ag_solanago.PublicKey,
	collectionMetadata ag_solanago.PublicKey,
	collectionMasterEdition ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	collectionAuthorityRecord ag_solanago.PublicKey) *SetCollectionDuringMint {
	return NewSetCollectionDuringMintInstructionBuilder().
		SetCandyMachineAccount(candyMachine).
		SetMetadataAccount(metadata).
		SetPayerAccount(payer).
		SetCollectionPdaAccount(collectionPda).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetInstructionsAccount(instructions).
		SetCollectionMintAccount(collectionMint).
		SetCollectionMetadataAccount(collectionMetadata).
		SetCollectionMasterEditionAccount(collectionMasterEdition).
		SetAuthorityAccount(authority).
		SetCollectionAuthorityRecordAccount(collectionAuthorityRecord)
}
