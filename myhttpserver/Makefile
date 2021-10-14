export tag=v1.0
#myHttpserver,要改成小写的。否则docker的时候会报错：
#invalid argument "golang-master/myHttpserver:v1.0" for "-t, --tag" flag: invalid reference format: repository name must be lowercase

root:
	export ROOT=github.com/golang-master/myhttpserver

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t golang-master/myhttpserver:${tag} .

push: release
	echo "modify tag name"
	#1、直接用上面的镜像包上传不成功的，要规范命名镜像
	#beigegede:username;golang-master:repository;image name: myhttpserver${tag}
	docker tag golang-master/myhttpserver:${tag} beigegede/golang-master:myhttpserver${tag}
	echo "pushing golang-master/myhttpserver"
	#2、要push成功，需要先登录hub.docker。命令是docker login
	docker push beigegede/golang-master:myhttpserver${tag}
