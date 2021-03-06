// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tokencmd

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	root "github.com/BOXFoundation/boxd/commands/box/root"
	"github.com/BOXFoundation/boxd/config"
	"github.com/BOXFoundation/boxd/core"
	"github.com/BOXFoundation/boxd/core/txlogic"
	"github.com/BOXFoundation/boxd/core/types"
	"github.com/BOXFoundation/boxd/crypto"
	"github.com/BOXFoundation/boxd/rpc/rpcutil"
	"github.com/BOXFoundation/boxd/util"
	"github.com/BOXFoundation/boxd/wallet"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile          string
	walletDir        string
	defaultWalletDir = path.Join(util.HomeDir(), ".box_keystore")
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "token",
	Short: "Token subcommand",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Init adds the sub command to the root command.
func init() {
	root.RootCmd.AddCommand(rootCmd)
	rootCmd.PersistentFlags().StringVar(&walletDir, "wallet_dir", defaultWalletDir, "Specify directory to search keystore files")
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "issue issuer issuee name symbol supply decimal",
			Short: "issue a new token",
			Run:   createTokenCmdFunc,
		},
		&cobra.Command{
			Use:   "transfer from tokenID addr1 amount1 addr2 amount2 ...",
			Short: "transfer tokens",
			Run:   transferTokenCmdFunc,
		},
		&cobra.Command{
			Use:   "getbalance tokenID addr1 addr2 ...",
			Short: "get token balance",
			Run:   getTokenBalanceCmdFunc,
		},
	)
}

func createTokenCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println("createToken called")
	if len(args) != 6 {
		fmt.Println("Invalid argument number")
		return
	}

	toAddr, err1 := types.NewAddress(args[1])
	tokenName := args[2]
	tokenSymbol := args[3]
	tokenTotalSupply, err2 := strconv.Atoi(args[4])
	tokenDecimals, err3 := strconv.Atoi(args[5])
	if err1 != nil && err2 != nil && err3 != nil {
		fmt.Println("Invalid argument format")
		return
	}
	wltMgr, err := wallet.NewWalletManager(walletDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	// from pub key hash
	account, exists := wltMgr.GetAccount(args[0])
	if !exists {
		fmt.Printf("Account %s not managed\n", args[0])
		return
	}
	passphrase, err := wallet.ReadPassphraseStdin()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := account.UnlockWithPassphrase(passphrase); err != nil {
		fmt.Println("Fail to unlock account", err)
		return
	}
	_, err = types.NewAddress(args[0])
	if err != nil {
		fmt.Printf("Invalid address: %s, error: %s\n", args[0], err)
		return
	}
	conn, err := rpcutil.GetGRPCConn(getRPCAddr())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	tag := txlogic.NewTokenTag(tokenName, tokenSymbol, uint32(tokenDecimals),
		uint64(tokenTotalSupply))
	tx, _, _, err := rpcutil.NewIssueTokenTx(account, toAddr.String(), tag,
		uint64(tokenTotalSupply), conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	hashStr, err := rpcutil.SendTransaction(conn, tx)
	if err != nil && !strings.Contains(err.Error(), core.ErrOrphanTransaction.Error()) {
		fmt.Println(err)
		return
	}
	hash := new(crypto.HashType)
	hash.SetString(hashStr)

	tid := txlogic.NewPbOutPoint(hash, 0)
	fmt.Println("Created Token Address: ", txlogic.EncodeOutPoint(tid))
}

func transferTokenCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println("transferToken called")
	if len(args) < 5 {
		fmt.Println("Invalid argument number")
		return
	}
	// from account
	wltMgr, err := wallet.NewWalletManager(walletDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	account, exists := wltMgr.GetAccount(args[0])
	if !exists {
		fmt.Printf("Account %s not managed\n", args[0])
		return
	}
	passphrase, err := wallet.ReadPassphraseStdin()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := account.UnlockWithPassphrase(passphrase); err != nil {
		fmt.Println("Fail to unlock account", err)
		return
	}
	// token id
	hashStr := args[1]
	tHash := new(crypto.HashType)
	if err := tHash.SetString(hashStr); err != nil {
		fmt.Println(err)
		return
	}
	tIdx, err := strconv.ParseUint(args[2], 10, 64)
	if err != nil {
		fmt.Printf("token index %s invalid\n", args[1])
		return
	}
	// to address
	to, amounts := make([]string, 0), make([]uint64, 0)
	for i := 2; i < len(args)-1; i += 2 {
		to = append(to, args[i])
		a, err := strconv.ParseUint(args[i+1], 10, 64)
		if err != nil {
			fmt.Printf("Invalid amount %s", args[i+1])
			return
		}
		amounts = append(amounts, a)
	}
	if err := types.ValidateAddr(to...); err != nil {
		fmt.Println(err)
		return
	}
	//
	conn, err := rpcutil.GetGRPCConn(getRPCAddr())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	tx, _, _, err := rpcutil.NewTokenTx(account, to, amounts, hashStr, uint32(tIdx), conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	hashStr, err = rpcutil.SendTransaction(conn, tx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Tx Hash:", hashStr)
	fmt.Println(util.PrettyPrint(tx))
}

func getTokenBalanceCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println("getTokenBalance called")
	if len(args) < 3 {
		fmt.Println("Invalid argument number")
		return
	}
	// tokenID
	hashStr := args[0]
	tokenIndex, err := strconv.ParseUint(args[1], 10, 64)
	if err != nil {
		fmt.Printf("token index %s invalid\n", args[1])
		return
	}
	// address
	addrs := args[2:]
	for _, addr := range addrs {
		_, err := types.NewAddress(addr)
		if err != nil {
			fmt.Println("Invalid address: ", addr)
			return
		}
	}
	if err := types.ValidateAddr(addrs...); err != nil {
		fmt.Println(err)
		return
	}
	// call rpc
	conn, err := rpcutil.GetGRPCConn(getRPCAddr())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	balances, err := rpcutil.GetTokenBalance(conn, addrs, hashStr, uint32(tokenIndex))
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, b := range balances {
		fmt.Printf("%s: %d\n", addrs[i], b)
	}
}

func getRPCAddr() string {
	var cfg config.Config
	viper.Unmarshal(&cfg)
	return fmt.Sprintf("%s:%d", cfg.RPC.Address, cfg.RPC.Port)
}
