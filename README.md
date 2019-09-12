# Githubhook (for server use)

[![Run on Google Cloud](https://storage.googleapis.com/cloudrun/button.svg)](https://console.cloud.google.com/cloudshell/editor?shellonly=true&cloudshell_image=gcr.io/cloudrun/button&cloudshell_git_repo=https://github.com/bowwowxx/githubhook.git)

### 1.Use ngrok tunnels to localhost

### 2.Run githubhook service
 

- go build
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o githubhook main.go
```

- run on server
```
chmod +x githubhook && ./githubhook &
```

### 3.Setting github repository webhook

 **Github page setting：**
  ![mole](https://github.com/bowwowxx/githubhook/blob/master/01.png)  
  ![mole](https://github.com/bowwowxx/githubhook/blob/master/02.png)

 **Run github hook server:**  
  ![mole](https://github.com/bowwowxx/githubhook/blob/master/04.png)  
  ![mole](https://github.com/bowwowxx/githubhook/blob/master/05.png)    
 
 **Check Log:**
 ```
 tail -f /tmp/log
 ```
  ![mole](https://github.com/bowwowxx/githubhook/blob/master/03.png) 

### 4.Verification Content
HTTP POST、RequestUrl、 X-Hub-Signature、 x-github-event
Please refer to the config gile

