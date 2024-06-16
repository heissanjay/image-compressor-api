# Image Compression API


Steps to build locally.

> Required: Docker Installation in local machine 

1. Make sure you are in the project directory or if not `cd /image-compressor-api`
2. Run `docker build -t <image-name>:<image-tag> .` to build the docker image
3. Run `docker run -p 8080:8080 <image-name>:<image-tag>` to the application as container

Example Usage:

```
curl --location 'localhost:8080/compress' --form 'file=@"path/to/the/image"' --output 'compressed.png'
```