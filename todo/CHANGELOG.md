# Changelog

## 2026-10-02

### Added

- Initial Todo CLI implementation with `add`, `list`, `completed`, and `delete` commands.
- CSV-backed storage through `DB_FILE` or `$HOME/data.csv`.
- README documentation for building and running the CLI, its data format, commands, output, and logging.
- This changelog.

### Documented current behavior

- `completed ID` currently marks rows whose ID differs from the supplied ID as `Completed`, rather than the matching row. The behavior is documented only; no logic was changed.
