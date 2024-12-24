COPY . /app
WORKDIR /app
RUN go mod tidy
CMD ["air"]

