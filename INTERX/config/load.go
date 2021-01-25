package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	functions "github.com/KiraCore/sekai/INTERX/functions"
	sekaiapp "github.com/KiraCore/sekai/app"
	functionmeta "github.com/KiraCore/sekai/function_meta"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bytesize "github.com/inhies/go-bytesize"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/sr25519"
	"github.com/tyler-smith/go-bip39"
)

var (
	// Config is a configuration for interx
	Config = InterxConfig{}
	// EncodingCg is a configuration for Amino Encoding
	EncodingCg = sekaiapp.MakeEncodingConfig()
)

func parseSizeString(size string) int64 {
	b, _ := bytesize.Parse(size)
	return int64(b)
}

func mnemonicFromFile(filename string) string {
	if len(filename) == 0 {
		return ""
	}

	mnemonic, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(mnemonic)
}

func loadMnemonic(mnemonic string) string {
	if bip39.IsMnemonicValid(mnemonic) {
		return mnemonic
	}

	return mnemonicFromFile(mnemonic)
}

// LoadConfig is a function to load interx configurations from a given file
func LoadConfig(configFilePath string) InterxConfig {
	functions.RegisterInterxFunctions()
	functionmeta.RegisterStdMsgs()
	sekaiapp.SetConfig()

	Config = InterxConfig{}

	file, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		fmt.Println("Invalid configuration: {}", err)
		panic(err)
	}

	configFromFile := InterxConfigFromFile{}

	err = json.Unmarshal([]byte(file), &configFromFile)
	if err != nil {
		fmt.Println("Invalid configuration: {}", err)
		panic(err)
	}

	// Interx Main Configuration
	Config.GRPC = configFromFile.GRPC
	Config.RPC = configFromFile.RPC
	Config.PORT = configFromFile.PORT
	Config.Mnemonic = loadMnemonic(configFromFile.MnemonicFile)
	Config.Cache.StatusSync = configFromFile.Cache.StatusSync
	Config.Cache.CacheDir = configFromFile.Cache.CacheDir
	Config.Cache.MaxCacheSize = parseSizeString(configFromFile.Cache.MaxCacheSize)
	Config.Cache.CachingDuration = configFromFile.Cache.CachingDuration
	Config.Cache.DownloadFileSizeLimitation = parseSizeString(configFromFile.Cache.DownloadFileSizeLimitation)

	if !bip39.IsMnemonicValid(Config.Mnemonic) {
		fmt.Println("Invalid Interx Mnemonic: ", Config.Mnemonic)
		panic("Invalid Interx Mnemonic")
	}
	Config.PrivKey = secp256k1.GenPrivKeyFromSecret(bip39.NewSeed(Config.Mnemonic, ""))
	Config.PubKey = Config.PrivKey.PubKey()
	Config.Address = sdk.MustBech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), Config.PubKey.Address())

	// Display mnemonic and keys
	fmt.Println("Interx Mnemonic   : ", Config.Mnemonic)
	fmt.Println("Interx Address    : ", Config.Address)
	fmt.Println("Interx Public Key : ", sdk.MustBech32ifyPubKey(sdk.Bech32PubKeyTypeAccPub, Config.PubKey))

	fmt.Println("Max Cache Size    : ", Config.Cache.MaxCacheSize)
	fmt.Println("Caching Duration  : ", Config.Cache.CachingDuration)
	fmt.Println("Download File Size Limitation  : ", Config.Cache.DownloadFileSizeLimitation)

	// Faucet Configuration
	Config.Faucet = FaucetConfig{
		Mnemonic:             loadMnemonic(configFromFile.Faucet.MnemonicFile),
		FaucetAmounts:        configFromFile.Faucet.FaucetAmounts,
		FaucetMinimumAmounts: configFromFile.Faucet.FaucetMinimumAmounts,
		FeeAmounts:           configFromFile.Faucet.FeeAmounts,
		TimeLimit:            configFromFile.Faucet.TimeLimit,
	}

	if !bip39.IsMnemonicValid(Config.Faucet.Mnemonic) {
		fmt.Println("Invalid Faucet Mnemonic: ", Config.Faucet.Mnemonic)
		panic("Invalid Faucet Mnemonic")
	}
	Config.Faucet.PrivKey = secp256k1.GenPrivKeyFromSecret(bip39.NewSeed(Config.Faucet.Mnemonic, ""))
	Config.Faucet.PubKey = Config.Faucet.PrivKey.PubKey()
	Config.Faucet.Address = sdk.MustBech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), Config.Faucet.PubKey.Address())

	// Display mnemonic and keys
	fmt.Println("Faucet Mnemonic   : ", Config.Faucet.Mnemonic)
	fmt.Println("Faucet Address    : ", Config.Faucet.Address)
	fmt.Println("Faucet Public Key : ", sdk.MustBech32ifyPubKey(sdk.Bech32PubKeyTypeAccPub, Config.Faucet.PubKey))

	// RPC Configuration
	Config.RPCMethods = configFromFile.RPCMethods

	if _, err := os.Stat(Config.Cache.CacheDir); os.IsNotExist(err) {
		os.Mkdir(Config.Cache.CacheDir, os.ModePerm)
	}
	if _, err := os.Stat(Config.Cache.CacheDir + "/reference/"); os.IsNotExist(err) {
		os.Mkdir(Config.Cache.CacheDir+"/reference/", os.ModePerm)
	}
	if _, err := os.Stat(Config.Cache.CacheDir + "/response/"); os.IsNotExist(err) {
		os.Mkdir(Config.Cache.CacheDir+"/response/", os.ModePerm)
	}
	if _, err := os.Stat(Config.Cache.CacheDir + "/db/"); os.IsNotExist(err) {
		os.Mkdir(Config.Cache.CacheDir+"/db/", os.ModePerm)
	}

	return Config
}

// GenPrivKey is a function to generate a privKey
func GenPrivKey() crypto.PrivKey {
	return sr25519.GenPrivKey()
}

// GetReferenceCacheDir is a function to get reference directory
func GetReferenceCacheDir() string {
	return Config.Cache.CacheDir + "/reference/"
}

// GetResponseCacheDir is a function to get reference directory
func GetResponseCacheDir() string {
	return Config.Cache.CacheDir + "/response/"
}

// GetDbCacheDir is a function to get db directory
func GetDbCacheDir() string {
	return Config.Cache.CacheDir + "/db/"
}