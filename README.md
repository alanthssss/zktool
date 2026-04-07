# zktool

A Golang CLI for **safe export, import, and bulk update of Zookeeper configuration data** across environments.

`zktool` is built for teams who still rely on Zookeeper-backed configuration but want to move away from fragile manual edits, temporary scripts, and one-off operational workflows.

It turns repetitive configuration operations into a more controlled, repeatable, and automation-friendly CLI workflow.

---

## Why this exists

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

`zktool` was created to reduce that operational overhead and make Zookeeper configuration workflows safer and more standardized.

---

## What problem it solves

`zktool` focuses on a practical platform/DevOps problem:

> How do you safely move, restore, and bulk-update configuration data in Zookeeper without turning every change into a custom script or a risky manual operation?

It currently supports three core workflows:

1. **Export** — extract Zookeeper configuration data into a portable file
2. **Import** — restore or migrate configuration data into a target environment
3. **Update** — apply structured updates from Excel or JSON into existing Zookeeper nodes

This makes it useful for scenarios such as:

- environment migration
- backup / restore of config data
- batch updates across many nodes
- configuration standardization in platform workflows
- replacing hand-edited Zookeeper operations with a reusable CLI process

---

## Positioning

`zktool` is not just a small utility script.

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

### Export
Export Zookeeper data to a JSON file.

Useful for:
- backup
- inspection
- migration preparation
- offline review before changes

```bash
export SOURCE_ZK=1.92.157.216:2181
export EXPORT_FILE=zookeeper_export.json
./zktool export
```

---

### Import
Import configuration data from a previously exported JSON file into a target environment.

Useful for:
- environment bootstrap
- controlled restore
- migration between environments

```bash
export TARGET_ZK=localhost:3000
export IMPORT_FILE=zookeeper_export.json
./zktool import
```

---

### Update
Apply structured updates from Excel (`.xlsx`) or JSON (`.json`).

Useful for:
- batch updates
- non-code-driven config maintenance
- structured config synchronization

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

---

## Example use cases

### 1. Backup before risky configuration changes
Before touching live Zookeeper data, export the current config state so the team has a recovery point.

### 2. Migrate `/config/product` data across environments
Move configuration data from one environment to another using a structured import/export workflow instead of manual recreation.

### 3. Apply batch changes from structured business input
When configuration updates are maintained in Excel or JSON, use `zktool update` to standardize bulk writes instead of editing nodes one by one.

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

Run the full test suite with:

```bash
go test ./...
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

Because `zktool` changes live configuration data, it should be used with environment-aware review and backup practices.

Recommended precautions:

- export before high-risk updates
- validate target environment carefully
- avoid treating Excel input as inherently correct
- test structured updates in a lower-risk environment first

---

## Why this repo matters

This repository is useful not only as a Zookeeper CLI, but also as an example of how platform/DevOps work can be turned into reusable tooling.

It shows a transition from:

- manual operations
- environment-specific scripts
- one-off fixes

into:

- repeatable workflows
- structured input/output
- reusable operational tooling

That is the real value of `zktool`.

---

## Roadmap ideas

Possible future improvements:

- dry-run mode for updates
- diff preview before import/update
- validation rules for Excel / JSON input
- audit-friendly output / change summary
- safer rollback helpers based on exported snapshots

---

## Summary

`zktool` is a practical configuration governance tool for teams working with Zookeeper-backed systems.

Its value is not just that it can export, import, and update data.
Its real value is that it turns risky, repetitive configuration work into a more repeatable and automation-friendly platform operation.
