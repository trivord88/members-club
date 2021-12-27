## members-club 

Latest release has been deployed to Heroku: https://members-club-test-app.herokuapp.com/

* App is written in Golang
* None of third-party libs used
* Be advised, frontend is not my strong suit, so better don't peek in
* Haven't realised how to pass errors from backend to smth like `alert()`, any try resulted in template crash

## How to run app frow scratch
``git clone -b dev git@github.com:trivord88/members-club.git`` <br>
``cd members-club`` <br>
``go run membersClub.go`` <br>
..or simply use Heroky instance :) 

Here's example of app logs from Heroku
```2021-12-27T15:19:26.470570+00:00 heroku[web.1]: State changed from down to starting
2021-12-27T15:19:26.891413+00:00 heroku[web.1]: Starting process with command `membersClub`
2021-12-27T15:19:27.697554+00:00 app[web.1]: Listening on port: 53161
2021-12-27T15:19:28.673357+00:00 heroku[web.1]: State changed from starting to up
2021-12-27T15:19:29.595220+00:00 app[web.1]: 2021/12/27 15:19:29 GET		*******:35495		/		
2021-12-27T15:19:29.595274+00:00 app[web.1]: 2021/12/27 15:19:29 "GET / HTTP/1.1\r\nHost: members-club-test-app.herokuapp.com\r\nConnection: close\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9\r\nAccept-Encoding: gzip, deflate, br\r\nAccept-Language: en-US,en;q=0.9\r\nConnect-Time: 0\r\nConnection: close\r\nDnt: 1\r\nReferer: https://dashboard.heroku.com/\r\nSec-Ch-Ua: \" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"\r\nSec-Ch-Ua-Mobile: ?0\r\nSec-Ch-Ua-Platform: \"Windows\"\r\nSec-Fetch-Dest: document\r\nSec-Fetch-Mode: navigate\r\nSec-Fetch-Site: cross-site\r\nSec-Fetch-User: ?1\r\nTotal-Route-Time: 3190\r\nUpgrade-Insecure-Requests: 1\r\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36\r\nVia: 1.1 vegur\r\nX-Forwarded-For: *******\r\nX-Forwarded-Port: 443\r\nX-Forwarded-Proto: https\r\nX-Request-Id: 987f65a6-7c78-40bb-*******c0a1a3c\r\nX-Request-Start: 1640618366403\r\n\r\n"
2021-12-27T15:19:29.595937+00:00 heroku[router]: at=info method=GET path="/" host=members-club-test-app.herokuapp.com request_id=987f65a6-7c78-40bb-88*******0a1a3c fwd="*******" dyno=web.1 connect=0ms service=0ms status=200 bytes=1333 protocol=https
2021-12-27T15:19:29.724505+00:00 app[web.1]: 2021/12/27 15:19:29 GET		*******:10796		/favicon.ico	```