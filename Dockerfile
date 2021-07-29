FROM alpine:latest as build
RUN echo "This is a docker image for landers1037/blogo"
# RUN cd html && npm install && npm run build
COPY ./blogo /app/
VOLUME ["/app/data", "/app/conf"]

EXPOSE 5000
WORKDIR /app
CMD ./blogo web start
