# macOS

## Installation

All three methods work on macOS. npm is the fastest:

```bash
npm install -g /anthropool
anthropool setup
```

---

## nvm / fnm PATH ordering

If you installed Node.js via **nvm** or **fnm**, the npm global bin may not be on your `PATH` in non-interactive shells (e.g. when hooks run inside Claude Code).

**Check:**

```bash
which anthropool
# should print something like /Users/<you>/.nvm/versions/node/v22.x.x/bin/anthropool
```

If `which anthropool` returns nothing, add the nvm init to your login shell profile:

**For zsh (default on macOS 10.15+):**

```bash
# Add to ~/.zprofile (login shells) — NOT only ~/.zshrc (interactive only)
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"
```

Then open a new terminal and re-run `which anthropool`.

**Why `~/.zprofile` and not `~/.zshrc`?**
Claude Code hooks are run in a non-interactive shell. `~/.zshrc` is only sourced for interactive shells. `~/.zprofile` is sourced for login shells, which covers the hook environment.

---

## Gatekeeper blocks the binary

If you used the manual binary download and macOS quarantines it:

```
"anthropool" cannot be opened because the developer cannot be verified.
```

**Fix:**

```bash
xattr -d com.apple.quarantine /path/to/anthropool
```

Or via **System Settings** → **Privacy & Security** → scroll to the blocked app → click **Allow Anyway**.

The npm and shell installer methods are not affected because the binary is downloaded by a trusted process (Node.js / curl), not opened directly from a browser download.

---

## Keychain prompts on first `anthropool add`

On macOS, `anthropool add` stores credentials in the **macOS Keychain**. The first time, macOS may ask:

```
"anthropool" wants to use your confidential information stored in "anthropool/<email>" in your keychain.
```

Click **Always Allow** to avoid being prompted on every anthropool invocation.

If you want to remove anthropool credentials from Keychain manually: open **Keychain Access** → search for `anthropool` → delete the matching entries.

---

## Homebrew-installed Node.js

If Node.js was installed via Homebrew, the global bin is typically at:

```bash
$(brew --prefix)/bin
```

Which is `/usr/local/bin` on Intel Macs or `/opt/homebrew/bin` on Apple Silicon. Both are usually on `PATH` already. Run `which anthropool` to confirm.

---

## Apple Silicon (M1 / M2 / M3)

The arm64 binary (`anthropool-darwin-arm64`) is downloaded automatically. No Rosetta needed.

---

## Shell installer

```bash
curl -fsSL https://raw.githubusercontent.com/Dave-Nguyen-PM/anthropool/main/scripts/install.sh | sh
```

The installer places the binary at `~/.local/bin/anthropool`. Ensure `~/.local/bin` is on your `PATH`:

```bash
# Add to ~/.zprofile
export PATH="$HOME/.local/bin:$PATH"
```
