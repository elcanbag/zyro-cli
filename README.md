# ⚡ Zyro CLI

**Zyro** is a lightweight and ultra-fast command-line tool written in **Go**.  
It offers a powerful set of security and developer utilities like password generation, hashing, bcrypt, hash verification, and fun extras — all from your terminal.

---

## 🎯 Features

| Command                          | Description                                     |
|----------------------------------|-------------------------------------------------|
| `zyro pass [length]`             | Generate a secure random password (default: 12) |
| `zyro hash [text]`               | Generate a SHA-256 hash from the input string   |
| `zyro bcrypt [text]`             | Generate a bcrypt hash from the input string    |
| `zyro verify-hash [text] [hash]` | Verify text against a given bcrypt hash         |
| `zyro flip`                      | Flip a virtual coin (🪙 Heads or Tails)         |
| `zyro -v` / `--version`          | Show Zyro version                               |
| `zyro help`                      | Show built-in help menu                         |

---

## 🚀 Usage Examples

```bash
# Generate a 16-character password
zyro pass 16

# SHA256 hash
zyro hash HelloWorld

# Bcrypt hash
zyro bcrypt test123

# Verify bcrypt hash (use single quotes in PowerShell!)
zyro verify-hash test123 '$2a$10$k8F...'

# Flip a coin
zyro flip

# Check version
zyro -v
```

---

## 💡 Why Use Zyro?

- 🧱 **Single Binary** — no dependencies  
- ⚡ **Instant Execution** — written in Go  
- 🔐 **Secure** — strong password and hashing tools  
- 🧩 **Extensible** — easy to add new commands  
- 🖥 **Cross-Terminal Support** — tested in PowerShell, CMD  

---

## 💻 Installation Guide (Windows Only)

### 1. ✅ Build the CLI  
Make sure you have [Go](https://go.dev/dl/) installed.

```bash
git clone https://github.com/your-username/zyro-cli.git
cd zyro-cli
go build -o zyro.exe
```

After building, `zyro.exe` will be created in your current directory.

---

### 2. 📁 Move Binary to a System-Wide Path

Recommended: Create your own `bin` folder in your user directory.

```powershell
mkdir "$env:USERPROFILE\bin" -Force
Move-Item -Path ".\zyro.exe" -Destination "$env:USERPROFILE\bin\zyro.exe"
```

---

### 3. 🔗 Add to PATH

#### Option 1: With PowerShell

```powershell
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";$env:USERPROFILE\bin", "User")
```

#### Option 2: Manually (Recommended)

1. Press `Win + R`, type:
   ```
   rundll32 sysdm.cpl,EditEnvironmentVariables
   ```
2. Under **User variables**, select `Path` → click **Edit**
3. Click **New** and add:
   ```
   C:\Users\YOUR_USERNAME\bin
   ```
4. Click **OK** and close all dialogs.

---

### 4. 🔄 Restart Terminal & Test

Close and reopen your terminal (PowerShell or CMD) and type:

```bash
zyro -v
```

You should see:

```
ZYRO 1.0.0
```

✅ Now Zyro works from anywhere on your system!

---

## 🧠 Internal Structure

All commands are defined in a single `zyro.go` (or `main.go`) file using Go’s `os.Args` array.

### 🛠 Adding New Commands

To add a new command like `zyro uuid`, inside `main()`:

```go
case "uuid":
    fmt.Println(generateUUID())
```

Then define the helper function:

```go
func generateUUID() string
```

---

## 📌 Planned Features

- `zyro uuid` — generate UUID v4  
- `zyro timestamp` — current UNIX timestamp  
- `zyro quote` — print a random inspirational quote  
- `zyro lorem` — dummy lorem ipsum text  
- `zyro clip` — copy output to clipboard  

---

## 🧪 Build from Source

```bash
git clone https://github.com/your-username/zyro-cli.git
cd zyro-cli
go mod tidy
go build -o zyro.exe
```

Then add your features in `zyro.go`.

---

## 🤝 Contributing

Pull requests are welcome!  
Feel free to fork this repo, build your own commands, and submit your ideas to improve Zyro.

---

## 📄 License

MIT License © 2025 Elcan Bagirov  
Zyro is free to use, modify, and distribute.

---

## 🙌 Credits

Built with ❤️ using Go  
