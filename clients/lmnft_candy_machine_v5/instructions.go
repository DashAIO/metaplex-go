// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

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

const ProgramName = "NftCandyMachine"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_MintEmbed = ag_binary.TypeID([8]byte{133, 221, 27, 96, 4, 65, 206, 186})

	Instruction_MintV5 = ag_binary.TypeID([8]byte{96, 114, 250, 235, 203, 205, 188, 36})

	Instruction_MintV4 = ag_binary.TypeID([8]byte{79, 71, 207, 32, 178, 102, 21, 19})

	Instruction_MintV2 = ag_binary.TypeID([8]byte{120, 121, 23, 146, 173, 110, 199, 205})

	Instruction_EditCmV3 = ag_binary.TypeID([8]byte{85, 87, 21, 33, 241, 184, 109, 132})

	Instruction_EditCmV2 = ag_binary.TypeID([8]byte{12, 115, 248, 70, 118, 97, 136, 163})

	Instruction_EditCmV4 = ag_binary.TypeID([8]byte{113, 108, 161, 132, 26, 55, 205, 242})

	Instruction_EditCmV5 = ag_binary.TypeID([8]byte{43, 25, 89, 26, 165, 209, 229, 92})

	Instruction_RevealV5 = ag_binary.TypeID([8]byte{57, 46, 3, 81, 70, 219, 168, 49})

	Instruction_Reveal = ag_binary.TypeID([8]byte{9, 35, 59, 190, 167, 249, 76, 115})

	Instruction_AllowUnfreeze = ag_binary.TypeID([8]byte{8, 32, 46, 11, 209, 211, 248, 81})

	Instruction_AllowReveal = ag_binary.TypeID([8]byte{215, 62, 72, 171, 41, 172, 244, 151})

	Instruction_AllowUnfreezeV5 = ag_binary.TypeID([8]byte{156, 1, 166, 135, 249, 139, 95, 54})

	Instruction_AllowRevealV5 = ag_binary.TypeID([8]byte{13, 155, 159, 197, 5, 211, 116, 91})

	Instruction_BurnSupplyV5 = ag_binary.TypeID([8]byte{107, 75, 172, 207, 53, 177, 22, 255})

	Instruction_BurnSupplyV3 = ag_binary.TypeID([8]byte{99, 9, 108, 6, 55, 0, 161, 118})

	Instruction_BurnFrozen = ag_binary.TypeID([8]byte{115, 215, 82, 186, 58, 93, 69, 24})

	Instruction_Thaw = ag_binary.TypeID([8]byte{226, 249, 34, 57, 189, 21, 177, 101})

	Instruction_ThawV5 = ag_binary.TypeID([8]byte{181, 47, 178, 189, 46, 187, 145, 248})

	Instruction_InitCmV5 = ag_binary.TypeID([8]byte{179, 164, 16, 182, 42, 27, 249, 75})

	Instruction_InitCmV4 = ag_binary.TypeID([8]byte{109, 204, 198, 181, 178, 68, 11, 239})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_MintEmbed:
		return "MintEmbed"
	case Instruction_MintV5:
		return "MintV5"
	case Instruction_MintV4:
		return "MintV4"
	case Instruction_MintV2:
		return "MintV2"
	case Instruction_EditCmV3:
		return "EditCmV3"
	case Instruction_EditCmV2:
		return "EditCmV2"
	case Instruction_EditCmV4:
		return "EditCmV4"
	case Instruction_EditCmV5:
		return "EditCmV5"
	case Instruction_RevealV5:
		return "RevealV5"
	case Instruction_Reveal:
		return "Reveal"
	case Instruction_AllowUnfreeze:
		return "AllowUnfreeze"
	case Instruction_AllowReveal:
		return "AllowReveal"
	case Instruction_AllowUnfreezeV5:
		return "AllowUnfreezeV5"
	case Instruction_AllowRevealV5:
		return "AllowRevealV5"
	case Instruction_BurnSupplyV5:
		return "BurnSupplyV5"
	case Instruction_BurnSupplyV3:
		return "BurnSupplyV3"
	case Instruction_BurnFrozen:
		return "BurnFrozen"
	case Instruction_Thaw:
		return "Thaw"
	case Instruction_ThawV5:
		return "ThawV5"
	case Instruction_InitCmV5:
		return "InitCmV5"
	case Instruction_InitCmV4:
		return "InitCmV4"
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
			"mint_embed", (*MintEmbed)(nil),
		},
		{
			"mint_v5", (*MintV5)(nil),
		},
		{
			"mint_v4", (*MintV4)(nil),
		},
		{
			"mint_v2", (*MintV2)(nil),
		},
		{
			"edit_cm_v3", (*EditCmV3)(nil),
		},
		{
			"edit_cm_v2", (*EditCmV2)(nil),
		},
		{
			"edit_cm_v4", (*EditCmV4)(nil),
		},
		{
			"edit_cm_v5", (*EditCmV5)(nil),
		},
		{
			"reveal_v5", (*RevealV5)(nil),
		},
		{
			"reveal", (*Reveal)(nil),
		},
		{
			"allow_unfreeze", (*AllowUnfreeze)(nil),
		},
		{
			"allow_reveal", (*AllowReveal)(nil),
		},
		{
			"allow_unfreeze_v5", (*AllowUnfreezeV5)(nil),
		},
		{
			"allow_reveal_v5", (*AllowRevealV5)(nil),
		},
		{
			"burn_supply_v5", (*BurnSupplyV5)(nil),
		},
		{
			"burn_supply_v3", (*BurnSupplyV3)(nil),
		},
		{
			"burn_frozen", (*BurnFrozen)(nil),
		},
		{
			"thaw", (*Thaw)(nil),
		},
		{
			"thaw_v5", (*ThawV5)(nil),
		},
		{
			"init_cm_v5", (*InitCmV5)(nil),
		},
		{
			"init_cm_v4", (*InitCmV4)(nil),
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