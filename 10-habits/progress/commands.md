# Commands

## Create a habit

```shell
grpcurl \
-import-path api/proto/ \
-proto service.proto \
-plaintext -d '{"name":"read a few pages", "weekly_frequency":3}' \
localhost:28710 \
habits.Habits/CreateHabit
```


```json
{
  "habit": {
    "id": "df990e0b-5825-460d-86d2-18dcec19adeb",
    "name": "read a few pages",
    "weeklyFrequency": 3
  }
}
```


## Tick a habit

```shell
grpcurl \
-import-path api/proto/ \
-proto service.proto \
-plaintext -d '{"id":"df990e0b-5825-460d-86d2-18dcec19adeb"}' \
localhost:28710 \
habits.Habits/TickHabit
```

```json
{}
```

## Get the status

```shell
grpcurl \
-import-path api/proto/ \
-proto service.proto \
-plaintext -d '{"id":"df990e0b-5825-460d-86d2-18dcec19adeb"}' \
localhost:28710 \
habits.Habits/GetHabitStatus
```

```json
{
  "habit": {
    "id": "df990e0b-5825-460d-86d2-18dcec19adeb",
    "name": "read a few pages",
    "weeklyFrequency": 3
  },
  "ticksCount": 2
}
```