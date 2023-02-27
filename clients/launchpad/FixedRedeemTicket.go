// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package launchpad

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FixedRedeemTicket is the `fixedRedeemTicket` instruction.
type FixedRedeemTicket struct {

	// [0] = [WRITE, SIGNER] initializer
	//
	// [1] = [WRITE] ticket
	//
	// [2] = [WRITE] collection
	//
	// [3] = [WRITE] stage
	//
	// [4] = [] nftMint
	//
	// [5] = [WRITE] initializerNftTa
	//
	// [6] = [WRITE] collectionNftTa
	//
	// [7] = [WRITE] nftMetadata
	//
	// [8] = [WRITE] nftMasterEdition
	//
	// [9] = [] mplTokenMetadata
	//
	// [10] = [] tokenProgram
	//
	// [11] = [] associatedTokenProgram
	//
	// [12] = [] systemProgram
	//
	// [13] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewFixedRedeemTicketInstructionBuilder creates a new `FixedRedeemTicket` instruction builder.
func NewFixedRedeemTicketInstructionBuilder() *FixedRedeemTicket {
	nd := &FixedRedeemTicket{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 14),
	}
	return nd
}

// SetInitializerAccount sets the "initializer" account.
func (inst *FixedRedeemTicket) SetInitializerAccount(initializer ag_solanago.PublicKey) *FixedRedeemTicket {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(initializer).WRITE().SIGNER()
	return inst
}

// GetInitializerAccount gets the "initializer" account.
func (inst *FixedRedeemTicket) GetInitializerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTicketAccount sets the "ticket" account.
func (inst *FixedRedeemTicket) SetTicketAccount(ticket ag_solanago.PublicKey) *FixedRedeemTicket {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(ticket).WRITE()
	return inst
}

// GetTicketAccount gets the "ticket" account.
func (inst *FixedRedeemTicket) GetTicketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetCollectionAccount sets the "collection" account.
func (inst *FixedRedeemTicket) SetCollectionAccount(collection ag_solanago.PublicKey) *FixedRedeemTicket {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(collection).WRITE()
	return inst
}

// GetCollectionAccount gets the "collection" account.
func (inst *FixedRedeemTicket) GetCollectionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetStageAccount sets the "stage" account.
func (inst *FixedRedeemTicket) SetStageAccount(stage ag_solanago.PublicKey) *FixedRedeemTicket {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(stage).WRITE()
	return inst
}

// GetStageAccount gets the "stage" account.
func (inst *FixedRedeemTicket) GetStageAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetNftMintAccount sets the "nftMint" account.
func (inst *FixedRedeemTicket) SetNftMintAccount(nftMint ag_solanago.PublicKey) *FixedRedeemTicket {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(nftMint)
	return inst
}

// GetNftMintAccount gets the "nftMint" account.
func (inst *FixedRedeemTicket) GetNftMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetInitializerNftTaAccount sets the "initializerNftTa" account.
func (inst *FixedRedeemTicket) SetInitializerNftTaAccount(initializerNftTa ag_solanago.PublicKey) *FixedRedeemTicket {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(initializerNftTa).WRITE()
	return inst
}

// GetInitializerNftTaAccount gets the "initializerNftTa" account.
func (inst *FixedRedeemTicket) GetInitializerNftTaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetCollectionNftTaAccount sets the "collectionNftTa" account.
func (inst *FixedRedeemTicket) SetCollectionNftTaAccount(collectionNftTa ag_solanago.PublicKey) *FixedRedeemTicket {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(collectionNftTa).WRITE()
	return inst
}

// GetCollectionNftTaAccount gets the "collectionNftTa" account.
func (inst *FixedRedeemTicket) GetCollectionNftTaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetNftMetadataAccount sets the "nftMetadata" account.
func (inst *FixedRedeemTicket) SetNftMetadataAccount(nftMetadata ag_solanago.PublicKey) *FixedRedeemTicket {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(nftMetadata).WRITE()
	return inst
}

// GetNftMetadataAccount gets the "nftMetadata" account.
func (inst *FixedRedeemTicket) GetNftMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetNftMasterEditionAccount sets the "nftMasterEdition" account.
func (inst *FixedRedeemTicket) SetNftMasterEditionAccount(nftMasterEdition ag_solanago.PublicKey) *FixedRedeemTicket {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(nftMasterEdition).WRITE()
	return inst
}

// GetNftMasterEditionAccount gets the "nftMasterEdition" account.
func (inst *FixedRedeemTicket) GetNftMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetMplTokenMetadataAccount sets the "mplTokenMetadata" account.
func (inst *FixedRedeemTicket) SetMplTokenMetadataAccount(mplTokenMetadata ag_solanago.PublicKey) *FixedRedeemTicket {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(mplTokenMetadata)
	return inst
}

// GetMplTokenMetadataAccount gets the "mplTokenMetadata" account.
func (inst *FixedRedeemTicket) GetMplTokenMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *FixedRedeemTicket) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *FixedRedeemTicket {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *FixedRedeemTicket) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *FixedRedeemTicket) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *FixedRedeemTicket {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *FixedRedeemTicket) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *FixedRedeemTicket) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *FixedRedeemTicket {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *FixedRedeemTicket) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetRentAccount sets the "rent" account.
func (inst *FixedRedeemTicket) SetRentAccount(rent ag_solanago.PublicKey) *FixedRedeemTicket {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *FixedRedeemTicket) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

func (inst FixedRedeemTicket) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FixedRedeemTicket,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FixedRedeemTicket) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FixedRedeemTicket) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Initializer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Ticket is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Collection is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Stage is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.NftMint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.InitializerNftTa is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.CollectionNftTa is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.NftMetadata is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.NftMasterEdition is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.MplTokenMetadata is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *FixedRedeemTicket) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FixedRedeemTicket")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=14]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("           initializer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                ticket", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("            collection", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                 stage", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("               nftMint", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("      initializerNftTa", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("       collectionNftTa", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("           nftMetadata", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("      nftMasterEdition", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("      mplTokenMetadata", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("                  rent", inst.AccountMetaSlice.Get(13)))
					})
				})
		})
}

func (obj FixedRedeemTicket) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *FixedRedeemTicket) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewFixedRedeemTicketInstruction declares a new FixedRedeemTicket instruction with the provided parameters and accounts.
func NewFixedRedeemTicketInstruction(
	// Accounts:
	initializer ag_solanago.PublicKey,
	ticket ag_solanago.PublicKey,
	collection ag_solanago.PublicKey,
	stage ag_solanago.PublicKey,
	nftMint ag_solanago.PublicKey,
	initializerNftTa ag_solanago.PublicKey,
	collectionNftTa ag_solanago.PublicKey,
	nftMetadata ag_solanago.PublicKey,
	nftMasterEdition ag_solanago.PublicKey,
	mplTokenMetadata ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *FixedRedeemTicket {
	return NewFixedRedeemTicketInstructionBuilder().
		SetInitializerAccount(initializer).
		SetTicketAccount(ticket).
		SetCollectionAccount(collection).
		SetStageAccount(stage).
		SetNftMintAccount(nftMint).
		SetInitializerNftTaAccount(initializerNftTa).
		SetCollectionNftTaAccount(collectionNftTa).
		SetNftMetadataAccount(nftMetadata).
		SetNftMasterEditionAccount(nftMasterEdition).
		SetMplTokenMetadataAccount(mplTokenMetadata).
		SetTokenProgramAccount(tokenProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}