# PvP Queue Pop Notification

Tired of being tethered to your desk with a wired headset while waiting for a World of Warcraft match? This lightweight application sends a Discord notification the moment your queue pops, so you'll never miss a match again!

As of right now due to the restrictions from the QueueNotify addon you only get notifications for Solo-Shuffle and BG Blitz.

## Features

- [x] notifications for Solo-Shuffle and BG Blitz
- [ ] companion addon which provides further ingame information like rating, 2v2 and 3v3 notifications
- [ ] 2v2 and 3v3 notifications
- [ ] rating information

## Prerequisites

This tool relies on the **QueueNotify** World of Warcraft addon. This addon automatically takes a screenshot (`.tga` file) when a queue pops, which this application detects.

- **Install QueueNotify Addon**: [Download from CurseForge](https://www.curseforge.com/wow/addons/queuenotify)

## Configuration

1.  Create a `config.yaml` file in the same directory as the `qpn.exe` executable.
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
Navigate to the `cmd\qpn` subdirectory and run the build command:
```sh
cd cmd\qpn
go build .
```
This will create the `qpn.exe` executable.

### Start the app
You have two options. Click on the executable which assumes you have a `config.yaml` in the same directory.

Alternatively run start the app manually by executing the compiled file from your terminal:
```
# windows
qpn.exe 

# macos
./qpn-macos-intel 
./qpn-macos-arm64

# linux
./qpn-linux-amd64
```

## Alternatives
* [Queue Notify Client](https://github.com/rudikiaz/queue-notify-client)
 
