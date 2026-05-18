# Windows

## The most common problem: `anthropool` not recognised after npm install

```
C:\Users\Admin>anthropool setup
'anthropool' is not recognized as an internal or external command,
operable program or batch file.
```

**Root cause:** npm's global bin directory is not on your `Path`. This is common on fresh Windows installs — Windows does not add it automatically.

### Step 1 — find your npm global bin directory

Open cmd or PowerShell:

```cmd
npm prefix -g
```

Typical output:

```
C:\Users\Admin\AppData\Roaming\npm
```

### Step 2 — add it to Path

Pick one method:

**Option A — Current session only (quick test)**

```powershell
$env:Path += ";$(npm prefix -g)"
anthropool --version
```

If `anthropool --version` prints the version, the fix works. Open a new terminal to make it permanent (Option B or C).

**Option B — GUI (permanent)**

1. Press `Win + R`, type `sysdm.cpl`, press Enter.
2. Go to **Advanced** → **Environment Variables**.
3. Under **User variables**, select `Path` and click **Edit**.
4. Click **New** and paste the path from Step 1 (e.g. `C:\Users\Admin\AppData\Roaming\npm`).
5. Click OK on all dialogs.
6. Open a **new** terminal window and run `anthropool --version`.

**Option C — PowerShell one-liner (permanent)**

```powershell
[Environment]::SetEnvironmentVariable(
  "Path",
  "$([Environment]::GetEnvironmentVariable('Path','User'));$(npm prefix -g)",
  "User"
)
```

Open a new terminal after running this — existing windows keep the old `Path`.

---

## cmd.exe vs PowerShell vs Windows Terminal

All three work once `Path` is set. There is no anthropool-specific difference between them.

If you use **Windows Terminal** with multiple profiles (cmd, PowerShell, Git Bash), each profile reads `Path` from the environment at startup, so you only need to set it once via Option B or C.

---

## Native Windows vs WSL

| | Native Windows (cmd / PowerShell) | WSL (Ubuntu, Debian, etc.) |
|---|---|---|
| Install via npm | `npm install -g /anthropool` | `npm install -g /anthropool` |
| Binary downloaded | `anthropool-windows-amd64.exe` | `anthropool-linux-amd64` or `anthropool-linux-arm64` |
| PATH to fix | Windows `Path` (see above) | Linux `PATH` (see [Linux guide](./linux.md)) |
| Credential storage | Windows Credential Manager | File (`~/.local/share/anthropool/`) |
| `curl … \| sh` installer | Does not work | Works |

**Do not mix** a Windows npm install with WSL usage. Install anthropool inside WSL with the Linux npm or the shell installer — do not try to call the Windows binary from WSL.

---

## Manual binary install on Windows

If you cannot use npm or the shell installer:

1. Download `anthropool-windows-amd64.exe` from the [releases page](https://github.com/Dave-Nguyen-PM/anthropool/releases/latest).
2. Rename it to `anthropool.exe`.
3. Move it to a directory that is already on your `Path`, for example:
   - `C:\Windows\System32\` (system-wide, requires admin)
   - `C:\Users\<YourName>\bin\` (create this folder and add it to your user `Path`)
4. Open a new terminal and run `anthropool --version`.

---

## Postinstall failed to download the binary

If npm install succeeds but running `anthropool` prints:

```
anthropool: binary not found at C:\...\anthropool.exe.
Postinstall may have failed; try `npm install -g /anthropool --force`
```

Re-trigger the postinstall:

```cmd
npm install -g /anthropool --force
```

If you are behind a corporate proxy that blocks GitHub, set the proxy first:

```cmd
set HTTPS_PROXY=http://proxy.corp.example.com:8080
npm install -g /anthropool --force
```

Or download the binary manually from the releases page and drop it at the path shown in the error.

---

## Windows Credential Manager

On Windows, `anthropool add` stores Claude credentials in **Windows Credential Manager** (not a plain file). You can view or remove entries there:

1. Open **Control Panel** → **Credential Manager** → **Windows Credentials**.
2. Look for entries named `anthropool/<email>`.

Removing a credential here is equivalent to `anthropool remove <email>`.

---

## Auto-resume on native Windows

On native Windows, `anthropool` hard-terminates `claude.exe` on a swap because Go cannot send `SIGINT` cross-process on Windows. The `Stop` hook still flushes the transcript before the wrapper acts, so the resumed conversation is intact. If you see unexpected behaviour, [open an issue](https://github.com/Dave-Nguyen-PM/anthropool/issues).

On WSL, normal `SIGINT` is used and behaviour matches Linux/macOS exactly.
