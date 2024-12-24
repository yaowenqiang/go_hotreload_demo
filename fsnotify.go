package main

import (
    "log"
    "os"
    "os/exec"

    "github.com/fsnotify/fsnotify"
)

func main() {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    // 监控的目录
    watchDir := "./"
    err = watcher.Add(watchDir)
    if err != nil {
        log.Fatal(err)
    }

    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                // 检测到文件变更，重启服务
                if event.Op&(fsnotify.Write|fsnotify.Create) != 0 {
                    log.Printf("File changed: %s, restarting server...\n", event.Name)
                    restartServer()
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("Error:", err)
            }
        }
    }()

    // 保持主程序运行
    done := make(chan bool)
    <-done
}

func restartServer() {
    cmd := exec.Command("go", "run", "main.go")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        log.Println("Error restarting server:", err)
    }
}

