# GCM Workflow Diagram

```
┌─────────────────────────────────────────────────────────────────────┐
│                         START: gcm                                   │
└─────────────────────────────────────────────────────────────────────┘
                                  │
                                  ▼
                    ┌─────────────────────────┐
                    │  Check for Changes      │
                    │  (git status)           │
                    └─────────────────────────┘
                                  │
                     ┌────────────┴────────────┐
                     │                         │
                     ▼                         ▼
         ┌─────────────────────┐   ┌─────────────────────┐
         │  No Changes         │   │  Changes Detected   │
         │  "Working tree      │   │                     │
         │   clean"            │   │  Continue...        │
         └─────────────────────┘   └─────────────────────┘
                     │                         │
                     ▼                         ▼
                  [EXIT]         ┌─────────────────────────┐
                                 │  Get Current Branch     │
                                 └─────────────────────────┘
                                              │
                              ┌───────────────┴───────────────┐
                              │                               │
                              ▼                               ▼
                  ┌──────────────────────┐      ┌──────────────────────┐
                  │  On main/master      │      │  On other branch     │
                  │                      │      │                      │
                  │   ️ Cannot commit    │      │  Use current or      │
                  │  Must create branch  │      │  create new? (y/n)   │
                  └──────────────────────┘      └──────────────────────┘
                              │                               │
                              │         ┌─────────────────────┤
                              │         │                     │
                              │         ▼                     ▼
                              │  ┌─────────────┐    ┌──────────────┐
                              │  │  'y' Use    │    │  'n' Create  │
                              │  │  current    │    │  new branch  │
                              │  └─────────────┘    └──────────────┘
                              │         │                     │
                              ▼         ▼                     ▼
                        ┌──────────────────────────────────────┐
                        │  Branch Name Input & Validation      │
                        │  - No spaces, max 50 chars           │
                        │  - Format: type/description          │
                        └──────────────────────────────────────┘
                                        │
                              ┌─────────┴─────────┐
                              │                   │
                              ▼                   ▼
                      ┌──────────────┐    ┌─────────────┐
                      │  Valid       │    │  Invalid    │
                      │  Continue    │    │  Try again  │
                      └──────────────┘    └─────────────┘
                              │                   │
                              │                   └──────┐
                              ▼                          │
          ┌────────────────────────────────────────┐     │
          │  Create/Switch to Branch               │◄────┘
          └────────────────────────────────────────┘
                              │
                              ▼
          ╔════════════════════════════════════════╗
          ║         COMMIT LOOP (START)            ║
          ╚════════════════════════════════════════╝
                              │
                              ▼
          ┌────────────────────────────────────────┐
          │  File Selection UI                     │
          │  ┌──────────────────────────────────┐  │
          │  │ MODIFIED:                        │  │
          │  │  [x] src/file1.go                │  │
          │  │  [ ] src/file2.go                │  │
          │  │                                  │  │
          │  │ ADDED:                           │  │
          │  │  [x] src/new.go                  │  │
          │  │                                  │  │
          │  │ Shortcuts: a/d/i/SPACE/ENTER     │  │
          │  └──────────────────────────────────┘  │
          └────────────────────────────────────────┘
                              │
                    ┌─────────┴─────────┐
                    │                   │
                    ▼                   ▼
            ┌──────────────┐    ┌─────────────┐
            │  Files       │    │  No files   │
            │  selected    │    │  selected   │
            └──────────────┘    └─────────────┘
                    │                   │
                    │                   ▼
                    │               [EXIT]
                    │
                    ▼
        ┌───────────────────────────┐
        │  Commit Type Selection    │
        │  ┌────────────────────┐   │
        │  │ > 1. feat          │   │
        │  │   2. fix           │   │
        │  │   3. docs          │   │
        │  │   ... (10 types)   │   │
        │  │   [c] custom       │   │
        │  └────────────────────┘   │
        └───────────────────────────┘
                    │
                    ▼
        ┌───────────────────────────┐
        │  Commit Message - Title   │
        │                           │
        │  > add user auth_         │
        │                           │
        │  Rules:                   │
        │  - 10-72 chars            │
        │  - lowercase start        │
        │  - no period at end       │
        └───────────────────────────┘
                    │
                    ▼
        ┌───────────────────────────┐
        │  Description (Optional)   │
        │                           │
        │  Press ENTER to skip or   │
        │  type description...      │
        │  Press CTRL+D when done   │
        └───────────────────────────┘
                    │
                    ▼
        ┌───────────────────────────┐
        │  Commit Preview           │
        │  ┌─────────────────────┐  │
        │  │ feat: add user auth │  │
        │  │                     │  │
        │  │ Implements JWT...   │  │
        │  │                     │  │
        │  │ Files (2):          │  │
        │  │ - src/auth.go       │  │
        │  │ - src/jwt.go        │  │
        │  └─────────────────────┘  │
        │                           │
        │  Confirm? (y/n/e)         │
        └───────────────────────────┘
                    │
        ┌───────────┼───────────┐
        │           │           │
        ▼           ▼           ▼
    ┌─────┐   ┌─────┐    ┌─────┐
    │  y  │   │  e  │    │  n  │
    └─────┘   └─────┘    └─────┘
        │         │          │
        │         │          ▼
        │         │      [EXIT]
        │         │
        │         └──────────┐
        │                    │
        ▼                    ▼
    ┌─────────────────┐  [Go back to title]
    │  Git Add        │
    │  Git Commit     │
    └─────────────────┘
            │
            ▼
    ┌─────────────────┐
    │  ✓ Commit       │
    │  Created!       │
    └─────────────────┘
            │
            ▼
    ┌─────────────────────────┐
    │  Check for more changes │
    └─────────────────────────┘
            │
    ┌───────┴───────┐
    │               │
    ▼               ▼
┌──────────────┐  ┌──────────────┐
│  More files  │  │  All done    │
│  uncommitted │  │              │
└──────────────┘  └──────────────┘
    │               │
    ▼               ▼
┌──────────────────────────┐  ╔════════════════════╗
│  Create another commit?  │  ║  COMMIT LOOP (END) ║
│  (y/n)                   │  ╚════════════════════╝
└──────────────────────────┘            │
    │                                   │
┌───┴────┐                              │
│        │                              │
▼        ▼                              │
y        n                              │
│        │                              │
│        └──────────────────────────────┘
│                                       │
└───────────────────┐                   │
                    │                   │
                    └──────┐            │
                           │            │
                           ▼            ▼
                    [Loop back]   ┌─────────────────────┐
                                  │  Session Summary    │
                                  │  ┌───────────────┐  │
                                  │  │ Commits:      │  │
                                  │  │ 1. [feat]...  │  │
                                  │  │ 2. [fix]...   │  │
                                  │  └───────────────┘  │
                                  └─────────────────────┘
                                            │
                                            ▼
                                  ┌─────────────────────┐
                                  │  Push to remote?    │
                                  │  (y/n)              │
                                  └─────────────────────┘
                                            │
                                    ┌───────┴───────┐
                                    │               │
                                    ▼               ▼
                            ┌──────────────┐  ┌──────────────┐
                            │  y - Push    │  │  n - Skip    │
                            └──────────────┘  └──────────────┘
                                    │               │
                                    ▼               ▼
                            ┌──────────────┐  ┌──────────────┐
                            │ Check remote │  │ Show help    │
                            │ branch       │  │ message      │
                            └──────────────┘  └──────────────┘
                                    │               │
                        ┌───────────┴────────┐      │
                        │                    │      │
                        ▼                    ▼      │
                ┌──────────────┐    ┌──────────────┐│
                │ New branch   │    │ Exists on    ││
                │ (use -u)     │    │ remote       ││
                └──────────────┘    └──────────────┘│
                        │                    │      │
                        ▼                    ▼      │
                ┌──────────────────────────────┐    │
                │  git push [-u] origin <br>   │    │
                └──────────────────────────────┘    │
                        │                           │
                        ▼                           │
                ┌──────────────────────────────┐    │
                │  ✓ Pushed to remote          │    │
                └──────────────────────────────┘    │
                        │                           │
                        └───────────────────────────┘
                                    │
                                    ▼
                        ┌───────────────────────┐
                        │    Done!              │
                        └───────────────────────┘
                                    │
                                    ▼
                                 [EXIT]


═══════════════════════════════════════════════════════════════════

LEGEND:

  ┌─────┐
  │     │   Step/Action
  └─────┘

  ╔═════╗
  ║     ║   Important Section
  ╚═════╝

      ▼       Flow Direction

      │       Continues
      ├       Branches

  [EXIT]      Program ends

═══════════════════════════════════════════════════════════════════

KEY DECISION POINTS:

1. Branch Check: main/master → Force new branch
                 other → Option to use or create

2. File Selection: Selected → Continue
                   None → Exit

3. Commit Preview: y → Commit
                   e → Edit
                   n → Cancel

4. Continue Loop: y → Create another commit
                  n → Show summary

5. Push: y → Push to remote
         n → Skip with helpful message

═══════════════════════════════════════════════════════════════════

KEYBOARD SHORTCUTS AT EACH STEP:

File Selection:
  ↑↓ / jk    Navigate
  SPACE      Toggle
  a          Select all
  d          Deselect all
  i          Invert
  ENTER      Confirm
  q/ESC      Cancel

Type Selection:
  ↑↓ / jk    Navigate
  ENTER      Select
  c          Custom type
  q/ESC      Cancel

All Steps:
  CTRL+C     Emergency exit
  ESC        Cancel/Go back

═══════════════════════════════════════════════════════════════════
```

