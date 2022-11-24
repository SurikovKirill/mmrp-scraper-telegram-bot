# Инструкция по запуску
## 1. Для начала нужно установить движок для rod

### На Ubuntu:

`wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb`

`apt install ./google-chrome-stable_current_amd64.deb`

### На alpine:

`apk add chromium`

## 2. Прописать в env данные

### Для тестирования в ***.env_test***

### Для прода в ***.env***

## 3. Выполнить make - инструкции

# Инструкция для тестирования

## Если запускать через go run

### В mapm.go раскоментить логопас для тестов

### В конфиге для телеги раскоментить логопас для тестового бота

## Если запускать через докер

### Создать скрытый файл ***.env_test***

### Выполнить make-инструкцию для сборки образа

### Запустить контейнер через ***make start-test-container***


