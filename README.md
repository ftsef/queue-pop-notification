# PvP Queue Pop Notification

Tired of being tethered to your desk with a wired headset while waiting for a World of Warcraft match? This lightweight application sends a Discord notification the moment your queue pops, so you'll never miss a match again!

As of right now due to the restrictions from the QueueNotify addon you only get notifications for Solo-Shuffle and BG Blitz.
---
I'll probably create a new addon to also receive notifications for 2v2 and 3v3 + some extras like Rating etc.


## Prerequisites

This tool relies on the **QueueNotify** World of Warcraft addon. This addon automatically takes a screenshot (`.tga` file) when a queue pops, which this application detects.

- **Install QueueNotify Addon**: [Download from CurseForge](https://www.curseforge.com/wow/addons/queuenotify)

## Configuration

1.  Create a `config.yaml` file in the same directory as the `solo-queue-pop.exe` executable.
2.  Copy the example below and modify it with your own settings.

-   **`url`**: Your Discord webhook URL. For more information, see the [Discord Webhook API documentation](https://discord.com/developers/docs/resources/webhook).
-   **`user-id`**: (Optional) Your Discord user ID if you want to be mentioned in the notification. Here’s [how to find your user ID](https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID).
-   **`base_path`**: The installation path for your World of Warcraft game.

**Example `config.yaml`:**
```yaml
webhook:
  url: "https://discord.com/api/webhooks/.../..."
  body: |
    {
      "content": "<@user-id> BG-Blitz match found!!",
      "allowed_mentions": {
        "users": [
          "user-id"
        ]
      }
    }
wow:
  base_path: "E:\\games\\World of Warcraft"
```

## Usage

### Build the Application
Navigate to the `solo-queue-pop` subdirectory and run the build command:
```sh
cd solo-queue-pop
go build
```
This will create the `solo-queue-pop.exe` executable.

### Run Manually
To run the application directly, execute the compiled file from your terminal:
```sh
./solo-queue-pop.exe
```

### Run Automatically on Startup (Windows)
A batch script is provided to create a Windows Scheduled Task that runs the application automatically when you log in.

1.  Place the `create_task.bat` script in the same directory as `solo-queue-pop.exe`.
2.  Right-click `create_task.bat` and select **"Run as administrator"**.

The script will handle the rest.

## Alternatives
* [Queue Notify Client](https://github.com/rudikiaz/queue-notify-client)
 
