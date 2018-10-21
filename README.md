# Getting Starting
 ## About the Project
This project is Application Word collection app.
Word reference is `応用情報技術者試験ドットコム https://www.ap-siken.com/`.
 ## Usage
 ```
git clone https://github.com/kosa3/app-info-words.git
```
 ```
// start mysql container & phpMyAdmin container
$ docker-compose up -d
```
 ```
// set GOPATH
$ export GOPATH=$(pwd)/app-info-words/goapi
```
 ```
$ cd goapi
$ go build
$ ./goapi
```
 First, you should make database data by sql.
But, you must not need typing sql command.
You should access this url.
 `http://localhost:8060/initialize`
 If you accessed this url, you see the display font `initialize!`
 Second, You access this url.
This url is words data from Rest API.
 `http://localhost:8060/api/words`
 Finally, You should set another project.
 ```
$ cd ../vue-app
$ npm install
$ npm run serve
```
 You can be confirmed this project running.
Let's see the this url `http://localhost:8081/`