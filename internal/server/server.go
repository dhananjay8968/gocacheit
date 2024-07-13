package server

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "gocacheit/internal/cache"
)

type CacheServer struct {
    cache cache.Cache
}

func NewCacheServer(c cache.Cache) *CacheServer {
    return &CacheServer{cache: c}
}

func (cs *CacheServer) handleGet(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")
    if key == "" {
        http.Error(w, "Key parameter is required", http.StatusBadRequest)
        return
    }
    value, found := cs.cache.Get(key)
    if !found {
        http.Error(w, "Key not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(map[string]interface{}{"key": key, "value": value})
}

func (cs *CacheServer) handlePut(w http.ResponseWriter, r *http.Request) {
    var data map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    key, ok1 := data["key"].(string)
    value, ok2 := data["value"].(interface{})
    if !ok1 || !ok2 {
        http.Error(w, "Invalid key or value", http.StatusBadRequest)
        return
    }

    cs.cache.Put(key, value)
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Key %s stored successfully", key)
}

func (cs *CacheServer) Start(port string) {
    http.HandleFunc("/get", cs.handleGet)
    http.HandleFunc("/put", cs.handlePut)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
