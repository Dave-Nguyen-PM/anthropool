# /anthropool

Multi-account switcher for Claude Code with hook-driven auto-resume.

This package is a thin npm wrapper. On install it downloads the right
prebuilt `anthropool` binary for your platform (Linux x64/arm64, macOS
x64/arm64, Windows x64) from the
[GitHub release matching its version](https://github.com/Dave-Nguyen-PM/anthropool/releases)
and exposes it as the `anthropool` command on your `PATH`.

```bash
npm install -g /anthropool
anthropool setup
anthropool add        # or /anthropool:add inside Claude while logged in
```

For full documentation (auto-swap on rate-limit, threshold-based
swaps, strategies, history, configuration), see the main repository:

- **Docs:** https://anthropool.dave.com/docs
- **Website:** https://anthropool.dave.com
- **README:** https://github.com/Dave-Nguyen-PM/anthropool#readme
- **Releases:** https://github.com/Dave-Nguyen-PM/anthropool/releases
- **Issues:** https://github.com/Dave-Nguyen-PM/anthropool/issues

## What it does

- `/switch` from inside a Claude session to swap accounts and resume
  the same conversation.
- `/anthropool:switch`, `/anthropool:add`, `/anthropool:list`, `/anthropool:support`, `/anthropool:config`, and related commands for in-session
  account management.
- `anthropool support` and `/anthropool:support` show the support URL.
- Automatic swap when an account hits a rate limit, with auto-resume.
- Threshold-based pre-emptive swap before a window caps.
- One Go binary, no Python, no Bash version requirement, no `jq`.

## How the npm wrapper works

`postinstall` runs `install.js`, which:

1. Reads `process.platform` and `process.arch` to pick the right
   release artefact name (e.g. `anthropool-darwin-arm64`).
2. Downloads the binary from
   `https://github.com/Dave-Nguyen-PM/anthropool/releases/download/v<version>/<asset>`.
3. Verifies the sidecar `<asset>.sha256` if present.
4. `chmod +x` on POSIX, then exposes it as the `anthropool` command via `bin`.

Override the source repo with `ANTHROPOOL_RELEASE_REPO=<owner>/<repo>` if you
maintain a fork. To skip the network entirely, install with
`npm install -g /anthropool --ignore-scripts` and drop the binary at
`node_modules//anthropool/bin/anthropool` (or `anthropool.exe` on Windows) yourself.

## Troubleshooting

**`anthropool` not recognized after install (Windows)**

npm's global bin directory is likely missing from `Path`. Find it:

```cmd
npm prefix -g
```

Then add the output (e.g. `C:\Users\Admin\AppData\Roaming\npm`) to your `Path` via System Properties → Environment Variables → User variables → Path → New.

Open a new terminal window after saving. Full instructions: https://anthropool.dave.com/docs

**Postinstall failed to download the binary**

```bash
npm install -g /anthropool --force
```

If you are behind a corporate proxy, set `HTTPS_PROXY` first. Or download the binary directly from the [releases page](https://github.com/Dave-Nguyen-PM/anthropool/releases/latest) and drop it at the path shown in the error.

**`anthropool` not found inside Claude Code hooks (macOS / Linux with nvm)**

Hooks run in non-interactive shells. Add your PATH setup to `~/.zprofile` (macOS) or `~/.profile` (Linux), not just `~/.zshrc` / `~/.bashrc`.

More: https://anthropool.dave.com/docs

---

## License

GPL-3.0-only — see https://github.com/Dave-Nguyen-PM/anthropool/blob/main/LICENSE.

---

If anthropool saves you time, you can support development at
[support.dave.com](https://support.dave.com).
