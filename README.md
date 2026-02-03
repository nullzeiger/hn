# hn Hacker News TUI
A modern Terminal User Interface (TUI) for browsing **Hacker News** top stories, built with **Go** and the **Bubble Tea** framework.

## Features
* **Real-time Top Stories:** Automatically fetches the top 30 trending stories.
* **Split-Pane Layout:** View the story list on the left and detailed metadata (score, author, link) on the right.
* **Dynamic Filtering:** Search through story titles instantly as you type.
* **Responsive Design:** The layout adapts fluidly to your terminal window size.

---

##  Project Structure
The project follows the **Model-View-Update (MVU)** pattern:

```text
hn/
├── main.go          # Application entry point
├── go.mod           # Dependency management
├── internal/api/
│   └── api.go       # HTTP client and Hacker News API integration
└── internal/ui/
   ├── model.go     # State and Message definitions
   ├── style.go     # Lipgloss styling
   ├── update.go    # Business logic and event handling
   └── view.go      # Interface rendering
```

## Getting Started
### Prerequisites
Go 1.25 or higher

### Installation
Clone the repository:
```bash
git clone https://github.com/nullzeiger/hn.git
cd hn
```

Initialize modules and download dependencies:
```bash
go mod tidy
```

Run the application:
```bash
go run main.go
```

## Controls (Keybindings)
| Key            | Action                         |
| -------------- | ------------------------------ |
| `↑`/`↓`        | Navigate through stories       |
| `Enter`        | Select a story to view details |
| `/`            | Activate search/filter mode    |
| `Esc`          | Clear filter or return to list |
| `q` / `Ctrl+C` | Quit application               |


## Built With
Bubble Tea - The TUI runtime based on the Elm Architecture

Lipgloss - For terminal layout and terminal-based CSS

Bubbles - Reusable UI components (list, spinner)

Hacker News API - Official Firebase-backed API

## License
Distributed under the BSD 3-Clause License. See LICENSE for more information.
