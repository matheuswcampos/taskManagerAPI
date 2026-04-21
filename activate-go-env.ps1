$root = Split-Path -Parent $MyInvocation.MyCommand.Path
$env:GOPATH = Join-Path $root '.goenv\gopath'
$env:GOMODCACHE = Join-Path $env:GOPATH 'pkg\mod'
$env:GOCACHE = Join-Path $root '.goenv\gocache'

Write-Host "Go env ativado para esta sessao:" -ForegroundColor Cyan
Write-Host "GOPATH=$env:GOPATH"
Write-Host "GOMODCACHE=$env:GOMODCACHE"
Write-Host "GOCACHE=$env:GOCACHE"
