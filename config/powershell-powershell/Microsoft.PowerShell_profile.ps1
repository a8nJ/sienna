# this file goes in C:\Users\Steven\Documents\PowerShell
# make sure to unblock file if downloading
$env:EDITOR = 'gvim.exe'

$env:LESS = -join @(
   # Quit if entire file fits on first screen.
   'F'
   # Output "raw" control characters.
   'R'
   # Don't use termcap init/deinit strings.
   'X'
   # Ignore case in searches that do not contain uppercase.
   'i'
)

$env:PATH = @(
   'C:\Users\Steven\go\bin'
   'C:\dart-sdk\bin'
   'C:\go\bin'
   'C:\ldc2\bin'
   'C:\nim\bin'
   'C:\path'
   'C:\php'
   'C:\python'
   'C:\python\Scripts'
   'C:\rubyinstaller\bin'
   'C:\sienna\msys2\mingw64\bin'
   'C:\sienna\msys2\usr\bin'
   'C:\sienna\rust\bin'
   'C:\sienna\vim'
) -join ';'

$env:RIPGREP_CONFIG_PATH = $env:USERPROFILE + '\ripgrep.txt'
$env:UMBER = 'D:\Git\umber\umber.json'
$env:WINTER = 'D:\Music\Backblaze\winter.db'

Set-PSReadLineKeyHandler Ctrl+UpArrow {
   Set-Location ..
   [Microsoft.PowerShell.PSConsoleReadLine]::InvokePrompt()
}

[Console]::OutputEncoding = [System.Text.UTF8Encoding]::new()
