# JARVIS OS
### Just A Rather Very Intelligent Operating System

> A custom Linux-based OS where AI is the primary interface. No desktop. No bash prompt. Just you and JARVIS.

---

## What is this?

JARVIS OS is an open source, AI-native operating system built on top of the Linux kernel. Instead of a traditional desktop environment or shell, JARVIS replaces the entire user interface with an AI layer — powered by [Ollama](https://ollama.com) locally when offline, and a user-chosen cloud provider (Kimi, OpenAI, etc.) when connected.

Built entirely in **Go** (daemon + TUI) and **Rust** (Wayland compositor), using open source software throughout.

---

## Architecture

```
Linux Kernel (standard — drivers included)
        ↓
systemd (init + service manager)
        ↓
aria-daemon (Go) ← Fiber API on :7777
        ↓
Provider Switcher → Ollama (offline) / Cloud AI (online)
        ↓
Wayland Compositor (Rust + smithay)  [planned]
        ↓
AI Interface — bubbletea TUI (Go)
        ↓
Observability — Grafana + Loki + OpenTelemetry  [planned]
```

---

## Status

### ✅ Done

| Component | Description |
|---|---|
| Project scaffold | Folder structure, Go modules initialized |
| `aria-daemon` | Go binary — Fiber REST API on `localhost:7777` |
| `/status` endpoint | Returns daemon health + current provider |
| `/chat` endpoint | Accepts user message, returns AI reply |
| Ollama integration | Local LLM via official Go client (`phi3` model) |
| Provider switcher | Detects WiFi — routes to Ollama if offline |
| bubbletea TUI | Terminal UI — user types, ARIA responds |
| GitHub repo | Public, MIT licensed, open source from day one |

---

### 🔧 In Progress

| Component | Description |
|---|---|
| OS tool registry | AI can read/write files, run commands, watch filesystem |
| Cloud provider | Kimi / OpenAI fallback when WiFi is detected |
| Streaming responses | SSE stream from daemon → TUI renders token by token |

---

### 📋 Planned

| Component | Description |
|---|---|
| `systemd` service | `aria-daemon.service` — auto-starts on boot |
| Filesystem tools | `fs.list`, `fs.read`, `fs.write` exposed to LLM as tool calls |
| Exec tools | Sandboxed command execution with user confirmation |
| dbus integration | System events — battery, network, notifications |
| Wayland compositor | Rust + smithay — replaces GNOME/KDE entirely |
| Custom boot screen | Branded Plymouth splash on startup |
| Grafana dashboard | System metrics + AI request observability |
| Loki log collection | AI request logs, provider usage, error tracking |
| OpenTelemetry | Latency, token count, provider used per request |
| ISO packaging | Bootable Ubuntu-based ISO via `cubic` / `live-build` |
| Multi-provider config | YAML config — user sets preferred cloud provider |
| Security sandbox | Whitelist for exec calls, confirmation for destructive ops |

---

## Tech Stack

| Layer | Technology | License |
|---|---|---|
| Daemon / API | Go + [Fiber](https://github.com/gofiber/fiber) | MIT |
| Local LLM | [Ollama](https://github.com/ollama/ollama) | MIT |
| TUI | [bubbletea](https://github.com/charmbracelet/bubbletea) + [lipgloss](https://github.com/charmbracelet/lipgloss) | MIT |
| Wayland compositor | Rust + [smithay](https://github.com/Smithay/smithay) | MIT |
| Observability | [Grafana](https://github.com/grafana/grafana) + [Loki](https://github.com/grafana/loki) | AGPL |
| Tracing | [OpenTelemetry Go](https://github.com/open-telemetry/opentelemetry-go) | Apache 2.0 |
| Init system | systemd | LGPL |
| Base kernel | Linux | GPL v2 |

---

## Getting Started

### Prerequisites

- Ubuntu 22.04+ (dev environment)
- Go 1.22+
- Ollama

### Install Go

```bash
wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### Install Ollama

```bash
curl -fsSL https://ollama.com/install.sh | sh
ollama pull phi3
```

### Run ARIA

```bash
git clone https://github.com/YOURUSERNAME/aria-os
cd aria-os

# Terminal 1 — start the daemon
cd daemon && go run main.go

# Terminal 2 — start the TUI
cd interface/tui && go run main.go
```

---

## Project Structure

```
aria-os/
├── daemon/
│   ├── main.go               # entry point
│   ├── api/
│   │   └── server.go         # Fiber REST API
│   ├── provider/
│   │   ├── switcher.go       # WiFi detection + routing
│   │   └── ollama.go         # Ollama Go client
│   └── tools/
│       ├── fs.go             # filesystem tools [planned]
│       └── exec.go           # command execution [planned]
├── interface/
│   └── tui/
│       └── main.go           # bubbletea terminal UI
└── README.md
```

---

## Roadmap

| Phase | Goal |
|---|---|
| v0.1 | ✅ Daemon + TUI + Ollama offline chat |
| v0.2 | OS tools — AI reads/writes files, runs commands |
| v0.3 | Cloud provider switching (Kimi / OpenAI) + streaming |
| v0.4 | systemd service — auto-start on boot |
| v0.5 | Grafana + Loki observability |
| v0.6 | Wayland compositor — replace GNOME entirely |
| v1.0 | Bootable ISO — installable ARIA OS |

---

## Contributing

Contributions welcome. Please open an issue before submitting a PR so we can discuss the approach.

- Fork the repo
- Create a branch: `git checkout -b feat/your-feature`
- Commit with clear messages: `feat:`, `fix:`, `docs:`
- Open a PR against `main`

---

## License

MIT — see [LICENSE](./LICENSE)

---

> Built with Go, Rust, and open source software throughout.
> ARIA OS is not affiliated with Ubuntu or Canonical.