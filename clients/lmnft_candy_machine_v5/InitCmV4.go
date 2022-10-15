// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitCmV4 is the `initCmV4` instruction.
type InitCmV4 struct {
	Data     *CandyMachineDataV2
	Seed     *ag_solanago.PublicKey
	ThawDate *int64 `bin:"optional"`

	// [0] = [WRITE] candyMachine
	//
	// [1] = [] wallet
	//
	// [2] = [SIGNER] authority
	//
	// [3] = [WRITE, SIGNER] payer
	//
	// [4] = [] systemProgram
	//
	// [5] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitCmV4InstructionBuilder creates a new `InitCmV4` instruction builder.
func NewInitCmV4InstructionBuilder() *InitCmV4 {
	nd := &InitCmV4{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetData sets the "data" parameter.
func (inst *InitCmV4) SetData(data CandyMachineDataV2) *InitCmV4 {
	inst.Data = &data
	return inst
}

// SetSeed sets the "seed" parameter.
func (inst *InitCmV4) SetSeed(seed ag_solanago.PublicKey) *InitCmV4 {
	inst.Seed = &seed
	return inst
}

// SetThawDate sets the "thawDate" parameter.
func (inst *InitCmV4) SetThawDate(thawDate int64) *InitCmV4 {
	inst.ThawDate = &thawDate
	return inst
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *InitCmV4) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *InitCmV4 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *InitCmV4) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetWalletAccount sets the "wallet" account.
func (inst *InitCmV4) SetWalletAccount(wallet ag_solanago.PublicKey) *InitCmV4 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(wallet)
	return inst
}

// GetWalletAccount gets the "wallet" account.
func (inst *InitCmV4) GetWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *InitCmV4) SetAuthorityAccount(authority ag_solanago.PublicKey) *InitCmV4 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *InitCmV4) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
func (inst *InitCmV4) SetPayerAccount(payer ag_solanago.PublicKey) *InitCmV4 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *InitCmV4) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitCmV4) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitCmV4 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitCmV4) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetRentAccount sets the "rent" account.
func (inst *InitCmV4) SetRentAccount(rent ag_solanago.PublicKey) *InitCmV4 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *InitCmV4) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst InitCmV4) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitCmV4,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitCmV4) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitCmV4) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Data == nil {
			return errors.New("Data parameter is not set")
		}
		if inst.Seed == nil {
			return errors.New("Seed parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Wallet is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *InitCmV4) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitCmV4")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("    Data", *inst.Data))
						paramsBranch.Child(ag_format.Param("    Seed", *inst.Seed))
						paramsBranch.Child(ag_format.Param("ThawDate (OPT)", inst.ThawDate))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta(" candyMachine", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       wallet", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("        payer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("         rent", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj InitCmV4) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	// Serialize `Seed` param:
	err = encoder.Encode(obj.Seed)
	if err != nil {
		return err
	}
	// Serialize `ThawDate` param (optional):
	{
		if obj.ThawDate == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.ThawDate)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (obj *InitCmV4) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	// Deserialize `Seed`:
	err = decoder.Decode(&obj.Seed)
	if err != nil {
		return err
	}
	// Deserialize `ThawDate` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.ThawDate)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewInitCmV4Instruction declares a new InitCmV4 instruction with the provided parameters and accounts.
func NewInitCmV4Instruction(
	// Parameters:
	data CandyMachineDataV2,
	seed ag_solanago.PublicKey,
	thawDate int64,
	// Accounts:
	candyMachine ag_solanago.PublicKey,
	wallet ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *InitCmV4 {
	return NewInitCmV4InstructionBuilder().
		SetData(data).
		SetSeed(seed).
		SetThawDate(thawDate).
		SetCandyMachineAccount(candyMachine).
		SetWalletAccount(wallet).
		SetAuthorityAccount(authority).
		SetPayerAccount(payer).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
