# Polyglot Compilers Lab

Welcome to my central laboratory for language engineering. This repository is a unified effort to integrate several compiler and interpreter implementations I previously developed in separate environments. 

The goal of this project is to showcase different translation strategies—from hardware-level assemblers to high-level functional interpreters—implemented across various programming paradigms (**C++, Go, Python, Clojure, Elixir, and more**).

---

## 🚀 Project Status: Integration Phase
> **Note:** This repository is currently under active migration. I am consolidating codebases, refactoring for consistency, and establishing a unified testing and benchmarking framework.

### 🛠 To-Do List
- [ ] **Code Migration:** Moving implementations from legacy repositories.
- [ ] **Build Automation:** Ensuring every implementation compiles via automated scripts (Make, CMake, Go Modules).
- [ ] **Standardized Documentation:** Adding a dedicated `README.md` for every language and tool combination.
- [ ] **Testing Suite:** Integrating standard test suites (e.g., Nand2Tetris `.tst` files, Lox test suites).
- [ ] **Benchmarks:** Implementing a cross-language performance comparison (Execution time and Memory usage).

---

## 📂 Repository Structure

The project is organized into two primary tracks, covering the spectrum from low-level translation to high-level language semantics.

### 1. Nand2Tetris (The Translation Stack)
A journey through the Hack computer architecture translation layers.
* **Assembler:** Translates Hack Assembly to binary code.
* **VM Translator:** Stack-based translation to Hack Assembly.
* **Jack Compiler:** A syntax-directed compiler for the Jack language.
* *Implementations: C++, Go, Python (and more to come).*

### 2. Interpreters & Compilers (Language Theory)
A horizontal exploration of modern language design and implementation.
* **Lox:** Tree-walking and Bytecode VM implementations (from *Crafting Interpreters*).
* **Eva:** A Lisp-style language focusing on closures, S-expressions, and functional semantics.
* **SubC (C--):** A classic procedural compiler focusing on semantic analysis, nested scopes, and type checking.

---

## 📊 Engineering & Quality Standards

To maintain a high-quality polyglot environment, each implementation is being refactored to meet these standards:

* **Isolated Builds:** Every project contains its own build manifest (e.g., `go.mod`, `CMakeLists.txt`, `Makefile`).
* **Local Documentation:** Each sub-folder includes a local README explaining specific design decisions and how to run it.
* **Automated Validation:** Implementation of scripts to run standard test programs against each tool.
* **Performance Metrics:** Comparative benchmarks across different languages to analyze efficiency.

---

## 👨‍🏫 About the Author
I am a **Computer Engineering Professor** with a deep interest in software engineering, functional programming, and compiler design. This repository serves as both my personal engineering portfolio and a pedagogical resource for my students.

---

## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.