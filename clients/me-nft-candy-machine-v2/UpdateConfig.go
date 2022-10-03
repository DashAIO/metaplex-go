// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateConfig is the `updateConfig` instruction.
type UpdateConfig struct {
	Params *UpdateConfigArgs

	// [0] = [WRITE] config
	//
	// [1] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateConfigInstructionBuilder creates a new `UpdateConfig` instruction builder.
func NewUpdateConfigInstructionBuilder() *UpdateConfig {
	nd := &UpdateConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetParams sets the "params" parameter.
func (inst *UpdateConfig) SetParams(params UpdateConfigArgs) *UpdateConfig {
	inst.Params = &params
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *UpdateConfig) SetConfigAccount(config ag_solanago.PublicKey) *UpdateConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *UpdateConfig) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *UpdateConfig) SetAuthorityAccount(authority ag_solanago.PublicKey) *UpdateConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *UpdateConfig) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst UpdateConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateConfig) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("Params parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *UpdateConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj UpdateConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateConfigInstruction declares a new UpdateConfig instruction with the provided parameters and accounts.
func NewUpdateConfigInstruction(
	// Parameters:
	params UpdateConfigArgs,
	// Accounts:
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *UpdateConfig {
	return NewUpdateConfigInstructionBuilder().
		SetParams(params).
		SetConfigAccount(config).
		SetAuthorityAccount(authority)
}
