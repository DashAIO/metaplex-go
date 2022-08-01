// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// MintNft is the `mintNft` instruction.
type MintNft struct {
	WalletLimitBump *uint8
	InOrder         *bool
	UserLimit       *uint8 `bin:"optional"`
	CurrTime	*int64

	// [0] = [] config
	//
	// [1] = [WRITE] candyMachine
	//
	// [2] = [] mintReceiver
	//
	// [3] = [WRITE, SIGNER] payer
	//
	// [4] = [WRITE] launchStagesInfo
	//
	// [5] = [WRITE] payFrom
	//
	// [6] = [WRITE] payTo
	//
	// [7] = [] notary
	//
	// [8] = [WRITE] metadata
	//
	// [9] = [WRITE, SIGNER] mint
	//
	// [10] = [WRITE] tokenAta
	//
	// [11] = [WRITE] masterEdition
	//
	// [12] = [WRITE] walletLimitInfo
	//
	// [13] = [WRITE] orderInfo
	//
	// [14] = [] slotHashes
	//
	// [15] = [] tokenMetadataProgram
	//
	// [16] = [] tokenProgram
	//
	// [17] = [] systemProgram
	//
	// [18] = [] associatedTokenProgram
	//
	// [19] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewMintNftInstructionBuilder creates a new `MintNft` instruction builder.
func NewMintNftInstructionBuilder() *MintNft {
	nd := &MintNft{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 21),
	}
	return nd
}

// SetWalletLimitBump sets the "walletLimitBump" parameter.
func (inst *MintNft) SetWalletLimitBump(walletLimitBump uint8) *MintNft {
	inst.WalletLimitBump = &walletLimitBump
	return inst
}

// SetInOrder sets the "inOrder" parameter.
func (inst *MintNft) SetInOrder(inOrder bool) *MintNft {
	inst.InOrder = &inOrder
	return inst
}

// SetUserLimit sets the "userLimit" parameter.
func (inst *MintNft) SetUserLimit(userLimit uint8) *MintNft {
	inst.UserLimit = &userLimit
	return inst
}
// SetCurrTime sets the "currTime" parameter.
func (inst *MintNft) SetCurrTime(currTime int64) *MintNft {
	inst.CurrTime = &currTime
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *MintNft) SetConfigAccount(config ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *MintNft) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *MintNft) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *MintNft) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMintReceiverAccount sets the "mintReceiver" account.
func (inst *MintNft) SetMintReceiverAccount(mintReceiver ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(mintReceiver)
	return inst
}

// GetMintReceiverAccount gets the "mintReceiver" account.
func (inst *MintNft) GetMintReceiverAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
func (inst *MintNft) SetPayerAccount(payer ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *MintNft) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetLaunchStagesInfoAccount sets the "launchStagesInfo" account.
func (inst *MintNft) SetLaunchStagesInfoAccount(launchStagesInfo ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(launchStagesInfo).WRITE()
	return inst
}

// GetLaunchStagesInfoAccount gets the "launchStagesInfo" account.
func (inst *MintNft) GetLaunchStagesInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetPayFromAccount sets the "payFrom" account.
func (inst *MintNft) SetPayFromAccount(payFrom ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(payFrom).WRITE()
	return inst
}

// GetPayFromAccount gets the "payFrom" account.
func (inst *MintNft) GetPayFromAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPayToAccount sets the "payTo" account.
func (inst *MintNft) SetPayToAccount(payTo ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(payTo).WRITE()
	return inst
}

// GetPayToAccount gets the "payTo" account.
func (inst *MintNft) GetPayToAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetNotaryAccount sets the "notary" account.
func (inst *MintNft) SetNotaryAccount(notary ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(notary)
	return inst
}

// GetNotaryAccount gets the "notary" account.
func (inst *MintNft) GetNotaryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *MintNft) SetMetadataAccount(metadata ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *MintNft) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetMintAccount sets the "mint" account.
func (inst *MintNft) SetMintAccount(mint ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(mint).WRITE().SIGNER()
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *MintNft) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenAtaAccount sets the "tokenAta" account.
func (inst *MintNft) SetTokenAtaAccount(tokenAta ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenAta).WRITE()
	return inst
}

// GetTokenAtaAccount gets the "tokenAta" account.
func (inst *MintNft) GetTokenAtaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetMasterEditionAccount sets the "masterEdition" account.
func (inst *MintNft) SetMasterEditionAccount(masterEdition ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(masterEdition).WRITE()
	return inst
}

// GetMasterEditionAccount gets the "masterEdition" account.
func (inst *MintNft) GetMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetWalletLimitInfoAccount sets the "walletLimitInfo" account.
func (inst *MintNft) SetWalletLimitInfoAccount(walletLimitInfo ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(walletLimitInfo).WRITE()
	return inst
}

// GetWalletLimitInfoAccount gets the "walletLimitInfo" account.
func (inst *MintNft) GetWalletLimitInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetOrderInfoAccount sets the "orderInfo" account.
func (inst *MintNft) SetOrderInfoAccount(orderInfo ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(orderInfo).WRITE()
	return inst
}

// GetOrderInfoAccount gets the "orderInfo" account.
func (inst *MintNft) GetOrderInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetSlotHashesAccount sets the "slotHashes" account.
func (inst *MintNft) SetSlotHashesAccount(slotHashes ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(slotHashes)
	return inst
}

// GetSlotHashesAccount gets the "slotHashes" account.
func (inst *MintNft) GetSlotHashesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
func (inst *MintNft) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
func (inst *MintNft) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *MintNft) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *MintNft) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *MintNft) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *MintNft) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *MintNft) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *MintNft) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(18)
}

// SetRentAccount sets the "rent" account.
func (inst *MintNft) SetRentAccount(rent ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[19] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *MintNft) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(19)
}

func (inst *MintNft) SetRemainingAccounts(pk []ag_solanago.AccountMeta) *MintNft {
	amount := len(pk)
	if amount == 0 {
		return inst
	}
	for i := 0; i < amount; i++ {
		inst.AccountMetaSlice[20+i] = &pk[i]
	}
	return inst
}


func (inst MintNft) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_MintNft,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MintNft) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MintNft) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.WalletLimitBump == nil {
			return errors.New("WalletLimitBump parameter is not set")
		}
		if inst.InOrder == nil {
			return errors.New("InOrder parameter is not set")
		}
		if inst.CurrTime == nil {
			return errors.New("CurrTime parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.MintReceiver is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.LaunchStagesInfo is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.PayFrom is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.PayTo is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Notary is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TokenAta is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.MasterEdition is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.WalletLimitInfo is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.OrderInfo is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.SlotHashes is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[19] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *MintNft) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MintNft")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("WalletLimitBump", *inst.WalletLimitBump))
						paramsBranch.Child(ag_format.Param("        InOrder", *inst.InOrder))
						paramsBranch.Child(ag_format.Param("      UserLimit (OPT)", inst.UserLimit))
						paramsBranch.Child(ag_format.Param("CurrTime", *inst.CurrTime))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=20]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          candyMachine", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("          mintReceiver", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                 payer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("      launchStagesInfo", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("               payFrom", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                 payTo", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                notary", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("              metadata", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                  mint", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("              tokenAta", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("         masterEdition", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("       walletLimitInfo", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("             orderInfo", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("            slotHashes", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("  tokenMetadataProgram", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(17)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(18)))
						accountsBranch.Child(ag_format.Meta("                  rent", inst.AccountMetaSlice.Get(19)))
					})
				})
		})
}

func (obj MintNft) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `WalletLimitBump` param:
	err = encoder.Encode(obj.WalletLimitBump)
	if err != nil {
		return err
	}
	// Serialize `InOrder` param:
	err = encoder.Encode(obj.InOrder)
	if err != nil {
		return err
	}
	// Serialize `UserLimit` param (optional):
	{
		if obj.UserLimit == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.UserLimit)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `CurrTime` param:
	err = encoder.Encode(obj.CurrTime)
	if err != nil {
		return err
	}
	return nil
}
func (obj *MintNft) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `WalletLimitBump`:
	err = decoder.Decode(&obj.WalletLimitBump)
	if err != nil {
		return err
	}
	// Deserialize `InOrder`:
	err = decoder.Decode(&obj.InOrder)
	if err != nil {
		return err
	}
	// Deserialize `UserLimit` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.UserLimit)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `CurrTime`:
	err = decoder.Decode(&obj.CurrTime)
	if err != nil {
		return err
	}
	return nil
}

// NewMintNftInstruction declares a new MintNft instruction with the provided parameters and accounts.
func NewMintNftInstruction(
	// Parameters:
	walletLimitBump uint8,
	inOrder bool,
	userLimit uint8,
	currTime int64,
	// Accounts:
	config ag_solanago.PublicKey,
	candyMachine ag_solanago.PublicKey,
	mintReceiver ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	launchStagesInfo ag_solanago.PublicKey,
	payFrom ag_solanago.PublicKey,
	payTo ag_solanago.PublicKey,
	notary ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	tokenAta ag_solanago.PublicKey,
	masterEdition ag_solanago.PublicKey,
	walletLimitInfo ag_solanago.PublicKey,
	orderInfo ag_solanago.PublicKey,
	slotHashes ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	remainingAccounts []ag_solanago.AccountMeta) *MintNft {
	return NewMintNftInstructionBuilder().
		SetWalletLimitBump(walletLimitBump).
		SetInOrder(inOrder).
		SetUserLimit(userLimit).
		SetCurrTime(currTime).
		SetConfigAccount(config).
		SetCandyMachineAccount(candyMachine).
		SetMintReceiverAccount(mintReceiver).
		SetPayerAccount(payer).
		SetLaunchStagesInfoAccount(launchStagesInfo).
		SetPayFromAccount(payFrom).
		SetPayToAccount(payTo).
		SetNotaryAccount(notary).
		SetMetadataAccount(metadata).
		SetMintAccount(mint).
		SetTokenAtaAccount(tokenAta).
		SetMasterEditionAccount(masterEdition).
		SetWalletLimitInfoAccount(walletLimitInfo).
		SetOrderInfoAccount(orderInfo).
		SetSlotHashesAccount(slotHashes).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetRentAccount(rent).
		SetRemainingAccounts(remainingAccounts)
}
