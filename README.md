# gocacheit

gocacheit is a Go-based HTTP cache server with support for LRU caching and consistent hashing.

## Features

- **LRU Caching:** Efficient eviction of least recently used items.
- **Consistent Hashing:** Distributes data across nodes for load balancing.
- **HTTP Server:** Exposes endpoints for caching operations.
- **Concurrency:** Handles concurrent requests using goroutines.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Endpoints](#endpoints)
    - [GET /get](#get-get)
    - [POST /put](#post-put)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)
