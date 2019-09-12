# Githubhook (for server use)

[![Run on Google Cloud](https://storage.googleapis.com/cloudrun/button.svg)](https://console.cloud.google.com/cloudshell/editor?shellonly=true&cloudshell_image=gcr.io/cloudrun/button&cloudshell_git_repo=https://github.com/bowwowxx/githubhook.git)

### 1.Use ngrok tunnels to localhost  

### 2.Run githubhook service  
 
**go build**  
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o githubhook main.go
```

**run on server**
```
chmod +x githubhook && ./githubhook &
```

### 3.Setting github repository webhook  

 **Github page setting：**  
 <img src="./01.png" width="40%" height="40%"><img src="./02.png" width="40%" height="50%">

 **Run github hook server:**  
 <img src="./04.png" width="40%" height="40%"><img src="./05.png" width="40%" height="40%">
 
 **Check Log:**  
 ```
 tail -f /tmp/log
 ```
 <img src="./03.png" width="50%" height="50%">  


### 4.Verification Content  
HTTP POST、RequestUrl、 X-Hub-Signature、 x-github-event
Please refer to the config gile

