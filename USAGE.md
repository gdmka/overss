# 📖 Overss Usage Guide

This guide will walk you through setting up and using Overss to sync your audiobooks to Overcast.

## 🚀 Getting Started

### Step 1: Start the Server

**On macOS/Linux:**
```bash
./start.sh
```

**On Windows:**
```bash
start.bat
```

**Or manually:**
```bash
./overss
```

The server will start on `http://localhost:8083`

### Step 2: Access the Web Interface

Open your browser and navigate to:
```
http://localhost:8083
```

You'll see the Overss web interface with three tabs:
- **Files**: Manage your audiobook selection
- **Settings**: Configure your feed
- **Feed**: Get your RSS feed URL

## ⚙️ Configuration

### Initial Setup

1. Click on the **Settings** tab
2. Fill in your feed information:

   - **Feed Title**: Give your feed a name (e.g., "My Audiobooks")
   - **Description**: Brief description of your feed
   - **Author Name**: Your name
   - **Email**: Your email address
   - **Base URL**: 
     - For local use: `http://localhost:8083`
     - For network access: `http://YOUR_IP:8083` (e.g., `http://192.168.1.100:8083`)
   - **Audio Directory**: Path to your audiobooks
     - Default: `./audiobooks`
     - Can be absolute: `/Users/username/Audiobooks`
     - Can be network path: `/Volumes/NAS/Audiobooks`
   - **Cover Image URL**: (Optional) URL to a cover image for your feed

3. Click **Save Settings**

### Finding Your IP Address

**macOS:**
```bash
ifconfig | grep "inet " | grep -v 127.0.0.1
```

**Linux:**
```bash
ip addr show | grep "inet " | grep -v 127.0.0.1
```

**Windows:**
```bash
ipconfig
```

Look for your local IP address (usually starts with 192.168.x.x or 10.x.x.x)

## 📁 Managing Your Audiobooks

### Adding Audiobooks

1. Place your audiobook files in the configured audio directory
2. Supported formats:
   - MP3 (`.mp3`)
   - M4A/M4B (`.m4a`, `.m4b`)
   - OGG (`.ogg`)
   - OPUS (`.opus`)
   - FLAC (`.flac`)
   - WAV (`.wav`)
   - AAC (`.aac`)

### Organizing Files

You can organize your audiobooks in subdirectories:

```
audiobooks/
├── Fiction/
│   ├── Book1.m4b
│   └── Book2.m4b
├── Non-Fiction/
│   ├── Book3.mp3
│   └── Book4.mp3
└── Series/
    ├── Book1.m4b
    ├── Book2.m4b
    └── Book3.m4b
```

### Selecting Files for Your Feed

1. Go to the **Files** tab
2. You'll see all your audiobook files listed
3. Check the boxes next to the files you want in your feed
4. Click **Save Selection**

**Tip**: Use "Select All" to add all files, or "Deselect All" to start fresh

### File Information

Each file shows:
- **Name**: The filename
- **Path**: Relative path from audio directory
- **Size**: File size in MB
- **Date**: Last modified date

## 📡 Using Your RSS Feed

### Getting Your Feed URL

1. Go to the **Feed** tab
2. Copy the RSS Feed URL shown
3. It will look like: `http://YOUR_IP:8083/feed.xml`

### Adding to Overcast

1. Open Overcast on your iPhone/iPad
2. Tap the **+** button (Add Podcast)
3. Select **Add URL**
4. Paste your RSS feed URL
5. Tap **Subscribe**

Your audiobooks will now appear as podcast episodes!

### Adding to Other Podcast Apps

The RSS feed works with any podcast app that supports custom feeds:

**Pocket Casts:**
1. Tap Search → Enter URL
2. Paste your feed URL

**Apple Podcasts:**
1. Library → Edit → Add a Show by URL
2. Paste your feed URL

**Podcast Addict:**
1. Add Podcast → RSS Feed
2. Paste your feed URL

## 🔄 Updating Your Feed

### Adding New Books

1. Add new audiobook files to your audio directory
2. Go to the **Files** tab
3. Click **Refresh** to see new files
4. Select the new files
5. Click **Save Selection**

Your podcast app will automatically detect the new episodes on next refresh!

### Removing Books

1. Go to the **Files** tab
2. Uncheck the files you want to remove
3. Click **Save Selection**

The files will be removed from your feed but remain on disk.

## 🌐 Remote Access

### Accessing from Other Devices on Your Network

1. Find your computer's IP address (see above)
2. Update the **Base URL** in Settings to use your IP
3. Make sure port 8083 is not blocked by your firewall
4. Access from other devices using: `http://YOUR_IP:8083`

### Accessing from the Internet

**Option 1: Port Forwarding**
1. Configure your router to forward port 8083 to your computer
2. Find your public IP address (google "what is my ip")
3. Update Base URL to: `http://YOUR_PUBLIC_IP:8083`

**Option 2: Using ngrok (Recommended for testing)**
```bash
# Install ngrok from https://ngrok.com
ngrok http 8083
```

This will give you a public URL like: `https://abc123.ngrok.io`

Update your Base URL to this ngrok URL.

**Security Note**: For internet access, consider:
- Using HTTPS (via reverse proxy like Caddy or nginx)
- Adding authentication
- Using a VPN instead

## 🎯 Tips & Best Practices

### File Naming

Use descriptive filenames - they become episode titles:
- ✅ Good: `Harry Potter 1 - Philosophers Stone.m4b`
- ❌ Bad: `book1.mp3`

### Organizing Large Collections

For large collections, use subdirectories:
```
audiobooks/
├── Author Name/
│   ├── Series Name/
│   │   ├── Book 1.m4b
│   │   └── Book 2.m4b
│   └── Standalone Book.m4b
```

### Network Storage

Overss works great with network storage:

**NAS/Network Drive:**
```
Audio Directory: /Volumes/NAS/Audiobooks
```

**External Drive:**
```
Audio Directory: /Volumes/External/Audiobooks
```

### Backup Your Configuration

Keep a backup of `config.json`:
```bash
cp config.json config.backup.json
```

### Performance

- The server scans files on demand
- Large directories (1000+ files) may take a moment to load
- Consider organizing into subdirectories for better performance

## 🐛 Troubleshooting

### Files Not Showing Up

**Problem**: No files appear in the Files tab

**Solutions**:
1. Check the Audio Directory path in Settings
2. Verify files have supported extensions
3. Click the **Refresh** button
4. Check file permissions (files must be readable)

### Can't Access from Other Devices

**Problem**: Can't connect from phone/tablet

**Solutions**:
1. Verify Base URL uses your IP, not `localhost`
2. Check both devices are on the same network
3. Disable firewall temporarily to test
4. Try accessing from browser first: `http://YOUR_IP:8083`

### RSS Feed Not Updating

**Problem**: Changes don't appear in Overcast

**Solutions**:
1. Verify you clicked **Save Selection**
2. Force refresh in Overcast (pull down)
3. Check the feed URL in browser: `http://YOUR_IP:8083/feed.xml`
4. Some apps cache feeds - wait 5-10 minutes

### Audio Won't Play

**Problem**: Episodes appear but won't play

**Solutions**:
1. Verify Base URL is accessible from your device
2. Test by opening audio URL in browser
3. Check file isn't corrupted
4. Verify file format is supported by your app

### Server Won't Start

**Problem**: Error when starting server

**Solutions**:
1. Check if port 8083 is already in use
2. Try a different port in config.json
3. Check file permissions
4. Verify Go is installed correctly

## 📊 Understanding the Stats

The Files tab shows three statistics:

- **Total Files**: All audio files found in your directory
- **Selected Files**: Files currently in your RSS feed
- **Total Size**: Combined size of all files

## 🔐 Security Considerations

### Local Network Only

By default, Overss is designed for local network use:
- No authentication required
- Files served directly
- Suitable for home networks

### For Internet Access

If exposing to the internet:
1. Use a reverse proxy (nginx, Caddy) with HTTPS
2. Add authentication (basic auth, OAuth)
3. Consider using a VPN instead
4. Regularly update the software

### File Access

- Overss can only serve files from the configured audio directory
- It cannot access files outside this directory
- Subdirectories within audio directory are accessible

## 🎓 Advanced Usage

### Custom Port

Edit `config.json`:
```json
{
  "port": ":3000",
  ...
}
```

### Multiple Feeds

Run multiple instances with different configs:
```bash
./overss -config feed1.json &
./overss -config feed2.json &
```

### Automation

Create a systemd service (Linux) or LaunchAgent (macOS) to start automatically.

See README.md for examples.

### Docker

Run in a container for isolation:
```bash
docker run -p 8083:8083 -v /path/to/audiobooks:/root/audiobooks overss
```

## 📞 Getting Help

If you encounter issues:

1. Check this guide first
2. Review the README.md
3. Check the server logs for errors
4. Verify your configuration in config.json
5. Test with a simple setup first (few files, local access)

## 🎉 Enjoy Your Audiobooks!

You're all set! Your audiobooks are now available as a podcast feed that you can access from any podcast app.

Happy listening! 📚🎧
