## Desc
поднятие в докере 
```shell
go mod vendor # зависимости
bash dev.sh
```

За основу взял пример с ботом, который заказывает блузки на магазе. Для примера с ценником можно сделать отдельный путь. Если планируется много разных схожих метрик - нужно определить возможные варианты структур

Компоненты у апп без жесткой связки, используются интерфейсы - позволит проще расширять/выносить функционал

### gathering app (GA)
В gathering app (GA) влетают данные и помещаются в mysql (т events_count).

Влетает такое сообщение
```go
type PayloadMessage struct {
	Label    string 'json:"label"' // описание события
	EntityID int32  'json:"id"' // id клиента, приславшего мтерику
	Value    int32  'json:"value"' // счетчик события
}
```

Структура таблицы

| entity_id | event_date | event_label | event_counter | 
| --- | --- | --- | --- |
| 5 | 2021-02-10 11:00:01 | tshirt | 123 |

Каждые 2 сек группируем. В примере меньше дня выборок нет, можно увеличить интервал группировки до минуты, немного упростит житие базы.

### analyse app (AA)
analyse app (AA) стартует и в фоне, в цикле каждые 2 сек скачивает данные из gathering.

Скачивание данных можно вынести в отдельный урл (что бы примерилось по какому-то событию извне, например cron или калькуляция какая). Каждый запрос в GA требует from, to, iterator

1. from - строка в формате "Y-m-d H:i:s". Указываем дату начала, выгрузка
1. to - опциональная строка в формате "Y-m-d H:i:s". Например, если выкачивается за незавершенный день или все, что есть до текущего момента
1. iterator - int, смещение по PK в БД. Выгрузка лимитирована в 500 строк. Если потребуется выгрузить быстро, можно отправить пачку запросов с разным смещением.

Агрегированную инфу и прогресс выгрузки можно хранить в tanantoole (не реализовал). Строк там много не ожидается, места много не займет. Будет быстрый доступ и позволит перезапускать сервис без потери состояния. 

* http://localhost:31115 - аппа по сбору данных
* http://localhost:31116 - аппа аналитическая
* http://localhost:31116/api/v1/state - посмотреть на инфу от сборщика. получение via json(лучше вкорячить protobuf), "проанализировано" через мок (internal/repository/anomaly_analysator/analysator.go:analyseData). Логин/пароль - a/b 

#### Тесты и бенчи
1. internal/repository/gather/gather_test.go - здесь задел на бенчмарки, можно написать для разных реализаций отдельный метод, на случай если текущей нагрузки много
1. internal/routes/analyser/base_test.go - задел тестирования для роутов, авторизация/ответы и т.д.

### Тестовые данные
Наполнить данным
```shell
wrk -s bench.lua -t12 -c400 -d30s http://localhost:31115/gather
hey -m POST -d '{"id":123,"label":"view","value":5}' -z 10s http://localhost:31115/gather
```
