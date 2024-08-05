## PSQL

- psql - консольная утилита.
- PostgresSQL - база данных.
- pgAdmin - GUI для работы с базой данных.

[Тренажёр для начинающих]()


___
N1
Найти самый поздний логин в 2020 году

```sql
SELECT user_id, MAX(time_stamp) AS last_stamp
FROM Logins
WHERE TO_CHAR(time_stamp, 'YYYY') = '2020'
GROUP BY user_id;
```
___
N2
CASE операция
```sql
SELECT DISTINCT stock_name, 
SUM(CASE 
WHEN operation = 'Sell' THEN price ELSE -price
END) as capital_gain_loss 
FROM Stocks
GROUP BY stock_name
```
___
N3
Записать значения в одну строку
```sql
SELECT DISTINCT sell_date, COUNT(DISTINCT product) AS num_sold, 
string_agg(DISTINCT(product), ',') as products
FROM Activities
GROUP BY sell_date
ORDER BY sell_date ASC
```
___
N4
Парная обработка данных

```sql
SELECT DISTINCT machine_id, ROUND(AVG(
    CASE 
    WHEN activity_type = 'start' THEN -timestamp ELSE timestamp
    END
)::decimal * 2, 3) AS processing_time
FROM Activity
GROUP BY machine_id
ORDER BY machine_id ASC
```

___
N5
Добавить элемент если столбец пустой

```sql
SELECT 
    CASE 
        WHEN MAX(num) IS NULL THEN NULL
        ELSE MAX(num)
    END AS num FROM ...
```
___
N6 Конкатенация строк + UPPER/lower

```sql
(UPPER(SUBSTRING(name, 1, 1)) || 
LOWER(SUBSTRING(name, 2, LENGTH(name)))) 
AS name
```
___
N7
Таблицы можно сшивать:

```sql
SELECT product_id, 'store1' as store, store1 as price
    FROM Products
    UNION
    SELECT product_id, 'store2' as store, store2 as price
    FROM Products
    UNION
    SELECT product_id, 'store3' as store, store3 as price
    FROM Products
```
___
N8 value в столбце
```sql
SELECT exists (
    SELECT 1 FROM Employees WHERE employee_id = 3 LIMIT 1
);
```
___