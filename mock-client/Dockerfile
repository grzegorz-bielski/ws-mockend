FROM node:latest

ENV FRONTEND_PATH /mock-client
WORKDIR ${FRONTEND_PATH}

COPY . .

RUN apt-get update -q && apt-get install -yq nginx

RUN npm install && npm run build
RUN chmod 755 ${FRONTEND_PATH} && chmod 755 ${FRONTEND_PATH}/dist -R

COPY dist /var/www/html
CMD nginx -g "daemon off;"