# 🚀 GemGo CLI

**GemGo** is a **simple, fast, and powerful** command-line interface (CLI) for interacting with the **Google Gemini API**, built with **Go**. It’s perfect for developers, automation, and scripting enthusiasts who want AI capabilities right from their terminal.

---

## 📝 Overview

GemGo lets you send prompts directly to the Gemini API from your terminal. You can pass input as an argument or via piped input — making it a highly flexible tool for daily tasks and scripting.

---

## ✨ Features

- 🧠 **Direct & Piped Input**  
  Use `gemgo "your prompt"` or pipe data like `echo "prompt" | gemgo`.

- ⚙️ **Simple Configuration**  
  Just set your API key using environment variables.

- 🚀 **Fast & Lightweight**  
  Built in Go for optimal performance and low overhead.

- 💻 **Cross-Platform Support**  
  Runs on macOS, Linux, and Windows.

---

## 🛠️ Installation

### ✅ Prerequisites

- Go 1.20 or later
- A [Google Gemini API Key](https://aistudio.google.com/app/apikey)
- Set the environment variable: `GEMINI_API_KEY`

---

### 📦 Steps to Install

1. **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/gemgo.git
    cd gemgo
    ```

2. **Build the binary:**

    ```bash
    go build -o gemgo .
    ```

3. **Move the binary to your PATH:**

    **macOS / Linux:**

    ```bash
    sudo mv gemgo /usr/local/bin/
    ```

    **Windows:**

    Move `gemgo.exe` to a folder included in your system’s `PATH`.

---

## ⚙️ Configuration

Set your **Gemini API key** as an environment variable:

1. **Add to your shell config file** (`.zshrc`, `.bashrc`, `.bash_profile`, etc.):

    ```bash
    export GEMINI_API_KEY="YOUR_API_KEY_HERE"
    ```

2. **Apply the changes:**

    ```bash
    source ~/.zshrc   # or source ~/.bashrc
    ```

---

## 🚀 Usage

### 📌 1. Prompt as an Argument

Send a prompt directly as a command-line argument:

```bash
gemgo "What is the Go programming language?"
```
---

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1.  Fork the Project
2.  Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3.  Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4.  Push to the Branch (`git push origin feature/AmazingFeature`)
5.  Open a Pull Request

---

## License

Distributed under the MIT License. See `LICENSE` for more information.

---
