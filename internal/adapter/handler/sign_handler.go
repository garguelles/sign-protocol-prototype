package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/pragma-collective/0xStarter-api/internal/application/contract"
)

func init() {
	godotenv.Load(`../../../.env`)
}

func Attest(c echo.Context) error {
	rpcUrl := os.Getenv(`RPC_URL`)
	fmt.Printf("RPC " + rpcUrl)
	org := c.QueryParam("org")

	err := createNotaryAttestation(org)
	if err != nil {
		log.Fatalf("Failed to create attestation: %v", err)
	}
	return c.JSON(http.StatusOK, map[string]bool{"Healthy": true})
}

func createNotaryAttestation(orgName string) error {
	// Get environment variables (you should set these in your environment)
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		rpcURL = "https://base-goerli.g.alchemy.com/v2/your-api-key" // Example for Base Goerli
	}

	schemaID := os.Getenv("SCHEMA_ID")
	if schemaID == "" {
		return fmt.Errorf("SCHEMA_ID environment variable is not set")
	}

	privateKeyHex := os.Getenv("ATTESTER_PRIVATE_KEY")
	if privateKeyHex == "" {
		return fmt.Errorf("PRIVATE_KEY environment variable is not set")
	}

	contractAddr := os.Getenv("SP_CONTRACT_ADDRESS")
	if contractAddr == "" {
		return fmt.Errorf("CONTRACT_ADDRESS environment variable is not set")
	}

	ctx := context.Background()

	// Connect to the Ethereum client
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	// Load the private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return fmt.Errorf("failed to load private key: %v", err)
	}

	// Get the chain ID
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %v", err)
	}

	log.Printf("CHAIN_ID: " + chainID.String())

	// Create the transaction auth
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create transaction auth: %v", err)
	}

	// Create the contract instance
	instance, err := contract.NewContract(common.HexToAddress(contractAddr), client)
	if err != nil {
		return fmt.Errorf("failed to create contract instance: %v", err)
	}

	recipientAddress := "0x8244c1645C1a7890Ef1F0E79AcCf817905Dbcba2"

	// Create the attestation
	attestation := contract.Attestation{
		SchemaId:            uint64(0x400),
		LinkedAttestationId: 0,
		AttestTimestamp:     0,
		RevokeTimestamp:     0,
		Attester:            common.HexToAddress(auth.From.Hex()),
		ValidUntil:          0,
		DataLocation:        0,
		Revoked:             false,
		Recipients:          [][]byte{common.HexToAddress(recipientAddress).Bytes()},
		Data:                []byte(orgName),
	}

	// Set up resolver fees parameters

	// Send the transaction
	tx, err := instance.Attest0(
		auth,
		attestation,
		recipientAddress, // indexing key
		[]byte{},         // no delegate signature
		[]byte{0x0},      // no extra data
	)
	if err != nil {
		return fmt.Errorf("failed to send attestation transaction: %v", err)
	}

	log.Printf("Transaction sent: %s", tx.Hash().Hex())

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return fmt.Errorf("failed while waiting for transaction: %v", err)
	}

	log.Printf("Transaction mined in block: %d", receipt.BlockNumber)
	log.Printf("Gas used: %d", receipt.GasUsed)

	// Check if the transaction was successful
	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed")
	}

	log.Printf("Attestation created successfully!")
	return nil
}
