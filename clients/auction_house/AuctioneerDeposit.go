// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction_house

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AuctioneerDeposit is the `auctioneerDeposit` instruction.
type AuctioneerDeposit struct {
	EscrowPaymentBump *uint8
	Amount            *uint64

	// [0] = [SIGNER] wallet
	//
	// [1] = [WRITE] paymentAccount
	//
	// [2] = [] transferAuthority
	//
	// [3] = [WRITE] escrowPaymentAccount
	//
	// [4] = [] treasuryMint
	//
	// [5] = [] authority
	//
	// [6] = [SIGNER] auctioneerAuthority
	//
	// [7] = [] auctionHouse
	//
	// [8] = [WRITE] auctionHouseFeeAccount
	//
	// [9] = [] ahAuctioneerPda
	//
	// [10] = [] tokenProgram
	//
	// [11] = [] systemProgram
	//
	// [12] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAuctioneerDepositInstructionBuilder creates a new `AuctioneerDeposit` instruction builder.
func NewAuctioneerDepositInstructionBuilder() *AuctioneerDeposit {
	nd := &AuctioneerDeposit{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 13),
	}
	return nd
}

// SetEscrowPaymentBump sets the "escrowPaymentBump" parameter.
func (inst *AuctioneerDeposit) SetEscrowPaymentBump(escrowPaymentBump uint8) *AuctioneerDeposit {
	inst.EscrowPaymentBump = &escrowPaymentBump
	return inst
}

// SetAmount sets the "amount" parameter.
func (inst *AuctioneerDeposit) SetAmount(amount uint64) *AuctioneerDeposit {
	inst.Amount = &amount
	return inst
}

// SetWalletAccount sets the "wallet" account.
func (inst *AuctioneerDeposit) SetWalletAccount(wallet ag_solanago.PublicKey) *AuctioneerDeposit {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(wallet).SIGNER()
	return inst
}

// GetWalletAccount gets the "wallet" account.
func (inst *AuctioneerDeposit) GetWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPaymentAccountAccount sets the "paymentAccount" account.
func (inst *AuctioneerDeposit) SetPaymentAccountAccount(paymentAccount ag_solanago.PublicKey) *AuctioneerDeposit {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(paymentAccount).WRITE()
	return inst
}

// GetPaymentAccountAccount gets the "paymentAccount" account.
func (inst *AuctioneerDeposit) GetPaymentAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTransferAuthorityAccount sets the "transferAuthority" account.
func (inst *AuctioneerDeposit) SetTransferAuthorityAccount(transferAuthority ag_solanago.PublicKey) *AuctioneerDeposit {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(transferAuthority)
	return inst
}

// GetTransferAuthorityAccount gets the "transferAuthority" account.
func (inst *AuctioneerDeposit) GetTransferAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetEscrowPaymentAccountAccount sets the "escrowPaymentAccount" account.
func (inst *AuctioneerDeposit) SetEscrowPaymentAccountAccount(escrowPaymentAccount ag_solanago.PublicKey) *AuctioneerDeposit {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(escrowPaymentAccount).WRITE()
	return inst
}

// GetEscrowPaymentAccountAccount gets the "escrowPaymentAccount" account.
func (inst *AuctioneerDeposit) GetEscrowPaymentAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTreasuryMintAccount sets the "treasuryMint" account.
func (inst *AuctioneerDeposit) SetTreasuryMintAccount(treasuryMint ag_solanago.PublicKey) *AuctioneerDeposit {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(treasuryMint)
	return inst
}

// GetTreasuryMintAccount gets the "treasuryMint" account.
func (inst *AuctioneerDeposit) GetTreasuryMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *AuctioneerDeposit) SetAuthorityAccount(authority ag_solanago.PublicKey) *AuctioneerDeposit {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *AuctioneerDeposit) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetAuctioneerAuthorityAccount sets the "auctioneerAuthority" account.
func (inst *AuctioneerDeposit) SetAuctioneerAuthorityAccount(auctioneerAuthority ag_solanago.PublicKey) *AuctioneerDeposit {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(auctioneerAuthority).SIGNER()
	return inst
}

// GetAuctioneerAuthorityAccount gets the "auctioneerAuthority" account.
func (inst *AuctioneerDeposit) GetAuctioneerAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetAuctionHouseAccount sets the "auctionHouse" account.
func (inst *AuctioneerDeposit) SetAuctionHouseAccount(auctionHouse ag_solanago.PublicKey) *AuctioneerDeposit {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(auctionHouse)
	return inst
}

// GetAuctionHouseAccount gets the "auctionHouse" account.
func (inst *AuctioneerDeposit) GetAuctionHouseAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetAuctionHouseFeeAccountAccount sets the "auctionHouseFeeAccount" account.
func (inst *AuctioneerDeposit) SetAuctionHouseFeeAccountAccount(auctionHouseFeeAccount ag_solanago.PublicKey) *AuctioneerDeposit {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(auctionHouseFeeAccount).WRITE()
	return inst
}

// GetAuctionHouseFeeAccountAccount gets the "auctionHouseFeeAccount" account.
func (inst *AuctioneerDeposit) GetAuctionHouseFeeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetAhAuctioneerPdaAccount sets the "ahAuctioneerPda" account.
func (inst *AuctioneerDeposit) SetAhAuctioneerPdaAccount(ahAuctioneerPda ag_solanago.PublicKey) *AuctioneerDeposit {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(ahAuctioneerPda)
	return inst
}

// GetAhAuctioneerPdaAccount gets the "ahAuctioneerPda" account.
func (inst *AuctioneerDeposit) GetAhAuctioneerPdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *AuctioneerDeposit) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *AuctioneerDeposit {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *AuctioneerDeposit) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *AuctioneerDeposit) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *AuctioneerDeposit {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *AuctioneerDeposit) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetRentAccount sets the "rent" account.
func (inst *AuctioneerDeposit) SetRentAccount(rent ag_solanago.PublicKey) *AuctioneerDeposit {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *AuctioneerDeposit) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

func (inst AuctioneerDeposit) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AuctioneerDeposit,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AuctioneerDeposit) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AuctioneerDeposit) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.EscrowPaymentBump == nil {
			return errors.New("EscrowPaymentBump parameter is not set")
		}
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Wallet is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.PaymentAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TransferAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.EscrowPaymentAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TreasuryMint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.AuctioneerAuthority is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.AuctionHouse is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.AuctionHouseFeeAccount is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.AhAuctioneerPda is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *AuctioneerDeposit) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AuctioneerDeposit")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("EscrowPaymentBump", *inst.EscrowPaymentBump))
						paramsBranch.Child(ag_format.Param("           Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=13]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             wallet", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("            payment", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("  transferAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      escrowPayment", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       treasuryMint", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("          authority", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("auctioneerAuthority", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("       auctionHouse", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("    auctionHouseFee", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("    ahAuctioneerPda", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("       tokenProgram", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("      systemProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("               rent", inst.AccountMetaSlice.Get(12)))
					})
				})
		})
}

func (obj AuctioneerDeposit) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `EscrowPaymentBump` param:
	err = encoder.Encode(obj.EscrowPaymentBump)
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
func (obj *AuctioneerDeposit) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `EscrowPaymentBump`:
	err = decoder.Decode(&obj.EscrowPaymentBump)
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

// NewAuctioneerDepositInstruction declares a new AuctioneerDeposit instruction with the provided parameters and accounts.
func NewAuctioneerDepositInstruction(
	// Parameters:
	escrowPaymentBump uint8,
	amount uint64,
	// Accounts:
	wallet ag_solanago.PublicKey,
	paymentAccount ag_solanago.PublicKey,
	transferAuthority ag_solanago.PublicKey,
	escrowPaymentAccount ag_solanago.PublicKey,
	treasuryMint ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	auctioneerAuthority ag_solanago.PublicKey,
	auctionHouse ag_solanago.PublicKey,
	auctionHouseFeeAccount ag_solanago.PublicKey,
	ahAuctioneerPda ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *AuctioneerDeposit {
	return NewAuctioneerDepositInstructionBuilder().
		SetEscrowPaymentBump(escrowPaymentBump).
		SetAmount(amount).
		SetWalletAccount(wallet).
		SetPaymentAccountAccount(paymentAccount).
		SetTransferAuthorityAccount(transferAuthority).
		SetEscrowPaymentAccountAccount(escrowPaymentAccount).
		SetTreasuryMintAccount(treasuryMint).
		SetAuthorityAccount(authority).
		SetAuctioneerAuthorityAccount(auctioneerAuthority).
		SetAuctionHouseAccount(auctionHouse).
		SetAuctionHouseFeeAccountAccount(auctionHouseFeeAccount).
		SetAhAuctioneerPdaAccount(ahAuctioneerPda).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}