---
description: Switch the active anthropool-managed Claude Code account
argument-hint: "[slot-number-or-email]"
allowed-tools: Bash(anthropool __slash-switch:*)
---

# /anthropool:switch

Switches the active anthropool-managed account in place — the conversation
continues on the new account with no restart. With no argument, rotates
to the next account.

```bash
anthropool __slash-switch $ARGUMENTS
```
