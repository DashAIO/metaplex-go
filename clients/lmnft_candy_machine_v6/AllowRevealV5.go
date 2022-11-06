// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AllowRevealV5 is the `allowRevealV5` instruction.
type AllowRevealV5 struct {
	NewUri *string

	// [0] = [SIGNER] authority
	//
	// [1] = [WRITE] candyMachine
	//
	// [2] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAllowRevealV5InstructionBuilder creates a new `AllowRevealV5` instruction builder.
func NewAllowRevealV5InstructionBuilder() *AllowRevealV5 {
	nd := &AllowRevealV5{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetNewUri sets the "newUri" parameter.
func (inst *AllowRevealV5) SetNewUri(newUri string) *AllowRevealV5 {
	inst.NewUri = &newUri
	return inst
}

// SetAuthorityAccount sets the "authority" account.
func (inst *AllowRevealV5) SetAuthorityAccount(authority ag_solanago.PublicKey) *AllowRevealV5 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *AllowRevealV5) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *AllowRevealV5) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *AllowRevealV5 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *AllowRevealV5) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *AllowRevealV5) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *AllowRevealV5 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *AllowRevealV5) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst AllowRevealV5) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AllowRevealV5,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AllowRevealV5) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AllowRevealV5) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.NewUri == nil {
			return errors.New("NewUri parameter is not set")
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

func (inst *AllowRevealV5) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AllowRevealV5")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("NewUri", *inst.NewUri))
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

func (obj AllowRevealV5) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewUri` param:
	err = encoder.Encode(obj.NewUri)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AllowRevealV5) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewUri`:
	err = decoder.Decode(&obj.NewUri)
	if err != nil {
		return err
	}
	return nil
}

// NewAllowRevealV5Instruction declares a new AllowRevealV5 instruction with the provided parameters and accounts.
func NewAllowRevealV5Instruction(
	// Parameters:
	newUri string,
	// Accounts:
	authority ag_solanago.PublicKey,
	candyMachine ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *AllowRevealV5 {
	return NewAllowRevealV5InstructionBuilder().
		SetNewUri(newUri).
		SetAuthorityAccount(authority).
		SetCandyMachineAccount(candyMachine).
		SetSystemProgramAccount(systemProgram)
}
