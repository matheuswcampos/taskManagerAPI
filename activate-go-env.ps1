$root = Split-Path -Parent $MyInvocation.MyCommand.Path
$env:GOPATH = Join-Path $root '.goenv\gopath'
$env:GOMODCACHE = Join-Path $env:GOPATH 'pkg\mod'
$env:GOCACHE = Join-Path $root '.goenv\gocache'
$envFile = Join-Path $root '.env'

if (Test-Path $envFile) {
  Get-Content $envFile | ForEach-Object {
    $line = $_.Trim()
    if ($line -eq '' -or $line.StartsWith('#')) {
      return
    }

    $parts = $line -split '=', 2
    if ($parts.Count -eq 2) {
      $name = $parts[0].Trim()
      $value = $parts[1]
      Set-Item -Path ("Env:{0}" -f $name) -Value $value
    }
  }
}

Write-Host "Go env ativado para esta sessao:" -ForegroundColor Cyan
Write-Host "GOPATH=$env:GOPATH"
Write-Host "GOMODCACHE=$env:GOMODCACHE"
Write-Host "GOCACHE=$env:GOCACHE"
if (Test-Path $envFile) {
  Write-Host ".env carregado: $envFile"
} else {
  Write-Host ".env nao encontrado; usando apenas variaveis ja definidas na sessao." -ForegroundColor Yellow
}
