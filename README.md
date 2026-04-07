# seed

A Golang CLI for **safe export, import, and bulk update of configuration data** across environments.

`seed` is a configuration governance tool designed to support multiple configuration management systems. Currently it includes a `zookeeper` module; future modules (e.g. `nacos`) can be added as additional subcommands.

`seed` is built for teams who still rely on configuration backends like Zookeeper but want to move away from fragile manual edits, temporary scripts, and one-off operational workflows.

---

## Why this exists

In many real production-like environments, configuration changes are still handled through:

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

`seed` was created to reduce that operational overhead and make configuration workflows safer and more standardized.

---

## What problem it solves

`seed` focuses on a practical platform/DevOps problem:

> How do you safely move, restore, and bulk-update configuration data without turning every change into a custom script or a risky manual operation?

The `zookeeper` module currently supports three core workflows:

1. **Export** — extract Zookeeper configuration data into a portable file
2. **Import** — restore or migrate configuration data into a target environment
3. **Update** — apply structured updates from Excel or JSON into existing Zookeeper nodes

---

## Project structure

```
seed
└── zookeeper    # Zookeeper configuration management
    ├── export
    ├── import
    └── update
```

Future modules (e.g. `nacos`) will be added as top-level subcommands alongside `zookeeper`.

---

## Key capabilities

### zookeeper export
Export Zookeeper data to a JSON file.

```bash
export SOURCE_ZK=1.92.157.216:2181
export EXPORT_FILE=zookeeper_export.json
./seed zookeeper export
```

---

### zookeeper import
Import configuration data from a previously exported JSON file into a target environment.

```bash
export TARGET_ZK=localhost:3000
export IMPORT_FILE=zookeeper_export.json
./seed zookeeper import
```

---

### zookeeper update
Apply structured updates from Excel (`.xlsx`) or JSON (`.json`).

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

1. **Export** current configuration
2. Review or prepare the update input
3. **Update** or **Import** into the target environment
4. Validate that target nodes match expectations
5. Keep export files as rollback references when appropriate

---

## Safety note

Because `seed` changes live configuration data, it should be used with environment-aware review and backup practices.

- export before high-risk updates
- validate target environment carefully
- avoid treating Excel input as inherently correct
- test structured updates in a lower-risk environment first

---

## Roadmap ideas

- `nacos` module for Nacos configuration management
- dry-run mode for updates
- diff preview before import/update
- validation rules for Excel / JSON input
- audit-friendly output / change summary
- safer rollback helpers based on exported snapshots

---

## Summary

`seed` is a practical configuration governance tool for teams working with configuration-backed systems.

Its value is that it turns risky, repetitive configuration work into a more repeatable and automation-friendly platform operation.
