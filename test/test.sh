#!/bin/bash

API_URL="http://localhost:8080/v1" # Замените на ваш URL

# Функция для проверки статуса ответа
check_status() {
    local expected_status=$1
    local actual_status=$2
    local message=$3
    if [ "$expected_status" -ne "$actual_status" ]; then
        echo "Ошибка: $message. Ожидалось $expected_status, получено $actual_status."
        exit 1
    fi
}

# Тест 1: Проверка пинга
echo "Тест 1: Ping"
response=$(curl -s -o /dev/null -w "%{http_code}" "$API_URL/ping")
check_status 200 "$response" "Ping не удался"

# Тест 2: Создание задачи
echo "Тест 2: Создание задачи"
create_response=$(curl -s -X POST "$API_URL/tasks" -H "Content-Type: application/json" -d '{"title": "Тестовая задача", "description": "Описание тестовой задачи", "completed": false}')
create_status=$(echo "$create_response" | jq -r '.id')
check_status 201 "$(curl -s -o /dev/null -w "%{http_code}" "$API_URL/tasks/$create_status")" "Создание задачи не удалось"

# Тест 3: Получение задач
echo "Тест 3: Получение задач"
response=$(curl -s -o /dev/null -w "%{http_code}" "$API_URL/tasks")
check_status 200 "$response" "Получение задач не удалось"

# Тест 4: Получение задачи по ID
echo "Тест 4: Получение задачи по ID"
response=$(curl -s -o /dev/null -w "%{http_code}" "$API_URL/tasks/$create_status")
check_status 200 "$response" "Получение задачи по ID не удалось"

# Тест 5: Обновление задачи
echo "Тест 5: Обновление задачи"
update_response=$(curl -s -X PUT "$API_URL/tasks/$create_status" -H "Content-Type: application/json" -d '{"title": "Обновленный заголовок", "description": "Обновленное описание", "completed": true}')
check_status 204 "$(echo "$update_response" | jq .)" "Обновление задачи не удалось"

# Тест 6: Удаление задачи
echo "Тест 6: Удаление задачи"
delete_response=$(curl -s -X DELETE "$API_URL/tasks/$create_status")
check_status 204 "$(echo "$delete_response" | jq .)" "Удаление задачи не удалось"

echo "Все тесты пройдены успешно!"