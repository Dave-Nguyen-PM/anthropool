#!/usr/bin/env node
// Bin shim: forward all argv + stdio to the postinstall-downloaded
// anthropool binary. Keeps the npm wrapper transparent — the user types
// `anthropool` and gets the native binary's behaviour (interactive TTY, exit
// codes, signals) untouched.

import { spawnSync } from "node:child_process";
import { existsSync } from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";

const here = dirname(fileURLToPath(import.meta.url));
const binName = process.platform === "win32" ? "anthropool.exe" : "anthropool";
const binPath = join(here, binName);

if (!existsSync(binPath)) {
  console.error(
    `anthropool: binary not found at ${binPath}.\n` +
    "Postinstall may have failed; try `npm install -g /anthropool --force`, " +
    "or download the binary manually from " +
    "https://github.com/Dave-Nguyen-PM/anthropool/releases",
  );
  process.exit(1);
}

const result = spawnSync(binPath, process.argv.slice(2), {
  stdio: "inherit",
  // Inherit env so ANTHROPOOL_*, CLAUDE_*, etc. propagate to the child.
});

if (result.error) {
  console.error("anthropool:", result.error.message);
  process.exit(1);
}
process.exit(result.status ?? 0);
