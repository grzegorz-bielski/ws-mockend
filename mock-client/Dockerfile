FROM node:latest

ENV FRONTEND_PATH /mock-client
WORKDIR ${FRONTEND_PATH}

COPY . .

RUN apt-get update -q && apt-get install -yq nginx
COPY nginx.conf /etc/nginx/conf.d/nginx.conf

RUN mkdir -p ${FRONTEND_PATH}/dist
RUN npm install && npm run build && chmod 755 ${FRONTEND_PATH} && chmod 755 ${FRONTEND_PATH}/dist -R
# COPY dist /var/www/html

CMD nginx -g "daemon off;"