# Template Update Guide

Этот репозиторий использует шаблон Практикума через remote `template`.

## Проверить подключение шаблона

```bash
git remote -v
```

Ожидаемый remote:

```bash
template  https://github.com/Yandex-Practicum/go-musthave-metrics-tpl.git
```

## Получить свежие изменения из шаблона

```bash
git fetch template
```

## Посмотреть доступные ветки шаблона

```bash
git branch -r | grep template/
```

Обычно использовать стоит `template/main`, если в материалах курса не указана другая ветка.

## Посмотреть файл из шаблона без изменения текущей ветки

```bash
git show template/main:.github/workflows/mertricstest.yml
```

Примеры:

```bash
git show template/main:README.md
git show template/main:cmd/agent/main.go
git show template/up_to_iter24:.github/workflows/mertricstest.yml
```

## Безопасно обновить только GitHub workflows

Сначала посмотреть разницу:

```bash
git diff template/main -- .github
```

Потом забрать файлы:

```bash
git checkout template/main -- .github
```

После этого проверить изменения:

```bash
git diff
```

## Обновить конкретные файлы или директории из шаблона

```bash
git checkout template/main -- cmd/agent
git checkout template/main -- cmd/server
git checkout template/main -- .github
```

Это уже может перезаписать локальные изменения, поэтому сначала лучше смотреть diff.

## Как выбрать ветку шаблона

- `template/main`:
  основной и самый безопасный вариант
- специальные ветки вроде `template/up_to_iter24` или `template/v2`:
  использовать только если курс или наставник явно требует именно их
- если есть сомнения:
  сравнить нужные файлы через `git show`

## Рекомендуемый порядок обновления

```bash
git fetch template
git diff template/main -- .github
git checkout template/main -- .github
git diff
git add .github
git commit -m "Update workflows from template"
```

## Что важно в этом репозитории

- автотесты ожидают структуру, совместимую с шаблоном Практикума
- перед обновлением шаблона лучше убедиться, что локальные изменения в `.github` осознанные
- если обновление тянет внешние зависимости, стоит проверить совместимость с версией Go из CI
