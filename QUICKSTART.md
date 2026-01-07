# GCM Quick Start Guide

## What is GCM?

GCM (Git Commit Manager) is an interactive tool that helps you create well-structured, conventional commits with ease. It guides you through the entire commit process with validations and best practices built-in.

## Quick Install

```bash
# Build from source
cd /path/to/gcm
go build -o gcm .

# Optional: Make it globally available
sudo cp gcm /usr/local/bin/
```

## First Time Use

1. Navigate to any git repository with changes
2. Run `gcm`
3. Follow the interactive prompts

## The Workflow

### 1 Branch Check
- If on `main`/`master`: You'll be forced to create a new branch
- If on another branch: Choose to use it or create a new one

### 2 File Selection
```
Use arrow keys (↑↓) or vim keys (jk) to navigate
Press SPACE to select/deselect files
Press 'a' to select all
Press 'd' to deselect all
Press 'i' to invert selection
Press ENTER when ready
```

### 3 Commit Type
Choose from standard conventional commit types:
- `feat` - New feature
- `fix` - Bug fix
- `docs` - Documentation
- `style` - Code style changes
- `refactor` - Code refactoring
- `perf` - Performance improvements
- `test` - Adding tests
- `chore` - Maintenance
- `build` - Build system
- `ci` - CI/CD changes

Or press 'c' for a custom type.

### 4 Commit Message
**Title (required):**
- 10-72 characters
- Start with lowercase
- No period at the end
- Be descriptive but concise

**Description (optional):**
- Press ENTER to skip
- Or type a detailed explanation
- Press CTRL+D when done

### 5 Preview & Confirm
Review your commit before finalizing:
- `y` - Confirm and commit
- `n` - Cancel
- `e` - Edit (go back to title)

### 6 Continue or Push
After committing:
- If more files remain: Option to create another commit
- When done: Option to push to remote

## Tips & Best Practices

###  Writing Good Commit Messages
- **Title**: What was changed
- **Description**: Why it was changed

###  Grouping Changes
- Keep related changes together in one commit
- Separate unrelated changes into multiple commits
- Use the "create another commit" feature for multi-commit workflows

###  Branch Naming
Good examples:
- `feat/user-authentication`
- `fix/memory-leak`
- `docs/api-updates`
- `refactor/database-layer`

Bad examples:
- `my branch` (has spaces)
- `feature` (too vague)
- `.hidden` (starts with dot)

###  Keyboard Shortcuts Reference

**File Selection:**
- `↑/k` - Move up
- `↓/j` - Move down
- `SPACE` - Toggle selection
- `a` - Select all
- `d` - Deselect all
- `i` - Invert selection
- `ENTER` - Confirm
- `q/ESC` - Cancel

**General:**
- `CTRL+C` - Exit anytime
- `ESC` - Go back or cancel

## Common Scenarios

### Scenario 1: Single Feature Commit
```bash
$ gcm
# 1. Create/use feature branch
# 2. Select all related files
# 3. Choose 'feat' type
# 4. Write descriptive message
# 5. Confirm and push
```

### Scenario 2: Multiple Related Commits
```bash
$ gcm
# 1. Use existing branch
# 2. Select files for first commit (e.g., core changes)
# 3. Create commit
# 4. Say 'yes' to create another commit
# 5. Select files for second commit (e.g., tests)
# 6. Create commit
# 7. Push when done
```

### Scenario 3: Bug Fix with Tests
```bash
$ gcm
# First commit - the fix
# - Select only the fixed files
# - Type: 'fix'
# - Message: "resolve memory leak in data processor"

# Second commit - the tests
# - Select only test files
# - Type: 'test'
# - Message: "add regression tests for memory leak fix"
```

## Troubleshooting

### "Cannot commit directly to main"
This is intentional! Create a branch instead. Main/master should be protected.

### "Title too short/long"
- Minimum: 10 characters
- Maximum: 72 characters
- Recommended: Under 50 characters

### "Branch name contains invalid characters"
Avoid: spaces, `~`, `^`, `:`, `?`, `*`, `[`, `]`, `\`

### No files appear
Make sure you have:
1. Initialized a git repository (`git init`)
2. Made changes to tracked files
3. Or added new files (`git add <file>`)

## Advanced Usage

### Installing Globally
```bash
# macOS/Linux
sudo cp gcm /usr/local/bin/
gcm --version

# Or add to PATH
export PATH="$PATH:/path/to/gcm"
```

### Using with Git Aliases
```bash
# Add to ~/.gitconfig
[alias]
    cm = !gcm
    
# Now you can use
git cm
```

## What's Next?

Once you're comfortable with the basics:
1. Try the loop feature for multi-commit workflows
2. Experiment with different commit types
3. Use descriptions for complex changes
4. Practice good branch naming

---

**Need help?** Check the full README.md for detailed documentation.

**Found a bug?** Please report it on GitHub.

**Have a feature request?** Contributions are welcome!

