// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// BurnSupplyV5 is the `burnSupplyV5` instruction.
type BurnSupplyV5 struct {
	PercentToBurn *uint8

	// [0] = [SIGNER] authority
	//
	// [1] = [WRITE] candyMachine
	//
	// [2] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewBurnSupplyV5InstructionBuilder creates a new `BurnSupplyV5` instruction builder.
func NewBurnSupplyV5InstructionBuilder() *BurnSupplyV5 {
	nd := &BurnSupplyV5{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetPercentToBurn sets the "percentToBurn" parameter.
func (inst *BurnSupplyV5) SetPercentToBurn(percentToBurn uint8) *BurnSupplyV5 {
	inst.PercentToBurn = &percentToBurn
	return inst
}

// SetAuthorityAccount sets the "authority" account.
func (inst *BurnSupplyV5) SetAuthorityAccount(authority ag_solanago.PublicKey) *BurnSupplyV5 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *BurnSupplyV5) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *BurnSupplyV5) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *BurnSupplyV5 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *BurnSupplyV5) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *BurnSupplyV5) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *BurnSupplyV5 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *BurnSupplyV5) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst BurnSupplyV5) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_BurnSupplyV5,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst BurnSupplyV5) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *BurnSupplyV5) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.PercentToBurn == nil {
			return errors.New("PercentToBurn parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *BurnSupplyV5) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("BurnSupplyV5")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("PercentToBurn", *inst.PercentToBurn))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta(" candyMachine", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj BurnSupplyV5) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PercentToBurn` param:
	err = encoder.Encode(obj.PercentToBurn)
	if err != nil {
		return err
	}
	return nil
}
func (obj *BurnSupplyV5) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PercentToBurn`:
	err = decoder.Decode(&obj.PercentToBurn)
	if err != nil {
		return err
	}
	return nil
}

// NewBurnSupplyV5Instruction declares a new BurnSupplyV5 instruction with the provided parameters and accounts.
func NewBurnSupplyV5Instruction(
	// Parameters:
	percentToBurn uint8,
	// Accounts:
	authority ag_solanago.PublicKey,
	candyMachine ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *BurnSupplyV5 {
	return NewBurnSupplyV5InstructionBuilder().
		SetPercentToBurn(percentToBurn).
		SetAuthorityAccount(authority).
		SetCandyMachineAccount(candyMachine).
		SetSystemProgramAccount(systemProgram)
}
