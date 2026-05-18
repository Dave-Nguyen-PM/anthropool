# Troubleshooting

Full docs online: **https://anthropool.dave.com/docs**

Common problems and how to fix them. For platform-specific detail see [Windows](./windows.md), [macOS](./macos.md), [Linux](./linux.md).

---

## `anthropool` is not recognized after `npm install -g`

**Symptom**

```
C:\Users\Admin>anthropool setup
'anthropool' is not recognized as an internal or external command,
operable program or batch file.
```

**Cause**

npm's global bin directory is not on your `Path`. This is common on fresh Windows installs.

**Fix**

See [Windows PATH fix](./installation.md#windows-path-fix) in the installation guide.

---

## `anthropool` command exists but postinstall failed to download the binary

**Symptom**

```
anthropool: binary not found at C:\...\node_modules\@dave\anthropool\bin\anthropool.exe.
Postinstall may have failed; try `npm install -g /anthropool --force`
```

**Cause**

The npm postinstall script downloads the native binary from GitHub Releases. If your network blocked the request or GitHub was temporarily unavailable, the download was skipped.

**Fix**

Re-run the install with `--force` to trigger the postinstall again:

```bash
npm install -g /anthropool --force
```

If that still fails (e.g. corporate proxy, offline environment), download the binary manually from the [releases page](https://github.com/Dave-Nguyen-PM/anthropool/releases/latest) and place it at the path shown in the error message (rename it to `anthropool.exe` on Windows or `anthropool` on Linux/macOS).

---

## `anthropool setup` says hooks already installed but `/switch` does not work

**Fix**

Restart Claude Code after running `anthropool setup`. Hooks are only loaded at startup.

```bash
anthropool setup   # re-run to be sure
# then restart Claude Code
```

---

## `anthropool add` says "no active Claude login found"

**Cause**

You need to be logged into Claude Code before adding an account.

**Fix**

```bash
claude login   # log in first
anthropool add        # then register the account
```

---

## Rate-limit swap triggers but Claude does not resume

**Cause**

`auto_resume` may be disabled, or the session ID was not captured.

**Check**

```bash
anthropool config show
```

Look for `auto_resume: false`. If it is false, re-enable it:

```bash
anthropool config set auto_resume true
```

Also confirm the wrapper is running — swap and resume only work when Claude Code was started with `anthropool`, not `claude` directly.

---

## `/switch` responds with "not running under anthropool"

**Cause**

The slash command is gated on the `ANTHROPOOL_WRAPPED` environment variable. It only acts when Claude Code was started via `anthropool`.

**Fix**

Start Claude Code through the wrapper:

```bash
anthropool   # instead of: claude
```

---

## `anthropool upgrade` reports "unknown install method"

**Cause**

`anthropool upgrade` auto-detects whether you installed via npm or the shell installer. If neither marker is found (e.g. you installed a manual binary), it cannot upgrade automatically.

**Fix**

Download the latest binary manually from the [releases page](https://github.com/Dave-Nguyen-PM/anthropool/releases/latest) and replace your existing binary, or re-install via npm:

```bash
npm install -g /anthropool
```

---

## macOS: Gatekeeper blocks the binary

**Symptom**

```
"anthropool" cannot be opened because the developer cannot be verified.
```

**Fix**

```bash
xattr -d com.apple.quarantine /path/to/anthropool
```

Or: **System Settings** → **Privacy & Security** → scroll down → **Allow Anyway**.

Only affects manual binary downloads. npm and shell installer installs are not quarantined.

---

## macOS / Linux: `anthropool` not found inside Claude Code hooks

**Symptom**

`anthropool` works in your terminal but hooks fail with "command not found" or silently do nothing.

**Cause**

Claude Code hooks run in a non-interactive shell. Shell config files like `~/.zshrc` or `~/.bashrc` are not sourced. If `anthropool` lives in a directory that is only on PATH for interactive shells (e.g. via nvm), hooks cannot find it.

**Fix**

Add the PATH export to your login shell profile instead:

- macOS zsh: `~/.zprofile`
- Linux bash: `~/.profile` or `~/.bash_profile`

```bash
# Example for nvm — add to ~/.zprofile (macOS) or ~/.profile (Linux)
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"
```

Then open a new terminal and re-run `anthropool setup` to reinstall hooks.

---

## Linux: `anthropool` not found after shell installer

**Symptom**

```bash
$ anthropool setup
bash: anthropool: command not found
```

**Fix**

The shell installer places the binary at `~/.local/bin/anthropool`. Add `~/.local/bin` to PATH:

```bash
# Add to ~/.bashrc or ~/.zshrc
export PATH="$HOME/.local/bin:$PATH"
```

Then reload: `source ~/.bashrc` or open a new terminal.

---

## Still stuck?

Open an issue at [github.com/Dave-Nguyen-PM/anthropool/issues](https://github.com/Dave-Nguyen-PM/anthropool/issues) and include:

- Your OS and version
- Output of `anthropool --version` (if it runs)
- Output of `npm prefix -g` (if you installed via npm)
- The full error message
