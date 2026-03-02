# Технологии разработки приложений на базе фреймворков

Хранит данные в памяти процесса, предоставляет REST API и конвейер обработки запросов.

## Архитектурная идея
- **Конвейер middleware** (RequestID → Logging → Recovery → Performance) — каждый слой делает ровно одну вещь и не мешает остальным.  
- **Единый формат ошибок** с `requestId` для трассировки в логах.  
- **Валидация** прямо в обработчике создания (name не пустое, difficulty 1–5, description ≤ 500).  
- **Хранение** в памяти с защитой от параллельных запросов (RWMutex + atomic).  

## Первый запуск (локально)
```bash
# 1. Клонирование проекта
git clone https://github.com/aquamarine27/frameworks_first.git

# 2. Перейти в директорию проекта
cd frameworks_first

# 3. Установить зависимости
go mod tidy

# 4. Запустить
go run ./cmd
```

## Curl запросы
### 1. Создание задачи
```bash
curl -X POST http://localhost:8080/api/items \
  -H "Content-Type: application/json" \
  -d '{"name":"Алгоритмы","description":"Деревья и графы","difficulty":4,"isClosed":false}'
```

### 2. Список задач
```bash
curl http://localhost:8080/api/items
```

### 3. Задача по ID
```bash
curl http://localhost:8080/api/items/1
```
