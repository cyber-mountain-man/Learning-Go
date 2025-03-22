# Rename-Folders-And-Files.ps1

$basePath = "C:\Users\Guillermo Morrison\src\golang-book"

# Function to sanitize names by replacing spaces and symbols
function Get-SafeName {
    param ([string]$name)
    return ($name -replace '[ &''":]', '-') -replace '-{2,}', '-' -replace '^-|-$', ''
}

# Step 1: Rename folders (deepest first to avoid path conflicts)
Get-ChildItem -Path $basePath -Recurse -Directory |
    Sort-Object FullName -Descending |
    ForEach-Object {
        $oldPath = $_.FullName
        $parentPath = Split-Path $oldPath
        $newName = Get-SafeName -name $_.Name
        $newPath = Join-Path $parentPath $newName

        if ($oldPath -ne $newPath) {
            Write-Host "📁 Renaming folder:`n  $oldPath`n  → $newPath"
            Rename-Item -Path $oldPath -NewName $newName
        }
    }

# Step 2: Rename files (can go in any order)
Get-ChildItem -Path $basePath -Recurse -File |
    ForEach-Object {
        $oldPath = $_.FullName
        $parentPath = Split-Path $oldPath
        $newName = Get-SafeName -name $_.Name
        $newPath = Join-Path $parentPath $newName

        if ($oldPath -ne $newPath) {
            Write-Host "📄 Renaming file:`n  $oldPath`n  → $newPath"
            Rename-Item -Path $oldPath -NewName $newName
        }
    }
