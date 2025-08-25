package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-zookeeper/zk"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Zookeeper nodes from Excel or JSON under /config/product",
	Run: func(cmd *cobra.Command, args []string) {
		targetZK := os.Getenv("TARGET_ZK")
		filePath := os.Getenv("EXCEL_FILE")
		if targetZK == "" || filePath == "" {
			fmt.Println("TARGET_ZK and EXCEL_FILE env vars must be set")
			os.Exit(1)
		}

		conn, _, err := zk.Connect(strings.Split(targetZK, ","), time.Second*10)
		if err != nil {
			fmt.Println("Connect error:", err)
			os.Exit(1)
		}
		defer conn.Close()

		switch strings.ToLower(filepath.Ext(filePath)) {
		case ".xlsx":
			updateFromExcel(conn, filePath)
		case ".json":
			updateFromJSON(conn, filePath)
		default:
			fmt.Println("Unsupported file format. Use .xlsx or .json")
			os.Exit(1)
		}

		fmt.Println("Update done")
	},
}

func updateFromExcel(conn *zk.Conn, filePath string) {
	f, err := excelize.OpenFile(filePath)
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
		applyValue(conn, full, val)
	}
}

func updateFromJSON(conn *zk.Conn, filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("JSON file open error:", err)
		os.Exit(1)
	}
	defer f.Close()
	var data map[string]string
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		fmt.Println("JSON decode error:", err)
		os.Exit(1)
	}
	for path, val := range data {
		if !strings.HasPrefix(path, "/config/product") {
			continue
		}
		applyValue(conn, path, val)
	}
}

func applyValue(conn *zk.Conn, zkPath, value string) {
	data := []byte(value)
	if exists, _, _ := conn.Exists(zkPath); exists {
		if _, err := conn.Set(zkPath, data, -1); err == nil {
			fmt.Printf("✅ Updated: %s = %s\n", zkPath, value)
		} else {
			fmt.Printf("❌ Failed to update: %s (%v)\n", zkPath, err)
		}
	} else {
		ensureParentExists(conn, zkPath)
		if _, err := conn.Create(zkPath, data, 0, zk.WorldACL(zk.PermAll)); err == nil {
			fmt.Printf("➕ Created: %s = %s\n", zkPath, value)
		} else {
			fmt.Printf("❌ Failed to create: %s (%v)\n", zkPath, err)
		}
	}
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