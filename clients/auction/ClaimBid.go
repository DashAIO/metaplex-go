// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Move SPL tokens from winning bid to the destination account.
type ClaimBid struct {
	Args *ClaimBidArgs

	// [0] = [WRITE] destinationAccount
	// ··········· The destination account
	//
	// [1] = [WRITE] bidderPotTokenAccount
	// ··········· The bidder pot token account
	//
	// [2] = [] bidderPotPdaAccount
	// ··········· The bidder pot pda account [seed of ['auction', program_id, auction key, bidder key]]
	//
	// [3] = [SIGNER] authorityOnTheAuctionAccount
	// ··········· The authority on the auction
	//
	// [4] = [] auction
	// ··········· The auction
	//
	// [5] = [] bidderWallet
	// ··········· The bidder wallet
	//
	// [6] = [] tokenMintOfTheAuctionAccount
	// ··········· Token mint of the auction
	//
	// [7] = [] clockSysvar
	// ··········· Clock sysvar
	//
	// [8] = [] tokenProgram
	// ··········· Token program
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewClaimBidInstructionBuilder creates a new `ClaimBid` instruction builder.
func NewClaimBidInstructionBuilder() *ClaimBid {
	nd := &ClaimBid{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *ClaimBid) SetArgs(args ClaimBidArgs) *ClaimBid {
	inst.Args = &args
	return inst
}

// SetDestinationAccount sets the "destinationAccount" account.
// The destination account
func (inst *ClaimBid) SetDestinationAccount(destinationAccount ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(destinationAccount).WRITE()
	return inst
}

// GetDestinationAccount gets the "destinationAccount" account.
// The destination account
func (inst *ClaimBid) GetDestinationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetBidderPotTokenAccount sets the "bidderPotTokenAccount" account.
// The bidder pot token account
func (inst *ClaimBid) SetBidderPotTokenAccount(bidderPotTokenAccount ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(bidderPotTokenAccount).WRITE()
	return inst
}

// GetBidderPotTokenAccount gets the "bidderPotTokenAccount" account.
// The bidder pot token account
func (inst *ClaimBid) GetBidderPotTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetBidderPotPdaAccount sets the "bidderPotPdaAccount" account.
// The bidder pot pda account [seed of ['auction', program_id, auction key, bidder key]]
func (inst *ClaimBid) SetBidderPotPdaAccount(bidderPotPdaAccount ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(bidderPotPdaAccount)
	return inst
}

// GetBidderPotPdaAccount gets the "bidderPotPdaAccount" account.
// The bidder pot pda account [seed of ['auction', program_id, auction key, bidder key]]
func (inst *ClaimBid) GetBidderPotPdaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetAuthorityOnTheAuctionAccount sets the "authorityOnTheAuctionAccount" account.
// The authority on the auction
func (inst *ClaimBid) SetAuthorityOnTheAuctionAccount(authorityOnTheAuctionAccount ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(authorityOnTheAuctionAccount).SIGNER()
	return inst
}

// GetAuthorityOnTheAuctionAccount gets the "authorityOnTheAuctionAccount" account.
// The authority on the auction
func (inst *ClaimBid) GetAuthorityOnTheAuctionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetAuctionAccount sets the "auction" account.
// The auction
func (inst *ClaimBid) SetAuctionAccount(auction ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(auction)
	return inst
}

// GetAuctionAccount gets the "auction" account.
// The auction
func (inst *ClaimBid) GetAuctionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetBidderWalletAccount sets the "bidderWallet" account.
// The bidder wallet
func (inst *ClaimBid) SetBidderWalletAccount(bidderWallet ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(bidderWallet)
	return inst
}

// GetBidderWalletAccount gets the "bidderWallet" account.
// The bidder wallet
func (inst *ClaimBid) GetBidderWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetTokenMintOfTheAuctionAccount sets the "tokenMintOfTheAuctionAccount" account.
// Token mint of the auction
func (inst *ClaimBid) SetTokenMintOfTheAuctionAccount(tokenMintOfTheAuctionAccount ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenMintOfTheAuctionAccount)
	return inst
}

// GetTokenMintOfTheAuctionAccount gets the "tokenMintOfTheAuctionAccount" account.
// Token mint of the auction
func (inst *ClaimBid) GetTokenMintOfTheAuctionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetClockSysvarAccount sets the "clockSysvar" account.
// Clock sysvar
func (inst *ClaimBid) SetClockSysvarAccount(clockSysvar ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(clockSysvar)
	return inst
}

// GetClockSysvarAccount gets the "clockSysvar" account.
// Clock sysvar
func (inst *ClaimBid) GetClockSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *ClaimBid) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *ClaimBid {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *ClaimBid) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

func (inst ClaimBid) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_ClaimBid),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ClaimBid) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ClaimBid) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.DestinationAccount is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.BidderPotTokenAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.BidderPotPdaAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.AuthorityOnTheAuctionAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Auction is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.BidderWallet is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenMintOfTheAuctionAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.ClockSysvar is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *ClaimBid) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ClaimBid")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("          destination", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("       bidderPotToken", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("         bidderPotPda", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("authorityOnTheAuction", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("              auction", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("         bidderWallet", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("tokenMintOfTheAuction", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("          clockSysvar", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("         tokenProgram", inst.AccountMetaSlice[8]))
					})
				})
		})
}

func (obj ClaimBid) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ClaimBid) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewClaimBidInstruction declares a new ClaimBid instruction with the provided parameters and accounts.
func NewClaimBidInstruction(
	// Parameters:
	args ClaimBidArgs,
	// Accounts:
	destinationAccount ag_solanago.PublicKey,
	bidderPotTokenAccount ag_solanago.PublicKey,
	bidderPotPdaAccount ag_solanago.PublicKey,
	authorityOnTheAuctionAccount ag_solanago.PublicKey,
	auction ag_solanago.PublicKey,
	bidderWallet ag_solanago.PublicKey,
	tokenMintOfTheAuctionAccount ag_solanago.PublicKey,
	clockSysvar ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *ClaimBid {
	return NewClaimBidInstructionBuilder().
		SetArgs(args).
		SetDestinationAccount(destinationAccount).
		SetBidderPotTokenAccount(bidderPotTokenAccount).
		SetBidderPotPdaAccount(bidderPotPdaAccount).
		SetAuthorityOnTheAuctionAccount(authorityOnTheAuctionAccount).
		SetAuctionAccount(auction).
		SetBidderWalletAccount(bidderWallet).
		SetTokenMintOfTheAuctionAccount(tokenMintOfTheAuctionAccount).
		SetClockSysvarAccount(clockSysvar).
		SetTokenProgramAccount(tokenProgram)
}