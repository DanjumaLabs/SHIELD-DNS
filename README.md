# SHIELD-DNS

A lightweight DNS filtering server built with Go that blocks unwanted domains and forwards legitimate requests to upstream DNS servers.

## Features

- **DNS Filtering**: Block domains from a customizable blocklist
- **Multiple Upstream Servers**: Automatically fallback to alternate DNS servers if one fails
- **Domain Management**: Add new domains to the blocklist dynamically
- **Configurable**: Easy configuration via `config.json`
- **Logging Support**: Track DNS queries and activities
- **Fast & Lightweight**: Written in Go for high performance

## Requirements

- **Go**: Version 1.26.1 or higher
- **Dependencies**:
  - `github.com/miekg/dns` v1.1.72 - DNS library for Go

## Installation

### 1. Clone or download the project
```bash
cd SHIELD-DNS
```

### 2. Install dependencies
```bash
go mod download
```

### 3. Build the executable
```bash
go build -o SHIELD-DNS.exe
```

## Configuration

Configure the server using `config.json`:

```json
{
    "port": ":53",
    "upstream_servers": [
        "1.1.1.1:53",
        "8.8.8.8:53"
    ],
    "blocklist_file": "blocked.txt",
    "log_file": "shielddns.log"
}
```

### Configuration Options

| Option | Description | Example |
|--------|-------------|---------|
| `port` | Port to listen on (requires admin/root privileges for port 53) | `:53` or `:5353` |
| `upstream_servers` | List of upstream DNS servers to forward requests to | `["1.1.1.1:53", "8.8.8.8:53"]` |
| `blocklist_file` | Path to the blocklist file | `blocked.txt` |
| `log_file` | Path to the log file | `shielddns.log` |

## Blocklist Format

Create a `blocked.txt` file with one domain per line:

```
tiktok.com
youtube.com
facebook.com
ben.com
```

- Empty lines and lines starting with `#` are ignored
- Domain names are trimmed of whitespace

## Usage

### Start the DNS Server
```bash
.\SHIELD-DNS.exe start
```

**Output:**
```
SHIELD-DNS Starting Smoothly on port :53...
Blacklist loaded: [tiktok.com youtube.com ben.com]
ShieldDNS DNS server is running on 127.0.0.1:53
```

### View Stats
```bash
.\SHIELD-DNS.exe stats
```

### Add a Domain to Blocklist
```bash
.\SHIELD-DNS.exe add <domain>
```

**Example:**
```bash
.\SHIELD-DNS.exe add malicious-site.com
```

**Output:**
```
ADDED: malicious-site.com
```

## How It Works

1. **DNS Request Received**: Server receives a DNS query for a domain
2. **Blacklist Check**: Checks if the domain is in the blocklist
   - If **blocked**: Returns `NXDOMAIN` (domain does not exist)
   - If **allowed**: Forwards to upstream DNS servers
3. **Upstream Resolution**: Tries each upstream server in order
   - Uses first successful response
   - Automatically fails over to next server if needed
4. **Response**: Sends resolved IP back to client

### Blocked Domain Response
When a domain is blocked, the server responds with:
- **Response Code**: `NXDOMAIN` (Name Error)
- **Effect**: Browser/application treats domain as non-existent

### Allowed Domain Response
Valid domains are resolved through upstream servers:
- **Primary**: Cloudflare DNS (1.1.1.1:53)
- **Fallback**: Google DNS (8.8.8.8:53)

## Project Structure

```
SHIELD-DNS/
├── main.go              # Main entry point and command handling
├── config.json          # Configuration file
├── blocked.txt          # Blocklist of domains
├── go.mod              # Go module definition
├── server/
│   ├── config.go       # Configuration loading
│   ├── dns_server.go   # DNS server setup
│   ├── handler.go      # DNS request handler
│   └── log.go          # Logging functionality
├── filter/
│   └── blacklist.go    # Blocklist operations
└── start.go            # Startup functions
```

## Troubleshooting

### Server won't start on port 53
**Issue**: Permission denied
- **Solution**: Run with administrator/root privileges or use a different port (e.g., `:5353`)

### Domains not being blocked
**Issue**: Domains still resolve even though they're in `blocked.txt`
- **Solution**: 
  - Ensure your system is using this server as DNS (configure in network settings)
  - Restart the server after modifying `blocked.txt`
  - Check domain format (no extra spaces or special characters)

### Upstream server errors
**Issue**: `All upstream servers failed` message
- **Solution**: 
  - Check internet connection
  - Verify upstream server addresses in `config.json`
  - Try different DNS providers (OpenDNS: `208.67.222.222:53`)

### Cannot compile
**Issue**: `go: cannot find module`
- **Solution**: Run `go mod download` to fetch dependencies

## Development

### Running Tests
```bash
go test ./...
```

### Building with Debug Info
```bash
go build -v -o SHIELD-DNS.exe
```

## License

See LICENSE file in the project directory.

## Support

For issues or feature requests, please refer to the project repository or documentation.

---

**Built with ❤️ using Go**

