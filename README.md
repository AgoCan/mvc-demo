# demo
项目启动

```bash
go build -o bin/app cmd/app/main.go
./bin/app -c {config_file_path} 
```

```bash
docker build -f build/Dockerfile -t test:v-app .
docker run -it --rm -p 9000:9000 test:v-app
```