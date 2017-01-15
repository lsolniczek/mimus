## MIMUS

Mimus is a simulator for HTTP-base APIs. It was created with mobile testing in mind, but it can be used successfully with any other typo of application.

###Features
---

- **Simulate** HTTP response 
- **Light** CLI application
- **Easy configuration** via JSON files

### Installation
---

### Setup
---

Before you start the first simulation, you need to create your first project.

```bash
mumus new new-project
```

Inside your user directory, the application created a config folder ```mimus-config```. You will keep inside all of your API simulations. For now, there is only one folder named ```new-project```. 

When you open your first project simulation and you will see an example JSON file which represents single case template. Each one is a separate API call you'd like to simulate.

```json
{"request":{"method":"GET","url":"/api-stub/template","headers":{"Content-Type":"application/json"}},"response":{"status":200,"bodyJSON":"{\"name\": \"Adam\",\"age\": 23}","headers":{"Content-Type":"application/json"}}}
```

###Running
---
When you want to start simulation you have to start the server with project name:

```bash
mimus run new-project
```
Local server is running under ```localhost:8080```. When you type address ```localhost:8080/api-stub/template``` in your browser you should see a json provided under ```bodyJSON``` key of ```response``` part.

Every time you add new JSON file with a new case, or you modify old one, you have to close the server and run it again.

###Contact
---

If you have any questions or suggestion, I will be more that happy when you tell me about this.

Lukasz Solniczek, l dot solniczek at gmail dot com

 
