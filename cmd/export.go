package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/go-zookeeper/zk"
	"github.com/spf13/cobra"
)

var ExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export Zookeeper data under /config/product",
	Run: func(cmd *cobra.Command, args []string) {
		sourceZK := os.Getenv("SOURCE_ZK")
		exportFile := os.Getenv("EXPORT_FILE")
		if sourceZK == "" || exportFile == "" {
			fmt.Println("SOURCE_ZK and EXPORT_FILE env vars must be set")
			os.Exit(1)
		}
		conn, _, err := zk.Connect(strings.Split(sourceZK, ","), time.Second*10)
		if err != nil {
			fmt.Println("Connect error:", err)
			os.Exit(1)
		}
		defer conn.Close()
		data := make(map[string]map[string]interface{})
		exportNode("/config/product", conn, data)

		f, err := os.Create(exportFile)
		if err != nil {
			fmt.Println("File create error:", err)
			os.Exit(1)
		}
		defer f.Close()
		enc := json.NewEncoder(f)
		enc.SetIndent("", "  ")
		_ = enc.Encode(data)
		fmt.Println("Export done")
	},
}

func exportNode(p string, conn *zk.Conn, data map[string]map[string]interface{}) {
	children, _, err := conn.Children(p)
	if err != nil {
		return
	}
	val, stat, err := conn.Get(p)
	if err == nil {
		data[p] = map[string]interface{}{
			"data": string(val),
			"stat": map[string]interface{}{
				"version": stat.Version, "mtime": stat.Mtime,
				"numChildren": stat.NumChildren,
				"dataLength": stat.DataLength,
				"ephemeralOwner": stat.EphemeralOwner,
			},
		}
	}
	for _, c := range children {
		childPath := path.Join(p, c)
		exportNode(childPath, conn, data)
	}
}
