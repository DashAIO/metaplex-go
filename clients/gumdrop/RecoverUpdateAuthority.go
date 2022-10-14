// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package gumdrop

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RecoverUpdateAuthority is the `recoverUpdateAuthority` instruction.
type RecoverUpdateAuthority struct {
	Bump       *uint8
	WalletBump *uint8

	// [0] = [SIGNER] base
	//
	// [1] = [] distributor
	//
	// [2] = [] distributorWallet
	//
	// [3] = [] newUpdateAuthority
	//
	// [4] = [WRITE] metadata
	//
	// [5] = [] systemProgram
	//
	// [6] = [] tokenMetadataProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRecoverUpdateAuthorityInstructionBuilder creates a new `RecoverUpdateAuthority` instruction builder.
func NewRecoverUpdateAuthorityInstructionBuilder() *RecoverUpdateAuthority {
	nd := &RecoverUpdateAuthority{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetBump sets the "bump" parameter.
func (inst *RecoverUpdateAuthority) SetBump(bump uint8) *RecoverUpdateAuthority {
	inst.Bump = &bump
	return inst
}

// SetWalletBump sets the "walletBump" parameter.
func (inst *RecoverUpdateAuthority) SetWalletBump(walletBump uint8) *RecoverUpdateAuthority {
	inst.WalletBump = &walletBump
	return inst
}

// SetBaseAccount sets the "base" account.
func (inst *RecoverUpdateAuthority) SetBaseAccount(base ag_solanago.PublicKey) *RecoverUpdateAuthority {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(base).SIGNER()
	return inst
}

// GetBaseAccount gets the "base" account.
func (inst *RecoverUpdateAuthority) GetBaseAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetDistributorAccount sets the "distributor" account.
func (inst *RecoverUpdateAuthority) SetDistributorAccount(distributor ag_solanago.PublicKey) *RecoverUpdateAuthority {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(distributor)
	return inst
}

// GetDistributorAccount gets the "distributor" account.
func (inst *RecoverUpdateAuthority) GetDistributorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetDistributorWalletAccount sets the "distributorWallet" account.
func (inst *RecoverUpdateAuthority) SetDistributorWalletAccount(distributorWallet ag_solanago.PublicKey) *RecoverUpdateAuthority {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(distributorWallet)
	return inst
}

// GetDistributorWalletAccount gets the "distributorWallet" account.
func (inst *RecoverUpdateAuthority) GetDistributorWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetNewUpdateAuthorityAccount sets the "newUpdateAuthority" account.
func (inst *RecoverUpdateAuthority) SetNewUpdateAuthorityAccount(newUpdateAuthority ag_solanago.PublicKey) *RecoverUpdateAuthority {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(newUpdateAuthority)
	return inst
}

// GetNewUpdateAuthorityAccount gets the "newUpdateAuthority" account.
func (inst *RecoverUpdateAuthority) GetNewUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *RecoverUpdateAuthority) SetMetadataAccount(metadata ag_solanago.PublicKey) *RecoverUpdateAuthority {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *RecoverUpdateAuthority) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *RecoverUpdateAuthority) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *RecoverUpdateAuthority {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *RecoverUpdateAuthority) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
func (inst *RecoverUpdateAuthority) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *RecoverUpdateAuthority {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
func (inst *RecoverUpdateAuthority) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst RecoverUpdateAuthority) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RecoverUpdateAuthority,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RecoverUpdateAuthority) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RecoverUpdateAuthority) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Bump == nil {
			return errors.New("Bump parameter is not set")
		}
		if inst.WalletBump == nil {
			return errors.New("WalletBump parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Base is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Distributor is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.DistributorWallet is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.NewUpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
	}
	return nil
}

func (inst *RecoverUpdateAuthority) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RecoverUpdateAuthority")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("      Bump", *inst.Bump))
						paramsBranch.Child(ag_format.Param("WalletBump", *inst.WalletBump))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                base", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         distributor", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("   distributorWallet", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("  newUpdateAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("            metadata", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("       systemProgram", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("tokenMetadataProgram", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj RecoverUpdateAuthority) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `WalletBump` param:
	err = encoder.Encode(obj.WalletBump)
	if err != nil {
		return err
	}
	return nil
}
func (obj *RecoverUpdateAuthority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `WalletBump`:
	err = decoder.Decode(&obj.WalletBump)
	if err != nil {
		return err
	}
	return nil
}

// NewRecoverUpdateAuthorityInstruction declares a new RecoverUpdateAuthority instruction with the provided parameters and accounts.
func NewRecoverUpdateAuthorityInstruction(
	// Parameters:
	bump uint8,
	walletBump uint8,
	// Accounts:
	base ag_solanago.PublicKey,
	distributor ag_solanago.PublicKey,
	distributorWallet ag_solanago.PublicKey,
	newUpdateAuthority ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey) *RecoverUpdateAuthority {
	return NewRecoverUpdateAuthorityInstructionBuilder().
		SetBump(bump).
		SetWalletBump(walletBump).
		SetBaseAccount(base).
		SetDistributorAccount(distributor).
		SetDistributorWalletAccount(distributorWallet).
		SetNewUpdateAuthorityAccount(newUpdateAuthority).
		SetMetadataAccount(metadata).
		SetSystemProgramAccount(systemProgram).
		SetTokenMetadataProgramAccount(tokenMetadataProgram)
}