# Notion Pomodoro CLI

A secure, multi-user CLI tool to track your productivity using the **Pomodoro Technique**, with seamless integration to **Notion**. Each user has an encrypted local vault to store Notion credentials, ensuring privacy even on shared machines.

---

## Features

- Start and log Pomodoro sessions from your terminal
- Secure, password-protected user registration
- Multi-user support on the same machine
- Per-user encrypted vaults (AES-GCM with password-derived key)
- Automatic logging to your Notion database
- Designed for developers and productivity nerds

---

## Installation

Clone the repository:

```bash
git clone https://github.com/jsingh0402/notion-pomodoro.git
cd notion-pomodoro
go build -o pomodoro ./cmd
