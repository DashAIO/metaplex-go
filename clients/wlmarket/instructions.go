// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package wlmarket

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "Wlmarket"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_AddItem = ag_binary.TypeID([8]byte{225, 38, 79, 147, 116, 142, 147, 57})

	Instruction_UpdateItem = ag_binary.TypeID([8]byte{28, 222, 44, 175, 216, 228, 171, 184})

	Instruction_ListItem = ag_binary.TypeID([8]byte{174, 245, 22, 211, 228, 103, 121, 13})

	Instruction_DelistItem = ag_binary.TypeID([8]byte{243, 224, 175, 39, 170, 10, 179, 114})

	Instruction_CloseItem = ag_binary.TypeID([8]byte{232, 80, 56, 56, 194, 171, 176, 235})

	Instruction_BuyItem = ag_binary.TypeID([8]byte{80, 82, 193, 201, 216, 27, 70, 184})

	Instruction_KickItem = ag_binary.TypeID([8]byte{28, 209, 188, 254, 32, 170, 103, 12})

	Instruction_CreateOrder = ag_binary.TypeID([8]byte{141, 54, 37, 207, 237, 210, 250, 215})

	Instruction_CancelOrder = ag_binary.TypeID([8]byte{95, 129, 237, 240, 8, 49, 223, 132})

	Instruction_UpdateOrder = ag_binary.TypeID([8]byte{54, 8, 208, 207, 34, 134, 239, 168})

	Instruction_FillOrder = ag_binary.TypeID([8]byte{232, 122, 115, 25, 199, 143, 136, 162})

	Instruction_CloseOrder = ag_binary.TypeID([8]byte{90, 103, 209, 28, 7, 63, 168, 4})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_AddItem:
		return "AddItem"
	case Instruction_UpdateItem:
		return "UpdateItem"
	case Instruction_ListItem:
		return "ListItem"
	case Instruction_DelistItem:
		return "DelistItem"
	case Instruction_CloseItem:
		return "CloseItem"
	case Instruction_BuyItem:
		return "BuyItem"
	case Instruction_KickItem:
		return "KickItem"
	case Instruction_CreateOrder:
		return "CreateOrder"
	case Instruction_CancelOrder:
		return "CancelOrder"
	case Instruction_UpdateOrder:
		return "UpdateOrder"
	case Instruction_FillOrder:
		return "FillOrder"
	case Instruction_CloseOrder:
		return "CloseOrder"
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
	ag_binary.AnchorTypeIDEncoding,
	[]ag_binary.VariantType{
		{
			"add_item", (*AddItem)(nil),
		},
		{
			"update_item", (*UpdateItem)(nil),
		},
		{
			"list_item", (*ListItem)(nil),
		},
		{
			"delist_item", (*DelistItem)(nil),
		},
		{
			"close_item", (*CloseItem)(nil),
		},
		{
			"buy_item", (*BuyItem)(nil),
		},
		{
			"kick_item", (*KickItem)(nil),
		},
		{
			"create_order", (*CreateOrder)(nil),
		},
		{
			"cancel_order", (*CancelOrder)(nil),
		},
		{
			"update_order", (*UpdateOrder)(nil),
		},
		{
			"fill_order", (*FillOrder)(nil),
		},
		{
			"close_order", (*CloseOrder)(nil),
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
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
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
