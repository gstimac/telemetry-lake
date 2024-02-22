SELECT *
FROM s3(
             'http://localhost:9000/clickhouse/*.jsonl',
             'VWOEE93slRVE7kso9tCW', 'Erk6DlqnMAbYlrsXY9kmlvQUTHx3wIJT9w04Bwdu',
             'JSONEachRow'
     )
LIMIT 100;
