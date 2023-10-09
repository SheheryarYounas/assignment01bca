package main

import (
	"crypto/sha256"
	"encoding/hex" //to convert the hash to a string otherwise it gives gibberish
	"fmt"
	"math/rand"
	"strconv"
)

type Block struct {
	Hash         string
	PreviousHash string
	Transaction  string
	Nonce        int
	NextBlock    *Block
}

type BlockChain struct {
	GenesisBlock *Block
	CurrentBlock *Block
}

type Miner struct {
	Name          string
	Target        string
	reward_earned int
}

func CreateBlock(Transaction string, PreviousHash string, Nonce int) *Block { //Outputs a pointer to a new block and not the block itself. Why? Because we want to be able to modify the block later on.
	block := &Block{
		Transaction:  Transaction,
		Nonce:        Nonce,
		PreviousHash: PreviousHash,
	}

	data := block.PreviousHash + block.Transaction + string(block.Nonce)
	block.Hash = CalculateHash(data)
	return block
}

func CreateGenesisBlock() *Block {
	return CreateBlock("Genesis Block", "", 0)
}

func (BlockChain *BlockChain) NewBlock(Transaction string, PreviousHash string, Nonce int, miners [100]Miner) {
	// mined := false

	NewBlock := CreateBlock(Transaction, PreviousHash, Nonce)

	// for mined == false {

	// 	fmt.Printf("No miner has mined the block yet\n")

	// 	for i := 0; i < len(miners); i++ {

	// 		miners[i].Target = CalculateHash(Transaction + PreviousHash + string(rand.Intn(10)))

	// 		if NewBlock.Hash[:4] == miners[i].Target[:4] {
	// 			miners[i].reward_earned += 1
	// 			fmt.Printf("Miner %s has earned a reward\n", miners[i].Name)
	// 			mined = true
	// 		}

	// 	}

	// } //end of for loop
	BlockChain.CurrentBlock.NextBlock = NewBlock
	BlockChain.CurrentBlock = NewBlock
}

func CalculateHash(StringToHash string) string {
	hash := sha256.Sum256([]byte(StringToHash)) //sha256.Sum256() returns a slice of bytes. We convert it to a string using string()
	return string(hash[:])
}

func NewBlockChain() *BlockChain {
	firstBlock := CreateGenesisBlock()
	NewBlockChain := &BlockChain{
		GenesisBlock: firstBlock,
		CurrentBlock: firstBlock,
	}
	return NewBlockChain
}

func (BlockChain *BlockChain) DisplayBlocks() {
	currentBlock := BlockChain.GenesisBlock
	for currentBlock != nil {

		fmt.Printf("\n-----------------------------------------------\n")
		fmt.Printf("Transaction: %s\n", currentBlock.Transaction)
		fmt.Printf("Nonce: %d\n", currentBlock.Nonce)
		fmt.Printf("Previous Block Hash: %s\n", hex.EncodeToString([]byte(currentBlock.PreviousHash)))
		fmt.Printf("Current Block Hash: %s\n", hex.EncodeToString([]byte(currentBlock.Hash)))
		fmt.Printf("-----------------------------------------------\n")

		currentBlock = currentBlock.NextBlock

	}
}

func (BlockChain *BlockChain) DisplayBlocksForChange(Transaction string) *Block {
	currentBlock := BlockChain.GenesisBlock
	for currentBlock != nil {

		if currentBlock.Transaction == Transaction {
			return currentBlock
		}

		currentBlock = currentBlock.NextBlock

	}
	return nil
}

func ChangeBlock(block *Block, Transaction string) {
	block.Transaction = Transaction
	blockInformation := block.PreviousHash + block.Transaction + string(block.Nonce)
	block.Hash = CalculateHash(blockInformation)
}

func (BlockChain *BlockChain) VerifyBlockChain() bool {
	currentBlock := BlockChain.GenesisBlock
	for currentBlock != nil {

		if currentBlock.Hash != CalculateHash(currentBlock.PreviousHash+currentBlock.Transaction+string(currentBlock.Nonce)) {
			return false
		}

		currentBlock = currentBlock.NextBlock

	}

	return true
}

func main() {
	fmt.Printf("Welcome to the FAST NUCES High Budget Blockchain\n")
	BlockChain := NewBlockChain()

	fmt.Printf("Please enter number of miners: ")
	var numberOfMiners int
	fmt.Scanln(&numberOfMiners)

	var miners [100]Miner

	for i := 0; i < numberOfMiners; i++ {
		fmt.Printf("Enter name of miner %d: ", i+1)
		var name string
		fmt.Scanln(&name)
		miners[i].Name = name
		miners[i].Target = "0000"
		miners[i].reward_earned = 0
	}

	var choice int

	for choice != 9 {

		fmt.Printf("Press 1 to record your transaction\n")
		fmt.Printf("Press 2 to view all the blocks in the blockchain\n")
		fmt.Printf("Press 3 to change a block\n")
		fmt.Printf("Press 4 to verify the blockchain \n")
		fmt.Printf("Press 9 to exit \n")

		var choice int
		fmt.Scanln(&choice)

		if choice == 1 {
			fmt.Printf("We will first obtain some details from you \n")
			fmt.Printf("Enter your name: ")
			var name string
			fmt.Scanln(&name)
			fmt.Printf("Press 1 if you are sending money to someone else\n")
			fmt.Printf("Press 2 if you are receiving money from someone else\n")

			var choice2 int
			fmt.Scanln(&choice2)

			if choice2 == 1 {
				fmt.Printf("Enter the name of the person you are sending money to: ")
				var receiver string
				fmt.Scanln(&receiver)
				fmt.Printf("Enter the amount you are sending: ")
				var amount int
				fmt.Scanln(&amount)

				Transaction := name + " sent " + receiver + " " + strconv.Itoa(amount) + " rupees"

				//random integer
				Nonce := rand.Intn(10) //any random number between 0 and 10

				//Now to create a new block
				BlockChain.NewBlock(Transaction, BlockChain.CurrentBlock.Hash, Nonce, miners)

				fmt.Printf("Your transaction has been recorded\n")

			}

			if choice2 == 2 {
				fmt.Printf("Enter the name of the person you are receiving money from: ")
				var sender string
				fmt.Scanln(&sender)
				fmt.Printf("Enter the amount you are receiving: ")
				var amount int
				fmt.Scanln(&amount)

				Transaction := name + " received " + strconv.Itoa(amount) + " rupees from " + sender

				//random integer
				Nonce := rand.Intn(10) //any random number between 0 and 10

				//Now to create a new block
				BlockChain.NewBlock(Transaction, BlockChain.CurrentBlock.Hash, Nonce, miners)

				fmt.Printf("Your transaction has been recorded\n")

			} else {
				fmt.Printf("Invalid choice\n")
			}

		} else if choice == 2 {

			BlockChain.DisplayBlocks()

		} else if choice == 3 {

			fmt.Printf("Enter the details of transaction for security: \n")
			fmt.Printf("What was the name of the sender/receiver? \n")
			var name string
			fmt.Scanln(&name)
			fmt.Printf("What was the amount? \n")
			var amount int
			fmt.Scanln(&amount)
			fmt.Printf("What was the name of the other person? \n")
			var other string
			fmt.Scanln(&other)
			fmt.Printf("Type 1 if person sent the money, Type 2 if person received the money \n")
			var choice2 int
			fmt.Scanln(&choice2)

			if choice2 == 1 {
				Transaction := name + " sent " + other + " " + strconv.Itoa(amount) + " rupees"

				block := BlockChain.DisplayBlocksForChange(Transaction)

				if block != nil {
					fmt.Printf("Block has been found \n")
					fmt.Printf("Enter the new transaction: \n")
					//Starting
					fmt.Printf("Enter your name: ")
					var name string
					fmt.Scanln(&name)
					fmt.Printf("Press 1 if you are sending money to someone else\n")
					fmt.Printf("Press 2 if you are receiving money from someone else\n")

					var choice2 int
					fmt.Scanln(&choice2)

					if choice2 == 1 {
						fmt.Printf("Enter the name of the person you are sending money to: ")
						var receiver string
						fmt.Scanln(&receiver)
						fmt.Printf("Enter the amount you are sending: ")
						var amount int
						fmt.Scanln(&amount)

						Transaction := name + " sent " + receiver + " " + strconv.Itoa(amount) + " rupees"

						//Now to create a new block
						ChangeBlock(block, Transaction)

						fmt.Printf("Block has been changed\n")

					}

					if choice2 == 2 {
						fmt.Printf("Enter the name of the person you are receiving money from: ")
						var sender string
						fmt.Scanln(&sender)
						fmt.Printf("Enter the amount you are receiving: ")
						var amount int
						fmt.Scanln(&amount)

						Transaction := name + " received " + strconv.Itoa(amount) + " rupees from " + sender

						ChangeBlock(block, Transaction)

						fmt.Printf("Block has been changed\n")

					} else {
						fmt.Printf("Invalid choice\n")
					}
					//Ending
				} else {
					fmt.Printf("Block not found\n")
				}

			}

			if choice2 == 2 {
				Transaction := name + " received " + strconv.Itoa(amount) + " rupees from " + other

				block := BlockChain.DisplayBlocksForChange(Transaction)

				if block != nil {
					fmt.Printf("Block has been found \n")
					fmt.Printf("Enter the new transaction in the same format as before: \n")
					//Starting
					fmt.Printf("Enter your name: ")
					var name string
					fmt.Scanln(&name)
					fmt.Printf("Press 1 if you are sending money to someone else\n")
					fmt.Printf("Press 2 if you are receiving money from someone else\n")

					var choice2 int
					fmt.Scanln(&choice2)

					if choice2 == 1 {
						fmt.Printf("Enter the name of the person you are sending money to: ")
						var receiver string
						fmt.Scanln(&receiver)
						fmt.Printf("Enter the amount you are sending: ")
						var amount int
						fmt.Scanln(&amount)

						Transaction := name + " sent " + receiver + " " + strconv.Itoa(amount) + " rupees"

						ChangeBlock(block, Transaction)

						fmt.Printf("Block has been changed\n")

					}

					if choice2 == 2 {
						fmt.Printf("Enter the name of the person you are receiving money from: ")
						var sender string
						fmt.Scanln(&sender)
						fmt.Printf("Enter the amount you are receiving: ")
						var amount int
						fmt.Scanln(&amount)

						Transaction := name + " received " + strconv.Itoa(amount) + " rupees from " + sender

						ChangeBlock(block, Transaction)

						fmt.Printf("Block has been changed\n")

					} else {
						fmt.Printf("Invalid choice\n")
					}
					//Ending
				} else {
					fmt.Printf("Block not found\n")
				}
			}

		} else if choice == 4 {

			if BlockChain.VerifyBlockChain() {
				fmt.Printf("The blockchain is valid\n")
			} else {
				fmt.Printf("The blockchain is invalid\n")
			}

		} else if choice == 9 {
			fmt.Printf("Thank you for using the FAST NUCES High Budget Blockchain\n")
		} else {
			fmt.Printf("Invalid choice\n")
		}

	}

}
