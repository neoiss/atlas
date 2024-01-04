package cmd

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/mapprotocol/atlas/cmd/new_marker/define"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strconv"
)

var (
	AccountSet   []cli.Command
	ValidatorSet []cli.Command
	VoterSet     []cli.Command
	ToolSet      []cli.Command
)

func init() {
	account := NewAccount()
	AccountSet = append(AccountSet, []cli.Command{
		{
			Name:   "getAccountMetadataURL",
			Usage:  "Get metadata url of account",
			Action: MigrateFlags(account.GetAccountMetadataURL),
			Flags:  define.BaseFlagCombination,
		},
		{
			Name:   "getAccountName",
			Usage:  "Get name of account",
			Action: MigrateFlags(account.GetAccountName),
			Flags:  define.BaseFlagCombination,
		},
		{
			Name:   "getAccountTotalLockedGold",
			Usage:  "Returns the total amount of locked gold for an account.",
			Action: MigrateFlags(account.GetAccountTotalLockedGold),
			Flags:  define.BaseFlagCombination,
		},
		{
			Name:   "getAccountNonvotingLockedGold",
			Usage:  "Returns the total amount of non-voting locked gold for an account",
			Action: MigrateFlags(account.GetAccountNonvotingLockedGold),
			Flags:  define.BaseFlagCombination,
		},
		{
			Name:   "getPendingVotesForValidatorByAccount",
			Usage:  "Returns the pending votes for `validator` made by `account`",
			Action: MigrateFlags(account.GetPendingVotesForValidatorByAccount),
			Flags:  define.BaseFlagCombination,
		},
		{
			Name:   "getActiveVotesForValidatorByAccount",
			Usage:  "Returns the active votes for `validator` made by `account`",
			Action: MigrateFlags(account.GetActiveVotesForValidatorByAccount),
			Flags:  define.BaseFlagCombination,
		},
		{
			Name:   "getValidatorsVotedForByAccount",
			Usage:  "Returns the validators that `account` has voted for.",
			Action: MigrateFlags(account.GetValidatorsVotedForByAccount),
			Flags:  define.BaseFlagCombination,
		},
		{
			Name:   "setAccountMetadataURL",
			Usage:  "Set metadata url of account",
			Action: MigrateFlags(account.SetAccountMetadataURL),
			Flags:  append(define.BaseFlagCombination, define.URLFlag),
		},
		{
			Name:   "setAccountName",
			Usage:  "Set name of account",
			Action: MigrateFlags(account.SetAccountName),
			Flags:  append([]cli.Flag{}, define.RPCAddrFlag, define.KeyStoreFlag, define.GasLimitFlag, define.NameFlag),
		},
		{
			Name:   "createAccount",
			Usage:  "Creat validator account",
			Action: MigrateFlags(account.CreateAccount),
			Flags:  append([]cli.Flag{}, define.RPCAddrFlag, define.KeyStoreFlag, define.GasLimitFlag, define.NameFlag),
		},
		{
			Name:   "signerToAccount",
			Usage:  "Returns the account associated with `signer`.",
			Action: MigrateFlags(account.SignerToAccount),
			Flags:  define.BaseFlagCombination,
		},
	}...)
	validator := NewValidator()
	ValidatorSet = append(ValidatorSet, []cli.Command{
		{
			Name:   "register",
			Usage:  "Register validator",
			Action: MigrateFlags(validator.RegisterValidator),
			Flags:  append([]cli.Flag{}, define.RPCAddrFlag, define.KeyStoreFlag, define.CommissionFlag, define.SignerPriFlag),
		},
		{
			Name:   "generateSignerProof",
			Usage:  "Generate proof of signer",
			Action: MigrateFlags(validator.GenerateSignerProof),
			Flags:  append([]cli.Flag{}, define.KeyStoreFlag, define.ValidatorAddressFlag, define.SignerPriFlag),
		},
		{
			Name:   "registerByProof",
			Usage:  "Register validator by signer proof",
			Action: MigrateFlags(validator.RegisterValidatorByProof),
			Flags:  append(define.MustFlagCombination, define.ProofFlag, define.CommissionFlag),
		},
		{
			Name:   "revertRegister",
			Usage:  "Register validator",
			Action: MigrateFlags(validator.RevertRegisterValidator),
			Flags:  define.MustFlagCombination,
		},
		{
			Name:   "deregister",
			Usage:  "Deregister Validator",
			Action: MigrateFlags(validator.DeregisterValidator),
			Flags:  define.MustFlagCombination,
		},
		{
			Name:   "quicklyRegister",
			Usage:  "Register validator",
			Action: MigrateFlags(validator.QuicklyRegisterValidator),
			Flags:  append(define.MustFlagCombination, define.CommissionFlag, define.LockedNumFlag, define.NameFlag),
		},
		{
			Name:   "authorizeValidatorSigner",
			Usage:  "Finish the process of authorizing an address to sign on behalf of the account.",
			Action: MigrateFlags(validator.AuthorizeValidatorSigner),
			Flags:  append(define.MustFlagCombination, define.SignerPriFlag),
		},
		{
			Name:   "authorizeValidatorSignerBySignature",
			Usage:  "Finish the process of authorizing an address to sign on behalf of the account.",
			Action: MigrateFlags(validator.AuthorizeValidatorSignerBySignature),
			Flags:  append(define.MustFlagCombination, define.SignatureFlag, define.SignerFlag),
		},
		{
			Name:   "updateValidatorSigner",
			Usage:  "Update signer account of validator.",
			Action: MigrateFlags(validator.UpdateValidatorSigner),
			Flags:  append(define.MustFlagCombination, define.SignerPriFlag),
		},
		{
			Name:   "makeECDSASignatureFromSigner",
			Usage:  "Print a ECDSASignature that signer sign the account(validator)",
			Action: MigrateFlags(validator.MakeECDSASignatureFromSigner),
			Flags:  append([]cli.Flag{}, define.RPCAddrFlag, define.KeyStoreFlag, define.SignerPriFlag, define.TargetAddressFlag),
		},
		{
			Name:   "makeBLSProofOfPossessionFromSigner",
			Usage:  "Print a BLSProofOfPossession that signer BLSSign the account(validator)",
			Action: MigrateFlags(validator.MakeBLSProofOfPossessionFromsigner),
			Flags:  append([]cli.Flag{}, define.RPCAddrFlag, define.KeyStoreFlag, define.SignerPriFlag, define.TargetAddressFlag),
		},
	}...)
	voter := NewVoter()
	VoterSet = append(VoterSet, []cli.Command{
		{
			Name:   "vote",
			Usage:  "Vote validator ",
			Action: MigrateFlags(voter.Vote),
			Flags:  append(define.BaseFlagCombination, define.VoteNumFlag),
		},
		{
			Name:   "quicklyVote",
			Usage:  "Vote validator ",
			Action: MigrateFlags(voter.QuicklyVote),
			Flags:  append(define.BaseFlagCombination, define.NameFlag, define.LockedNumFlag, define.VoteNumFlag),
		},
		{
			Name:   "activate",
			Usage:  "Converts `account`'s pending votes for `validator` to active votes.",
			Action: MigrateFlags(voter.Activate),
			Flags:  define.BaseFlagCombination,
		},
		{
			Name:   "getActiveVotesForValidator",
			Usage:  "Returns the total active vote units made for `validator`.",
			Action: MigrateFlags(voter.GetActiveVotesForValidator),
			Flags:  define.BaseFlagCombination,
		},
		{
			Name:   "getPendingVotersForValidator",
			Usage:  "Returns the total pending voters vote for target `validator`.",
			Action: MigrateFlags(voter.GetPendingVotersForValidator),
			Flags:  define.BaseFlagCombination,
		},
		{
			Name:   "getPendingInfoForValidator",
			Usage:  "Returns the  pending Info voters vote And Epoch for target `validator`.",
			Action: MigrateFlags(voter.GetPendingInfoForValidator),
			Flags:  define.BaseFlagCombination,
		},
		{
			Name:   "revokePending",
			Usage:  "Revokes `value` pending votes for `validator`",
			Action: MigrateFlags(voter.RevokePending),
			Flags:  append(define.BaseFlagCombination, define.LockedNumFlag),
		},
		{
			Name:   "revokeActive",
			Usage:  "Revokes `value` active votes for `validator`",
			Action: MigrateFlags(voter.RevokeActive),
			Flags:  append(define.BaseFlagCombination, define.LockedNumFlag),
		},
		{
			Name:   "lockedMAP",
			Usage:  "Locked MAP",
			Action: MigrateFlags(voter.LockedMAP),
			Flags:  append(define.MustFlagCombination, define.LockedNumFlag),
		},
		{
			Name:   "unlockMap",
			Usage:  "Unlocked MAP",
			Action: MigrateFlags(voter.UnlockedMAP),
			Flags:  append(define.MustFlagCombination, define.LockedNumFlag),
		},
		{
			Name:   "relockMAP",
			Usage:  "Unlocked MAP",
			Action: MigrateFlags(voter.RelockMAP),
			Flags:  append(define.MustFlagCombination, define.LockedNumFlag, define.ReLockIndexFlag),
		},
		{
			Name:   "withdrawMap",
			Usage:  "Withdraw MAP",
			Action: MigrateFlags(voter.Withdraw),
			Flags:  append(define.MustFlagCombination, define.WithdrawIndexFlag),
		},
		{
			Name:   "getTotalVotesForEligibleValidators",
			Usage:  "Vote validator ",
			Action: MigrateFlags(voter.GetTotalVotesForEligibleValidators),
			Flags:  define.MustFlagCombination,
		},
		{
			Name:   "getRegisteredValidatorSigners",
			Usage:  "Get Registered Validator Signers",
			Action: MigrateFlags(voter.GetRegisteredValidatorSigners),
			Flags:  define.MustFlagCombination,
		},
		{
			Name:   "getValidator",
			Usage:  "Validator Info",
			Action: MigrateFlags(voter.GetValidator),
			Flags:  define.BaseFlagCombination,
		},
		{
			Name:   "getValidatorRewardInfo",
			Usage:  "GetValidator Info",
			Action: MigrateFlags(voter.GetRewardInfo),
			Flags:  define.MustFlagCombination,
		},
		{
			Name:   "getVoterRewardInfo",
			Usage:  "Get Voter Reward Information about yourself",
			Action: MigrateFlags(voter.getVoterRewardInfo),
			Flags:  define.MustFlagCombination,
		},
		{
			Name:   "getNumRegisteredValidators",
			Usage:  "Get Num RegisteredValidators",
			Action: MigrateFlags(voter.getNumRegisteredValidators),
			Flags:  define.MustFlagCombination,
		},
		{
			Name:   "getTopValidators",
			Usage:  "Get Top Validators",
			Action: MigrateFlags(voter.getTopValidators),
			Flags:  append(define.MustFlagCombination, define.TopNumFlag),
		},
		{
			Name:   "getValidatorEligibility",
			Usage:  "Judge whether the verifier`s Eligibility",
			Action: MigrateFlags(voter.getValidatorEligibility),
			Flags:  define.BaseFlagCombination,
		},
		{
			Name:   "balanceOf",
			Usage:  "Gets the balance of the specified address using the presently stored inflation factor.",
			Action: MigrateFlags(voter.balanceOf),
			Flags:  define.BaseFlagCombination,
		},
		{
			Name:   "getTotalVotes",
			Usage:  "Returns the total votes received across all validators.",
			Action: MigrateFlags(voter.getTotalVotes),
			Flags:  define.MustFlagCombination,
		},
		{
			Name:   "getPendingWithdrawals",
			Usage:  "Returns the pending withdrawals from unlocked gold for an account.",
			Action: MigrateFlags(voter.getPendingWithdrawals),
			Flags:  define.BaseFlagCombination,
		},
		{
			Name:   "setValidatorLockedGoldRequirements",
			Usage:  "Updates the Locked Gold requirements for Validators.",
			Action: MigrateFlags(voter.setValidatorLockedGoldRequirements),
			Flags:  append(define.MustFlagCombination, define.DurationFlag, define.ValueFlag),
		},
		{
			Name:   "setImplementation",
			Usage:  "Sets the address of the implementation contract.",
			Action: MigrateFlags(voter.setImplementation),
			Flags:  append(define.MustFlagCombination, define.ContractAddressFlag, define.ImplementationAddressFlag),
		},
		{
			Name:   "setContractOwner",
			Usage:  "Transfers ownership of the contract to a new account (`newOwner`).",
			Action: MigrateFlags(voter.setContractOwner),
			Flags:  append(define.BaseFlagCombination, define.ContractAddressFlag),
		},
		{
			Name:   "setProxyContractOwner",
			Usage:  "Transfers ownership of the contract to a new account (`newOwner`).",
			Action: MigrateFlags(voter.setProxyContractOwner),
			Flags:  append(define.MustFlagCombination, define.ContractAddressFlag),
		},
		{
			Name:   "getProxyContractOwner",
			Usage:  "Transfers ownership of the contract to a new account (`newOwner`).",
			Action: MigrateFlags(voter.getProxyContractOwner),
			Flags:  append(define.MustFlagCombination, define.ContractAddressFlag),
		},
		{
			Name:   "getContractOwner",
			Usage:  "Transfers ownership of the contract to a new account (`newOwner`).",
			Action: MigrateFlags(voter.getContractOwner),
			Flags:  append(define.MustFlagCombination, define.ContractAddressFlag),
		},
		{
			Name:   "updateBlsPublicKey",
			Usage:  "UpdateBlsPublicKey",
			Action: MigrateFlags(voter.updateBlsPublicKey),
			Flags:  define.MustFlagCombination,
		},
		{
			Name:   "setNextCommissionUpdate",
			Usage:  "Set Next Commission Update",
			Action: MigrateFlags(voter.setNextCommissionUpdate),
			Flags:  append(define.MustFlagCombination, define.CommissionFlag),
		},
		{
			Name:   "updateCommission",
			Usage:  "UpdateCommission",
			Action: MigrateFlags(voter.updateCommission),
			Flags:  append(define.MustFlagCombination, define.CommissionFlag),
		},
		{
			Name:   "setValidatorEpochPayment",
			Usage:  "Sets the target per-epoch payment in MAP  for validators",
			Action: MigrateFlags(voter.setTargetValidatorEpochPayment),
			Flags:  append(define.MustFlagCombination, define.ValueFlag),
		},
		{
			Name:   "setEpochMaintainerPaymentFraction",
			Usage:  "Set Epoch Maintainer PaymentFraction",
			Action: MigrateFlags(voter.setEpochMaintainerPaymentFraction),
			Flags:  append(define.MustFlagCombination, define.RelayerFlag),
		},
		{
			Name:   "setMgrMaintainerAddress",
			Usage:  "Set manager maintainer address",
			Action: MigrateFlags(voter.setMgrMaintainerAddress),
			Flags:  define.BaseFlagCombination,
		},
		{
			Name:   "getMgrMaintainerAddress",
			Usage:  "Set manager maintainer address",
			Action: MigrateFlags(voter.getMgrMaintainerAddress),
			Flags:  define.MustFlagCombination,
		},
	}...)
	tool := NewTool()
	ToolSet = append(ToolSet, []cli.Command{
		{
			Name:      "genesis",
			Usage:     "Creates genesis.json from a template and overrides",
			Action:    tool.createGenesis,
			ArgsUsage: "",
			Flags: append(
				[]cli.Flag{
					define.BuildpathFlag,
					define.NewEnvFlag,
					define.MarkerCfgFlag,
				},
				define.TemplateFlags...),
		},
		{
			Name:   "transfer",
			Usage:  "Transfer",
			Action: MigrateFlags(tool.transfer),
			Flags:  append(define.MustFlagCombination, define.AmountFlag, define.TargetAddressFlag),
		},
		{
			Name:   "voterMonitor",
			Usage:  "Monitor the revenue of voter to a validator",
			Action: MigrateFlags(tool.voterMonitor),
			Flags:  define.MustFlagCombination,
		},
	}...)
}

func MigrateFlags(hdl func(ctx *cli.Context, cfg *define.Config) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		for _, name := range ctx.FlagNames() {
			if ctx.IsSet(name) {
				err := ctx.Set(name, ctx.String(name))
				if err != nil {
					log.Error("MigrateFlags", "=== err ===", err, ctx.IsSet(name))
				}
			}
		}
		_config, err := define.AssemblyConfig(ctx)
		if err != nil {
			cli.ShowAppHelpAndExit(ctx, 1)
			panic(err)
		}
		err = startLogger(ctx, _config)
		if err != nil {
			cli.ShowAppHelpAndExit(ctx, 1)
			panic(err)
		}
		return hdl(ctx, _config)
	}
}

func startLogger(_ *cli.Context, config *define.Config) error {
	logger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	var lvl log.Lvl
	if lvlToInt, err := strconv.Atoi(config.Verbosity); err == nil {
		lvl = log.Lvl(lvlToInt)
	} else if lvl, err = log.LvlFromString(config.Verbosity); err != nil {
		return err
	}
	logger.Verbosity(lvl)
	log.Root().SetHandler(log.LvlFilterHandler(lvl, logger))
	return nil
}
