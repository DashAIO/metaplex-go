// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// BuyRaffleTicket is the `buyRaffleTicket` instruction.
type BuyRaffleTicket struct {
	WalletLimitBump  *uint8
	RaffleTicketBump *uint8
	EscrowBump       *uint8

	// [0] = [] config
	//
	// [1] = [WRITE] candyMachine
	//
	// [2] = [WRITE, SIGNER] payer
	//
	// [3] = [] launchStagesInfo
	//
	// [4] = [WRITE] raffleTicket
	//
	// [5] = [WRITE] raffleEscrow
	//
	// [6] = [WRITE] walletLimitInfo
	//
	// [7] = [] slotHashes
	//
	// [8] = [SIGNER] notary
	//
	// [9] = [] systemProgram
	//
	// [10] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewBuyRaffleTicketInstructionBuilder creates a new `BuyRaffleTicket` instruction builder.
func NewBuyRaffleTicketInstructionBuilder() *BuyRaffleTicket {
	nd := &BuyRaffleTicket{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetWalletLimitBump sets the "walletLimitBump" parameter.
func (inst *BuyRaffleTicket) SetWalletLimitBump(walletLimitBump uint8) *BuyRaffleTicket {
	inst.WalletLimitBump = &walletLimitBump
	return inst
}

// SetRaffleTicketBump sets the "raffleTicketBump" parameter.
func (inst *BuyRaffleTicket) SetRaffleTicketBump(raffleTicketBump uint8) *BuyRaffleTicket {
	inst.RaffleTicketBump = &raffleTicketBump
	return inst
}

// SetEscrowBump sets the "escrowBump" parameter.
func (inst *BuyRaffleTicket) SetEscrowBump(escrowBump uint8) *BuyRaffleTicket {
	inst.EscrowBump = &escrowBump
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *BuyRaffleTicket) SetConfigAccount(config ag_solanago.PublicKey) *BuyRaffleTicket {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *BuyRaffleTicket) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *BuyRaffleTicket) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *BuyRaffleTicket {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *BuyRaffleTicket) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
func (inst *BuyRaffleTicket) SetPayerAccount(payer ag_solanago.PublicKey) *BuyRaffleTicket {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *BuyRaffleTicket) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetLaunchStagesInfoAccount sets the "launchStagesInfo" account.
func (inst *BuyRaffleTicket) SetLaunchStagesInfoAccount(launchStagesInfo ag_solanago.PublicKey) *BuyRaffleTicket {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(launchStagesInfo)
	return inst
}

// GetLaunchStagesInfoAccount gets the "launchStagesInfo" account.
func (inst *BuyRaffleTicket) GetLaunchStagesInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetRaffleTicketAccount sets the "raffleTicket" account.
func (inst *BuyRaffleTicket) SetRaffleTicketAccount(raffleTicket ag_solanago.PublicKey) *BuyRaffleTicket {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(raffleTicket).WRITE()
	return inst
}

// GetRaffleTicketAccount gets the "raffleTicket" account.
func (inst *BuyRaffleTicket) GetRaffleTicketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetRaffleEscrowAccount sets the "raffleEscrow" account.
func (inst *BuyRaffleTicket) SetRaffleEscrowAccount(raffleEscrow ag_solanago.PublicKey) *BuyRaffleTicket {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(raffleEscrow).WRITE()
	return inst
}

// GetRaffleEscrowAccount gets the "raffleEscrow" account.
func (inst *BuyRaffleTicket) GetRaffleEscrowAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetWalletLimitInfoAccount sets the "walletLimitInfo" account.
func (inst *BuyRaffleTicket) SetWalletLimitInfoAccount(walletLimitInfo ag_solanago.PublicKey) *BuyRaffleTicket {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(walletLimitInfo).WRITE()
	return inst
}

// GetWalletLimitInfoAccount gets the "walletLimitInfo" account.
func (inst *BuyRaffleTicket) GetWalletLimitInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSlotHashesAccount sets the "slotHashes" account.
func (inst *BuyRaffleTicket) SetSlotHashesAccount(slotHashes ag_solanago.PublicKey) *BuyRaffleTicket {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(slotHashes)
	return inst
}

// GetSlotHashesAccount gets the "slotHashes" account.
func (inst *BuyRaffleTicket) GetSlotHashesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetNotaryAccount sets the "notary" account.
func (inst *BuyRaffleTicket) SetNotaryAccount(notary ag_solanago.PublicKey) *BuyRaffleTicket {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(notary).SIGNER()
	return inst
}

// GetNotaryAccount gets the "notary" account.
func (inst *BuyRaffleTicket) GetNotaryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *BuyRaffleTicket) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *BuyRaffleTicket {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *BuyRaffleTicket) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetRentAccount sets the "rent" account.
func (inst *BuyRaffleTicket) SetRentAccount(rent ag_solanago.PublicKey) *BuyRaffleTicket {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *BuyRaffleTicket) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst BuyRaffleTicket) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_BuyRaffleTicket,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst BuyRaffleTicket) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *BuyRaffleTicket) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.WalletLimitBump == nil {
			return errors.New("WalletLimitBump parameter is not set")
		}
		if inst.RaffleTicketBump == nil {
			return errors.New("RaffleTicketBump parameter is not set")
		}
		if inst.EscrowBump == nil {
			return errors.New("EscrowBump parameter is not set")
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
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.LaunchStagesInfo is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.RaffleTicket is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.RaffleEscrow is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.WalletLimitInfo is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SlotHashes is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Notary is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *BuyRaffleTicket) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("BuyRaffleTicket")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" WalletLimitBump", *inst.WalletLimitBump))
						paramsBranch.Child(ag_format.Param("RaffleTicketBump", *inst.RaffleTicketBump))
						paramsBranch.Child(ag_format.Param("      EscrowBump", *inst.EscrowBump))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("          config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    candyMachine", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("           payer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("launchStagesInfo", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("    raffleTicket", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("    raffleEscrow", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta(" walletLimitInfo", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("      slotHashes", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("          notary", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("   systemProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("            rent", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj BuyRaffleTicket) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `WalletLimitBump` param:
	err = encoder.Encode(obj.WalletLimitBump)
	if err != nil {
		return err
	}
	// Serialize `RaffleTicketBump` param:
	err = encoder.Encode(obj.RaffleTicketBump)
	if err != nil {
		return err
	}
	// Serialize `EscrowBump` param:
	err = encoder.Encode(obj.EscrowBump)
	if err != nil {
		return err
	}
	return nil
}
func (obj *BuyRaffleTicket) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `WalletLimitBump`:
	err = decoder.Decode(&obj.WalletLimitBump)
	if err != nil {
		return err
	}
	// Deserialize `RaffleTicketBump`:
	err = decoder.Decode(&obj.RaffleTicketBump)
	if err != nil {
		return err
	}
	// Deserialize `EscrowBump`:
	err = decoder.Decode(&obj.EscrowBump)
	if err != nil {
		return err
	}
	return nil
}

// NewBuyRaffleTicketInstruction declares a new BuyRaffleTicket instruction with the provided parameters and accounts.
func NewBuyRaffleTicketInstruction(
	// Parameters:
	walletLimitBump uint8,
	raffleTicketBump uint8,
	escrowBump uint8,
	// Accounts:
	config ag_solanago.PublicKey,
	candyMachine ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	launchStagesInfo ag_solanago.PublicKey,
	raffleTicket ag_solanago.PublicKey,
	raffleEscrow ag_solanago.PublicKey,
	walletLimitInfo ag_solanago.PublicKey,
	slotHashes ag_solanago.PublicKey,
	notary ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *BuyRaffleTicket {
	return NewBuyRaffleTicketInstructionBuilder().
		SetWalletLimitBump(walletLimitBump).
		SetRaffleTicketBump(raffleTicketBump).
		SetEscrowBump(escrowBump).
		SetConfigAccount(config).
		SetCandyMachineAccount(candyMachine).
		SetPayerAccount(payer).
		SetLaunchStagesInfoAccount(launchStagesInfo).
		SetRaffleTicketAccount(raffleTicket).
		SetRaffleEscrowAccount(raffleEscrow).
		SetWalletLimitInfoAccount(walletLimitInfo).
		SetSlotHashesAccount(slotHashes).
		SetNotaryAccount(notary).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
