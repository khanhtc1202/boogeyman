FROM nginx:mainline-alpine
COPY ./deployments/web/default.nginx.conf /etc/nginx/conf.d/default.conf
COPY ./web/static /app/
