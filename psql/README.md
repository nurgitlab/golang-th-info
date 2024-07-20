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