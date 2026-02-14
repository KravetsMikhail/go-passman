# Как создать релиз на GitHub (через Actions)

В релизе выкладываются **архивы** (zip для Windows, tar.gz для Linux/macOS). В каждом архиве:
- приложение с обычным именем (`go-passman.exe` или `go-passman`);
- файл **README.txt** с краткой инструкцией: первый запуск, зачем и как зашифровать хранилище, основные команды (RU/EN).

## Важно: порядок действий

Workflow запускается **при пуше тега**. Тег привязан к **конкретному коммиту**. На этом коммите уже должен быть файл `.github/workflows/release.yml`, иначе Actions не найдёт workflow и ничего не запустит.

### Пошагово

1. **Убедитесь, что workflow в репозитории**
   - Файл `.github/workflows/release.yml` закоммичен и запушен в основную ветку (например `main`).
   - На GitHub в репозитории: вкладка **Actions** → слева должен быть workflow **Release**.

2. **Права для workflow**
   - Репозиторий → **Settings** → **Actions** → **General**.
   - В блоке **Workflow permissions** выберите **Read and write permissions**.
   - Сохраните (**Save**).

3. **Создайте тег и запушьте**
   ```bash
   git checkout main
   git pull
   git tag v0.3.0
   git push origin v0.3.0
   ```
   Тег должен стоять на том коммите, где уже есть `.github/workflows/release.yml`.

4. **Проверьте запуск**
   - Откройте репозиторий на GitHub → вкладка **Actions**.
   - Должен появиться запуск **Release** по тегу `v0.3.0`. Дождитесь зелёной галочки (успех) или откройте запуск и посмотрите ошибки (красный крестик).

5. **Где искать релиз**
   - Репозиторий → **Releases** (справа от **Packages**) или `https://github.com/USER/REPO/releases`.
   - Должна появиться запись **v0.3.0** с прикреплёнными бинарниками.

## Если тег уже был запушен раньше

Тогда workflow мог не запуститься (например, на том коммите ещё не было workflow). Сделайте так:

```bash
# Удалить тег локально и на GitHub
git tag -d v0.3.0
git push origin :refs/tags/v0.3.0

# Убедиться, что workflow в репо и на нужном коммите
git status
git log -1 --oneline

# Создать тег заново и запушить
git tag v0.3.0
git push origin v0.3.0
```

После этого снова откройте **Actions** и проверьте, что **Release** запустился.

## Если workflow упал с ошибкой

- Зайдите в **Actions** → выберите упавший запуск → откройте job **release** и шаг с ошибкой.
- Частые причины:
  - **Permission denied** при создании release → включите **Read and write permissions** (шаг 2 выше).
  - **No such file** при загрузке файлов → проверьте, что шаг **Build** отработал и в `dist/release-0.3.0/` есть бинарники (логи шага Build).

## Ручная сборка (без Actions)

Локально можно собрать те же бинарники и загрузить их в релиз вручную:

```bash
./build.sh release 0.3.0
# или на Windows: build.bat release 0.3.0
```

Артефакты появятся в `dist/release-0.3.0/`. Дальше на GitHub: **Releases** → **Draft a new release** → выберите тег `v0.3.0` (или создайте его) → прикрепите эти файлы → **Publish release**.
