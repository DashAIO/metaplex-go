// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package wlmarket

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreateOrder is the `createOrder` instruction.
type CreateOrder struct {
	Bump   *uint8
	Cost   *uint64
	Amount *uint64
	Foxy   *bool
	Expiry *uint64

	// [0] = [WRITE] order
	//
	// [1] = [WRITE, SIGNER] owner
	//
	// [2] = [] mint
	//
	// [3] = [] systemProgram
	//
	// [4] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateOrderInstructionBuilder creates a new `CreateOrder` instruction builder.
func NewCreateOrderInstructionBuilder() *CreateOrder {
	nd := &CreateOrder{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetBump sets the "bump" parameter.
func (inst *CreateOrder) SetBump(bump uint8) *CreateOrder {
	inst.Bump = &bump
	return inst
}

// SetCost sets the "cost" parameter.
func (inst *CreateOrder) SetCost(cost uint64) *CreateOrder {
	inst.Cost = &cost
	return inst
}

// SetAmount sets the "amount" parameter.
func (inst *CreateOrder) SetAmount(amount uint64) *CreateOrder {
	inst.Amount = &amount
	return inst
}

// SetFoxy sets the "foxy" parameter.
func (inst *CreateOrder) SetFoxy(foxy bool) *CreateOrder {
	inst.Foxy = &foxy
	return inst
}

// SetExpiry sets the "expiry" parameter.
func (inst *CreateOrder) SetExpiry(expiry uint64) *CreateOrder {
	inst.Expiry = &expiry
	return inst
}

// SetOrderAccount sets the "order" account.
func (inst *CreateOrder) SetOrderAccount(order ag_solanago.PublicKey) *CreateOrder {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(order).WRITE()
	return inst
}

// GetOrderAccount gets the "order" account.
func (inst *CreateOrder) GetOrderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnerAccount sets the "owner" account.
func (inst *CreateOrder) SetOwnerAccount(owner ag_solanago.PublicKey) *CreateOrder {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *CreateOrder) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMintAccount sets the "mint" account.
func (inst *CreateOrder) SetMintAccount(mint ag_solanago.PublicKey) *CreateOrder {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *CreateOrder) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CreateOrder) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateOrder {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CreateOrder) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetRentAccount sets the "rent" account.
func (inst *CreateOrder) SetRentAccount(rent ag_solanago.PublicKey) *CreateOrder {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *CreateOrder) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst CreateOrder) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateOrder,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateOrder) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateOrder) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Bump == nil {
			return errors.New("Bump parameter is not set")
		}
		if inst.Cost == nil {
			return errors.New("Cost parameter is not set")
		}
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
		if inst.Foxy == nil {
			return errors.New("Foxy parameter is not set")
		}
		if inst.Expiry == nil {
			return errors.New("Expiry parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Order is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *CreateOrder) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateOrder")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=5]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  Bump", *inst.Bump))
						paramsBranch.Child(ag_format.Param("  Cost", *inst.Cost))
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
						paramsBranch.Child(ag_format.Param("  Foxy", *inst.Foxy))
						paramsBranch.Child(ag_format.Param("Expiry", *inst.Expiry))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        order", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        owner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         mint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("         rent", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj CreateOrder) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `Cost` param:
	err = encoder.Encode(obj.Cost)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	// Serialize `Foxy` param:
	err = encoder.Encode(obj.Foxy)
	if err != nil {
		return err
	}
	// Serialize `Expiry` param:
	err = encoder.Encode(obj.Expiry)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateOrder) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `Cost`:
	err = decoder.Decode(&obj.Cost)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	// Deserialize `Foxy`:
	err = decoder.Decode(&obj.Foxy)
	if err != nil {
		return err
	}
	// Deserialize `Expiry`:
	err = decoder.Decode(&obj.Expiry)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateOrderInstruction declares a new CreateOrder instruction with the provided parameters and accounts.
func NewCreateOrderInstruction(
	// Parameters:
	bump uint8,
	cost uint64,
	amount uint64,
	foxy bool,
	expiry uint64,
	// Accounts:
	order ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *CreateOrder {
	return NewCreateOrderInstructionBuilder().
		SetBump(bump).
		SetCost(cost).
		SetAmount(amount).
		SetFoxy(foxy).
		SetExpiry(expiry).
		SetOrderAccount(order).
		SetOwnerAccount(owner).
		SetMintAccount(mint).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
