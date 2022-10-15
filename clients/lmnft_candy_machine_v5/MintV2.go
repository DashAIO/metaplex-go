// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// MintV2 is the `mintV2` instruction.
type MintV2 struct {
	CreatorBump    *uint8
	TotalMintsBump *uint8
	Proof          *[][32]uint8 `bin:"optional"`

	// [0] = [WRITE] candyMachine
	//
	// [1] = [WRITE, SIGNER] payer
	//
	// [2] = [] candyMachineCreator
	//
	// [3] = [WRITE] wallet
	//
	// [4] = [WRITE] wallet2
	//
	// [5] = [WRITE] metadata
	//
	// [6] = [WRITE] mint
	//
	// [7] = [WRITE] masterEdition
	//
	// [8] = [WRITE] totalMints
	//
	// [9] = [] tokenMetadataProgram
	//
	// [10] = [] tokenProgram
	//
	// [11] = [] systemProgram
	//
	// [12] = [] rent
	//
	// [13] = [] clock
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewMintV2InstructionBuilder creates a new `MintV2` instruction builder.
func NewMintV2InstructionBuilder() *MintV2 {
	nd := &MintV2{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 14),
	}
	return nd
}

// SetCreatorBump sets the "creatorBump" parameter.
func (inst *MintV2) SetCreatorBump(creatorBump uint8) *MintV2 {
	inst.CreatorBump = &creatorBump
	return inst
}

// SetTotalMintsBump sets the "totalMintsBump" parameter.
func (inst *MintV2) SetTotalMintsBump(totalMintsBump uint8) *MintV2 {
	inst.TotalMintsBump = &totalMintsBump
	return inst
}

// SetProof sets the "proof" parameter.
func (inst *MintV2) SetProof(proof [][32]uint8) *MintV2 {
	inst.Proof = &proof
	return inst
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *MintV2) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *MintV2 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *MintV2) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPayerAccount sets the "payer" account.
func (inst *MintV2) SetPayerAccount(payer ag_solanago.PublicKey) *MintV2 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *MintV2) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetCandyMachineCreatorAccount sets the "candyMachineCreator" account.
func (inst *MintV2) SetCandyMachineCreatorAccount(candyMachineCreator ag_solanago.PublicKey) *MintV2 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(candyMachineCreator)
	return inst
}

// GetCandyMachineCreatorAccount gets the "candyMachineCreator" account.
func (inst *MintV2) GetCandyMachineCreatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetWalletAccount sets the "wallet" account.
func (inst *MintV2) SetWalletAccount(wallet ag_solanago.PublicKey) *MintV2 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(wallet).WRITE()
	return inst
}

// GetWalletAccount gets the "wallet" account.
func (inst *MintV2) GetWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetWallet2Account sets the "wallet2" account.
func (inst *MintV2) SetWallet2Account(wallet2 ag_solanago.PublicKey) *MintV2 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(wallet2).WRITE()
	return inst
}

// GetWallet2Account gets the "wallet2" account.
func (inst *MintV2) GetWallet2Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *MintV2) SetMetadataAccount(metadata ag_solanago.PublicKey) *MintV2 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *MintV2) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMintAccount sets the "mint" account.
func (inst *MintV2) SetMintAccount(mint ag_solanago.PublicKey) *MintV2 {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(mint).WRITE()
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *MintV2) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetMasterEditionAccount sets the "masterEdition" account.
func (inst *MintV2) SetMasterEditionAccount(masterEdition ag_solanago.PublicKey) *MintV2 {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(masterEdition).WRITE()
	return inst
}

// GetMasterEditionAccount gets the "masterEdition" account.
func (inst *MintV2) GetMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTotalMintsAccount sets the "totalMints" account.
func (inst *MintV2) SetTotalMintsAccount(totalMints ag_solanago.PublicKey) *MintV2 {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(totalMints).WRITE()
	return inst
}

// GetTotalMintsAccount gets the "totalMints" account.
func (inst *MintV2) GetTotalMintsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
func (inst *MintV2) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *MintV2 {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
func (inst *MintV2) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *MintV2) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *MintV2 {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *MintV2) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *MintV2) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *MintV2 {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *MintV2) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetRentAccount sets the "rent" account.
func (inst *MintV2) SetRentAccount(rent ag_solanago.PublicKey) *MintV2 {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *MintV2) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetClockAccount sets the "clock" account.
func (inst *MintV2) SetClockAccount(clock ag_solanago.PublicKey) *MintV2 {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(clock)
	return inst
}

// GetClockAccount gets the "clock" account.
func (inst *MintV2) GetClockAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

func (inst MintV2) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_MintV2,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MintV2) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MintV2) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.CreatorBump == nil {
			return errors.New("CreatorBump parameter is not set")
		}
		if inst.TotalMintsBump == nil {
			return errors.New("TotalMintsBump parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.CandyMachineCreator is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Wallet is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Wallet2 is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.MasterEdition is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TotalMints is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.Clock is not set")
		}
	}
	return nil
}

func (inst *MintV2) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MintV2")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("   CreatorBump", *inst.CreatorBump))
						paramsBranch.Child(ag_format.Param("TotalMintsBump", *inst.TotalMintsBump))
						paramsBranch.Child(ag_format.Param("         Proof (OPT)", inst.Proof))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=14]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        candyMachine", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("               payer", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta(" candyMachineCreator", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("              wallet", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("             wallet2", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("            metadata", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                mint", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("       masterEdition", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("          totalMints", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("tokenMetadataProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("        tokenProgram", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("       systemProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("                rent", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("               clock", inst.AccountMetaSlice.Get(13)))
					})
				})
		})
}

func (obj MintV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `CreatorBump` param:
	err = encoder.Encode(obj.CreatorBump)
	if err != nil {
		return err
	}
	// Serialize `TotalMintsBump` param:
	err = encoder.Encode(obj.TotalMintsBump)
	if err != nil {
		return err
	}
	// Serialize `Proof` param (optional):
	{
		if obj.Proof == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Proof)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (obj *MintV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `CreatorBump`:
	err = decoder.Decode(&obj.CreatorBump)
	if err != nil {
		return err
	}
	// Deserialize `TotalMintsBump`:
	err = decoder.Decode(&obj.TotalMintsBump)
	if err != nil {
		return err
	}
	// Deserialize `Proof` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Proof)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewMintV2Instruction declares a new MintV2 instruction with the provided parameters and accounts.
func NewMintV2Instruction(
	// Parameters:
	creatorBump uint8,
	totalMintsBump uint8,
	proof [][32]uint8,
	// Accounts:
	candyMachine ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	candyMachineCreator ag_solanago.PublicKey,
	wallet ag_solanago.PublicKey,
	wallet2 ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	masterEdition ag_solanago.PublicKey,
	totalMints ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	clock ag_solanago.PublicKey) *MintV2 {
	return NewMintV2InstructionBuilder().
		SetCreatorBump(creatorBump).
		SetTotalMintsBump(totalMintsBump).
		SetProof(proof).
		SetCandyMachineAccount(candyMachine).
		SetPayerAccount(payer).
		SetCandyMachineCreatorAccount(candyMachineCreator).
		SetWalletAccount(wallet).
		SetWallet2Account(wallet2).
		SetMetadataAccount(metadata).
		SetMintAccount(mint).
		SetMasterEditionAccount(masterEdition).
		SetTotalMintsAccount(totalMints).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent).
		SetClockAccount(clock)
}
