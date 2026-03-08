# zktool

A CLI tool to export, import, and update Zookeeper nodes under `/config/product`.

## Build

```bash
bash build.sh
```

Builds binaries for `linux/amd64`, `windows/amd64`, `darwin/amd64`, and `darwin/arm64` into the `build/` directory.

## Export

Export Zookeeper data to a JSON file.

```bash
export SOURCE_ZK=1.92.157.216:2181
export EXPORT_FILE=zookeeper_export.json
./zktool export
```

## Import

Import Zookeeper data from a previously exported JSON file.

```bash
export TARGET_ZK=localhost:3000
export IMPORT_FILE=zookeeper_export.json
./zktool import
```

## Update

Update Zookeeper nodes from an Excel (`.xlsx`) or JSON (`.json`) file.

**Using Excel:**

```bash
export TARGET_ZK=localhost:3000
export UPDATE_FILE=./zk.xlsx
./zktool update
```

**Using JSON:**

```bash
export TARGET_ZK=localhost:3000
export UPDATE_FILE=./zk_temp_data.json
./zktool update
```
