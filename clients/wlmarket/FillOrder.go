// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package wlmarket

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FillOrder is the `fillOrder` instruction.
type FillOrder struct {
	Amount *uint64
	Price  *uint64

	// [0] = [WRITE] order
	//
	// [1] = [] mint
	//
	// [2] = [WRITE] owner
	//
	// [3] = [WRITE, SIGNER] seller
	//
	// [4] = [WRITE] sellerAccount
	//
	// [5] = [WRITE] buyerAccount
	//
	// [6] = [WRITE] payment
	//
	// [7] = [] associatedTokenProgram
	//
	// [8] = [] tokenProgram
	//
	// [9] = [] systemProgram
	//
	// [10] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewFillOrderInstructionBuilder creates a new `FillOrder` instruction builder.
func NewFillOrderInstructionBuilder() *FillOrder {
	nd := &FillOrder{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
func (inst *FillOrder) SetAmount(amount uint64) *FillOrder {
	inst.Amount = &amount
	return inst
}

// SetPrice sets the "price" parameter.
func (inst *FillOrder) SetPrice(price uint64) *FillOrder {
	inst.Price = &price
	return inst
}

// SetOrderAccount sets the "order" account.
func (inst *FillOrder) SetOrderAccount(order ag_solanago.PublicKey) *FillOrder {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(order).WRITE()
	return inst
}

// GetOrderAccount gets the "order" account.
func (inst *FillOrder) GetOrderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMintAccount sets the "mint" account.
func (inst *FillOrder) SetMintAccount(mint ag_solanago.PublicKey) *FillOrder {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *FillOrder) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOwnerAccount sets the "owner" account.
func (inst *FillOrder) SetOwnerAccount(owner ag_solanago.PublicKey) *FillOrder {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(owner).WRITE()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *FillOrder) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSellerAccount sets the "seller" account.
func (inst *FillOrder) SetSellerAccount(seller ag_solanago.PublicKey) *FillOrder {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(seller).WRITE().SIGNER()
	return inst
}

// GetSellerAccount gets the "seller" account.
func (inst *FillOrder) GetSellerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSellerAccountAccount sets the "sellerAccount" account.
func (inst *FillOrder) SetSellerAccountAccount(sellerAccount ag_solanago.PublicKey) *FillOrder {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(sellerAccount).WRITE()
	return inst
}

// GetSellerAccountAccount gets the "sellerAccount" account.
func (inst *FillOrder) GetSellerAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetBuyerAccountAccount sets the "buyerAccount" account.
func (inst *FillOrder) SetBuyerAccountAccount(buyerAccount ag_solanago.PublicKey) *FillOrder {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(buyerAccount).WRITE()
	return inst
}

// GetBuyerAccountAccount gets the "buyerAccount" account.
func (inst *FillOrder) GetBuyerAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPaymentAccount sets the "payment" account.
func (inst *FillOrder) SetPaymentAccount(payment ag_solanago.PublicKey) *FillOrder {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(payment).WRITE()
	return inst
}

// GetPaymentAccount gets the "payment" account.
func (inst *FillOrder) GetPaymentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *FillOrder) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *FillOrder {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *FillOrder) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *FillOrder) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *FillOrder {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *FillOrder) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *FillOrder) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *FillOrder {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *FillOrder) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetRentAccount sets the "rent" account.
func (inst *FillOrder) SetRentAccount(rent ag_solanago.PublicKey) *FillOrder {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *FillOrder) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst FillOrder) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FillOrder,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FillOrder) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FillOrder) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
		if inst.Price == nil {
			return errors.New("Price parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Order is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Seller is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SellerAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.BuyerAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Payment is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *FillOrder) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FillOrder")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
						paramsBranch.Child(ag_format.Param(" Price", *inst.Price))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                 order", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                  mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                 owner", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                seller", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                seller", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                 buyer", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("               payment", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                  rent", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj FillOrder) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	return nil
}
func (obj *FillOrder) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	return nil
}

// NewFillOrderInstruction declares a new FillOrder instruction with the provided parameters and accounts.
func NewFillOrderInstruction(
	// Parameters:
	amount uint64,
	price uint64,
	// Accounts:
	order ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	seller ag_solanago.PublicKey,
	sellerAccount ag_solanago.PublicKey,
	buyerAccount ag_solanago.PublicKey,
	payment ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *FillOrder {
	return NewFillOrderInstructionBuilder().
		SetAmount(amount).
		SetPrice(price).
		SetOrderAccount(order).
		SetMintAccount(mint).
		SetOwnerAccount(owner).
		SetSellerAccount(seller).
		SetSellerAccountAccount(sellerAccount).
		SetBuyerAccountAccount(buyerAccount).
		SetPaymentAccount(payment).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
