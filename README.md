```
  █████╗ ███╗   ██╗████████╗██╗  ██╗██████╗  ██████╗ ██████╗  ██████╗  ██████╗ ██╗
 ██╔══██╗████╗  ██║╚══██╔══╝██║  ██║██╔══██╗██╔═══██╗██╔══██╗██╔═══██╗██╔═══██╗██║
 ███████║██╔██╗ ██║   ██║   ███████║██████╔╝██║   ██║██████╔╝██║   ██║██║   ██║██║
 ██╔══██║██║╚██╗██║   ██║   ██╔══██║██╔══██╗██║   ██║██╔═══╝ ██║   ██║██║   ██║██║
 ██║  ██║██║ ╚████║   ██║   ██║  ██║██║  ██║╚██████╔╝██║     ╚██████╔╝╚██████╔╝███████╗
 ╚═╝  ╚═╝╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═╝      ╚═════╝  ╚═════╝ ╚══════╝
```

<p align="center">
  Claude Code account switcher for uninterrupted sessions.
</p>

<p align="center">
  <a href="https://github.com/Dave-Nguyen-PM/anthropool/releases/latest">
    <img alt="latest release"
         src="https://img.shields.io/github/v/release/Dave-Nguyen-PM/anthropool?style=for-the-badge&label=Release&color=c8763a&labelColor=0e1116&logo=github&logoColor=f0ead6">
  </a>
  &nbsp;
  <a href="https://www.npmjs.com/package//anthropool">
    <img alt="npm version"
         src="https://img.shields.io/npm/v//anthropool?style=for-the-badge&label=npm&color=c8763a&labelColor=0e1116&logo=npm&logoColor=f0ead6">
  </a>
  &nbsp;
  <a href="https://github.com/Dave-Nguyen-PM/anthropool/blob/main/LICENSE">
    <img alt="license"
         src="https://img.shields.io/badge/License-GPL--3.0-c8763a?style=for-the-badge&labelColor=0e1116&logo=gnu&logoColor=f0ead6">
  </a>
  &nbsp;
  <a href="https://support.dave.com">
    <img alt="support"
         src="https://img.shields.io/badge/Support-%E2%99%A5-c8763a?style=for-the-badge&labelColor=0e1116&logo=githubsponsors&logoColor=f0ead6">
  </a>
  &nbsp;
  <a href="https://anthropool.dave.com/docs">
    <img alt="docs"
         src="https://img.shields.io/badge/Docs-anthropool.dave.com-c8763a?style=for-the-badge&labelColor=0e1116&logo=readthedocs&logoColor=f0ead6">
  </a>
</p>

---

**`anthropool`** is a CLI tool for Claude Code that pools multiple Pro/Max
accounts behind a single live session. When the active account hits
a rate limit, `anthropool` switches to a healthy account and continues the
same conversation. For proactive threshold swaps, it waits for the
current turn to finish first. No logout, no reload, no lost context.


```text
$ anthropool
anthropool: rate limit on alice@example.com → swapped to bob@example.com, resuming…
> What number did I tell you to remember?
4729.
```

---

## Contents

- [Contents](#contents)
- [Install](#install)
  - [Option 1 — npm](#option-1--npm)
  - [Option 2 — shell installer](#option-2--shell-installer)
  - [Option 3 — manual binary](#option-3--manual-binary)
  - [After install](#after-install)
  - [What works on which platform](#what-works-on-which-platform)
- [Quick start](#quick-start)
  - [Verify your setup once](#verify-your-setup-once)
- [How it works](#how-it-works)
- [Daily usage](#daily-usage)
- [Configuration](#configuration)
  - [Strategies](#strategies)
- [Swap history](#swap-history)
- [Data layout](#data-layout)
- [Security](#security)
- [Building from source](#building-from-source)
- [License](#license)
- [Socials](#socials)

## Install

Three install methods. Pick the one that fits your platform — they
all install the same `anthropool` binary.

### Option 1 — npm

Works on Linux, macOS and Windows. Requires Node.js 18 or newer.

```bash
npm install -g /anthropool
```

### Option 2 — shell installer

Works on Linux, macOS, WSL and Git Bash on Windows.

```bash
curl -fsSL https://raw.githubusercontent.com/Dave-Nguyen-PM/anthropool/main/scripts/install.sh | sh
```

### Option 3 — manual binary

Works everywhere. Useful if you don't want Node.js and can't run
shell scripts (e.g. native Windows PowerShell or cmd.exe).

1. Download the matching artefact from the
   [releases page](https://github.com/Dave-Nguyen-PM/anthropool/releases):
   - `anthropool-linux-amd64`, `anthropool-linux-arm64`
   - `anthropool-darwin-amd64`, `anthropool-darwin-arm64`
   - `anthropool-windows-amd64.exe`
2. On Linux/macOS, `chmod +x anthropool-<os>-<arch>` and rename to `anthropool`.
3. Move it somewhere on your `PATH`:
   - Linux/macOS: `~/.local/bin/anthropool`
   - Windows: any directory listed in your `Path` environment variable.

### After install

Run `anthropool setup` once. That installs the `/switch` and `/anthropool:*` slash
commands plus the three Claude Code hooks. Restart Claude Code
afterwards so it picks them up.

### What works on which platform

| | Linux | macOS | WSL / Git Bash | native Windows |
|---|:---:|:---:|:---:|:---:|
| Account commands (`add`, `list`, `switch`, `status`, …) | ✅ | ✅ | ✅ | ✅ |
| Credential storage | file (0600) | Keychain | file (0600) | Credential Manager |
| Hooks + `/switch` and `/anthropool:*` slash commands | ✅ | ✅ | ✅ | ✅ |
| Auto-resume on swap | ✅ | ✅ | ✅ | ✅ † |
| `npm install -g /anthropool` | ✅ | ✅ | ✅ | ✅ |
| `curl … \| sh` shell installer | ✅ | ✅ | ✅ | ❌ |
| Manual binary download | ✅ | ✅ | ✅ | ✅ |

† On native Windows the wrapper hard-terminates `claude` on swap (Go
can't send `SIGINT` cross-process there). The `Stop` hook still
flushes the transcript before the wrapper acts, so the resumed
conversation is intact — but if you see anything unexpected, please
[open an issue](https://github.com/Dave-Nguyen-PM/anthropool/issues).

## Quick start

Website: https://anthropool.dave.com

```bash
anthropool setup           # install /switch, /anthropool:* + Claude Code hooks
anthropool add             # register the currently-logged-in account
claude logout && claude login   # log into your second account
anthropool add             # register it
anthropool                 # launch claude under anthropool instead of `claude` directly
```

After `anthropool setup`, restart Claude Code so it picks up the newly
installed hooks. From then on:

- `/switch` from inside a Claude Code session rotates accounts.
- `/switch <slot|email>` switches to a specific one.
- `/anthropool:add`, `/anthropool:list`, `/anthropool:status`, `/anthropool:support`,
  `/anthropool:remove`, and `/anthropool:switch`, `/anthropool:config`, `/anthropool:usage-refresh` run
  account-management commands in-session.
- A rate-limit response from the API auto-triggers the same flow and
  does not wait for another Stop hook before reconnecting.

### Verify your setup once

A 30-second check that proves end-to-end context preservation:

1. Send: *"Please remember the number 4729."*
2. Wait for the reply.
3. Send `/switch`.
4. After the ~2-second reconnect, ask: *"What number did I tell you?"*

If the answer is `4729`, swap-and-resume is working.

## How it works

```
   user types     ┌────── claude (running, account A) ──────┐
   /switch ──────►│  hooks: UserPromptSubmit, Stop,         │
   or rate-limit  │         SessionStart, PostToolUseFailure│
   ───────────────┴──┬──────────────────────────────────────┘
                     │ writes signal files
                     │ runtime/signals/{wrapperPID}-{name}
                     ▼
             ┌──────────────────────────────────────┐
             │  anthropool wrapper                         │  polls signals
             │   on rate-limit OR threshold OR      │  every 100 ms
             │   /switch:                           │
             │     rate-limit/manual: exit now      │  avoids hard-limit stall
             │     threshold: wait for Stop signal  │  guarantees flush
             │     ask claude to exit cleanly       │
             │     swap creds (transactional)       │
             │     append history.Entry             │
             │     relaunch claude --resume <id>    │
             │       [optional auto_message]        │
             └──────────────────────────────────────┘
```

`anthropool` writes its hooks into `~/.claude/settings.json` by signature,
so it never modifies entries owned by other tools and
`anthropool uninstall-hooks` removes only its own. Every anthropool-owned file goes
through atomic writes (`tmp + fsync + rename`) and state mutations
are serialised with file locks (`flock` / `LockFileEx`).

## Daily usage

```bash
anthropool                          # launch claude under the wrapper
anthropool list                     # accounts with 5h / 7d utilisation
anthropool list --refresh           # refresh usage before listing
anthropool status                   # current login + anthropool state
anthropool switch <slot|email>      # manual swap (no auto-resume)
anthropool remove <slot|email>      # forget an account
anthropool history                  # recent swaps with reasons
anthropool usage refresh            # poll all account usage
anthropool config show              # current settings
anthropool config edit              # interactive settings editor
anthropool upgrade                  # update anthropool (npm or installer; auto-detects)
```

From inside a session started with `anthropool`:

```text
/switch                      # rotate per the configured strategy
/switch 2                    # by slot number
/switch alt@example.com      # by email
/anthropool:switch 2                # same switch flow under the /anthropool namespace
/anthropool:add                     # add/refresh the current login
/anthropool:list --refresh          # list accounts from inside Claude Code
/anthropool:status                  # show live login + anthropool state
/anthropool:support                 # show support URL
/anthropool:config show             # show anthropool configuration
/anthropool:remove 2                # remove an account
/anthropool:usage-refresh           # refresh account usage
```

If Claude is already hard-blocked, `/switch` is handled by anthropool's
`UserPromptSubmit` hook before Claude processes the prompt. If you are
on an older session that was started before that hook was installed,
run this from another terminal:

```bash
anthropool force-switch             # rotate the active anthropool-wrapped session
anthropool force-switch 2           # force a specific slot/email
```

## Configuration

```bash
anthropool config keys                                      # discover everything
anthropool config show
anthropool config set thresholds.five_hour 85
anthropool config set strategy.kind balanced
anthropool config set strategy.order alice@x,bob@x         # drain priority
anthropool config set auto_message ""                      # silent resume
anthropool config set update_check.enabled true            # opt in to update checks
```

| Key | Default | Description |
|---|---|---|
| `thresholds.five_hour`        | `100`          | Auto-swap when 5h utilisation hits this %. `100` = reactive only. |
| `thresholds.seven_day`        | `100`          | Auto-swap when 7d utilisation hits this %. `100` = reactive only. |
| `strategy.kind`               | `drain`        | `drain` / `balanced` / `manual` |
| `strategy.order`              | `[]`           | Drain mode priority (emails); empty = auto by highest 7d |
| `auto_switch_on_threshold`    | `true`         | Master toggle for pre-emptive threshold swap |
| `auto_switch_on_rate_limit`   | `true`         | Master toggle for swap on rate-limit hook |
| `auto_resume`                 | `true`         | Pass `--resume <id>` to the relaunched claude |
| `auto_message`                | `Go continue.` | First user turn after auto-swap; `""` = silent |
| `update_check.enabled`        | `false`        | Check GitHub for newer anthropool releases on startup |
| `update_check.cadence_hours`  | `6`            | Minimum hours between update checks (cached locally) |

Config file: `~/.config/anthropool/config.json` (XDG-aware).

### Strategies

- **drain** — Use one account until its 7-day cap is near, then move
  on. Set `order` for explicit priority, or leave empty to auto-drain
  the highest-7d account first.
- **balanced** — Always pick the account with the lowest 7-day
  utilisation (tiebreak by lowest 5h).
- **manual** — Never swap automatically. `/switch` and `anthropool switch`
  still work.

## Swap history

```text
$ anthropool history
2026-05-01 14:12:33  alice@x → bob@x  [threshold]
    reason: 7d utilization 95% ≥ threshold 95%
    usage: alice@x 5h:34% 7d:95% → bob@x 5h:8% 7d:30%
    session: 143eec0f-277e-4ce1-95f1-58eb56331874

2026-05-01 13:55:08  bob@x → alice@x  [manual]
    reason: user requested via /switch
```

`anthropool history -n 5` for the last five, `anthropool history --json` to pipe,
`anthropool history --clear` to wipe. History is capped at 1000 entries.

## Data layout

```
~/.local/share/anthropool/                  (~/.anthropool/ on macOS/Windows)
├── state.json                      # accounts, sequence, active slot
├── .lock                           # flock target for state mutations
├── accounts/
│   └── 01-user@example.com/
│       ├── credentials.json        # Linux only; macOS/Win uses keystore
│       └── oauth.json              # the oauthAccount block, raw JSON
└── runtime/
    ├── signals/                    # hook → wrapper signal files
    ├── usage-cache.json            # per-account 5h / 7d snapshot
    ├── update-cache.json           # update-check cadence cache
    └── swap-history.json           # capped at 1000 entries

~/.config/anthropool/config.json           # XDG_CONFIG_HOME-aware
~/.claude/settings.json             # hooks upserted here
~/.claude/commands/switch.md        # /switch slash command
~/.claude/commands/anthropool/*.md         # /anthropool:* account commands
```

## Security

- **Tokens are never logged.** Credential blobs are opaque to
  logging; the helper that extracts a token never surfaces it in any
  error message.
- **All anthropool-owned directories and files are 0700 / 0600.**
- **The installer refuses to run as root** unless inside a container.
- **`/switch` is gated by `ANTHROPOOL_WRAPPED`** — the slash command refuses
  to act unless anthropool is the parent process, so it cannot accidentally
  signal an unrelated `claude`.
- **Hook upsert is signature-keyed.** `anthropool install-hooks` only ever
  modifies entries whose `command` field contains the literal string
  `anthropool ` or `/anthropool ` — every other tool's hooks are preserved.

## Building from source

```bash
git clone https://github.com/Dave-Nguyen-PM/anthropool
cd anthropool
go build -o anthropool ./cmd/anthropool
./anthropool help
```

Requires Go 1.21+.

## License

[GPL-3.0-only](./LICENSE). Modifying and redistributing `anthropool` is
welcome; if you do, your changes need to ship under GPL-3.0 too.

## Socials

All dave social channels live at
**[socials.dave.com](https://socials.dave.com)** — one place
for updates, other projects, and how to reach me.

---

If `anthropool` saves you time, you can support development at
[support.dave.com](https://support.dave.com).
