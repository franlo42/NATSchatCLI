
<div align="center"><a name="readme-top"></a>
  
  <img alt="To-Do List banner image" src="https://github.com/user-attachments/assets/532387dc-3148-414b-a991-843758d2d7e1">

# NATS chat CLI
  
  ![GitHub Created At](https://img.shields.io/github/created-at/franlo42/NATSchatCLI%20?color=%234F1787)
  ![GitHub contributors](https://img.shields.io/github/contributors/franlo42/NATSchatCLI?COLOR=%23FF6500)
  ![GitHub top language](https://img.shields.io/github/languages/top/franlo42/NATSchatCLI?color=%231230AE)
  ![Last commit](https://img.shields.io/github/last-commit/franlo42/NATSchatCLI?color=%23005B41)
  ![GitHub repo size](https://img.shields.io/github/repo-size/franlo42/NATSchatCLI?color=%23704264)

-description-
</div>

<details>
<summary><kbd>Table of Contents</kbd></summary>

#### ToC

- [Objective](#-objective)
- [Requirements](#-requirements)
- [Quick Setup](#-quick-setup)
- [API Test](#-api-test)
- [Stopping the Application](#-stopping-the-application)

</details>

## 🎯 Objective

Create a basic chat CLI using Go, talking to a NATS server.

## 📋 Requirements

1. **🦫 Go Application**
   - Accept three parameters: the address of the NATS server, the name of the chat channel, and a name (your name in the chat).
   - When run, it waits for
       - Text user input, to publish in the chat.
       - Chat messages from other users, that it shows in the command line, with an indication of the name of the user sending the message.

2. **📨 NATS server**
   - Launch a simple NATS server using docker.

3. **✨ Extra Features** (Optional)
   - Ensure that when joining a channel, we can get all messages sent during the last hour.

## ⚡ Quick Setup

## 🛑 Stopping the Application
