FROM node:16.18.1-alpine3.17 AS frontend

WORKDIR /tetua

# Копирование исходного кода фронтенда
COPY ../packages/editor .
RUN echo "Building editor..."
RUN yarn install
RUN yarn build

# Этап 2: Сборка приложения (бэкенд)
FROM golang:1.18-alpine3.17 AS backend

WORKDIR /tetua
COPY --from=frontend /tetua/dist ./packages/editor/dist
# Копирование исходного кода бэкенда
COPY ../go.mod go.sum ./
#RUN  go mod download
RUN go mod download
COPY . .
RUN echo "package views" > views/views.go
RUN go run packages/prebuild/prebuild.go --force
RUN go run . bundlestatic

RUN go build -o /bin/tetua


# Этап 3: Создание минимального образа для запуска приложения
FROM alpine:3.13

WORKDIR /tetua

# Копирование файлов бэкенда и необходимых файлов из этапа 2
COPY --from=backend /bin/tetua /tetua

# Запуск приложения
ENTRYPOINT ["./tetua"]
