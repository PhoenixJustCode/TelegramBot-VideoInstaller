# 🎮 TelegramBot InstallerVideo

## 🚀 Описание

**TelegramBot InstallerVideo** — это мультифункциональный Telegram-бот на **Go**, который принимает видео (**mp4** или ссылку) и автоматически конвертирует его в:

- 🎵 Аудио-файл (`.mp3`)
- 🔣 Голосовое сообщение (`.ogg` / Opus)

После обработки оригинальное видео удаляется для экономии пространства.  
Проект построен с учётом масштабируемости, многопоточности и развёртки в Docker.

---

## 🛠 Технологии

- **Go** — основной язык разработки
- **Telegram Bot API** — интеграция с Telegram
- **FFmpeg** — обработка и конвертация медиа
- **Logrus** — мощное логирование
- **goroutines** — многопоточность
- **envconfig** — конфигурация через `.env`
- **Docker & Docker Compose** — контейнеризация
- **GitHub Actions** — CI/CD

---


## 🔗 Возможности

✅ Приём mp4 или ссылки на видео  
✅ Скачивание видео
✅ Конвертация mp4 → mp3  
✅ Конвертация mp4 → ogg (Opus) для Telegram voice  
✅ Автоматическая отправка пользователю  
✅ Удаление временных файлов  
✅ Логирование через Logrus  
✅ Docker-ready  
✅ CI/CD pipeline  

---

## ⚙️ Установка и запуск

### 1️⃣ Клонирование проекта

```bash
git clone https://github.com/YourUsername/telegramBot-installervideo.git
cd telegramBot-installervideo  
```

---

### 2️⃣ Сборка и запуск через Makefile

make build        # Собрать приложение
make run          # Запустить приложение
make lint         # Проверка кода
make clean         # Очистка бинарника  
make deploy       # CI/CD деплой


### 3️⃣Docker запуск

```bash
docker-compose up --build -d
```

---


💡 Как использовать

    Открой Telegram и отправь боту mp4-файл или ссылку на видео.

    Получи mp3-файл и голосовое сообщение.

    Готово!

---

👨‍💻 Авторы и вклад

Разработка: PhoenixJustCode(Alexandr)
PR и issue welcome! 🙌