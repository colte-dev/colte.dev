FROM nginx:alpine

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]