package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	dir := "/etc/config"
	file := "app.yaml"
	path := filepath.Join(dir, file)

	log.Println("starting watcher, file:", path)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	if err := watcher.Add(dir); err != nil {
		log.Fatal(err)
	}

	// 启动时读一次
	b, _ := os.ReadFile(path)
	log.Println("initial load:", string(b))

	for {
		select {
		case event := <-watcher.Events:
			log.Printf("EVENT: %-10s %s\n", event.Op.String(), event.Name)

			if event.Op&(fsnotify.Write|
				fsnotify.Create|
				fsnotify.Remove|
				fsnotify.Rename) != 0 {

				time.Sleep(200 * time.Millisecond)
				b, err := os.ReadFile(path)
				if err != nil {
					log.Println("reload error:", err)
					continue
				}
				log.Println("🔥 reload:", string(b))
			}
		case err := <-watcher.Errors:
			log.Println("watcher error:", err)
		}
	}
}
