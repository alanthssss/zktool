package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-zookeeper/zk"
	"github.com/spf13/cobra"
)

var ImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import Zookeeper data into /config/product",
	Run: func(cmd *cobra.Command, args []string) {
		targetZK := os.Getenv("TARGET_ZK")
		importFile := os.Getenv("IMPORT_FILE")
		if targetZK == "" || importFile == "" {
			fmt.Println("TARGET_ZK and IMPORT_FILE env vars must be set")
			os.Exit(1)
		}
		conn, _, err := zk.Connect(strings.Split(targetZK, ","), time.Second*10)
		if err != nil {
			fmt.Println("Connect error:", err)
			os.Exit(1)
		}
		defer conn.Close()
		f, err := os.Open(importFile)
		if err != nil {
			fmt.Println("File open error:", err)
			os.Exit(1)
		}
		defer f.Close()
		var nodes map[string]map[string]interface{}
		if err := json.NewDecoder(f).Decode(&nodes); err != nil {
			fmt.Println("JSON decode error:", err)
			os.Exit(1)
		}
		for path, node := range nodes {
			if !strings.HasPrefix(path, "/config/product") {
				continue
			}
			conn.Create(path, []byte(node["data"].(string)), 0, zk.WorldACL(zk.PermAll))
			conn.Set(path, []byte(node["data"].(string)), -1)
		}
		fmt.Println("Import done")
	},
}
