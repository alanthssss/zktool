# zktool

## Export

```bash
export SOURCE_ZK=1.92.157.216:2181
export EXPORT_FILE=zookeeper_export.json
./zktool export
```

## Import

```bash
export TARGET_ZK=localhost:3000
export IMPORT_FILE=zookeeper_export.json
./zktool import
```

## Update

```bash
export TARGET_ZK=localhost:3000
export EXCEL_FILE=./zk.xlsx
./zktool update
```
