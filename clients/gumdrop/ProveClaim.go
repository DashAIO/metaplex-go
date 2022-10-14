// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package gumdrop

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ProveClaim is the `proveClaim` instruction.
type ProveClaim struct {
	ClaimPrefix    *[]byte
	ClaimBump      *uint8
	Index          *uint64
	Amount         *uint64
	ClaimantSecret *ag_solanago.PublicKey
	Resource       *ag_solanago.PublicKey
	ResourceNonce  *[]byte
	Proof          *[][32]uint8

	// [0] = [WRITE] distributor
	//
	// [1] = [WRITE] claimProof
	//
	// [2] = [SIGNER] temporal
	//
	// [3] = [WRITE, SIGNER] payer
	//
	// [4] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewProveClaimInstructionBuilder creates a new `ProveClaim` instruction builder.
func NewProveClaimInstructionBuilder() *ProveClaim {
	nd := &ProveClaim{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetClaimPrefix sets the "claimPrefix" parameter.
func (inst *ProveClaim) SetClaimPrefix(claimPrefix []byte) *ProveClaim {
	inst.ClaimPrefix = &claimPrefix
	return inst
}

// SetClaimBump sets the "claimBump" parameter.
func (inst *ProveClaim) SetClaimBump(claimBump uint8) *ProveClaim {
	inst.ClaimBump = &claimBump
	return inst
}

// SetIndex sets the "index" parameter.
func (inst *ProveClaim) SetIndex(index uint64) *ProveClaim {
	inst.Index = &index
	return inst
}

// SetAmount sets the "amount" parameter.
func (inst *ProveClaim) SetAmount(amount uint64) *ProveClaim {
	inst.Amount = &amount
	return inst
}

// SetClaimantSecret sets the "claimantSecret" parameter.
func (inst *ProveClaim) SetClaimantSecret(claimantSecret ag_solanago.PublicKey) *ProveClaim {
	inst.ClaimantSecret = &claimantSecret
	return inst
}

// SetResource sets the "resource" parameter.
func (inst *ProveClaim) SetResource(resource ag_solanago.PublicKey) *ProveClaim {
	inst.Resource = &resource
	return inst
}

// SetResourceNonce sets the "resourceNonce" parameter.
func (inst *ProveClaim) SetResourceNonce(resourceNonce []byte) *ProveClaim {
	inst.ResourceNonce = &resourceNonce
	return inst
}

// SetProof sets the "proof" parameter.
func (inst *ProveClaim) SetProof(proof [][32]uint8) *ProveClaim {
	inst.Proof = &proof
	return inst
}

// SetDistributorAccount sets the "distributor" account.
func (inst *ProveClaim) SetDistributorAccount(distributor ag_solanago.PublicKey) *ProveClaim {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(distributor).WRITE()
	return inst
}

// GetDistributorAccount gets the "distributor" account.
func (inst *ProveClaim) GetDistributorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetClaimProofAccount sets the "claimProof" account.
func (inst *ProveClaim) SetClaimProofAccount(claimProof ag_solanago.PublicKey) *ProveClaim {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(claimProof).WRITE()
	return inst
}

// GetClaimProofAccount gets the "claimProof" account.
func (inst *ProveClaim) GetClaimProofAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTemporalAccount sets the "temporal" account.
func (inst *ProveClaim) SetTemporalAccount(temporal ag_solanago.PublicKey) *ProveClaim {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(temporal).SIGNER()
	return inst
}

// GetTemporalAccount gets the "temporal" account.
func (inst *ProveClaim) GetTemporalAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
func (inst *ProveClaim) SetPayerAccount(payer ag_solanago.PublicKey) *ProveClaim {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *ProveClaim) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *ProveClaim) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ProveClaim {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *ProveClaim) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst ProveClaim) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ProveClaim,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ProveClaim) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ProveClaim) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ClaimPrefix == nil {
			return errors.New("ClaimPrefix parameter is not set")
		}
		if inst.ClaimBump == nil {
			return errors.New("ClaimBump parameter is not set")
		}
		if inst.Index == nil {
			return errors.New("Index parameter is not set")
		}
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
		if inst.ClaimantSecret == nil {
			return errors.New("ClaimantSecret parameter is not set")
		}
		if inst.Resource == nil {
			return errors.New("Resource parameter is not set")
		}
		if inst.ResourceNonce == nil {
			return errors.New("ResourceNonce parameter is not set")
		}
		if inst.Proof == nil {
			return errors.New("Proof parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Distributor is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ClaimProof is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Temporal is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *ProveClaim) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ProveClaim")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=8]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("   ClaimPrefix", *inst.ClaimPrefix))
						paramsBranch.Child(ag_format.Param("     ClaimBump", *inst.ClaimBump))
						paramsBranch.Child(ag_format.Param("         Index", *inst.Index))
						paramsBranch.Child(ag_format.Param("        Amount", *inst.Amount))
						paramsBranch.Child(ag_format.Param("ClaimantSecret", *inst.ClaimantSecret))
						paramsBranch.Child(ag_format.Param("      Resource", *inst.Resource))
						paramsBranch.Child(ag_format.Param(" ResourceNonce", *inst.ResourceNonce))
						paramsBranch.Child(ag_format.Param("         Proof", *inst.Proof))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("  distributor", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   claimProof", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     temporal", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("        payer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj ProveClaim) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ClaimPrefix` param:
	err = encoder.Encode(obj.ClaimPrefix)
	if err != nil {
		return err
	}
	// Serialize `ClaimBump` param:
	err = encoder.Encode(obj.ClaimBump)
	if err != nil {
		return err
	}
	// Serialize `Index` param:
	err = encoder.Encode(obj.Index)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	// Serialize `ClaimantSecret` param:
	err = encoder.Encode(obj.ClaimantSecret)
	if err != nil {
		return err
	}
	// Serialize `Resource` param:
	err = encoder.Encode(obj.Resource)
	if err != nil {
		return err
	}
	// Serialize `ResourceNonce` param:
	err = encoder.Encode(obj.ResourceNonce)
	if err != nil {
		return err
	}
	// Serialize `Proof` param:
	err = encoder.Encode(obj.Proof)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ProveClaim) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ClaimPrefix`:
	err = decoder.Decode(&obj.ClaimPrefix)
	if err != nil {
		return err
	}
	// Deserialize `ClaimBump`:
	err = decoder.Decode(&obj.ClaimBump)
	if err != nil {
		return err
	}
	// Deserialize `Index`:
	err = decoder.Decode(&obj.Index)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	// Deserialize `ClaimantSecret`:
	err = decoder.Decode(&obj.ClaimantSecret)
	if err != nil {
		return err
	}
	// Deserialize `Resource`:
	err = decoder.Decode(&obj.Resource)
	if err != nil {
		return err
	}
	// Deserialize `ResourceNonce`:
	err = decoder.Decode(&obj.ResourceNonce)
	if err != nil {
		return err
	}
	// Deserialize `Proof`:
	err = decoder.Decode(&obj.Proof)
	if err != nil {
		return err
	}
	return nil
}

// NewProveClaimInstruction declares a new ProveClaim instruction with the provided parameters and accounts.
func NewProveClaimInstruction(
	// Parameters:
	claimPrefix []byte,
	claimBump uint8,
	index uint64,
	amount uint64,
	claimantSecret ag_solanago.PublicKey,
	resource ag_solanago.PublicKey,
	resourceNonce []byte,
	proof [][32]uint8,
	// Accounts:
	distributor ag_solanago.PublicKey,
	claimProof ag_solanago.PublicKey,
	temporal ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *ProveClaim {
	return NewProveClaimInstructionBuilder().
		SetClaimPrefix(claimPrefix).
		SetClaimBump(claimBump).
		SetIndex(index).
		SetAmount(amount).
		SetClaimantSecret(claimantSecret).
		SetResource(resource).
		SetResourceNonce(resourceNonce).
		SetProof(proof).
		SetDistributorAccount(distributor).
		SetClaimProofAccount(claimProof).
		SetTemporalAccount(temporal).
		SetPayerAccount(payer).
		SetSystemProgramAccount(systemProgram)
}