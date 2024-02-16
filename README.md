# golang-upload-file-project

## Specification of application
```
can't upload file more than 10MB
```
## Download repository

```
git clone https://github.com/khowpchom/golang-upload-file-project.git
```

## Create .env on root patch of golang-upload-file-project and copy content below
```
PORT=8000
MONGOURI=mongodb+srv://poonyaveechom:hflektGl07wsbTqd@cluster0.757cikp.mongodb.net
DB_NAME=go

MAILER_HOST=smtp.gmail.com
MAILER_USERNAME=panforsendmail@gmail.com
MAILER_PASSWORD=xoea xivn oqtj utix

SECRET=DJAnQTK2NmEc4SWwS1sFNgmtTJTy0oyE
```


## Run golang-upload-file-project as docker by command below

```
cd golang-upload-file-project
docker build --tag golang-upload-file-project . && docker run --restart=always --name golang-app -d -p 8000:8000 --env-file .env golang-upload-file-project
```

## Then can play with Postman
```
https://www.postman.com/khowpan/workspace/public-workspace/collection/12982810-a4a80068-a1b7-4e70-b8c6-0bcc816e62da?action=share&creator=12982810&active-environment=12982810-f818cc2a-b482-4b10-9867-8ff5acfe9254
```