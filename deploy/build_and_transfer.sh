# Собрать Docker-образ
#если произойдет ошибка выполнение скрипта прервется
set -e
docker build -t my-golang-app -f deploy/Dockerfile .
#docker build -t my-golang-app -f deploy/Dockerfile --build-arg project=../ .
# Сохранить Docker-образ в архив
docker save -o my-golang-app.tar my-golang-app
# Копировать архив на удаленный сервер
scp  deploy/.env kate@192.168.1.4:
scp deploy/deploy-docker-compose.yml kate@192.168.1.4:
scp my-golang-app.tar kate@192.168.1.4:
# Загружаем в доккер образ на виртуальной машине
#остановка доккера с приложением
#запуск доккер компоуза в фоновом режиме
#Получаем ответ о том что команды загружены успешно,
#или сообщение об ошибке
ssh kate@192.168.1.4 '
    set -e
    docker load -i my-golang-app.tar;
    docker compose -f deploy-docker-compose.yml down tetua;
    docker compose -f deploy-docker-compose.yml up -d;
    '
