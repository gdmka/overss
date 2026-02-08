package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/feeds"
	"github.com/gorilla/mux"
)

type Config struct {
	Port          string   `json:"port"`
	BaseURL       string   `json:"base_url"`
	AudioDir      string   `json:"audio_dir"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Author        string   `json:"author"`
	Email         string   `json:"email"`
	ImageURL      string   `json:"image_url"`
	SelectedFiles []string `json:"selected_files"`
}

type AudioFile struct {
	Name     string    `json:"name"`
	Path     string    `json:"path"`
	Size     int64     `json:"size"`
	ModTime  time.Time `json:"mod_time"`
	Selected bool      `json:"selected"`
}

var config Config

func main() {
	configFile := flag.String("config", "config.json", "Path to configuration file")
	flag.Parse()

	// Load or create default config
	if err := loadConfig(*configFile); err != nil {
		log.Printf("Creating default config: %v", err)
		createDefaultConfig(*configFile)
	}

	r := mux.NewRouter()

	// API endpoints
	r.HandleFunc("/api/files", listFilesHandler).Methods("GET")
	r.HandleFunc("/api/config", getConfigHandler).Methods("GET")
	r.HandleFunc("/api/config", updateConfigHandler).Methods("POST")
	r.HandleFunc("/api/selection", updateSelectionHandler).Methods("POST")

	// RSS feed endpoint
	r.HandleFunc("/feed.xml", rssFeedHandler).Methods("GET")
	r.HandleFunc("/rss", rssFeedHandler).Methods("GET")

	// Serve audio files
	r.PathPrefix("/audio/").Handler(http.StripPrefix("/audio/", http.FileServer(http.Dir(config.AudioDir))))

	// Serve static UI files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	log.Printf("Starting Overss RSS Server on port %s", config.Port)
	log.Printf("RSS Feed URL: %s/feed.xml", config.BaseURL)
	log.Printf("")
	log.Printf("Access the server at:")
	log.Printf("  Local:   http://localhost%s", config.Port)

	// Get and display all network interfaces
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					log.Printf("  Network: http://%s%s", ipnet.IP.String(), config.Port)
				}
			}
		}
	}
	log.Printf("")
	log.Printf("Press Ctrl+C to stop the server")
	log.Printf("")

	log.Fatal(http.ListenAndServe(config.Port, r))
}

func loadConfig(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &config)
}

func createDefaultConfig(filename string) {
	config = Config{
		Port:          ":8083",
		BaseURL:       "http://localhost:8083",
		AudioDir:      "./audiobooks",
		Title:         "My Audiobook Feed",
		Description:   "Personal audiobook RSS feed for Overcast",
		Author:        "Overss",
		Email:         "user@example.com",
		ImageURL:      "",
		SelectedFiles: []string{},
	}

	// Create audiobooks directory if it doesn't exist
	os.MkdirAll(config.AudioDir, 0755)

	data, _ := json.MarshalIndent(config, "", "  ")
	os.WriteFile(filename, data, 0644)
}

func saveConfig(filename string) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func listFilesHandler(w http.ResponseWriter, r *http.Request) {
	files, err := scanAudioFiles(config.AudioDir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
}

func getConfigHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}

func updateConfigHandler(w http.ResponseWriter, r *http.Request) {
	var newConfig Config
	if err := json.NewDecoder(r.Body).Decode(&newConfig); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	config.Title = newConfig.Title
	config.Description = newConfig.Description
	config.Author = newConfig.Author
	config.Email = newConfig.Email
	config.ImageURL = newConfig.ImageURL
	config.BaseURL = newConfig.BaseURL
	config.AudioDir = newConfig.AudioDir

	if err := saveConfig("config.json"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func updateSelectionHandler(w http.ResponseWriter, r *http.Request) {
	var selection struct {
		Files []string `json:"files"`
	}

	if err := json.NewDecoder(r.Body).Decode(&selection); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	config.SelectedFiles = selection.Files

	if err := saveConfig("config.json"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func scanAudioFiles(dir string) ([]AudioFile, error) {
	var files []AudioFile
	supportedExts := map[string]bool{
		".mp3": true, ".m4a": true, ".m4b": true,
		".ogg": true, ".opus": true, ".flac": true,
		".wav": true, ".aac": true,
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(path))
			if supportedExts[ext] {
				relPath, _ := filepath.Rel(dir, path)
				selected := false
				for _, sf := range config.SelectedFiles {
					if sf == relPath {
						selected = true
						break
					}
				}

				files = append(files, AudioFile{
					Name:     info.Name(),
					Path:     relPath,
					Size:     info.Size(),
					ModTime:  info.ModTime(),
					Selected: selected,
				})
			}
		}
		return nil
	})

	return files, err
}

func rssFeedHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	feed := &feeds.Feed{
		Title:       config.Title,
		Link:        &feeds.Link{Href: config.BaseURL},
		Description: config.Description,
		Author:      &feeds.Author{Name: config.Author, Email: config.Email},
		Created:     now,
	}

	if config.ImageURL != "" {
		feed.Image = &feeds.Image{
			Url:   config.ImageURL,
			Title: config.Title,
			Link:  config.BaseURL,
		}
	}

	// Add selected files to feed
	for _, relPath := range config.SelectedFiles {
		fullPath := filepath.Join(config.AudioDir, relPath)
		info, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		// Create enclosure URL
		enclosureURL := fmt.Sprintf("%s/audio/%s", config.BaseURL, strings.ReplaceAll(relPath, "\\", "/"))

		// Determine MIME type
		mimeType := getMimeType(filepath.Ext(relPath))

		item := &feeds.Item{
			Title:       strings.TrimSuffix(filepath.Base(relPath), filepath.Ext(relPath)),
			Link:        &feeds.Link{Href: enclosureURL},
			Description: fmt.Sprintf("Audiobook: %s", filepath.Base(relPath)),
			Created:     info.ModTime(),
			Enclosure: &feeds.Enclosure{
				Url:    enclosureURL,
				Length: fmt.Sprintf("%d", info.Size()),
				Type:   mimeType,
			},
		}

		feed.Items = append(feed.Items, item)
	}

	// Generate RSS 2.0 feed
	rss, err := feed.ToRss()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/rss+xml; charset=utf-8")
	w.Write([]byte(rss))
}

func getMimeType(ext string) string {
	mimeTypes := map[string]string{
		".mp3":  "audio/mpeg",
		".m4a":  "audio/mp4",
		".m4b":  "audio/mp4",
		".ogg":  "audio/ogg",
		".opus": "audio/opus",
		".flac": "audio/flac",
		".wav":  "audio/wav",
		".aac":  "audio/aac",
	}

	if mime, ok := mimeTypes[strings.ToLower(ext)]; ok {
		return mime
	}
	return "audio/mpeg"
}
