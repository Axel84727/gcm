# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.2.0] - 2026-01-07

### Added

#### Core Features
- **Intelligent Branch Management**
  - Detects current branch and prevents direct commits to main/master
  - Interactive prompt to use current branch or create new one
  - Branch name validation with comprehensive rules
  - Suggested format: `type/short-description`
  
- **Enhanced File Selection UI**
  - Categorized display by change type (MODIFIED, ADDED, DELETED, RENAMED, UNTRACKED)
  - Keyboard shortcuts: `a` (select all), `d` (deselect all), `i` (invert selection)
  - Visual improvements with styled output using lipgloss
  - Selection counter showing selected/total files
  - Helpful tips for grouping related changes

- **Conventional Commits Support**
  - Interactive commit type selection menu
  - 10 predefined types: feat, fix, docs, style, refactor, perf, test, chore, build, ci
  - Custom commit type option
  - Full conventional commit message formatting

- **Commit Message Validation**
  - Title requirements: 10-72 characters, lowercase start, no period at end
  - Warning for titles over 50 characters
  - Optional multi-line description support
  - Interactive commit preview before confirmation
  - Edit capability from preview screen

- **Loop Until Clean Workflow**
  - After commit, checks for remaining changes
  - Option to create another commit on the same branch
  - Shows status: files committed vs. still uncommitted
  - Enables multi-commit workflows without restarting

- **Intelligent Push Support**
  - Summary of all commits created in session
  - Optional push to remote after commits
  - Auto-detection of remote branch existence
  - Uses `-u` flag for first-time branch push
  - Helpful message if user skips push

#### New UI Components
- `internal/ui/branch.go` - Branch management interface
- `internal/ui/commit_type.go` - Commit type selection
- `internal/ui/commit_message.go` - Title and description input with validation
- `internal/ui/confirm.go` - Yes/no confirmation dialogs
- `internal/ui/input.go` - Generic text input for custom types

#### Enhanced Git Operations
- `GetCurrentBranch()` - Detects current branch
- `IsMainBranch()` - Checks if branch is main/master
- `CreateBranch()` - Creates and switches to new branch
- `CheckoutBranch()` - Switches branches
- `CommitWithDescription()` - Commits with multi-line messages
- `HasRemoteBranch()` - Checks if branch exists on remote
- `Push()` - Pushes to remote with optional upstream flag

#### Model Enhancements
- `DisplayType()` method for categorizing changes
- `CommitType` struct for conventional commit types
- `CommitInfo` struct for structured commit data
- Predefined list of conventional commit types

### Changed
- **File Selection Menu**
  - Changed from "Enter to toggle" to "Space to toggle, Enter to confirm"
  - Added categorization by file status
  - Improved visual styling and layout
  - Added cancel detection to distinguish between quit and confirm

- **Main Workflow**
  - Completely redesigned to support multi-step process
  - Added comprehensive error handling
  - Enhanced user feedback with styled messages
  - Added session tracking for commits

### Documentation
- Comprehensive README.md with all features documented
- QUICKSTART.md for new users
- Examples and common scenarios
- Troubleshooting section
- Keyboard shortcuts reference

### Fixed
- Variable name collision in `commit_type.go` (model vs imported package)
- Improved error messages throughout the application
- Better validation feedback to users

## [0.1.0] - Initial Release

### Added
- Basic file change detection using `git status --porcelain`
- Simple file parsing for git changes
- Interactive file selection using Bubble Tea
- Basic categorization by status
- Simple commit workflow
- Git add and commit operations

---

## Future Roadmap

### Planned for 0.3.0
- Configuration file support (`~/.gcm/config.json`)
- Customizable commit types
- Default settings (auto-push, branch prefix, etc.)

### Planned for 0.4.0
- GitHub CLI integration for PR creation
- Issue/ticket number linking
- Co-author support

### Planned for 0.5.0
- Git hooks support
- Pre-commit validation
- Commit message templates
- History search and amend

### Ideas for Future Versions
- AI-powered commit message suggestions
- Jira/Linear integration
- Team conventions enforcement
- Commit statistics and insights

