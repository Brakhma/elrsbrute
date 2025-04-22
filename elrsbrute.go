package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	// Define and parse command line flags
	searchUidFlag := flag.String("uid", "", "UID in quotes (e.g., \"79, 4, 253, 130, 33, 85\")")
	flag.Parse()

	// Validate UID input
	if *searchUidFlag == "" {
		flag.PrintDefaults()
		return
	}

	// Convert comma-separated UID string to bytes
	parts := strings.Split(*searchUidFlag, ",")
	var uidBytes []byte

	for _, part := range parts {
		numStr := strings.TrimSpace(part) // Remove surrounding whitespace
		num, err := strconv.Atoi(numStr)  // Convert string to integer
		if err != nil {
			fmt.Printf("Error converting number %s: %v\n", numStr, err)
			return
		}
		uidBytes = append(uidBytes, byte(num))
	}

	// Verify UID length (MD5 hash will be truncated to 6 bytes)
	if len(uidBytes) != 6 {
		fmt.Println("UID must be exactly 6 bytes long")
		return
	}

	// Convert UID bytes to hex string for comparison
	uidHexString := hex.EncodeToString(uidBytes)

	// Read potential binding phrases from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// scanner.Text() automatically handles line endings (\n, \r\n)
		inputString := scanner.Text() 
		
		// Format the binding phrase as it would appear in compilation flags
		bindingPhraseFlag := "-DMY_BINDING_PHRASE=\"" + inputString + "\""
		
		// Calculate MD5 hash and take first 6 bytes
		hash := md5.Sum([]byte(bindingPhraseFlag))
		hashStr := hex.EncodeToString(hash[:6])

		// Compare with target UID
		if hashStr == uidHexString {
			fmt.Printf("Bindphrase found: %s\n", inputString)
			return
		}
	}

	// Handle potential scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	} else {
		fmt.Println("No matching bindphrase found")
	}
}