
<div align="center"><a name="readme-top"></a>
  
  <img alt="To-Do List banner image" src="https://github.com/user-attachments/assets/532387dc-3148-414b-a991-843758d2d7e1">

# NATS chat CLI
  
  ![GitHub Created At](https://img.shields.io/github/created-at/franlo42/NATSchatCLI%20?color=%234F1787)
  ![GitHub contributors](https://img.shields.io/github/contributors/franlo42/NATSchatCLI?COLOR=%23FF6500)
  ![GitHub top language](https://img.shields.io/github/languages/top/franlo42/NATSchatCLI?color=%231230AE)
  ![Last commit](https://img.shields.io/github/last-commit/franlo42/NATSchatCLI?color=%23005B41)
  ![GitHub repo size](https://img.shields.io/github/repo-size/franlo42/NATSchatCLI?color=%23704264)

The NATS Chat CLI application allows users to join chat rooms and exchange messages in real time via terminal. Messages are persisted for one hour using NATS JetStream, ensuring they can be replayed when a user joins a channel. It’s designed to be minimal yet powerful, leveraging NATS for efficient and reliable communication.
</div>

<details>
<summary><kbd>Table of Contents</kbd></summary>

#### ToC

- [Objective](#-objective)
- [Requirements](#-project-requirements)
- [Quick Setup](#-quick-setup)
- [Stopping the Application](#-stopping-the-application)

</details>

## 🎯 Objective

Create a basic chat CLI using Go, talking to a NATS server.

## 📋 Project Requirements

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

1. Clone the repository.
```shell
git clone https://github.com/franlo42/NATSchatCLI.git
cd NATSchatCLI
```

2. Run the setup shell script for an automatic deploy of the environment.
```shell
./build-and-extract.sh
```
> [!NOTE]  
> Optionally you can also compile the main Go program by hand if you prefer rather than running the setup shell script.
> ```bash
> docker-compose down -v 
> docker-compose up --build
> ```
> You will also need to run the NATS server manually.
> ```bash
> docker-compose up -d
> ```

3. Start the client chat application.
```shell
./nats_chat_linux -a nats://localhost:4222 -c chatRoom1 -n User1
```
> [!IMPORTANT]  
> You must provide the application with the following parameters.
> 
> -a \<adress of the NATS server>
> 
> -c \<channel name>
>
> -n \<your user name>

## 🛑 Stopping the Application

1. Stop the client application

You can easily stop the process by pressing the key combination `CTRL+c` in the terminal

2. Stop the NATS server container.
```bash
docker-compose down
```
> [!TIP]
> If you want to restart the NATS server:
> ```bash
> docker-compose up --d
> ```
