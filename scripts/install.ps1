# Gitty Installation Script for Windows
param(
    [switch]$Help,
    [switch]$Version
)

# Configuration
$Repo = "Sunchiii/gitty"
$InstallDir = "$env:LOCALAPPDATA\gitty"
$BinaryName = "gitty.exe"

# Get the latest version
function Get-LatestVersion {
    $response = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases/latest"
    return $response.tag_name
}

# Get the appropriate asset URL for Windows
function Get-AssetUrl {
    param($Version)
    $assetName = "gitty-$Version-windows-amd64.exe"
    return "https://github.com/$Repo/releases/download/$Version/$assetName"
}

# Download and install
function Install-Gitty {
    param($Version, $DownloadUrl)
    
    Write-Host "Downloading Gitty $Version..." -ForegroundColor Blue
    
    # Create installation directory
    if (!(Test-Path $InstallDir)) {
        New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
    }
    
    $tempFile = Join-Path $env:TEMP "gitty-temp.exe"
    $installPath = Join-Path $InstallDir $BinaryName
    
    try {
        # Download the binary
        Invoke-WebRequest -Uri $DownloadUrl -OutFile $tempFile
        
        # Move to installation directory
        Move-Item -Path $tempFile -Destination $installPath -Force
        
        # Add to PATH if not already there
        $currentPath = [Environment]::GetEnvironmentVariable("PATH", "User")
        if ($currentPath -notlike "*$InstallDir*") {
            $newPath = "$currentPath;$InstallDir"
            [Environment]::SetEnvironmentVariable("PATH", $newPath, "User")
            Write-Host "Added $InstallDir to PATH" -ForegroundColor Yellow
        }
        
        Write-Host "Gitty $Version installed successfully!" -ForegroundColor Green
    }
    catch {
        Write-Host "Failed to install Gitty: $($_.Exception.Message)" -ForegroundColor Red
        exit 1
    }
}

# Verify installation
function Test-Installation {
    $gittyPath = Join-Path $InstallDir $BinaryName
    if (Test-Path $gittyPath) {
        Write-Host "Installation verified!" -ForegroundColor Green
        Write-Host "Run 'gitty --help' to get started" -ForegroundColor Blue
    } else {
        Write-Host "Installation failed" -ForegroundColor Red
        exit 1
    }
}

# Main installation function
function Main {
    Write-Host "Installing Gitty..." -ForegroundColor Blue
    
    # Check if gitty is already installed
    $gittyPath = Join-Path $InstallDir $BinaryName
    if (Test-Path $gittyPath) {
        Write-Host "Gitty is already installed." -ForegroundColor Yellow
        $response = Read-Host "Do you want to update it? (y/N)"
        if ($response -notmatch "^[Yy]$") {
            Write-Host "Installation cancelled."
            exit 0
        }
    }
    
    # Get latest version
    Write-Host "Fetching latest version..." -ForegroundColor Blue
    $latestVersion = Get-LatestVersion
    
    if (!$latestVersion) {
        Write-Host "Failed to get latest version" -ForegroundColor Red
        exit 1
    }
    
    Write-Host "Latest version: $latestVersion" -ForegroundColor Green
    
    # Get download URL
    $downloadUrl = Get-AssetUrl $latestVersion
    
    if (!$downloadUrl) {
        Write-Host "Failed to get download URL" -ForegroundColor Red
        exit 1
    }
    
    # Install
    Install-Gitty $latestVersion $downloadUrl
    
    # Verify
    Test-Installation
}

# Handle command line arguments
if ($Help) {
    Write-Host "Gitty Installation Script for Windows"
    Write-Host ""
    Write-Host "Usage:"
    Write-Host "  .\install.ps1"
    Write-Host "  .\install.ps1 --version"
    Write-Host ""
    Write-Host "Options:"
    Write-Host "  --version    Show latest version"
    Write-Host "  --help       Show this help"
    exit 0
}

if ($Version) {
    Get-LatestVersion
    exit 0
}

# Run main installation
Main 