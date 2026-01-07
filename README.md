# GCM - Git Commit Manager

An interactive Git commit tool that enforces best practices and makes creating well-structured commits a breeze.

## Features

###  Core Functionalities

#### 1. **Intelligent Change Detection**
- Automatically detects all tracked files with changes
- Categorizes changes visually:
  - **MODIFIED** - Changed files
  - **DELETED** - Removed files
  - **RENAMED** - Renamed/moved files
  - **ADDED** - New tracked files
  - **UNTRACKED** - New untracked files
- Clean working tree message when nothing to commit

#### 2. **Smart Branch Management**
- **Prevents direct commits to main/master** - Forces branch creation
- For other branches: Option to use current branch or create new one
- Branch name validation:
  - No spaces or invalid characters
  - Max 50 characters
  - Suggested format: `type/short-description`
  - Examples: `feat/login`, `fix/button-crash`, `chore/deps`

#### 3. **Interactive File Selection**
- Beautiful TUI with categorized file display
- Navigate with arrow keys or vim bindings (j/k)
- Keyboard shortcuts:
  - `SPACE` - Toggle file selection
  - `a` - Select all files
  - `d` - Deselect all files
  - `i` - Invert selection
  - `ENTER` - Confirm selection
  - `q` / `ESC` - Cancel and exit
- Visual indicator of selected files
- Tip: Group related changes in the same commit

#### 4. **Conventional Commits Support**
- **Type selection** with predefined options:
  1. `feat` - New feature
  2. `fix` - Bug fix
  3. `docs` - Documentation
  4. `style` - Formatting/styles (no logic changes)
  5. `refactor` - Refactoring (no functional changes)
  6. `perf` - Performance improvement
  7. `test` - Tests
  8. `chore` - Maintenance tasks
  9. `build` - Build system
  10. `ci` - CI/CD
  - Custom type option available

#### 5. **Commit Message Validation**
- **Title requirements:**
  - Minimum 10 characters, maximum 72
  - Must start with lowercase letter (conventional commits)
  - Cannot end with a period
  - Warning if over 50 characters
- **Optional description:**
  - Multi-line support
  - Best practice tip: Explain the "why", not the "what"
- **Live preview** before committing

#### 6. **Loop Until Working Tree is Clean**
- After each commit, checks for remaining changes
- Option to create another commit on the same branch
- Shows current status:
  - Files committed
  - Files still uncommitted
- Continue committing related changes without restarting

#### 7. **Intelligent Push Support**
- After all commits, option to push to remote
- Automatically detects if branch exists on remote
- Uses `-u` flag for first-time push
- Summary of all commits created in session
- Helpful message if skipping push

## Installation

```bash
# Clone the repository
git clone <repo-url>
cd gcm

# Build the binary
go build -o gcm .

# Optional: Install globally
sudo mv gcm /usr/local/bin/
```

## Usage

Simply run `gcm` in any git repository:

```bash
gcm
```

The tool will guide you through:
1.  Detecting changes
2.  Branch management
3.  File selection
4.  Commit type selection
5. ️ Commit message creation
6.  Preview and confirmation
7.  Loop for additional commits
8.  Optional push to remote

## Examples

### Creating a Feature Commit

```bash
$ gcm

 Branch Management
Current branch: main

️ Cannot commit directly to 'main'.
Please create a new branch:

Branch name: feat/user-authentication_

 Created and switched to branch 'feat/user-authentication'

 File Selection
Navigate with ↑/↓, SPACE to mark, ENTER to continue

MODIFIED:
 [x] src/auth/login.go
 [x] src/auth/register.go
 [ ] src/utils/helpers.go

Selected: 2/3

 Select Commit Type
> 1. feat     - New feature

 Commit Title
Type: feat

Commit title (required):
> add user authentication with jwt_

 Commit Preview
+----------------------------------------------------------------------+
|  feat: add user authentication with jwt                             |
+----------------------------------------------------------------------+

Confirm commit message? (y/n/e to edit): y

 Commit created: [feat] add user authentication with jwt

Want to create another commit on the same branch? (y/n): n

 Commits created in this session:
  1. [feat] add user authentication with jwt

Push to remote? (y/n): y

 Pushing to origin/feat/user-authentication...
 Successfully pushed to remote

```

## Validations

###  Required
- Cannot commit directly to `main`/`master` without creating a branch
- Commit title cannot be empty
- At least 1 file must be selected
- Branch name must be valid

### ⚠ Warnings (don't block)
- Title over 50 characters (warns but allows)
- Large commits (>10 files)
- Branch name doesn't follow suggested convention

###  Prohibitions
- Branch names with spaces or invalid characters
- Commit titles over 72 characters
- Commit titles ending with period
- Commit titles starting with uppercase

## Dependencies

- Go 1.25+
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

A lightweight git wrapper focused on atomic commits and reduced friction.
