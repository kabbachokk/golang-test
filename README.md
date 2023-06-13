Необходимо реализовать CLI утилиту и HTTP-сервер.

CLI утилита должна принимать на вход строку (без пробелов) и url, куда надо слать запрос. Оба аргумента обязательны.

HTTP-сервер обслуживает только один endpoint (пускай это будет /api/substring). В этот endpoint через CLI надо отправлять строки без пробелов. При получении данных, HTTP-сервер должен найти длину самой длинной подстроки без повторяющихся символов и затем отправить ответ, который выведет CLI утилита. Пример: abcabcbb -> abc, bbbb -> b, pwwkew -> wke.

Использовать язык программирования Golang. Допускается использование сторонних пакетов, но без них будет больше баллов за выполнение задания.

БОНУС (доп баллы):

1.  Makefile (что туда добавить нужно придумать самому)
2.  Unit-тесты
3.  Использование Docker
4.  Без использования сторонних пакетов

```sh
PS > make
build                          сборка
help                           помощь
run-client                     запуск клиента
run-server                     запуск сервера
test                           тесты
```

```sh
make build cmd=apiserver
make build cmd=apiserver-client
```

```sh
make run-client string=abcabcbb url=http://localhost:8080/api/substring
make run-server
```

```sh
PS > apiserver-client -s abcabcbb -u http://localhost:8080/api/substring
3
```
