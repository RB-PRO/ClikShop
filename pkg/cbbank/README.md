# Центральный банк РФ

Прослойка для получения курса валют

### Получить курс доллара $

```go
cb, ErrorCB := cbbank.New()
if ErrorCB != nil {
	t.Error(ErrorCB)
}
fmt.Println("Курс доллара", cb.Data.Valute.Usd.Value)
```

[PostMan](https://app.getpostman.com/join-team?invite_code=ba77616c0d844d73f78a8e1fddfa56a2&target_code=c50fa27f70575cf6b245a692b47b9be8)
