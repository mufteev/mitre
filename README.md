# mitre

**mitre** — парсер [Enterprise-Attack](https://github.com/mitre-attack/attack-stix-data/tree/master/enterprise-attack) в STIX-формате


2 вида представления информации из STIX:
- **Иерархическая** - получение вложенной структуры Тактика -> Техника -> Подтехника, подойдёт для программной логики
  ```go
  import "github.com/mufteev/mitre/attack/stix/hierarchy"
  // ...
  // Получение данных из памяти
  hierarchy.LoadFromByteAssociate(ctx context.Context, data []byte) ([]*Tactic, error)
  // ...
  // Получение даных из stream
  hierarchy.LoadFromReaderAssociate(ctx context.Context, r io.Reader) ([]*Tactic, error)
  // ...
  ```
- **Плоская** - получение списочной структуры Тактика + Техника + СвязьТехникаТактика, подойдёт для сохранения в БД
  ```go
  import "github.com/mufteev/mitre/attack/stix/flat"
  // ...
  // Получение данных из памяти
  flat.LoadFromByteAssociate(ctx context.Context, data []byte) ([]*stix.Tactic, []*Technique, []*TechniqueTactic, error)
  // ...
  // Получение даных из stream
  flat.LoadFromReaderAssociate(ctx context.Context, r io.Reader) ([]*stix.Tactic, []*Technique, []*TechniqueTactic, error)
  // ...
  ```