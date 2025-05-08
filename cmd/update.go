package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/go-zookeeper/zk"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Zookeeper nodes from Excel under /config/product",
	Run: func(cmd *cobra.Command, args []string) {
		targetZK := os.Getenv("TARGET_ZK")
		excelPath := os.Getenv("EXCEL_FILE")
		if targetZK == "" || excelPath == "" {
			fmt.Println("TARGET_ZK and EXCEL_FILE env vars must be set")
			os.Exit(1)
		}

		conn, _, err := zk.Connect(strings.Split(targetZK, ","), time.Second*10)
		if err != nil {
			fmt.Println("Connect error:", err)
			os.Exit(1)
		}
		defer conn.Close()

		f, err := excelize.OpenFile(excelPath)
		if err != nil {
			fmt.Println("Excel open error:", err)
			os.Exit(1)
		}

		rows, err := f.GetRows("Sheet1")
		if err != nil || len(rows) < 2 {
			fmt.Println("Read sheet error or empty")
			os.Exit(1)
		}

		header := map[string]int{}
		for i, h := range rows[0] {
			header[strings.TrimSpace(h)] = i
		}

		for _, row := range rows[1:] {
			if len(row) <= header["华为云压测环境"] {
				continue
			}
			base := strings.TrimSpace(row[header["路径"]])
			key := strings.TrimSpace(row[header["参数"]])
			val := strings.TrimSpace(row[header["华为云压测环境"]])
			full := fmt.Sprintf("%s/%s", base, key)

			if !strings.HasPrefix(full, "/config/product") {
				continue
			}

			if exists, _, _ := conn.Exists(full); exists {
				if _, err := conn.Set(full, []byte(val), -1); err == nil {
					fmt.Printf("✅ Updated: %s = %s\n", full, val)
				} else {
					fmt.Printf("❌ Failed to update: %s (%v)\n", full, err)
				}
			} else {
				ensureParentExists(conn, full)
				_, err := conn.Create(full, []byte(val), 0, zk.WorldACL(zk.PermAll))
				if err == nil {
					fmt.Printf("➕ Created: %s = %s\n", full, val)
				} else {
					fmt.Printf("❌ Failed to create: %s (%v)\n", full, err)
				}
			}
		}

		fmt.Println("Update done")
	},
}

func ensureParentExists(conn *zk.Conn, fullPath string) {
	parent := path.Dir(fullPath)
	parts := strings.Split(parent, "/")
	cur := ""
	for _, part := range parts {
		if part == "" {
			continue
		}
		cur += "/" + part
		exists, _, _ := conn.Exists(cur)
		if !exists {
			conn.Create(cur, []byte{}, 0, zk.WorldACL(zk.PermAll))
		}
	}
}