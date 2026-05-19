# CLAUDE.md — anthropool

## What it does
Wraps the `claude` binary to pool multiple Claude Code Pro/Max accounts. Runs
`claude` as a child process, listens for hook-written signal files, and on a
rate-limit, usage-threshold, or manual `/switch` event it swaps credentials and
relaunches `claude --resume <session_id>` so the conversation continues on the
new account without a terminal restart.

## Build & run
```bash
go build ./cmd/anthropool   # produces ./anthropool
go test ./...
go vet ./...
```
Requires `claude` on PATH or `ANTHROPOOL_CLAUDE_BIN` set to the absolute path.
Dev install: `go install ./cmd/anthropool`.

## First-time setup
```bash
anthropool setup   # installs hooks in ~/.claude/settings.json + /switch slash command
anthropool add     # run once per account while that account is logged in
```

## Project layout
```
cmd/anthropool/     CLI dispatch, all subcommands, UI rendering (main.go, setup.go)
internal/
  wrapper/          Core loop: launch claude, poll signals, swap, re-launch
  signals/          File-based hook→wrapper bus ($RUNTIME/signals/{pid}-{name})
  hooks/            Hook handlers (stop, session-start, rate-limit, prompt-submit)
  hookinstall/      Install/uninstall entries in ~/.claude/settings.json
  switcher/         add/switch/remove accounts under lockfile
  store/            state.json: accounts, active slot, rotation metadata
  creds/            Credential read/write (keyring mac/win; file linux)
  claudecfg/        Parse/write oauthAccount block in ~/.claude/.claude.json
  strategy/         drain / balanced / manual rotation algorithms
  monitor/          Fetch + cache usage from api.claude.ai
  usage/            Usage cache types + threshold evaluation
  config/           anthropool config (thresholds, auto-switch flags, strategy)
  history/          Append-only swap log
  paths/            All on-disk locations — single source of truth
  lockfile/         Cross-platform advisory lock
  atomicfile/       Atomic write helper (temp + rename)
  updater/          GitHub release check + upgrade command
npm/                npm package (postinstall downloads the binary)
scripts/install.sh  Curl installer
slashcmd/           /switch markdown source
```

## Architecture
1. **Wrapper** starts claude with `ANTHROPOOL_WRAPPED=1` and `ANTHROPOOL_WRAPPER_PID=<pid>`.
2. **Poll loop** checks `$BACKUP_ROOT/runtime/signals/` every 100 ms for four events:
   `session-started`, `stopped`, `rate-limited`, `switch-requested`.
3. Rate-limit and manual swap: send SIGINT immediately (5 s grace, then SIGKILL).
   Threshold swap: wait for a `stopped` signal so the turn transcript is flushed first.
4. After claude exits: `switcher.SwitchTo` backs up current creds, writes target creds
   + oauthAccount, updates state.json, then relaunches with `--resume <sessionId>`.
5. **Signal files**: `{pid}-{name}`, written atomically by hooks, deleted by wrapper on
   consume. PID namespacing isolates concurrent terminals.
6. **Data roots**: `~/.anthropool` (mac/win) or `$XDG_DATA_HOME/anthropool` (linux).
   Claude Code state stays in `~/.claude` — anthropool never invents paths there.

## Conventions
- All paths via `internal/paths` — never build paths inline elsewhere.
- All credential/state mutations under `lockfile.Acquire`.
- All concurrent writes via `atomicfile.Write`.
- Hook handlers always exit 0; log errors to stderr, never to stdout.
- `os.Exit` only in `cmd/anthropool`; `internal/` packages return errors.
- No output from hook handlers — Claude Code captures stdout; stray text corrupts tool results.

## Testing
```bash
go test ./...                       # all packages
go test -v ./internal/wrapper/...   # wrapper signal/swap logic
go test -v ./internal/hooks/...     # hook JSON parsing
```
Tests use temp dirs; no network calls, no real claude binary required.
End-to-end: run `anthropool setup`, then use `/switch` inside a live session.
