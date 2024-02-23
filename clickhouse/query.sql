SELECT *
FROM s3(
             'http://localhost:9000/github/*.jsonl',
             'VWOEE93slRVE7kso9tCW', 'Erk6DlqnMAbYlrsXY9kmlvQUTHx3wIJT9w04Bwdu',
             'JSONEachRow'
     )
LIMIT 1;
---
SELECT action, workflow_run.actor.events_url, workflow_run.run_started_at
FROM s3(
    'http://localhost:9000/github/*.jsonl',
    'VWOEE93slRVE7kso9tCW', 'Erk6DlqnMAbYlrsXY9kmlvQUTHx3wIJT9w04Bwdu',
    'JSONEachRow'
    ) SETTINGS output_format_json_named_tuples_as_objects=1;