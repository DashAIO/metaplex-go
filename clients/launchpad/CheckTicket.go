// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package launchpad

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CheckTicket is the `checkTicket` instruction.
type CheckTicket struct {
	Result *bool

	// [0] = [WRITE, SIGNER] launchpadManager
	//
	// [1] = [] initializer
	//
	// [2] = [WRITE] ticket
	//
	// [3] = [WRITE] collection
	//
	// [4] = [WRITE] stage
	//
	// [5] = [] nftMint
	//
	// [6] = [] paymentMint
	//
	// [7] = [] whitelistMint
	//
	// [8] = [WRITE] initializerWhitelistTa
	//
	// [9] = [WRITE] collectionWhitelistTa
	//
	// [10] = [WRITE] ticketNftTa
	//
	// [11] = [WRITE] collectionNftTa
	//
	// [12] = [WRITE] initializerTa
	//
	// [13] = [WRITE] stageTa
	//
	// [14] = [WRITE] nftMetadata
	//
	// [15] = [WRITE] nftMasterEdition
	//
	// [16] = [] mplTokenMetadata
	//
	// [17] = [] tokenProgram
	//
	// [18] = [] associatedTokenProgram
	//
	// [19] = [] systemProgram
	//
	// [20] = [] rent
	//
	// [21] = [] clock
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCheckTicketInstructionBuilder creates a new `CheckTicket` instruction builder.
func NewCheckTicketInstructionBuilder() *CheckTicket {
	nd := &CheckTicket{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 22),
	}
	return nd
}

// SetResult sets the "result" parameter.
func (inst *CheckTicket) SetResult(result bool) *CheckTicket {
	inst.Result = &result
	return inst
}

// SetLaunchpadManagerAccount sets the "launchpadManager" account.
func (inst *CheckTicket) SetLaunchpadManagerAccount(launchpadManager ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(launchpadManager).WRITE().SIGNER()
	return inst
}

// GetLaunchpadManagerAccount gets the "launchpadManager" account.
func (inst *CheckTicket) GetLaunchpadManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetInitializerAccount sets the "initializer" account.
func (inst *CheckTicket) SetInitializerAccount(initializer ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(initializer)
	return inst
}

// GetInitializerAccount gets the "initializer" account.
func (inst *CheckTicket) GetInitializerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTicketAccount sets the "ticket" account.
func (inst *CheckTicket) SetTicketAccount(ticket ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(ticket).WRITE()
	return inst
}

// GetTicketAccount gets the "ticket" account.
func (inst *CheckTicket) GetTicketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetCollectionAccount sets the "collection" account.
func (inst *CheckTicket) SetCollectionAccount(collection ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(collection).WRITE()
	return inst
}

// GetCollectionAccount gets the "collection" account.
func (inst *CheckTicket) GetCollectionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetStageAccount sets the "stage" account.
func (inst *CheckTicket) SetStageAccount(stage ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(stage).WRITE()
	return inst
}

// GetStageAccount gets the "stage" account.
func (inst *CheckTicket) GetStageAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetNftMintAccount sets the "nftMint" account.
func (inst *CheckTicket) SetNftMintAccount(nftMint ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(nftMint)
	return inst
}

// GetNftMintAccount gets the "nftMint" account.
func (inst *CheckTicket) GetNftMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPaymentMintAccount sets the "paymentMint" account.
func (inst *CheckTicket) SetPaymentMintAccount(paymentMint ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(paymentMint)
	return inst
}

// GetPaymentMintAccount gets the "paymentMint" account.
func (inst *CheckTicket) GetPaymentMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetWhitelistMintAccount sets the "whitelistMint" account.
func (inst *CheckTicket) SetWhitelistMintAccount(whitelistMint ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(whitelistMint)
	return inst
}

// GetWhitelistMintAccount gets the "whitelistMint" account.
func (inst *CheckTicket) GetWhitelistMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetInitializerWhitelistTaAccount sets the "initializerWhitelistTa" account.
func (inst *CheckTicket) SetInitializerWhitelistTaAccount(initializerWhitelistTa ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(initializerWhitelistTa).WRITE()
	return inst
}

// GetInitializerWhitelistTaAccount gets the "initializerWhitelistTa" account.
func (inst *CheckTicket) GetInitializerWhitelistTaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetCollectionWhitelistTaAccount sets the "collectionWhitelistTa" account.
func (inst *CheckTicket) SetCollectionWhitelistTaAccount(collectionWhitelistTa ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(collectionWhitelistTa).WRITE()
	return inst
}

// GetCollectionWhitelistTaAccount gets the "collectionWhitelistTa" account.
func (inst *CheckTicket) GetCollectionWhitelistTaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTicketNftTaAccount sets the "ticketNftTa" account.
func (inst *CheckTicket) SetTicketNftTaAccount(ticketNftTa ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(ticketNftTa).WRITE()
	return inst
}

// GetTicketNftTaAccount gets the "ticketNftTa" account.
func (inst *CheckTicket) GetTicketNftTaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetCollectionNftTaAccount sets the "collectionNftTa" account.
func (inst *CheckTicket) SetCollectionNftTaAccount(collectionNftTa ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(collectionNftTa).WRITE()
	return inst
}

// GetCollectionNftTaAccount gets the "collectionNftTa" account.
func (inst *CheckTicket) GetCollectionNftTaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetInitializerTaAccount sets the "initializerTa" account.
func (inst *CheckTicket) SetInitializerTaAccount(initializerTa ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(initializerTa).WRITE()
	return inst
}

// GetInitializerTaAccount gets the "initializerTa" account.
func (inst *CheckTicket) GetInitializerTaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetStageTaAccount sets the "stageTa" account.
func (inst *CheckTicket) SetStageTaAccount(stageTa ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(stageTa).WRITE()
	return inst
}

// GetStageTaAccount gets the "stageTa" account.
func (inst *CheckTicket) GetStageTaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetNftMetadataAccount sets the "nftMetadata" account.
func (inst *CheckTicket) SetNftMetadataAccount(nftMetadata ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(nftMetadata).WRITE()
	return inst
}

// GetNftMetadataAccount gets the "nftMetadata" account.
func (inst *CheckTicket) GetNftMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetNftMasterEditionAccount sets the "nftMasterEdition" account.
func (inst *CheckTicket) SetNftMasterEditionAccount(nftMasterEdition ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(nftMasterEdition).WRITE()
	return inst
}

// GetNftMasterEditionAccount gets the "nftMasterEdition" account.
func (inst *CheckTicket) GetNftMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetMplTokenMetadataAccount sets the "mplTokenMetadata" account.
func (inst *CheckTicket) SetMplTokenMetadataAccount(mplTokenMetadata ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(mplTokenMetadata)
	return inst
}

// GetMplTokenMetadataAccount gets the "mplTokenMetadata" account.
func (inst *CheckTicket) GetMplTokenMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *CheckTicket) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *CheckTicket) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *CheckTicket) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *CheckTicket) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(18)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CheckTicket) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[19] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CheckTicket) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(19)
}

// SetRentAccount sets the "rent" account.
func (inst *CheckTicket) SetRentAccount(rent ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[20] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *CheckTicket) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(20)
}

// SetClockAccount sets the "clock" account.
func (inst *CheckTicket) SetClockAccount(clock ag_solanago.PublicKey) *CheckTicket {
	inst.AccountMetaSlice[21] = ag_solanago.Meta(clock)
	return inst
}

// GetClockAccount gets the "clock" account.
func (inst *CheckTicket) GetClockAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(21)
}

func (inst CheckTicket) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CheckTicket,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CheckTicket) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CheckTicket) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Result == nil {
			return errors.New("Result parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.LaunchpadManager is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Initializer is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Ticket is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Collection is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Stage is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.NftMint is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.PaymentMint is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.WhitelistMint is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.InitializerWhitelistTa is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.CollectionWhitelistTa is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TicketNftTa is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.CollectionNftTa is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.InitializerTa is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.StageTa is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.NftMetadata is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.NftMasterEdition is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.MplTokenMetadata is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[19] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[20] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[21] == nil {
			return errors.New("accounts.Clock is not set")
		}
	}
	return nil
}

func (inst *CheckTicket) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CheckTicket")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Result", *inst.Result))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=22]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      launchpadManager", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("           initializer", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                ticket", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("            collection", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                 stage", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("               nftMint", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           paymentMint", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("         whitelistMint", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("initializerWhitelistTa", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta(" collectionWhitelistTa", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("           ticketNftTa", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("       collectionNftTa", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("         initializerTa", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("               stageTa", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("           nftMetadata", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("      nftMasterEdition", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("      mplTokenMetadata", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(17)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(18)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(19)))
						accountsBranch.Child(ag_format.Meta("                  rent", inst.AccountMetaSlice.Get(20)))
						accountsBranch.Child(ag_format.Meta("                 clock", inst.AccountMetaSlice.Get(21)))
					})
				})
		})
}

func (obj CheckTicket) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Result` param:
	err = encoder.Encode(obj.Result)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CheckTicket) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Result`:
	err = decoder.Decode(&obj.Result)
	if err != nil {
		return err
	}
	return nil
}

// NewCheckTicketInstruction declares a new CheckTicket instruction with the provided parameters and accounts.
func NewCheckTicketInstruction(
	// Parameters:
	result bool,
	// Accounts:
	launchpadManager ag_solanago.PublicKey,
	initializer ag_solanago.PublicKey,
	ticket ag_solanago.PublicKey,
	collection ag_solanago.PublicKey,
	stage ag_solanago.PublicKey,
	nftMint ag_solanago.PublicKey,
	paymentMint ag_solanago.PublicKey,
	whitelistMint ag_solanago.PublicKey,
	initializerWhitelistTa ag_solanago.PublicKey,
	collectionWhitelistTa ag_solanago.PublicKey,
	ticketNftTa ag_solanago.PublicKey,
	collectionNftTa ag_solanago.PublicKey,
	initializerTa ag_solanago.PublicKey,
	stageTa ag_solanago.PublicKey,
	nftMetadata ag_solanago.PublicKey,
	nftMasterEdition ag_solanago.PublicKey,
	mplTokenMetadata ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	clock ag_solanago.PublicKey) *CheckTicket {
	return NewCheckTicketInstructionBuilder().
		SetResult(result).
		SetLaunchpadManagerAccount(launchpadManager).
		SetInitializerAccount(initializer).
		SetTicketAccount(ticket).
		SetCollectionAccount(collection).
		SetStageAccount(stage).
		SetNftMintAccount(nftMint).
		SetPaymentMintAccount(paymentMint).
		SetWhitelistMintAccount(whitelistMint).
		SetInitializerWhitelistTaAccount(initializerWhitelistTa).
		SetCollectionWhitelistTaAccount(collectionWhitelistTa).
		SetTicketNftTaAccount(ticketNftTa).
		SetCollectionNftTaAccount(collectionNftTa).
		SetInitializerTaAccount(initializerTa).
		SetStageTaAccount(stageTa).
		SetNftMetadataAccount(nftMetadata).
		SetNftMasterEditionAccount(nftMasterEdition).
		SetMplTokenMetadataAccount(mplTokenMetadata).
		SetTokenProgramAccount(tokenProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent).
		SetClockAccount(clock)
}
