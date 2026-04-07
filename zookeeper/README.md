# zookeeper

The `zookeeper` module is part of [seed](../README.md) — a configuration governance CLI for platform and DevOps teams.

It provides safe, repeatable workflows for **exporting, importing, and bulk-updating Zookeeper configuration data** across environments.

---

## Why this module exists

In many real production-like environments, Zookeeper configuration changes are still handled through:

- manual console edits
- temporary migration scripts
- Excel-driven batch changes
- environment-specific ad-hoc commands

This creates familiar problems:

- low repeatability
- high operator risk
- weak auditability
- environment migration friction
- inconsistent update behavior across teams

The `zookeeper` module was created to reduce that operational overhead and make Zookeeper configuration workflows safer and more standardized.

---

## Positioning

The `zookeeper` module is not just a small utility.

A better way to understand it is:

- **Configuration Governance CLI**
- **Platform Operations Utility**
- **Zookeeper Config Migration / Update Tool**
- **DevOps Automation Tooling**

It reflects a platform-engineering mindset:

- reduce manual edits
- standardize repeated operations
- make risky workflows scriptable
- improve consistency across environments

---

## Key capabilities

### export
Export Zookeeper data to a JSON file.

Useful for:
- backup
- inspection
- migration preparation
- offline review before changes

```bash
export SOURCE_ZK=1.92.157.216:2181
export EXPORT_FILE=zookeeper_export.json
./seed zookeeper export
```

---

### import
Import configuration data from a previously exported JSON file into a target environment.

Useful for:
- environment bootstrap
- controlled restore
- migration between environments

```bash
export TARGET_ZK=localhost:3000
export IMPORT_FILE=zookeeper_export.json
./seed zookeeper import
```

---

### update
Apply structured updates from Excel (`.xlsx`) or JSON (`.json`).

Useful for:
- batch updates
- non-code-driven config maintenance
- structured config synchronization

**Using Excel:**

```bash
export TARGET_ZK=localhost:3000
export UPDATE_FILE=./zk.xlsx
./seed zookeeper update
```

**Using JSON:**

```bash
export TARGET_ZK=localhost:3000
export UPDATE_FILE=./zk_temp_data.json
./seed zookeeper update
```

---

## Example use cases

### 1. Backup before risky configuration changes
Before touching live Zookeeper data, export the current config state so the team has a recovery point.

### 2. Migrate `/config/product` data across environments
Move configuration data from one environment to another using a structured import/export workflow instead of manual recreation.

### 3. Apply batch changes from structured business input
When configuration updates are maintained in Excel or JSON, use `seed zookeeper update` to standardize bulk writes instead of editing nodes one by one.

---

## Build

```bash
bash build.sh
```

This builds binaries for:

- `linux/amd64`
- `windows/amd64`
- `darwin/amd64`
- `darwin/arm64`

Artifacts are written into the `build/` directory.

---

## Testing

Run the module tests with:

```bash
go test ./zookeeper/...
```

---

## Suggested workflow

A typical safe workflow looks like this:

1. **Export** current configuration
2. Review or prepare the update input
3. **Update** or **Import** into the target environment
4. Validate that target nodes match expectations
5. Keep export files as rollback references when appropriate

---

## Safety note

Because this module changes live Zookeeper configuration data, it should be used with environment-aware review and backup practices.

Recommended precautions:

- export before high-risk updates
- validate target environment carefully
- avoid treating Excel input as inherently correct
- test structured updates in a lower-risk environment first

---

## Roadmap ideas

Possible future improvements:

- dry-run mode for updates
- diff preview before import/update
- validation rules for Excel / JSON input
- audit-friendly output / change summary
- safer rollback helpers based on exported snapshots
