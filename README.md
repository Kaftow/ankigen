# Ankigen

Ankigen is a desktop application built with **Wails** (Go + WebView) and **Vue**. It extracts text from files (TXT, PDF, DOCX, etc.), optionally splits or paginates the content, and generates **Anki-ready cards** using LLM-driven or rule-based logic.

[![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white&style=flat-square)](https://golang.org/)
[![Wails](https://img.shields.io/badge/Wails-000000?style=flat-square)](https://wails.io/)
[![Vue](https://img.shields.io/badge/Vue-35495E?logo=vue.js&logoColor=4FC08D&style=flat-square)](https://vuejs.org/)
[![Node.js](https://img.shields.io/badge/Node.js-339933?logo=node.js&logoColor=white&style=flat-square)](https://nodejs.org/)
[![Vite](https://img.shields.io/badge/Vite-646CFF?logo=vite&logoColor=white&style=flat-square)](https://vitejs.dev/)
[![Anki](https://img.shields.io/badge/Anki-2EA1F6?style=flat-square)](https://apps.ankiweb.net/)


## Development

Start development server:

Prerequisites
- Go (>= 1.18)
- Node.js (>= 16) and npm or yarn
- Wails CLI: go install github.com/wailsapp/wails/v2/cmd/wails@latest

Quick start
```bash
go mod download
wails dev
```

Build the app
```bash
wails build
```

## License

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.


