@echo off
REM Build script for go-passman on Windows

setlocal enabledelayedexpansion

set BINARY_NAME=go-passman
set OUTPUT_DIR=dist

if not exist %OUTPUT_DIR% mkdir %OUTPUT_DIR%

if "%1"=="" goto build_current
if "%1"=="all" goto build_all
if "%1"=="release" goto build_release
if "%1"=="clean" goto clean
if "%1"=="help" goto help
goto unknown

:build_current
echo Building for current platform (slim)...
go build -ldflags="-s -w" -trimpath -o %OUTPUT_DIR%\%BINARY_NAME%.exe
echo Build complete: %OUTPUT_DIR%\%BINARY_NAME%.exe
goto end

:build_all
echo Building for Windows (amd64)...
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w" -trimpath -o %OUTPUT_DIR%\%BINARY_NAME%-windows-amd64.exe
echo.

echo Building for Linux (amd64)...
set GOOS=linux
set GOARCH=amd64
go build -ldflags="-s -w" -trimpath -o %OUTPUT_DIR%\%BINARY_NAME%-linux-amd64
echo.

echo Building for macOS (Intel)...
set GOOS=darwin
set GOARCH=amd64
go build -ldflags="-s -w" -trimpath -o %OUTPUT_DIR%\%BINARY_NAME%-darwin-amd64
echo.

echo Building for macOS (Apple Silicon)...
set GOOS=darwin
set GOARCH=arm64
go build -ldflags="-s -w" -trimpath -o %OUTPUT_DIR%\%BINARY_NAME%-darwin-arm64
echo.

echo All builds complete!
goto end

:build_release
set VERSION=%~2
if "%VERSION%"=="" set VERSION=0.3.0
set RELEASE_DIR=%OUTPUT_DIR%\release-%VERSION%
if not exist %RELEASE_DIR% mkdir %RELEASE_DIR%
echo Building release %VERSION% for GitHub...
echo.

set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w" -trimpath -o %RELEASE_DIR%\%BINARY_NAME%-%VERSION%-windows-amd64.exe
echo Built windows/amd64

set GOOS=windows
set GOARCH=arm64
go build -ldflags="-s -w" -trimpath -o %RELEASE_DIR%\%BINARY_NAME%-%VERSION%-windows-arm64.exe
echo Built windows/arm64

set GOOS=linux
set GOARCH=amd64
go build -ldflags="-s -w" -trimpath -o %RELEASE_DIR%\%BINARY_NAME%-%VERSION%-linux-amd64
echo Built linux/amd64

set GOOS=linux
set GOARCH=arm64
go build -ldflags="-s -w" -trimpath -o %RELEASE_DIR%\%BINARY_NAME%-%VERSION%-linux-arm64
echo Built linux/arm64

set GOOS=linux
set GOARCH=386
go build -ldflags="-s -w" -trimpath -o %RELEASE_DIR%\%BINARY_NAME%-%VERSION%-linux-386
echo Built linux/386

set GOOS=darwin
set GOARCH=amd64
go build -ldflags="-s -w" -trimpath -o %RELEASE_DIR%\%BINARY_NAME%-%VERSION%-darwin-amd64
echo Built darwin/amd64 (Intel Mac)

set GOOS=darwin
set GOARCH=arm64
go build -ldflags="-s -w" -trimpath -o %RELEASE_DIR%\%BINARY_NAME%-%VERSION%-darwin-arm64
echo Built darwin/arm64 (Apple Silicon)

echo.
echo Creating archives (binary + README)...
set STAGING=%OUTPUT_DIR%\.stage
if not exist %STAGING% mkdir %STAGING%
copy /Y docs\RELEASE_README.txt %STAGING%\README.txt >nul

copy /Y %RELEASE_DIR%\%BINARY_NAME%-%VERSION%-windows-amd64.exe %STAGING%\go-passman.exe >nul
powershell -NoProfile -Command "Compress-Archive -Path '%STAGING%\go-passman.exe','%STAGING%\README.txt' -DestinationPath '%RELEASE_DIR%\%BINARY_NAME%-%VERSION%-windows-amd64.zip' -Force"
del %STAGING%\go-passman.exe

copy /Y %RELEASE_DIR%\%BINARY_NAME%-%VERSION%-windows-arm64.exe %STAGING%\go-passman.exe >nul
powershell -NoProfile -Command "Compress-Archive -Path '%STAGING%\go-passman.exe','%STAGING%\README.txt' -DestinationPath '%RELEASE_DIR%\%BINARY_NAME%-%VERSION%-windows-arm64.zip' -Force"
del %STAGING%\go-passman.exe

copy /Y %RELEASE_DIR%\%BINARY_NAME%-%VERSION%-linux-amd64 %STAGING%\go-passman >nul
powershell -NoProfile -Command "Compress-Archive -Path '%STAGING%\go-passman','%STAGING%\README.txt' -DestinationPath '%RELEASE_DIR%\%BINARY_NAME%-%VERSION%-linux-amd64.zip' -Force"
del %STAGING%\go-passman

copy /Y %RELEASE_DIR%\%BINARY_NAME%-%VERSION%-linux-arm64 %STAGING%\go-passman >nul
powershell -NoProfile -Command "Compress-Archive -Path '%STAGING%\go-passman','%STAGING%\README.txt' -DestinationPath '%RELEASE_DIR%\%BINARY_NAME%-%VERSION%-linux-arm64.zip' -Force"
del %STAGING%\go-passman

copy /Y %RELEASE_DIR%\%BINARY_NAME%-%VERSION%-linux-386 %STAGING%\go-passman >nul
powershell -NoProfile -Command "Compress-Archive -Path '%STAGING%\go-passman','%STAGING%\README.txt' -DestinationPath '%RELEASE_DIR%\%BINARY_NAME%-%VERSION%-linux-386.zip' -Force"
del %STAGING%\go-passman

copy /Y %RELEASE_DIR%\%BINARY_NAME%-%VERSION%-darwin-amd64 %STAGING%\go-passman >nul
powershell -NoProfile -Command "Compress-Archive -Path '%STAGING%\go-passman','%STAGING%\README.txt' -DestinationPath '%RELEASE_DIR%\%BINARY_NAME%-%VERSION%-darwin-amd64.zip' -Force"
del %STAGING%\go-passman

copy /Y %RELEASE_DIR%\%BINARY_NAME%-%VERSION%-darwin-arm64 %STAGING%\go-passman >nul
powershell -NoProfile -Command "Compress-Archive -Path '%STAGING%\go-passman','%STAGING%\README.txt' -DestinationPath '%RELEASE_DIR%\%BINARY_NAME%-%VERSION%-darwin-arm64.zip' -Force"
del %STAGING%\go-passman

rmdir /s /q %STAGING% 2>nul
echo.
echo Release archives: %RELEASE_DIR%\
echo Each archive contains the app (simple name) + README.txt
echo Upload the .zip files to GitHub Releases.
goto end

:clean
echo Cleaning...
if exist %OUTPUT_DIR% rmdir /s /q %OUTPUT_DIR%
echo Clean complete
goto end

:help
echo Usage: build.bat [target]
echo.
echo Targets:
echo   (none)    - Build for current platform
echo   all       - Build for all platforms
echo   release   - Build for GitHub release (all platforms, version in name)
echo   release 0.3.0  - Same with explicit version
echo   clean     - Clean build artifacts
echo   help      - Show this help message
goto end

:unknown
echo Unknown target: %1
call :help
exit /b 1

:end
endlocal
