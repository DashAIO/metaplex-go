// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package wlmarket

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// BuyItem is the `buyItem` instruction.
type BuyItem struct {
	Sixtynine *uint64
	Amount    *uint64
	Price     *uint64
	Fuckoff   *uint64

	// [0] = [WRITE] item
	//
	// [1] = [WRITE, SIGNER] signer
	//
	// [2] = [WRITE] owner
	//
	// [3] = [] mint
	//
	// [4] = [WRITE] payment
	//
	// [5] = [WRITE] fuckoff
	//
	// [6] = [WRITE] mintMarketAccount
	//
	// [7] = [WRITE] mintUserAccount
	//
	// [8] = [WRITE] foxy
	//
	// [9] = [WRITE] sellerFoxyAccount
	//
	// [10] = [WRITE] buyerFoxyAccount
	//
	// [11] = [] associatedTokenProgram
	//
	// [12] = [] tokenProgram
	//
	// [13] = [] systemProgram
	//
	// [14] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewBuyItemInstructionBuilder creates a new `BuyItem` instruction builder.
func NewBuyItemInstructionBuilder() *BuyItem {
	nd := &BuyItem{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 15),
	}
	return nd
}

// SetSixtynine sets the "sixtynine" parameter.
func (inst *BuyItem) SetSixtynine(sixtynine uint64) *BuyItem {
	inst.Sixtynine = &sixtynine
	return inst
}

// SetAmount sets the "amount" parameter.
func (inst *BuyItem) SetAmount(amount uint64) *BuyItem {
	inst.Amount = &amount
	return inst
}

// SetPrice sets the "price" parameter.
func (inst *BuyItem) SetPrice(price uint64) *BuyItem {
	inst.Price = &price
	return inst
}

// SetFuckoff sets the "fuckoff" parameter.
func (inst *BuyItem) SetFuckoff(fuckoff uint64) *BuyItem {
	inst.Fuckoff = &fuckoff
	return inst
}

// SetItemAccount sets the "item" account.
func (inst *BuyItem) SetItemAccount(item ag_solanago.PublicKey) *BuyItem {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(item).WRITE()
	return inst
}

// GetItemAccount gets the "item" account.
func (inst *BuyItem) GetItemAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSignerAccount sets the "signer" account.
func (inst *BuyItem) SetSignerAccount(signer ag_solanago.PublicKey) *BuyItem {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(signer).WRITE().SIGNER()
	return inst
}

// GetSignerAccount gets the "signer" account.
func (inst *BuyItem) GetSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOwnerAccount sets the "owner" account.
func (inst *BuyItem) SetOwnerAccount(owner ag_solanago.PublicKey) *BuyItem {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(owner).WRITE()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *BuyItem) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMintAccount sets the "mint" account.
func (inst *BuyItem) SetMintAccount(mint ag_solanago.PublicKey) *BuyItem {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *BuyItem) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPaymentAccount sets the "payment" account.
func (inst *BuyItem) SetPaymentAccount(payment ag_solanago.PublicKey) *BuyItem {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(payment).WRITE()
	return inst
}

// GetPaymentAccount gets the "payment" account.
func (inst *BuyItem) GetPaymentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetFuckoffAccount sets the "fuckoff" account.
func (inst *BuyItem) SetFuckoffAccount(fuckoff ag_solanago.PublicKey) *BuyItem {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(fuckoff).WRITE()
	return inst
}

// GetFuckoffAccount gets the "fuckoff" account.
func (inst *BuyItem) GetFuckoffAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMintMarketAccountAccount sets the "mintMarketAccount" account.
func (inst *BuyItem) SetMintMarketAccountAccount(mintMarketAccount ag_solanago.PublicKey) *BuyItem {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(mintMarketAccount).WRITE()
	return inst
}

// GetMintMarketAccountAccount gets the "mintMarketAccount" account.
func (inst *BuyItem) GetMintMarketAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetMintUserAccountAccount sets the "mintUserAccount" account.
func (inst *BuyItem) SetMintUserAccountAccount(mintUserAccount ag_solanago.PublicKey) *BuyItem {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(mintUserAccount).WRITE()
	return inst
}

// GetMintUserAccountAccount gets the "mintUserAccount" account.
func (inst *BuyItem) GetMintUserAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetFoxyAccount sets the "foxy" account.
func (inst *BuyItem) SetFoxyAccount(foxy ag_solanago.PublicKey) *BuyItem {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(foxy).WRITE()
	return inst
}

// GetFoxyAccount gets the "foxy" account.
func (inst *BuyItem) GetFoxyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSellerFoxyAccountAccount sets the "sellerFoxyAccount" account.
func (inst *BuyItem) SetSellerFoxyAccountAccount(sellerFoxyAccount ag_solanago.PublicKey) *BuyItem {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(sellerFoxyAccount).WRITE()
	return inst
}

// GetSellerFoxyAccountAccount gets the "sellerFoxyAccount" account.
func (inst *BuyItem) GetSellerFoxyAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetBuyerFoxyAccountAccount sets the "buyerFoxyAccount" account.
func (inst *BuyItem) SetBuyerFoxyAccountAccount(buyerFoxyAccount ag_solanago.PublicKey) *BuyItem {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(buyerFoxyAccount).WRITE()
	return inst
}

// GetBuyerFoxyAccountAccount gets the "buyerFoxyAccount" account.
func (inst *BuyItem) GetBuyerFoxyAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *BuyItem) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *BuyItem {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *BuyItem) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *BuyItem) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *BuyItem {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *BuyItem) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *BuyItem) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *BuyItem {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *BuyItem) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetRentAccount sets the "rent" account.
func (inst *BuyItem) SetRentAccount(rent ag_solanago.PublicKey) *BuyItem {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *BuyItem) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

func (inst BuyItem) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_BuyItem,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst BuyItem) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *BuyItem) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Sixtynine == nil {
			return errors.New("Sixtynine parameter is not set")
		}
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
		if inst.Price == nil {
			return errors.New("Price parameter is not set")
		}
		if inst.Fuckoff == nil {
			return errors.New("Fuckoff parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Item is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Signer is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Payment is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Fuckoff is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.MintMarketAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.MintUserAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Foxy is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SellerFoxyAccount is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.BuyerFoxyAccount is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *BuyItem) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("BuyItem")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=4]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Sixtynine", *inst.Sixtynine))
						paramsBranch.Child(ag_format.Param("   Amount", *inst.Amount))
						paramsBranch.Child(ag_format.Param("    Price", *inst.Price))
						paramsBranch.Child(ag_format.Param("  Fuckoff", *inst.Fuckoff))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=15]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                  item", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                signer", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                 owner", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                  mint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("               payment", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("               fuckoff", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("            mintMarket", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("              mintUser", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                  foxy", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("            sellerFoxy", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("             buyerFoxy", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("                  rent", inst.AccountMetaSlice.Get(14)))
					})
				})
		})
}

func (obj BuyItem) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Sixtynine` param:
	err = encoder.Encode(obj.Sixtynine)
	if err != nil {
		return err
	}
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
	// Serialize `Fuckoff` param:
	err = encoder.Encode(obj.Fuckoff)
	if err != nil {
		return err
	}
	return nil
}
func (obj *BuyItem) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Sixtynine`:
	err = decoder.Decode(&obj.Sixtynine)
	if err != nil {
		return err
	}
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
	// Deserialize `Fuckoff`:
	err = decoder.Decode(&obj.Fuckoff)
	if err != nil {
		return err
	}
	return nil
}

// NewBuyItemInstruction declares a new BuyItem instruction with the provided parameters and accounts.
func NewBuyItemInstruction(
	// Parameters:
	sixtynine uint64,
	amount uint64,
	price uint64,
	fuckoff uint64,
	// Accounts:
	item ag_solanago.PublicKey,
	signer ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	payment ag_solanago.PublicKey,
	fuckoffAccount ag_solanago.PublicKey,
	mintMarketAccount ag_solanago.PublicKey,
	mintUserAccount ag_solanago.PublicKey,
	foxy ag_solanago.PublicKey,
	sellerFoxyAccount ag_solanago.PublicKey,
	buyerFoxyAccount ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *BuyItem {
	return NewBuyItemInstructionBuilder().
		SetSixtynine(sixtynine).
		SetAmount(amount).
		SetPrice(price).
		SetFuckoff(fuckoff).
		SetItemAccount(item).
		SetSignerAccount(signer).
		SetOwnerAccount(owner).
		SetMintAccount(mint).
		SetPaymentAccount(payment).
		SetFuckoffAccount(fuckoffAccount).
		SetMintMarketAccountAccount(mintMarketAccount).
		SetMintUserAccountAccount(mintUserAccount).
		SetFoxyAccount(foxy).
		SetSellerFoxyAccountAccount(sellerFoxyAccount).
		SetBuyerFoxyAccountAccount(buyerFoxyAccount).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
