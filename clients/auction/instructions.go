// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_text "github.com/desperatee/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey = ag_solanago.MustPublicKeyFromBase58("auctxRXPeJoc4817jDhf4HbjnhEcr1cCXenosMhK5R8")

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "Auction"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

const (
	// Cancel a bid on a running auction.
	Instruction_CancelBid uint8 = iota

	// Create a new auction account bound to a resource, initially in a pending state.
	Instruction_CreateAuction

	// Move SPL tokens from winning bid to the destination account.
	Instruction_ClaimBid

	// Ends an auction, regardless of end timing conditions
	Instruction_EndAuction

	// Start an inactive auction.
	Instruction_StartAuction

	// Update the authority for an auction account.
	Instruction_SetAuthority

	// Place a bid on a running auction.
	Instruction_PlaceBid

	// Create a new auction account bound to a resource, initially in a pending state.
	// The only one difference with above instruction it's additional parameters in CreateAuctionArgsV2
	Instruction_CreateAuctionV2
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id uint8) string {
	switch id {
	case Instruction_CancelBid:
		return "CancelBid"
	case Instruction_CreateAuction:
		return "CreateAuction"
	case Instruction_ClaimBid:
		return "ClaimBid"
	case Instruction_EndAuction:
		return "EndAuction"
	case Instruction_StartAuction:
		return "StartAuction"
	case Instruction_SetAuthority:
		return "SetAuthority"
	case Instruction_PlaceBid:
		return "PlaceBid"
	case Instruction_CreateAuctionV2:
		return "CreateAuctionV2"
	default:
		return ""
	}
}

type Instruction struct {
	ag_binary.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent ag_treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = ag_binary.NewVariantDefinition(
	ag_binary.Uint8TypeIDEncoding,
	[]ag_binary.VariantType{
		{
			"CancelBid", (*CancelBid)(nil),
		},
		{
			"CreateAuction", (*CreateAuction)(nil),
		},
		{
			"ClaimBid", (*ClaimBid)(nil),
		},
		{
			"EndAuction", (*EndAuction)(nil),
		},
		{
			"StartAuction", (*StartAuction)(nil),
		},
		{
			"SetAuthority", (*SetAuthority)(nil),
		},
		{
			"PlaceBid", (*PlaceBid)(nil),
		},
		{
			"CreateAuctionV2", (*CreateAuctionV2)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() ag_solanago.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*ag_solanago.AccountMeta) {
	return inst.Impl.(ag_solanago.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ag_binary.NewBorshEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *ag_text.Encoder, option *ag_text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst *Instruction) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	err := encoder.WriteUint8(inst.TypeID.Uint8())
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := ag_binary.NewBorshDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(ag_solanago.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}
