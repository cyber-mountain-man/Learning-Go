# Init-GoModules.ps1
# 🔧 Automatically initialize go.mod in folders with main.go but no go.mod

$root = Get-Location

Get-ChildItem -Recurse -Directory | ForEach-Object {
    $mainGo = Join-Path $_.FullName "main.go"
    $modFile = Join-Path $_.FullName "go.mod"

    if ((Test-Path $mainGo) -and (-not (Test-Path $modFile))) {
        Write-Host "📁 Initializing go.mod in:" $_.FullName -ForegroundColor Cyan
        Set-Location $_.FullName

        # Generate a module path based on folder name
        $folderName = Split-Path -Leaf $_.FullName
        $moduleName = "example.com/$folderName"

        go mod init $moduleName
        go mod tidy

        Set-Location $root
    }
}

Write-Host "`n✅ Done initializing go.mod files!" -ForegroundColor Green
