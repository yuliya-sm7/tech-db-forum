# tech-db-forum

Тестовое задание для реализации проекта "Форумы" на курсе по базам данных в Технопарке Mail.ru (https://park.mail.ru).

Суть задания заключается в реализации API к базе данных проекта «Форумы» по документации к этому API.

Таким образом, на входе:

 * документация к API;

На выходе:

 * репозиторий, содержащий все необходимое для разворачивания сервиса в Docker-контейнере.

## Как запустить
```
docker build -t smirnova -f Dockerfile.golang .
docker run -p 127.0.0.1:5000:5000 -ti smirnova
// Tests:
./tech-db-forum func -u http://localhost:5000/api -r report.html
```

## Документация к API
Документация к API предоставлена в виде спецификации [OpenAPI](https://ru.wikipedia.org/wiki/OpenAPI_%28%D1%81%D0%BF%D0%B5%D1%86%D0%B8%D1%84%D0%B8%D0%BA%D0%B0%D1%86%D0%B8%D1%8F%29): swagger.yml

Документацию можно читать как собственно в файле swagger.yml, так и через Swagger UI (там же есть возможность поиграться с запросами): https://tech-db-forum.bozaro.ru/

## Требования к проекту
Проект должен включать в себя все необходимое для разворачивания сервиса в Docker-контейнере.

При этом:

 * файл для сборки Docker-контейнера должен называться Dockerfile и располагаться в корне репозитория;
 * реализуемое API должно быть доступно на 5000-ом порту по протоколу http;
 * допускается использовать любой язык программирования;
 * крайне не рекомендуется использовать ORM.

Контейнер будет собираться из запускаться командами вида:
```
docker build -t a.navrotskiy https://github.com/bozaro/tech-db-forum-server.git
docker run -p 5000:5000 --name a.navrotskiy -t a.navrotskiy
```

В качестве отправной точки можно посмотреть на примеры реализации более простого API на различных языках программирования: https://github.com/bozaro/tech-db-hello/

## Функциональное тестирование
Корректность API будет проверяться при помощи автоматического функционального тестирования.

Методика тестирования:

 * собирается Docker-контейнер из репозитория;
 * запускается Docker-контейнер;
 * запускается скрипт на Go, который будет проводить тестирование;
 * останавливается Docker-контейнер.

Скомпилированные программы для тестирования можно скачать по ссылкам:

 * [darwin_amd64.zip](https://bozaro.github.io/tech-db-forum/darwin_amd64.zip)
 * [linux_amd64.zip](https://bozaro.github.io/tech-db-forum/linux_amd64.zip)
 * [windows_386.zip](https://bozaro.github.io/tech-db-forum/windows_386.zip)
 * [windows_amd64.zip](https://bozaro.github.io/tech-db-forum/windows_amd64.zip)

Для локальной сборки Go-скрипта достаточно выполнить команду:
```
go get -u -v github.com/bozaro/tech-db-forum
go build github.com/bozaro/tech-db-forum
```
После этого в текущем каталоге будет создан исполняемый файл `tech-db-forum`.

### Запуск функционального тестирования

Для запуска функционального тестирования нужно выполнить команду вида:
```
./tech-db-forum func -u http://localhost:5000/api -r report.html
```

Поддерживаются следующие параметры:

Параметр                              | Описание
---                                   | ---
-h, --help                            | Вывод списка поддерживаемых параметров
-u, --url[=http://localhost:5000/api] | Указание базовой URL тестируемого приложения
-k, --keep                            | Продолжить тестирование после первого упавшего теста
-t, --tests[=.*]                      | Маска запускаемых тестов (регулярное выражение)
-r, --report[=report.html]            | Имя файла для детального отчета о функциональном тестировании
