// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package transactioncmd

import (
	"encoding/hex"
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/BOXFoundation/boxd/commands/box/root"
	"github.com/BOXFoundation/boxd/config"
	"github.com/BOXFoundation/boxd/core"
	"github.com/BOXFoundation/boxd/core/types"
	"github.com/BOXFoundation/boxd/crypto"
	"github.com/BOXFoundation/boxd/rpc/rpcutil"
	"github.com/BOXFoundation/boxd/util"
	"github.com/BOXFoundation/boxd/wallet"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var walletDir string
var defaultWalletDir = path.Join(util.HomeDir(), ".box_keystore")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tx",
	Short: "A brief description of your application",
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
			Use:   "sendfrom [fromaccount] [toaddress] [amount]",
			Short: "Send coins from an account to an address",
			Run:   sendFromCmdFunc,
		},
		&cobra.Command{
			Use:   "sendtoaddress [address]",
			Short: "Send coins to an address",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("sendtoaddress called")
			},
		},
		&cobra.Command{
			Use:   "signrawtx [rawtx]",
			Short: "Sign a transaction with privatekey and send it to the network",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("signrawtx called")
			},
		},
		&cobra.Command{
			Use:   "createrawtx [from] [txhash1,txhash2...] [vout1,vout2...] [to1,to2...] [amount1,amount2...]",
			Short: "Create a raw transaction",
			Run:   createRawTransaction,
		},
		&cobra.Command{
			Use:   "decoderawtx",
			Short: "A brief description of your command",
			Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra  application.`,
			Run: decoderawtx,
		},
		&cobra.Command{
			Use:   "getrawtx [txhash]",
			Short: "Get the raw transaction for a txid",
			Run:   getRawTxCmdFunc,
		},
		&cobra.Command{
			Use:   "sendrawtx [rawtx]",
			Short: "Send a raw transaction to the network",
			Run:   sendrawtx,
		},
	)
}

func createRawTransaction(cmd *cobra.Command, args []string) {
	if len(args) < 4 {
		fmt.Println("Invalide argument number")
		return
	}
	fmt.Println("createRawTx called")
	from := args[0]
	//Cut characters around commas
	txHashStr := strings.Split(args[1], ",")
	txHash := make([]crypto.HashType, 0)
	for _, x := range txHashStr {
		tmp := crypto.HashType{}
		if err := tmp.SetString(x); err != nil {
			fmt.Println("set string error: ", err)
			return
		}
		txHash = append(txHash, tmp)
	}
	voutStr := strings.Split(args[2], ",")
	vout := make([]uint32, 0)
	for _, x := range voutStr {
		tmp, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println("Type conversion failed: ", err)
			return
		}
		vout = append(vout, uint32(tmp))
	}
	toStr := strings.Split(args[3], ",")
	to := make([]string, 0)
	for _, x := range toStr {
		to = append(to, x)
	}
	amountStr := strings.Split(args[4], ",")
	amounts := make([]uint64, 0)
	for _, x := range amountStr {
		tmp, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println("Type conversion failed: ", err)
			return
		}
		amounts = append(amounts, uint64(tmp))
	}
	if len(txHash) != len(vout) {
		fmt.Println(" The number of [txid] should be the same as the number of [vout]")
		return
	}
	if len(to) != len(amounts) {
		fmt.Println("The number of [to] should be the same as the number of [amount]")
		return
	}
	conn, err := rpcutil.GetGRPCConn(getRPCAddr())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	height, err := rpcutil.GetBlockCount(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	tx, err := rpcutil.CreateRawTransaction(from, txHash, vout, to, amounts, height)
	if err != nil {
		fmt.Println(err)
		return
	}
	mashalTx, err := tx.Marshal()
	if err != nil {
		fmt.Println(err)
		return
	}
	strTx := hex.EncodeToString(mashalTx)
	fmt.Println(strTx)
}

func sendrawtx(cmd *cobra.Command, args []string) {
	fmt.Println("sendrawtx called")
	if len(args) != 1 {
		fmt.Println("Can only enter one string for a transaction")
		return
	}
	conn, err := rpcutil.GetGRPCConn(getRPCAddr())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	txByte, err := hex.DecodeString(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	tx := new(types.Transaction)
	if err := tx.Unmarshal(txByte); err != nil {
		fmt.Println("unmarshal error: ", err)
		return
	}
	resp, err := rpcutil.SendTransaction(conn, tx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(util.PrettyPrint(resp))
}

func getRawTxCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println("getrawtx called")
	if len(args) < 1 {
		fmt.Println("Param txhash required")
		return
	}
	hash := crypto.HashType{}
	hash.SetString(args[0])
	conn, err := rpcutil.GetGRPCConn(getRPCAddr())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	tx, err := rpcutil.GetRawTransaction(conn, hash.GetBytes())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(util.PrettyPrint(tx))
	}
}

func decoderawtx(cmd *cobra.Command, args []string) {
	fmt.Println("decoderawtx called")
	if len(args) != 1 {
		fmt.Println("Invalide argument number")
		return
	}
	txByte, err := hex.DecodeString(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	tx := new(types.Transaction)
	if err := tx.Unmarshal(txByte); err != nil {
		fmt.Println("Unmarshal error: ", err)
		return
	}
	fmt.Println(util.PrettyPrint(tx))
}

func sendFromCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) < 3 {
		fmt.Println("Invalid argument number")
		return
	}
	// account
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
	// to address
	addrs, amounts := make([]string, 0), make([]uint64, 0)
	for i := 1; i < len(args)-1; i += 2 {
		addrs = append(addrs, args[i])
		a, err := strconv.ParseUint(args[i+1], 10, 64)
		if err != nil {
			fmt.Printf("Invalid amount %s", args[i+1])
			return
		}
		amounts = append(amounts, a)
	}
	if err := types.ValidateAddr(addrs...); err != nil {
		fmt.Println(err)
		return
	}
	//
	conn, err := rpcutil.GetGRPCConn(getRPCAddr())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	tx, _, _, err := rpcutil.NewTx(account, addrs, amounts, conn)
	if err != nil {
		fmt.Println("new tx error: ", err)
		return
	}
	hash, err := rpcutil.SendTransaction(conn, tx)
	if err != nil && !strings.Contains(err.Error(), core.ErrOrphanTransaction.Error()) {
		fmt.Println(err)
	}
	fmt.Println("Tx Hash:", hash)
	fmt.Println(util.PrettyPrint(tx))
}

func parseSendParams(args []string) (addrs []string, amounts []uint64, err error) {
	for i := 0; i < len(args)/2; i++ {
		_, err := types.NewAddress(args[i*2])
		if err != nil {
			return nil, nil, err
		}
		addrs = append(addrs, args[i*2])
		amount, err := strconv.Atoi(args[i*2+1])
		if err != nil {
			return nil, nil, err
		}
		amounts = append(amounts, uint64(amount))
	}
	return addrs, amounts, nil
}

func getRPCAddr() string {
	var cfg config.Config
	viper.Unmarshal(&cfg)
	return fmt.Sprintf("%s:%d", cfg.RPC.Address, cfg.RPC.Port)
}
