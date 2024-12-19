# Rate-limiter


1. **Run the application:**

```sh
go run main.go limiter.go
```

2.** Hit the '/api' endpoint:**

```sh
curl http://localhost:6666/api  
```

**it wil return you this response:**

```sh
StatusCode        : 200
StatusDescription : OK
Content           : {"status":"200","body":"Hello World"}

RawContent        : HTTP/1.1 200 OK
                    Content-Length: 38
                    Content-Type: application/json
                    Date: Thu, 19 Dec 2024 09:56:58 GMT

                    {"status":"200","body":"Hello World"}

Forms             : {}
Headers           : {[Content-Length, 38], [Content-Type, application/json], [Date, Thu, 19 Dec 2024 09:56:58 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 38
```

3. **Try hitting the endpoint 5 times:**

```sh
for i in {1..6}; do
  echo "Request #$i"
  curl http://localhost:6666/api
  echo ""
done
```

**It will accept request till 5 but not the 6th:**

```sh
Request #1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    38  100    38    0     0  22946      0 --:--:-- --:--:-- --:--:-- 38000{"status":"200","body":"Hello World"}


Request #2
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    38  100    38    0     0  28358      0 --:--:-- --:--:-- --:--:-- 38000{"status":"200","body":"Hello World"}


Request #3
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    38  100    38    0     0  29389      0 --:--:-- --:--:-- --:--:-- 38000{"status":"200","body":"Hello World"}


Request #4
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    38  100    38    0     0  36893      0 --:--:-- --:--:-- --:--:-- 38000{"status":"200","body":"Hello World"}


Request #5
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    38  100    38    0     0  13031      0 --:--:-- --:--:-- --:--:-- 19000{"status":"200","body":"Hello World"}


Request #6
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    82  100    82    0     0  87141      0 --:--:-- --:--:-- --:--:-- 82000{"status":"429","body":"I ain't coming into your DDOS attack you filthy animal!"}

```
