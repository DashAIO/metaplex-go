// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction_house

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateAuctioneer is the `updateAuctioneer` instruction.
type UpdateAuctioneer struct {
	Scopes *[]AuthorityScope

	// [0] = [WRITE] auctionHouse
	//
	// [1] = [WRITE, SIGNER] authority
	//
	// [2] = [] auctioneerAuthority
	//
	// [3] = [WRITE] ahAuctioneerPda
	//
	// [4] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateAuctioneerInstructionBuilder creates a new `UpdateAuctioneer` instruction builder.
func NewUpdateAuctioneerInstructionBuilder() *UpdateAuctioneer {
	nd := &UpdateAuctioneer{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetScopes sets the "scopes" parameter.
func (inst *UpdateAuctioneer) SetScopes(scopes []AuthorityScope) *UpdateAuctioneer {
	inst.Scopes = &scopes
	return inst
}

// SetAuctionHouseAccount sets the "auctionHouse" account.
func (inst *UpdateAuctioneer) SetAuctionHouseAccount(auctionHouse ag_solanago.PublicKey) *UpdateAuctioneer {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(auctionHouse).WRITE()
	return inst
}

// GetAuctionHouseAccount gets the "auctionHouse" account.
func (inst *UpdateAuctioneer) GetAuctionHouseAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *UpdateAuctioneer) SetAuthorityAccount(authority ag_solanago.PublicKey) *UpdateAuctioneer {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *UpdateAuctioneer) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuctioneerAuthorityAccount sets the "auctioneerAuthority" account.
func (inst *UpdateAuctioneer) SetAuctioneerAuthorityAccount(auctioneerAuthority ag_solanago.PublicKey) *UpdateAuctioneer {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(auctioneerAuthority)
	return inst
}

// GetAuctioneerAuthorityAccount gets the "auctioneerAuthority" account.
func (inst *UpdateAuctioneer) GetAuctioneerAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAhAuctioneerPdaAccount sets the "ahAuctioneerPda" account.
func (inst *UpdateAuctioneer) SetAhAuctioneerPdaAccount(ahAuctioneerPda ag_solanago.PublicKey) *UpdateAuctioneer {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(ahAuctioneerPda).WRITE()
	return inst
}

// GetAhAuctioneerPdaAccount gets the "ahAuctioneerPda" account.
func (inst *UpdateAuctioneer) GetAhAuctioneerPdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *UpdateAuctioneer) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UpdateAuctioneer {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *UpdateAuctioneer) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst UpdateAuctioneer) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateAuctioneer,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateAuctioneer) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateAuctioneer) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Scopes == nil {
			return errors.New("Scopes parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.AuctionHouse is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.AuctioneerAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.AhAuctioneerPda is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *UpdateAuctioneer) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateAuctioneer")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Scopes", *inst.Scopes))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       auctionHouse", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("auctioneerAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("    ahAuctioneerPda", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("      systemProgram", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj UpdateAuctioneer) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Scopes` param:
	err = encoder.Encode(obj.Scopes)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateAuctioneer) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Scopes`:
	err = decoder.Decode(&obj.Scopes)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateAuctioneerInstruction declares a new UpdateAuctioneer instruction with the provided parameters and accounts.
func NewUpdateAuctioneerInstruction(
	// Parameters:
	scopes []AuthorityScope,
	// Accounts:
	auctionHouse ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	auctioneerAuthority ag_solanago.PublicKey,
	ahAuctioneerPda ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *UpdateAuctioneer {
	return NewUpdateAuctioneerInstructionBuilder().
		SetScopes(scopes).
		SetAuctionHouseAccount(auctionHouse).
		SetAuthorityAccount(authority).
		SetAuctioneerAuthorityAccount(auctioneerAuthority).
		SetAhAuctioneerPdaAccount(ahAuctioneerPda).
		SetSystemProgramAccount(systemProgram)
}
