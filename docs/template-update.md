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

## Обязательная проверка после обновления `.github`

После команды:

```bash
git checkout template/main -- .github
```

нужно отдельно проверить [`mertricstest.yml`](../.github/workflows/mertricstest.yml).

В этом репозитории шаг:

```bash
go install golang.org/x/tools/cmd/goimports@latest
```

ломает CI, потому что `@latest` сейчас требует Go `>= 1.25`, а workflow Практикума запускается на Go `1.24`.

Рабочий вариант для этого репозитория:

```bash
go install golang.org/x/tools/cmd/goimports@v0.42.0
```

Если после обновления шаблона снова вернулся `@latest`, его нужно сразу заменить обратно на `@v0.42.0` перед коммитом.

## Что важно в этом репозитории

- автотесты ожидают структуру, совместимую с шаблоном Практикума
- перед обновлением шаблона лучше убедиться, что локальные изменения в `.github` осознанные
- если обновление тянет внешние зависимости, стоит проверить совместимость с версией Go из CI
