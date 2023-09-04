# go-musthave-shortener-tpl

# url-shortener

Сервис для сокращения URL

## Запуск проекта

Проект может работать с БД, с файлом или хранить данные в оперативной памяти.

### Запуск с использованием БД

Для запуска в моде с БД необходимо задать переменную окружения `DATABASE_DSN` – адрес БД, либо использовать флаг при запуске `-d`

Пример:
```
DATABASE_DSN="postgres://shortener@localhost:5432/shortener?sslmode=disable" go run ./cmd/shortener/
```

### Запуск с использованием файла

Для запуска в моде с файлом необходимо задать переменную окружения `FILE_STORAGE_PATH` – путь до файла, либо использовать флаг при запуске `-f`

Пример:
```
FILE_STORAGE_PATH="/storage/storage.json" go run ./cmd/shortener/
```

### Другие флаги для запуска

`SERVER_ADDRESS` или `-a` - адрес запуска приложения

`BASE_URL` или `-b` - URL для резолва сокращённой ссылки

`LOG_LEVEL` или `-l` - уровень логгирования

## Profiler

Чтобы запустить Profiler, необходимо:

1. Запустить приложение
2. Запустить профайлер:
   1. `curl -sK -v http://localhost:8080/debug/pprof/profile > ../../profiles/cpu_base.pprof`
   2. `go tool pprof -http=":9090" -seconds=30 ../../profiles/cpu_base.pprof`
3. Создать нагрузку на приложение

## Doc

Для генерации документации необходимо выполнить:

```
godoc -http:8080
```