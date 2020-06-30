package v2

import (
	"errors"
	"io"
	"unicode"

	"github.com/manifoldco/promptui"
	"github.com/prysmaticlabs/prysm/shared/bls"
	"github.com/prysmaticlabs/prysm/validator/flags"

	logrus "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	keystorev4 "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"
)

var log = logrus.WithField("prefix", "accounts-v2")

// WalletType defines an enum for either direct, derived, or remote-signing
// wallets as specified by a user during account creation.
type WalletType int

const (
	DirectWallet  WalletType = iota // Direct, on-disk wallet.
	DerivedWallet                   // Derived, hierarchical-deterministic wallet.
	RemoteWallet                    // Remote-signing wallet.
)

// Steps: ask for the path to store the validator datadir
// Ask for the type of wallet: direct, derived, remote
// Ask for where to store passwords: default path within datadir
// Ask them to enter a password for the account.
// If user already has account, instead just make a new account with a good name
// Allow for creating more than 1 validator??
// TODOS: mnemonic for withdrawal key, ensure they write it down.
func New(cliCtx *cli.Context) error {
	walletPath := inputWalletPath(cliCtx)
	_ = walletPath

	// Determine the type of wallet for a user (e.g.: Direct, Keystore, Derived).
	walletType := inputWalletType(cliCtx)
	_ = walletType

	// Read the account password from user input.
	password := inputAccountPassword(cliCtx)
	_ = password

	// Read the directory for password storage from user input.
	passwordDirPath := inputPasswordsDirectory(cliCtx)
	_ = passwordDirPath

	// Open the wallet and password directories for writing.
	//createWalletPath()
	return nil
}

func inputWalletPath(cliCtx *cli.Context) string {
	datadir := cliCtx.String(flags.WalletDirFlag.Name)
	prompt := promptui.Prompt{
		Label:    "Enter a wallet directory",
		Validate: validateDirectoryPath,
		Default:  datadir,
	}
	walletPath, err := prompt.Run()
	if err != nil {
		log.Fatalf("Could not determine wallet directory: %v", formatPromptError(err))
	}
	return walletPath
}

func inputWalletType(_ *cli.Context) WalletType {
	promptSelect := promptui.Select{
		Label: "Select a type of wallet",
		Items: []string{
			"Direct, On-Disk (Recommended)",
			"Derived (Advanced)",
			"Remote (Advanced)",
		},
	}
	selection, _, err := promptSelect.Run()
	if err != nil {
		log.Fatalf("Could not select wallet type: %v", formatPromptError(err))
	}
	return WalletType(selection)
}

func inputAccountPassword(_ *cli.Context) string {
	prompt := promptui.Prompt{
		Label:    "Strong password",
		Validate: validatePasswordInput,
		Mask:     '*',
	}

	walletPassword, err := prompt.Run()
	if err != nil {
		log.Fatalf("Could not read wallet password: %v", formatPromptError(err))
	}

	prompt = promptui.Prompt{
		Label: "Confirm password",
		Mask:  '*',
	}
	confirmPassword, err := prompt.Run()
	if err != nil {
		log.Fatalf("Could not read password confirmation: %v", formatPromptError(err))
	}
	if walletPassword != confirmPassword {
		log.Fatal("Passwords do not match")
	}
	return walletPassword
}

func inputPasswordsDirectory(cliCtx *cli.Context) string {
	passwordsDir := cliCtx.String(flags.WalletPasswordsDirFlag.Name)
	prompt := promptui.Prompt{
		Label:    "Enter the directory where passwords will be stored",
		Validate: validateDirectoryPath,
		Default:  passwordsDir,
	}
	passwordsPath, err := prompt.Run()
	if err != nil {
		log.Fatalf("Could not determine passwords directory: %v", formatPromptError(err))
	}
	return passwordsPath
}

func createWalletPath(walletWriter io.Writer, passwordsWriter io.Writer, password string) {
	key := bls.RandKey()
	encryptor := keystorev4.New()
	keystore, err := encryptor.Encrypt(key.Marshal(), []byte(password))
	if err != nil {
		log.Fatal(err)
	}
	log.Info(keystore)
	_ = keystore
}

func validatePasswordInput(input string) error {
	var (
		hasMinLen  = false
		hasLetter  = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(input) >= 8 {
		hasMinLen = true
	}
	for _, char := range input {
		switch {
		case unicode.IsLetter(char):
			hasLetter = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	if !(hasMinLen && hasLetter && hasNumber && hasSpecial) {
		return errors.New(
			"password must have more than 8 characters, at least 1 special character, and 1 number",
		)
	}
	return nil
}

func validateDirectoryPath(input string) error {
	if len(input) == 0 {
		return errors.New("directory path must not be empty")
	}
	return nil
}

func formatPromptError(err error) error {
	switch err {
	case promptui.ErrAbort:
		return errors.New("wallet creation aborted, closing")
	case promptui.ErrInterrupt:
		return errors.New("keyboard interrupt, closing")
	case promptui.ErrEOF:
		return errors.New("no input received, closing")
	default:
		return err
	}
}
