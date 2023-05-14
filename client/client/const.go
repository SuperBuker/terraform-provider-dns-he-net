package client

import "time"

const endpoint = "https://dns.he.net"

const retries = 3

const retryDelay = 30 * time.Second
