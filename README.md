## Ошибки и их исправления

---- 

[ ] .env
    [ ] Узнать что такое .env, почитать и прочее
    [ ] Добавить .env в проект
    [ ] Перенести порт / api-key в .env
[ ] Сделать понятный нейминг переменных, структур и функций
[ ] Добавить типизированную обработку ошибок
[ ] Разобрать и вынести в отдельный метод http client (https://pkg.go.dev/net/http) (ПОТОМУ ЧТО РЕИСПОЛЬЗУЕМЫЙ)
[ ] Почему мы постоянно считываем с диска general.html? Надо его 1 раз считать и держать в памяти - это же template
    [ ] Вынести client и template и создавать / читать 1 раз
[ ] Изучить data race | race condition | deadlock и прочие проблемы горутин и каналов 
[ ] Сделать валидацию response, типизировать его 
[ ] Изменить на html -> JSON 


store (1, 2, 3)

--------- [0 -> 6]
-> примитивы синхронизации -> | Mutex | RWmutex | Channels | Once | Atomic | Cond |
--------- [0 -> 8]