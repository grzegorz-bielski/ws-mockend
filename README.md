### ws echo server for your mocking needs

~ work in progress

Test your websocket based application with ease. This app is developed to be used in localhost environment as a tool during development process. Dockerized and ready to **Go**.


### How to
1. Install docker and docker-compose
2. Run
```
make build-prod
make prod
```
3. Go to 'localhost:3000'. 
4. Add new route and send JSON to each client listening on that socket

### todo
- add response preview to each handler
