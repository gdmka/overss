# 🚀 Overss Quick Start

Get your audiobooks on Overcast in 5 minutes!

## 1️⃣ Start the Server

**macOS/Linux:**
```bash
./start.sh
```

**Windows:**
```bash
start.bat
```

## 2️⃣ Open Web Interface

Open in browser: **http://localhost:8083**

## 3️⃣ Configure (Settings Tab)

- **Base URL**: `http://localhost:8083` (or your IP for remote access)
- **Audio Directory**: Path to your audiobooks (default: `./audiobooks`)
- Fill in feed title, author, etc.
- Click **Save Settings**

## 4️⃣ Select Files (Files Tab)

- Check the audiobooks you want in your feed
- Click **Save Selection**

## 5️⃣ Add to Overcast (Feed Tab)

1. Copy the RSS Feed URL
2. Open Overcast → Add Podcast → Add URL
3. Paste the URL
4. Done! 🎉

## 📱 Remote Access

To access from your phone:

1. Find your computer's IP: `ifconfig | grep inet` (macOS/Linux)
2. Update Base URL in Settings to: `http://YOUR_IP:8083`
3. Use this URL in Overcast

## 🆘 Need Help?

- Full documentation: [`README.md`](README.md)
- Detailed guide: [`USAGE.md`](USAGE.md)
- Example config: [`config.example.json`](config.example.json)

## 📁 Supported Formats

MP3, M4A, M4B, OGG, OPUS, FLAC, WAV, AAC

## 🎯 Quick Tips

- Use descriptive filenames (they become episode titles)
- Organize in subdirectories for better management
- Click Refresh in Files tab after adding new books
- Force refresh in Overcast to see updates

---

**That's it! Enjoy your audiobooks! 📚🎧**
