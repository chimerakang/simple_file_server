mkdir -p /home/data/simple_file_server/files
docker build -t chimerakang/simple_file_server .
docker stop simple_file_server
docker rm simple_file_server
docker run -d --name simple_file_server --restart=always -p 2077:2077 -v /home/data/simple_file_server/files/:/go/src/github.com/chimerakang/simple_file_server/downloads chimerakang/simeple_file_server
