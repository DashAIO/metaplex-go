// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine_v2

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AddConfigLines is the `addConfigLines` instruction.
type AddConfigLines struct {
	Index       *uint32
	ConfigLines *[]ConfigLine

	// [0] = [WRITE] candyMachine
	//
	// [1] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAddConfigLinesInstructionBuilder creates a new `AddConfigLines` instruction builder.
func NewAddConfigLinesInstructionBuilder() *AddConfigLines {
	nd := &AddConfigLines{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetIndex sets the "index" parameter.
func (inst *AddConfigLines) SetIndex(index uint32) *AddConfigLines {
	inst.Index = &index
	return inst
}

// SetConfigLines sets the "configLines" parameter.
func (inst *AddConfigLines) SetConfigLines(configLines []ConfigLine) *AddConfigLines {
	inst.ConfigLines = &configLines
	return inst
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *AddConfigLines) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *AddConfigLines {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *AddConfigLines) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *AddConfigLines) SetAuthorityAccount(authority ag_solanago.PublicKey) *AddConfigLines {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *AddConfigLines) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst AddConfigLines) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AddConfigLines,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AddConfigLines) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AddConfigLines) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Index == nil {
			return errors.New("Index parameter is not set")
		}
		if inst.ConfigLines == nil {
			return errors.New("ConfigLines parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *AddConfigLines) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AddConfigLines")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("      Index", *inst.Index))
						paramsBranch.Child(ag_format.Param("ConfigLines", *inst.ConfigLines))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("candyMachine", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   authority", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj AddConfigLines) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Index` param:
	err = encoder.Encode(obj.Index)
	if err != nil {
		return err
	}
	// Serialize `ConfigLines` param:
	err = encoder.Encode(obj.ConfigLines)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AddConfigLines) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Index`:
	err = decoder.Decode(&obj.Index)
	if err != nil {
		return err
	}
	// Deserialize `ConfigLines`:
	err = decoder.Decode(&obj.ConfigLines)
	if err != nil {
		return err
	}
	return nil
}

// NewAddConfigLinesInstruction declares a new AddConfigLines instruction with the provided parameters and accounts.
func NewAddConfigLinesInstruction(
	// Parameters:
	index uint32,
	configLines []ConfigLine,
	// Accounts:
	candyMachine ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *AddConfigLines {
	return NewAddConfigLinesInstructionBuilder().
		SetIndex(index).
		SetConfigLines(configLines).
		SetCandyMachineAccount(candyMachine).
		SetAuthorityAccount(authority)
}
