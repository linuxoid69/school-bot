# school

Данный бот получает оценки с сайта школы и отправляет их в телеграм.

Для работы бота требуется:

- Телеграм-токен, который можно получить в @BotFather
- Телеграм группа, куда будут отправляться сообщения
- Учетные данные от сайта школы.

Для работы бота требуется следующие переменные окружения:

```bash
SCHOOL_JWT             # jwt токен для сайта школы
SCHOOL_URL="https://dnevnik2.petersburgedu.ru/api/journal/estimate/table"  # url сайта школы
SCHOOL_EUCATION_ID     # id учащегося
SCHOOL_TOKEN           # телеграм токен
SCHOOL_CHAT_ID         # id чата телеграм
SCHOOL_CRON_WORK_WEEK  # cron выражение когда будут опрашиваться данные со школы
SCHOOL_USER_AGENT      # user-agent для запросов к сайту школы
```

В планах:

- Переделать получение JWT токена на автоматическое получение (так как сейчас он протухает)
- Сделать SCHOOL_USER_AGENT по умолчанию заданным "Mozilla/5.0 (X11; Linux x86_64)"
- Получение id учащегося по имени
