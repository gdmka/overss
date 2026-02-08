# 📚 Overss - RSS Feed Server for Audiobooks

Overss is a lightweight Go-based RSS feed server that allows you to sync your audiobook collection to podcast apps like Overcast. It provides a beautiful web UI for managing your audiobook files and generates a podcast-compatible RSS feed.

## ✨ Features

- 🎵 **Multiple Audio Format Support**: MP3, M4A, M4B, OGG, OPUS, FLAC, WAV, AAC
- 🖥️ **Beautiful Web UI**: Easy-to-use interface for file selection and configuration
- 📡 **RSS 2.0 Feed**: Compatible with all major podcast apps (Overcast, Pocket Casts, etc.)
- 📁 **Flexible Storage**: Works with local disk or network storage
- ⚙️ **Configurable**: Customize feed metadata, directories, and URLs
- 🚀 **Lightweight**: Single binary with no external dependencies
- 🔄 **Real-time Updates**: Changes reflect immediately in your feed

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher (for building from source)
- Your audiobook files in a directory

### Installation

1. **Clone the repository**:
```bash
git clone <repository-url>
cd overss
```

2. **Install dependencies**:
```bash
go mod tidy
```

3. **Build the application**:
```bash
go build -o overss
```

4. **Run the server**:

**Option A: With ngrok (Internet Access)**
```bash
./start.sh
```
This will start the server and automatically create an ngrok tunnel for internet access.

**Option B: Local Network Only**
```bash
./start-local.sh
```
This will start the server for local network access only.

**Option C: Manual Start**
```bash
./overss
```

The server will start on `http://localhost:8083` and display all available network addresses.

### First-Time Setup

1. Open your browser and navigate to `http://localhost:8083`
2. Go to the **Settings** tab
3. Configure your feed details:
   - **Feed Title**: Name of your audiobook feed
   - **Description**: Brief description
   - **Author Name**: Your name
   - **Base URL**: Your server's public URL (use your IP or domain for remote access)
   - **Audio Directory**: Path to your audiobook files (default: `./audiobooks`)
4. Click **Save Settings**

## 📖 Usage

### Adding Audiobooks to Your Feed

1. **Navigate to the Files tab**
2. Your audiobook files will be automatically scanned and listed
3. **Select the files** you want to include in your RSS feed
4. Click **Save Selection**

### Subscribing in Overcast (or any podcast app)

1. Go to the **Feed** tab in the web UI
2. **Copy the RSS Feed URL**
3. Open Overcast on your device
4. Tap the **+** button to add a new podcast
5. Select **Add URL**
6. Paste your RSS feed URL
7. Your audiobooks will appear as podcast episodes!

## 🔧 Configuration

### Configuration File

Overss uses a `config.json` file for persistent configuration. It's automatically created on first run with default values:

```json
{
  "port": ":8083",
  "base_url": "http://localhost:8083",
  "audio_dir": "./audiobooks",
  "title": "My Audiobook Feed",
  "description": "Personal audiobook RSS feed for Overcast",
  "author": "Overss",
  "email": "user@example.com",
  "image_url": "",
  "selected_files": []
}
```

### Command-Line Options

```bash
./overss -config /path/to/config.json
```

- `-config`: Specify a custom configuration file path (default: `config.json`)

### Startup Scripts

**`start.sh`** - Starts server with ngrok tunnel (automatic internet access)
- Builds the application if needed
- Starts the server in the background
- Launches ngrok for public internet access
- Perfect for accessing from anywhere

**`start-local.sh`** - Starts server for local network only
- Builds the application if needed
- Displays all local network addresses
- No internet exposure
- Perfect for home network use

**`start.bat`** - Windows version (local network only)

### Environment Setup for Remote Access

**Local Network Access:**
The server automatically displays all available network addresses on startup. Use any of the "Network" URLs shown to access from other devices on your network.

**Internet Access (via ngrok):**
1. Run `./start.sh` (automatically starts ngrok)
2. Copy the ngrok URL from the terminal (e.g., `https://abc123.ngrok.io`)
3. Update the **Base URL** in Settings to use the ngrok URL
4. Your feed is now accessible from anywhere!

**Manual ngrok setup:**
```bash
# In one terminal
./overss

# In another terminal
ngrok http 8083
```

## 📁 Directory Structure

```
overss/
├── main.go              # Main application code
├── go.mod               # Go module dependencies
├── config.json          # Configuration file (auto-generated)
├── audiobooks/          # Default audiobook directory (auto-created)
├── static/
│   └── index.html       # Web UI
└── README.md
```

## 🎯 API Endpoints

### Web UI
- `GET /` - Web interface

### API
- `GET /api/files` - List all audio files
- `GET /api/config` - Get current configuration
- `POST /api/config` - Update configuration
- `POST /api/selection` - Update file selection

### RSS Feed
- `GET /feed.xml` - RSS 2.0 feed
- `GET /rss` - Alternative RSS endpoint

### Audio Files
- `GET /audio/{filepath}` - Serve audio files

## 🔒 Security Considerations

- **Local Network Only**: By default, the server binds to all interfaces. For production use, consider:
  - Using a reverse proxy (nginx, Caddy) with HTTPS
  - Implementing authentication
  - Restricting access to specific IP ranges

- **File Access**: The server can only serve files from the configured audio directory

## 🐛 Troubleshooting

### Files Not Showing Up

1. Check that your audio directory path is correct in Settings
2. Verify that files have supported extensions (mp3, m4a, m4b, etc.)
3. Click the **Refresh** button in the Files tab

### Can't Access from Other Devices

1. Verify your Base URL uses your computer's IP address, not `localhost`
2. Check firewall settings allow connections on port 8083
3. Ensure both devices are on the same network

### RSS Feed Not Updating in Overcast

1. Make sure you clicked **Save Selection** after choosing files
2. Try refreshing the feed in Overcast (pull down to refresh)
3. Some podcast apps cache feeds - wait a few minutes or force refresh

### Audio Files Won't Play

1. Verify the Base URL is accessible from your device
2. Check that the audio files aren't corrupted
3. Ensure your podcast app supports the audio format

## 🛠️ Development

### Building from Source

```bash
# Clone the repository
git clone <repository-url>
cd overss

# Install dependencies
go mod tidy

# Run in development mode
go run main.go

# Build for production
go build -o overss

# Build for different platforms
GOOS=linux GOARCH=amd64 go build -o overss-linux
GOOS=windows GOARCH=amd64 go build -o overss.exe
GOOS=darwin GOARCH=arm64 go build -o overss-mac
```

### Project Structure

- **main.go**: Core application logic
  - HTTP server setup
  - RSS feed generation
  - File scanning and management
  - Configuration handling
  
- **static/index.html**: Single-page web application
  - File selection interface
  - Settings management
  - Feed information display

## 📝 Supported Audio Formats

| Format | Extension | MIME Type  |
| ------ | --------- | ---------- |
| MP3    | .mp3      | audio/mpeg |
| M4A    | .m4a      | audio/mp4  |
| M4B    | .m4b      | audio/mp4  |
| OGG    | .ogg      | audio/ogg  |
| OPUS   | .opus     | audio/opus |
| FLAC   | .flac     | audio/flac |
| WAV    | .wav      | audio/wav  |
| AAC    | .aac      | audio/aac  |

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is open source and available under the GNU General Public License v3.0 (GPLv3). See the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with [Gorilla Mux](https://github.com/gorilla/mux) for routing
- RSS generation powered by [Gorilla Feeds](https://github.com/gorilla/feeds)
- Designed for use with [Overcast](https://overcast.fm/) podcast app

## 💡 Tips & Best Practices

1. **Organize Your Files**: Use subdirectories to organize audiobooks by author or series
2. **File Naming**: Use descriptive names - they become episode titles in your podcast app
3. **Cover Art**: Add a cover image URL in settings for a professional look
4. **Regular Updates**: The feed updates immediately when you change selection
5. **Backup Config**: Keep a backup of your `config.json` file
6. **Network Storage**: Works great with NAS or network drives - just point the audio directory to your mount point

## 🚀 Advanced Usage

### Using with Docker (Optional)

Create a `Dockerfile`:

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o overss

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/overss .
COPY --from=builder /app/static ./static
EXPOSE 8083
CMD ["./overss"]
```

Build and run:
```bash
docker build -t overss .
docker run -p 8083:8083 -v /path/to/audiobooks:/root/audiobooks overss
```

### Systemd Service (Linux)

Create `/etc/systemd/system/overss.service`:

```ini
[Unit]
Description=Overss RSS Feed Server
After=network.target

[Service]
Type=simple
User=youruser
WorkingDirectory=/path/to/overss
ExecStart=/path/to/overss/overss
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

Enable and start:
```bash
sudo systemctl enable overss
sudo systemctl start overss
```

---

**Enjoy your audiobooks on the go! 📚🎧**
