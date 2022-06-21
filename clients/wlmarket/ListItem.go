// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package wlmarket

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ListItem is the `listItem` instruction.
type ListItem struct {
	Amount *uint64
	Cost   *uint64

	// [0] = [WRITE] item
	//
	// [1] = [WRITE, SIGNER] owner
	//
	// [2] = [] mint
	//
	// [3] = [WRITE] mintAccount
	//
	// [4] = [WRITE] mintMarketAccount
	//
	// [5] = [] associatedTokenProgram
	//
	// [6] = [] tokenProgram
	//
	// [7] = [] systemProgram
	//
	// [8] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewListItemInstructionBuilder creates a new `ListItem` instruction builder.
func NewListItemInstructionBuilder() *ListItem {
	nd := &ListItem{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
func (inst *ListItem) SetAmount(amount uint64) *ListItem {
	inst.Amount = &amount
	return inst
}

// SetCost sets the "cost" parameter.
func (inst *ListItem) SetCost(cost uint64) *ListItem {
	inst.Cost = &cost
	return inst
}

// SetItemAccount sets the "item" account.
func (inst *ListItem) SetItemAccount(item ag_solanago.PublicKey) *ListItem {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(item).WRITE()
	return inst
}

// GetItemAccount gets the "item" account.
func (inst *ListItem) GetItemAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnerAccount sets the "owner" account.
func (inst *ListItem) SetOwnerAccount(owner ag_solanago.PublicKey) *ListItem {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *ListItem) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMintAccount sets the "mint" account.
func (inst *ListItem) SetMintAccount(mint ag_solanago.PublicKey) *ListItem {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *ListItem) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMintAccountAccount sets the "mintAccount" account.
func (inst *ListItem) SetMintAccountAccount(mintAccount ag_solanago.PublicKey) *ListItem {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mintAccount).WRITE()
	return inst
}

// GetMintAccountAccount gets the "mintAccount" account.
func (inst *ListItem) GetMintAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMintMarketAccountAccount sets the "mintMarketAccount" account.
func (inst *ListItem) SetMintMarketAccountAccount(mintMarketAccount ag_solanago.PublicKey) *ListItem {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(mintMarketAccount).WRITE()
	return inst
}

// GetMintMarketAccountAccount gets the "mintMarketAccount" account.
func (inst *ListItem) GetMintMarketAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *ListItem) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *ListItem {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *ListItem) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *ListItem) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *ListItem {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *ListItem) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *ListItem) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ListItem {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *ListItem) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetRentAccount sets the "rent" account.
func (inst *ListItem) SetRentAccount(rent ag_solanago.PublicKey) *ListItem {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *ListItem) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst ListItem) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ListItem,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ListItem) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ListItem) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
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
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.MintAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.MintMarketAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *ListItem) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ListItem")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
						paramsBranch.Child(ag_format.Param("  Cost", *inst.Cost))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                  item", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                 owner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                  mint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                  mint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("            mintMarket", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                  rent", inst.AccountMetaSlice.Get(8)))
					})
				})
		})
}

func (obj ListItem) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	// Serialize `Cost` param:
	err = encoder.Encode(obj.Cost)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ListItem) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	// Deserialize `Cost`:
	err = decoder.Decode(&obj.Cost)
	if err != nil {
		return err
	}
	return nil
}

// NewListItemInstruction declares a new ListItem instruction with the provided parameters and accounts.
func NewListItemInstruction(
	// Parameters:
	amount uint64,
	cost uint64,
	// Accounts:
	item ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	mintAccount ag_solanago.PublicKey,
	mintMarketAccount ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *ListItem {
	return NewListItemInstructionBuilder().
		SetAmount(amount).
		SetCost(cost).
		SetItemAccount(item).
		SetOwnerAccount(owner).
		SetMintAccount(mint).
		SetMintAccountAccount(mintAccount).
		SetMintMarketAccountAccount(mintMarketAccount).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
