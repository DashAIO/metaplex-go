// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AllowReveal is the `allowReveal` instruction.
type AllowReveal struct {
	NewUri *string

	// [0] = [SIGNER] authority
	//
	// [1] = [WRITE] candyMachine
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAllowRevealInstructionBuilder creates a new `AllowReveal` instruction builder.
func NewAllowRevealInstructionBuilder() *AllowReveal {
	nd := &AllowReveal{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetNewUri sets the "newUri" parameter.
func (inst *AllowReveal) SetNewUri(newUri string) *AllowReveal {
	inst.NewUri = &newUri
	return inst
}

// SetAuthorityAccount sets the "authority" account.
func (inst *AllowReveal) SetAuthorityAccount(authority ag_solanago.PublicKey) *AllowReveal {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *AllowReveal) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *AllowReveal) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *AllowReveal {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *AllowReveal) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst AllowReveal) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AllowReveal,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AllowReveal) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AllowReveal) Validate() error {
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
	}
	return nil
}

func (inst *AllowReveal) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AllowReveal")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("NewUri", *inst.NewUri))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   authority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("candyMachine", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj AllowReveal) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewUri` param:
	err = encoder.Encode(obj.NewUri)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AllowReveal) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewUri`:
	err = decoder.Decode(&obj.NewUri)
	if err != nil {
		return err
	}
	return nil
}

// NewAllowRevealInstruction declares a new AllowReveal instruction with the provided parameters and accounts.
func NewAllowRevealInstruction(
	// Parameters:
	newUri string,
	// Accounts:
	authority ag_solanago.PublicKey,
	candyMachine ag_solanago.PublicKey) *AllowReveal {
	return NewAllowRevealInstructionBuilder().
		SetNewUri(newUri).
		SetAuthorityAccount(authority).
		SetCandyMachineAccount(candyMachine)
}