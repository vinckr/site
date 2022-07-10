To stop all Docker containers, simply run the following command in your terminal:
docker kill $(docker ps -q)
How to Remove All Docker Containers
docker rm $(docker ps -a -q)
How To Remove All Docker Images
docker rmi $(docker images -q)
