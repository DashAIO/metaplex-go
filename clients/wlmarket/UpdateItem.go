// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package wlmarket

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateItem is the `updateItem` instruction.
type UpdateItem struct {
	Cost *uint64

	// [0] = [WRITE] item
	//
	// [1] = [WRITE, SIGNER] owner
	//
	// [2] = [] mint
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateItemInstructionBuilder creates a new `UpdateItem` instruction builder.
func NewUpdateItemInstructionBuilder() *UpdateItem {
	nd := &UpdateItem{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetCost sets the "cost" parameter.
func (inst *UpdateItem) SetCost(cost uint64) *UpdateItem {
	inst.Cost = &cost
	return inst
}

// SetItemAccount sets the "item" account.
func (inst *UpdateItem) SetItemAccount(item ag_solanago.PublicKey) *UpdateItem {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(item).WRITE()
	return inst
}

// GetItemAccount gets the "item" account.
func (inst *UpdateItem) GetItemAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnerAccount sets the "owner" account.
func (inst *UpdateItem) SetOwnerAccount(owner ag_solanago.PublicKey) *UpdateItem {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *UpdateItem) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMintAccount sets the "mint" account.
func (inst *UpdateItem) SetMintAccount(mint ag_solanago.PublicKey) *UpdateItem {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *UpdateItem) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UpdateItem) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateItem,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateItem) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateItem) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Cost == nil {
			return errors.New("Cost parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Item is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Mint is not set")
		}
	}
	return nil
}

func (inst *UpdateItem) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateItem")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Cost", *inst.Cost))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta(" item", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("owner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta(" mint", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj UpdateItem) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Cost` param:
	err = encoder.Encode(obj.Cost)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateItem) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Cost`:
	err = decoder.Decode(&obj.Cost)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateItemInstruction declares a new UpdateItem instruction with the provided parameters and accounts.
func NewUpdateItemInstruction(
	// Parameters:
	cost uint64,
	// Accounts:
	item ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	mint ag_solanago.PublicKey) *UpdateItem {
	return NewUpdateItemInstructionBuilder().
		SetCost(cost).
		SetItemAccount(item).
		SetOwnerAccount(owner).
		SetMintAccount(mint)
}