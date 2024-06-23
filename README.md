# docker-micro-service
docker-micro-service


docker build -t gyanendra371/node-task-app  ./tasks-api

docker build -t gyanendra371/node-node-user ./users-api

docker run -d --name auth --network mynetwork -p 80:80 gyanendra371/node-auth-app

docker run -e TASKS_FOLDER=tasks -e AUTH_ADDRESS=auth -d --name task --network mynetwork -p 8000:8000 gyanendra371/node-task-app

docker run -e AUTH_SERVICE_SERVICE_HOST=auth -e AUTH_ADDRESS=auth -d --name user --network mynetwork -p 8080:8080 gyanendra371/node-user-app


docker run -it -p 8080:80 --network mynetwork gyanendra371/node-react-app


docker compose build

docker compose build


docker compose stop

docker comoose down -v