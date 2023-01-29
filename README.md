
# GiG - Core-tech 

This project involves **2 microservices** located in the same repository for ease of testing. The **communication between** **clients** and the **services** is done using the **WebSocket** **protocol**, and the **communication between** the **2 services** is done using **NATS as a message queue**. 

## Authors

- [@xavimg](https://github.com/xavimg)


## Run Locally

Clone the project

```bash
  git clone https://github.com/xavimg/The-Core-Tech-Challenge.git
```
Run this Makefile command:

```bash
  make run
```

## Test Locally

Run this Makefile command:

```bash
  make test
```

## Screenshots

![diagram](diagram.png)

## Demo using wscat

1. **IMPORTANT** Start your client subscriber demo. Go to the-core-tech-challenge/demo_subscriber and with VSCode extension *live server* open new one.
2. Open a terminal and run this command **wscat -c localhost:3010** to connect with service listener throught websocket connection for send messages.
3. Go to your terminal where we previous connect with **wscat -c localhost:3010** and you can start to write messages.
4. Receive all sended messages in client demo subscribers.

![demo](demo.png)


# Demo using browser console

1. **IMPORTANT** Start your client subscriber demo. Go to the-core-tech-challenge/demo_subscriber and with VSCode extension *live server* open new one.

2. Open Google Chrome browser and inside console write:
     ```
      let webSocket = new WebSocket('ws://localhost:3010');
      webSocket.send("test");
     ``` 
3. Receive all sended messages in client demo subscribers.

![demo](demo_browser.png)
