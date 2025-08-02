@echo off

REM This script creates a scheduled task to run the Solo Queue Pop application when you log in.
REM
REM IMPORTANT:
REM 1. This script assumes your project's executable will be located at 'd:\projects\solo-queue-pop\solo-queue-pop.exe'.
REM    You may need to build it first, for example by running 'go build' from the 'd:\projects\solo-queue-pop\solo-queue-pop' directory.
REM 2. Please run this script as an Administrator to ensure it has permission to create scheduled tasks.

set "EXECUTABLE=qpn.exe"
set "WORKING_DIR=%~dp0"
set "TASK_NAME=WoWQueuePopNotification"
set "EXECUTABLE_PATH=%WORKING_DIR%%EXECUTABLE%"

echo Creating Windows Scheduled Task to run Solo Queue Pop on logon...
echo.
echo Task Name: %TASK_NAME%
echo Executable: %EXECUTABLE_PATH%
echo Working Directory: %WORKING_DIR%
echo.

REM Check if the executable exists before creating the task
if not exist "%EXECUTABLE_PATH%" (
    echo ERROR: Executable not found at the specified path.
    echo Please make sure you have built the project and the path is correct.
    pause
    exit /b 1
)

REM Use schtasks to create the task.
REM /sc ONLOGON - Runs when the user logs in.
REM /tr - Specifies the task to run. We use 'start' to ensure the console window is visible.
REM /f  - Overwrites the task if it already exists.
schtasks /create /sc ONLOGON /tn "%TASK_NAME%" /tr "cmd /c 'cd %WORKING_DIR% && %EXECUTABLE_PATH% config.yaml'" /f

if %errorlevel% equ 0 (
    echo.
    echo SUCCESS: Scheduled task '%TASK_NAME%' was created successfully.
    echo It will run the next time you log into Windows.
) else (
    echo.
    echo ERROR: Failed to create the scheduled task.
    echo Please make sure you are running this script as an Administrator.
)

echo.
pause
