# Gitty - Git Policy CLI

A Git workflow enforcement tool built with Cobra that helps teams maintain consistent branching strategies and avoid common Git mistakes.

## 🎯 Purpose

This CLI tool enforces Git flow-like policies to prevent common team workflow issues:
- **Forgetting to pull** before starting work
- **Forgetting to rebase** with the latest changes
- **Direct commits to protected branches** (main, uat, develop)
- **Inconsistent branch naming**
- **Merge conflicts** due to outdated branches

## 🏗️ Branch Strategy

```
main (production) ← hotfix branches
├── uat (customer testing)
└── develop (team development) ← feature branches
    └── feature/* (new features)
    └── hotfix/* (production fixes)
```

## 🚀 Installation

### Quick Install (Recommended)

**macOS/Linux:**
```bash
# Install with one command
curl -fsSL https://raw.githubusercontent.com/Sunchiii/gitty/main/scripts/install.sh | bash

# Or download and run
curl -fsSL https://raw.githubusercontent.com/Sunchiii/gitty/main/install.sh | bash
```

**Windows (PowerShell):**
```powershell
# Download and run installation script
Invoke-Expression (Invoke-WebRequest -Uri "https://raw.githubusercontent.com/Sunchiii/gitty/main/scripts/install.ps1").Content
```

### Manual Installation
```bash
# Clone the repository
git clone <your-repo>
cd gitty

# Build the application
make build
# or
go build -o gitty

# Build with version information
make build-version
# or
./scripts/build.sh

# Install globally (optional)
make install
# or
sudo cp gitty /usr/local/bin/
```

### Uninstall

**macOS/Linux:**
```bash
# Remove gitty
curl -fsSL https://raw.githubusercontent.com/Sunchiii/gitty/main/scripts/uninstall.sh | bash
```

**Windows:**
```powershell
# Remove gitty from Windows
Remove-Item -Path "$env:LOCALAPPDATA\gitty" -Recurse -Force
```

## 📦 Building & Releasing

### Building with Version Info
```bash
# Build with current git commit and date
make build-version

# Build with specific version
VERSION=v1.0.0 make build-version
```

### Creating Releases
```bash
# Create a release for version v1.0.0
make release VERSION=v1.0.0

# This will:
# - Build for multiple platforms (Linux, macOS, Windows)
# - Create checksums for verification
# - Generate release notes
# - Prepare files for GitHub release
```

## 📦 Installation Features

### One-Line Installation
Users can install Gitty with a single command:
```bash
curl -fsSL https://raw.githubusercontent.com/Sunchiii/gitty/main/scripts/install.sh | bash
```

### Automatic Updates
Once installed, users can update to the latest version:
```bash
gitty update
```

### Platform Support
- **macOS**: Intel and Apple Silicon
- **Linux**: AMD64 and ARM64
- **Windows**: AMD64 (with .exe extension)

### Installation Locations
- **Global**: `/usr/local/bin/gitty` (default)
- **User**: `~/bin/gitty` (if `/usr/local/bin` is not writable)

## 🔄 Typical Workflow

### For New Features

1. **Start feature branch:**
   ```bash
   gitty start feature new-feature-name
   ```

2. **Make changes and commit:**
   ```bash
   git add .
   git commit -m "Add new feature"
   ```

3. **Sync with develop:**
   ```bash
   gitty sync
   ```

4. **Push and create PR:**
   ```bash
   git push origin feature/new-feature-name
   ```

5. **After merge, finish the branch:**
   ```bash
   gitty finish feature new-feature-name
   ```

### For Hotfixes

1. **Start hotfix branch:**
   ```bash
   gitty start hotfix critical-fix
   ```

2. **Make changes and commit:**
   ```bash
   git add .
   git commit -m "Fix critical issue"
   ```

3. **Finish hotfix:**
   ```bash
   gitty finish hotfix critical-fix
   ```

## 📋 Commands

### Core Commands

#### `gitty start <type> <name>`
Start a new feature or hotfix branch.

```bash
# Start a new feature
gitty start feature user-authentication

# Start a hotfix
gitty start hotfix critical-bug-fix
```

**What it does:**
- Validates branch type (feature/hotfix)
- Checks if branch already exists
- Switches to appropriate base branch (develop/main)
- Pulls latest changes
- Creates new branch

#### `gitty finish <type> <name>`
Finish a feature or hotfix branch.

```bash
# Finish a feature
gitty finish feature user-authentication

# Finish a hotfix
gitty finish hotfix critical-bug-fix
```

**What it does:**
- Validates branch exists
- Checks for uncommitted changes
- Switches to target branch (develop/main)
- Merges with --no-ff flag
- Deletes the feature/hotfix branch

#### `gitty sync`
Sync current branch with develop.

```bash
gitty sync
```

**What it does:**
- Checks for uncommitted changes
- Fetches latest from remote
- Rebase current branch with develop

### Validation Commands

#### `gitty check`
Check branch rules before push.

```bash
gitty check
```

**What it does:**
- Prevents direct pushes to protected branches
- Shows current status
- Shows commits ahead of develop

#### `gitty validate`
Validate current branch follows naming conventions.

```bash
gitty validate
```

**What it does:**
- Checks branch naming (feature/*, hotfix/*)
- Validates protected branches
- Shows if branch is up to date with develop

#### `gitty conflict`
Check for potential merge conflicts.

```bash
gitty conflict
```

**What it does:**
- Fetches latest changes
- Checks for potential conflicts with develop

### Utility Commands

#### `gitty status`
Show current branch status and pending changes.

```bash
gitty status
```

#### `gitty cleanup`
Clean up merged branches and sync with remote.

```bash
gitty cleanup
```

#### `gitty protect`
Install Git hook to protect critical branches.

```bash
gitty protect
```

**What it does:**
- Creates pre-push hook
- Prevents direct pushes to main/uat/develop

#### `gitty team`
Show team workflow information.

```bash
gitty team
```

#### `gitty update`
Check for updates and update gitty to the latest version.

```bash
gitty update
```

**What it does:**
- Checks for new versions on GitHub
- Downloads and installs the latest version
- Creates backup before updating
- Supports multiple platforms

#### `gitty version`
Show gitty version information.

```bash
gitty version
```

**What it shows:**
- Current version
- Build date
- Git commit hash
- Go version
- Platform information

## 🛡️ Protection Features

### Branch Protection
- **main**: Production branch (protected)
- **uat**: Customer testing branch (protected)
- **develop**: Team development branch (protected)

### Automatic Checks
- Prevents direct commits to protected branches
- Validates branch naming conventions
- Checks for uncommitted changes before operations
- Warns about potential conflicts

### Git Hooks
The `protect` command installs a pre-push hook that prevents direct pushes to protected branches.

## 🎯 Benefits for Your Team

### Solves Common Problems
1. **"Forgot to pull"** - `start` command automatically pulls latest changes
2. **"Forgot to rebase"** - `sync` command handles rebasing
3. **"Direct commits to main"** - Protection prevents this
4. **"Merge conflicts"** - `conflict` command detects issues early
5. **"Inconsistent naming"** - `validate` command enforces conventions

### Team Workflow Improvements
- **Consistent branch naming** across the team
- **Automated safety checks** before operations
- **Clear workflow guidance** with helpful messages
- **Conflict prevention** through early detection
- **Branch cleanup** to keep repository tidy


## 📁 Project Structure
```
gitty/
├── cmd/
│   ├── root.go      # Main CLI setup and command registration
│   ├── start.go     # Start feature/hotfix branches
│   ├── finish.go    # Finish feature/hotfix branches
│   ├── sync.go      # Sync with develop branch
│   ├── check.go     # Validate before push
│   ├── protect.go   # Install Git hooks
│   ├── status.go    # Show branch status
│   ├── cleanup.go   # Clean up merged branches
│   ├── validate.go  # Validate branch naming
│   ├── conflict.go  # Check for conflicts
│   ├── team.go      # Show workflow guide
│   ├── utils.go     # Shared utility functions
│   └── config.go    # Configuration management
├── hook/
│   └── pre-push/    # Sample Git hooks
├── main.go          # Application entry point
├── go.mod           # Go module definition
├── Makefile         # Build and install commands
└── README.md        # This file
```

## 🔧 Configuration

The tool works with standard Git repositories and doesn't require additional configuration. However, you can customize the protected branches by modifying the source code.

## 🚨 Error Handling

The tool provides clear error messages and prevents dangerous operations:
- Validates all inputs before execution
- Checks for uncommitted changes
- Prevents operations on non-existent branches
- Shows helpful guidance messages

## 📝 Contributing

1. Fork the repository
2. Create a feature branch: `gitty start feature your-feature`
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License. 