# Instruction

## Docker Image

Pull
```
[vagrant@localhost ~]$ docker pull swaggerapi/swagger-editor
Using default tag: latest
Trying to pull repository docker.io/swaggerapi/swagger-editor ... 
latest: Pulling from docker.io/swaggerapi/swagger-editor
acb474fa8956: Pull complete 
d85d35d2a58f: Pull complete 
22feacb2d7ab: Pull complete 
cac20b0fe542: Pull complete 
09ce9a13a8b4: Pull complete 
8094c52b1362: Pull complete 
8db039060181: Pull complete 
fb2c32a20725: Pull complete 
Digest: sha256:8fb0b96d1ea5b128005f43bbb005d4de4c205a2b3b721913738cf0ae364fe9aa
Status: Downloaded newer image for docker.io/swaggerapi/swagger-editor:latest
```

List
```
[vagrant@localhost ~]$ docker images swaggerapi/*
REPOSITORY                            TAG                 IMAGE ID            CREATED             SIZE
docker.io/swaggerapi/swagger-editor   latest              83a50f7be45c        16 hours ago        10.11 MB
```

## Docker run

Run
```
[vagrant@localhost ~]$ docker run -d -p 28080:8080 docker.io/swaggerapi/swagger-editor
427c58fe6ac3b97bc9e2df1b17102a1125e5ae9869487944f66765e56476af65
```

Show
```
[vagrant@localhost ~]$ docker ps -l
CONTAINER ID        IMAGE                                 COMMAND                  CREATED             STATUS              PORTS                     NAMES
427c58fe6ac3        docker.io/swaggerapi/swagger-editor   "sh /usr/share/nginx/"   4 minutes ago       Up 4 minutes        0.0.0.0:28080->8080/tcp   gloomy_darwin
```

## Edit

Import go-to-openstack-bootcamp/kopos/echopb/service.swagger.json

![屏幕快照 2017-06-24 下午2.20.25.png](./屏幕快照%202017-06-24%20下午2.20.25.png)
 
