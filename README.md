# gset

A dead-simple CLI to configure git settings following best practices.

Run one command to set up your identity and apply 18 curated git settings — globally for your machine, or per-repository for a different identity.

## Install

```bash
brew install theenclave-corp/tap/gset
```

Or build from source:

```bash
go install github.com/theenclave-corp/gset@latest
```

## Usage

### Configure your machine (global)

```bash
gset global
```

Prompts for your name, email, and editor, then applies your identity plus 18 best-practice settings to `~/.gitconfig`.

```bash
gset global --name "Jane Doe" --email "jane@example.com" --editor vscode --yes
```

Flags:
- `--name` — git user.name
- `--email` — git user.email
- `--editor` — one of: `vscode`, `vim`, `nano`, `emacs`
- `--yes`, `-y` — skip confirmation prompt
- `--dry-run` — show what would be applied, make no changes
- `--no-best-practices` — set identity only, skip best-practice settings

### Override identity for a repository (local)

```bash
gset local
```

Sets `user.name` and `user.email` in the current repository's `.git/config`. Useful when a work repo needs a different email than your global identity.

```bash
gset local --name "Jane Doe" --email "jane@work.com" --yes
```

### Inspect current configuration

```bash
gset status
```

Shows your effective git configuration (global + local merged), sorted alphabetically. Keys overridden locally are annotated with `(local)`.

## Best-practice settings

`gset global` applies these 18 settings on top of your identity:

| Setting | Value | Why |
|---------|-------|-----|
| `init.defaultBranch` | `main` | Modern default branch name |
| `branch.sort` | `-committerdate` | Sort branches by most recently used |
| `column.ui` | `auto` | Display branch list in columns |
| `pull.rebase` | `true` | Avoid accidental merge commits on pull |
| `push.default` | `simple` | Push current branch to same-named remote |
| `push.autoSetupRemote` | `true` | Auto-configure upstream on first push |
| `fetch.prune` | `true` | Remove stale remote-tracking branches |
| `fetch.pruneTags` | `true` | Remove stale remote-tracking tags |
| `diff.algorithm` | `histogram` | More readable diffs for moved/refactored code |
| `diff.colorMoved` | `plain` | Highlight moved code in a distinct color |
| `merge.conflictstyle` | `zdiff3` | Show original base in conflict markers |
| `rebase.autoSquash` | `true` | Auto-combine fixup! commits |
| `rebase.autoStash` | `true` | Stash/restore dirty working tree around rebase |
| `commit.verbose` | `true` | Show diff while writing commit message |
| `rerere.enabled` | `true` | Remember and reuse conflict resolutions |
| `help.autocorrect` | `prompt` | Ask before auto-correcting mistyped commands |
| `core.fsmonitor` | `true` | Faster git status via filesystem events |
| `core.untrackedCache` | `true` | Cache untracked file info for speed |

## Why global vs local?

Best-practice settings apply to your entire workflow — they're not project-specific. Your global config is where they belong.

Local config exists for one common case: you have a personal email globally but need a work email for a specific repository. `gset local` handles exactly that.

## License

MIT
