# 🤯 Code-LLM Helper 🚀  

Welcome to **Code-LLM Helper**, a mind-blowing CLI tool designed to help you crawl code files, merge instructions, and generate prompts for LLMs (Large Language Models). More context makes better responses, right buddy? 😏

---

## 🛠️ What It Does 🧩  

✅ **Crawl files** from a specific directory 🎯  
✅ **Filter files by extension** and ignore patterns 🛑  
✅ Merge code snippets with instructions to create `output.prompt` 🧾  
✅ Perfect for generating structured LLM prompts ✍️  

---

## 📦 `llm-helper-config.json` Explained 🔍  

The `llm-helper-config.json` file is the configuration heart of the Code Crawler. It defines:  

### **1. Directory (`dir`)**  
Specify the root directory to crawl.  

### **2. File Extension (`file_extension`)**  
Define the file types to include using patterns (`*.js|*.ts|*.kt`, etc.)  

### **3. Ignore Files (`ignore_files`)**  
Specify patterns for files or directories to ignore (supports glob patterns).  

### **Example Configuration**  

```json
{
  "dir" : "src",
  "file_extension" : "*.ts|*.js",
  "ignore_files" : "node_modules/**"
}
```

---

## 🔧 Installation 🛠️  

```bash
git clone https://github.com/your-repo/code-llm-helper.git
cd code-llm-helper
go mod tidy
```
---

## 🚀 How to Use It ⚡

**1. Prepare your config file**

Create `llm-helper-config.json` in the project directory:

```json
{
    "dir" : "src/",
    "file_extension" : "*.kt|*.ts|*.js",
    "ignore_files": "dist/**"
}
```

2. Add your instructions

Write prompt instructions in `instruct.prompt`. This file always should be top of your project.

3. Run the app like a champ 💥

```bash
go run main.go
```

4. Check the generated prompt 🎉

Open output.prompt to see the result!

### Use it as executable file

1. build main.go as executable file

```shell
go build -o code-crawler main.go
```

2. move executable file 

```shell
sudo mv code-crawler /usr/local/bin
```

3. Verify the move by running

```shell
ls -l /usr/local/bin/code-crawler
```

4. 🚀 Test the CLI Tool

```shell
code-crawler
```

5. 🧪 Optional: Create a Shell Alias (If Desired)

If you want a shorter command, you can create a shell alias in your ~/.zshrc file:

```bash
echo 'alias crawler="code-crawler"' >> ~/.zshrc
source ~/.zshrc
```

Then simply use:

```bash
crawler -projectPath /path/to/project
```
---

## 📝 Example Workflow

**Given Files and Prompts**
`llm-helper-config.json`

```json
{
  "dir" : "src",
  "file_extension" : "*.ts|*.js",
  "ignore_files" : "node_modules/**"
}
```

`instruct.prompt` 

```prompt
As a developer, analyze code! 
```

`src/index.js`
```javascript
console.log("hello");
```

`src/world.js`
```javascript
console.log("pop");
```

### 🛠 Command Execution

```bash
go run main.go --projectPath /path/to/project
```

### 📝 Generated Output (output.prompt)

```prompt
As a developer, analyze code!

```src/index.js
console.log("hello")
\```

\```
src/world.js
console.log("pop")
```

## 🤝 Contributing 👏
We love contributions! Here’s how you can help make this project even more awesome:

1. Fork the repo 🍴
2. Create your feature branch 🌵
```bash
git checkout -b my-awesome-feature
```
3. Commit your changes 🎉
```bash
git commit -m "Add awesome feature"
```
4. Push to your branch 🚀
```bash
git push origin my-awesome-feature
```
5. Open a Pull Request 🔥