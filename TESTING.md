# GCM Testing Checklist

Use this checklist to verify all features are working correctly.

## Pre-Test Setup

```bash
# Build the tool
go build -o gcm .

# Or run the install script
./install.sh
```

## Test Scenarios

###  Test 1: Clean Working Tree
**Scenario**: No changes to commit
```bash
# Make sure working tree is clean
git status
# Should show nothing to commit

# Run gcm
./gcm
```
**Expected**: " Working tree clean, nothing to commit"

---

###  Test 2: Branch Management - On Main
**Scenario**: Prevent direct commit to main
```bash
# Switch to main
git checkout main

# Make a change
echo "test" >> test.txt
git add test.txt

# Run gcm
./gcm
```
**Expected**: 
- Error message: "Cannot commit directly to main"
- Prompt to create new branch
- Branch name validation

**Test branch validation:**
- Try: "my branch" → Should reject (spaces)
- Try: "feature" → Should accept but warn
- Try: "feat/test-feature" → Should accept

---

###  Test 3: Branch Management - On Feature Branch
**Scenario**: Option to use current or create new branch
```bash
# Create and switch to feature branch
git checkout -b feat/existing-feature

# Make a change
echo "test2" >> test2.txt
git add test2.txt

# Run gcm
./gcm
```
**Expected**:
- Shows current branch: "feat/existing-feature"
- Asks: "Use current branch? (y/n)"
- Test both options:
  - 'y' → Uses current branch
  - 'n' → Prompts for new branch name

---

###  Test 4: File Selection - Multiple Files
**Scenario**: Select specific files
```bash
# Create multiple changes
echo "file1" >> file1.txt
echo "file2" >> file2.txt
echo "file3" >> file3.txt
git add file1.txt file2.txt file3.txt

# Run gcm
./gcm
```
**Expected**:
- Shows all 3 files
- Categorized by type (ADDED)
- Test keyboard shortcuts:
  - `↑↓` or `jk` → Navigate
  - `SPACE` → Toggle selection
  - `a` → Select all (3 files)
  - `d` → Deselect all (0 files)
  - `i` → Invert selection
  - `q` → Cancel and exit
  - `ENTER` → Confirm selection

---

###  Test 5: File Selection - Mixed Types
**Scenario**: Different change types
```bash
# Modify existing file
echo "modified" >> existing.txt

# Delete a file
git rm old-file.txt

# Add new file
echo "new" >> new-file.txt
git add .

# Run gcm
./gcm
```
**Expected**:
- Files grouped by type:
  - MODIFIED: existing.txt
  - DELETED: old-file.txt
  - ADDED: new-file.txt
- Each category clearly labeled

---

###  Test 6: Commit Type Selection
**Scenario**: Choose commit type
```bash
# Have some changes ready
echo "test" >> test.txt
git add test.txt

# Run gcm and get to type selection
./gcm
```
**Expected**:
- Shows 10 types with descriptions
- Navigate with `↑↓` or `jk`
- Highlight current selection
- Test:
  - Select 'feat' → Should proceed
  - Press 'c' → Should prompt for custom type

---

###  Test 7: Commit Message - Title Validation
**Scenario**: Test all title validations
```bash
# Run gcm and get to title input
./gcm
```
**Test these titles:**
- "short" → Error: too short (min 10)
- "This Starts With Uppercase" → Error: must be lowercase
- "this ends with period." → Error: no period at end
- "this is exactly ten characters long" → Should work
- "this is a very very very very very very very very very very very long title" → Error if >72 chars
- "this is over fifty characters long message here" → Warning but allows (>50 chars)
- "add email validation to registration form" → Should work perfectly

---

###  Test 8: Commit Message - Description
**Scenario**: Optional description
```bash
# Run gcm and get to description
./gcm
```
**Test:**
- Press ENTER on empty → Skip description
- Type text → Add description
- Press CTRL+D → Finish description
- Multi-line description

---

###  Test 9: Commit Preview
**Scenario**: Preview before commit
```bash
# Complete the flow to preview
./gcm
```
**Expected**:
- Shows formatted box with:
  - Type: title
  - Description (if provided)
- Options:
  - 'y' → Confirm and commit
  - 'n' → Cancel
  - 'e' → Edit (go back to title)

**Test all options:**
- Press 'e' → Should go back to title
- Press 'n' → Should cancel
- Press 'y' → Should create commit

---

###  Test 10: Loop - Multiple Commits
**Scenario**: Create multiple commits in one session
```bash
# Create multiple files
echo "file1" >> file1.txt
echo "file2" >> file2.txt
echo "file3" >> file3.txt
git add .

# Run gcm
./gcm
```
**Test:**
1. Select only file1.txt → Commit
2. Should ask: "Want to create another commit?"
3. Press 'y' → Should show remaining files (file2.txt, file3.txt)
4. Select file2.txt → Commit
5. Should ask again
6. Press 'y' → Should show remaining file (file3.txt)
7. Select file3.txt → Commit
8. Should show: "All files committed!"

---

###  Test 11: Loop - Working Tree Clean
**Scenario**: Commit all changes at once
```bash
# Create changes
echo "test" >> test.txt
git add test.txt

# Run gcm
./gcm
```
**Test:**
1. Select all files → Commit
2. Should show: "All files committed!"
3. Should proceed to push prompt

---

###  Test 12: Push - New Branch
**Scenario**: First push of a new branch
```bash
# Create new branch with commits
git checkout -b feat/new-feature
echo "test" >> test.txt
git add test.txt

# Run gcm and create commit
./gcm
```
**Expected**:
- Shows summary of commits
- Asks: "Push to remote?"
- If 'y' → Should use `git push -u origin feat/new-feature`

---

###  Test 13: Push - Existing Branch
**Scenario**: Push to existing remote branch
```bash
# Use branch that exists on remote
git checkout existing-branch
echo "test" >> test.txt
git add test.txt

# Run gcm and create commit
./gcm
```
**Expected**:
- Shows summary of commits
- Asks: "Push to remote?"
- If 'y' → Should use `git push origin existing-branch` (no -u)

---

###  Test 14: Push - Decline
**Scenario**: Choose not to push
```bash
# Create commit
# ... complete flow ...
```
**Expected**:
- At push prompt, press 'n'
- Should show: " You can push later with: git push origin <branch>"
- Should exit cleanly

---

###  Test 15: Cancel at Each Step
**Scenario**: Test cancellation
```bash
# Run gcm and test ESC/q at each step
./gcm
```
**Test pressing ESC or q at:**
1. Branch selection → Should exit
2. File selection → Should exit
3. Type selection → Should exit
4. Title input → Should exit
5. Description input → Should go back (or exit)
6. Preview → Should go back (with 'n')
7. Continue prompt → Should exit

---

### ✅ Test 16: Large Commit Warning
**Scenario**: More than 10 files
```bash
# Create many files
for i in {1..15}; do echo "file$i" >> "file$i.txt"; done
git add .

# Run gcm
./gcm
```
**Expected**:
- Shows warning: "️ Warning: Large commit with 15 files"
- Still allows commit

---

###  Test 17: Session Summary
**Scenario**: Multiple commits summary
```bash
# Create multiple files
echo "f1" >> f1.txt; echo "f2" >> f2.txt; echo "f3" >> f3.txt
git add .

# Run gcm, create 3 commits
./gcm
```
**Expected**:
- After all commits, shows:
  ```
   Commits created in this session:
    1. [feat] first commit
    2. [fix] second commit
    3. [docs] third commit
  ```

---

###  Test 18: Custom Commit Type
**Scenario**: Use custom type
```bash
echo "test" >> test.txt
git add test.txt

# Run gcm
./gcm
```
**Test:**
1. At type selection, press 'c'
2. Enter custom type: "hotfix"
3. Should use "hotfix" as commit type
4. Final commit: "hotfix: <your message>"

---

###  Test 19: Edge Cases
**Scenario**: Test boundary conditions

**Test 1**: Exactly 10 character title
- "0123456789" → Should accept

**Test 2**: Exactly 50 character title
- "this title is exactly fifty characters long here!" → Should work, no warning

**Test 3**: Exactly 51 character title
- "this title is exactly fifty one characters long ok!" → Should work with warning

**Test 4**: Exactly 72 character title
- "this title is exactly seventy two characters long and should work perfectly!" → Should accept

**Test 5**: 73 character title
- "this title is exactly seventy three characters long and should be rejected now!" → Should reject

---

###  Test 20: Keyboard Shortcuts
**Scenario**: Verify all shortcuts work

**File Selection:**
- `j` → Move down
- `k` → Move up
- `↓` → Move down
- `↑` → Move up
- `a` → Select all
- `d` → Deselect all
- `i` → Invert selection
- `SPACE` → Toggle current
- `ENTER` → Confirm
- `q` → Cancel
- `ESC` → Cancel
- `CTRL+C` → Force exit

**Confirmation:**
- `y` → Yes
- `n` → No

**Type Selection:**
- Numbers → Select type
- `c` → Custom
- `CTRL+C` → Exit

---

## Post-Testing Verification

### Check Git History
```bash
git log --oneline -n 5
```
Should show your conventional commits.

### Check Branch
```bash
git branch
```
Should show your feature branch.

### Check Remote
```bash
git log origin/your-branch --oneline -n 5
```
Should show commits if you pushed.

---

## Issues to Report

If you find any issues during testing, note:
1. What step you were on
2. What you expected
3. What actually happened
4. Error messages (if any)
5. How to reproduce

---

## Success Criteria

All tests should:
-  Run without crashes
-  Show appropriate messages
-  Validate input correctly
-  Create proper commits
-  Handle errors gracefully
-  Allow cancellation
-  Display styled output

---

## Quick Smoke Test

For a fast verification:
```bash
# 1. Build
go build -o gcm .

# 2. Create test changes
echo "test1" >> test1.txt
echo "test2" >> test2.txt
git add .

# 3. Run and complete full workflow
./gcm

# 4. Verify commit
git log -1 --pretty=format:"%s"
```

Should show a properly formatted conventional commit message!

---

**Happy Testing! **

