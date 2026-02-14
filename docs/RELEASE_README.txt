================================================================================
  go-passman — password manager (first run and security)
================================================================================

ON FIRST RUN
------------
When you run go-passman for the first time, it creates a vault file (vault.json)
in the SAME folder as the program. This vault is NOT encrypted — passwords
are stored in plain text until you encrypt it.

RECOMMENDED: Encrypt the vault right after adding your first password
------------
  1. Add a password:    go-passman add
     (or: go-passman add --generate  for a generated password)

  2. Encrypt the vault: go-passman encrypt
     Set a strong master password and confirm it.

  3. From now on, any command (list, add, copy, etc.) will ask for the master
     password once — the vault is protected.

QUICK COMMANDS
--------------
  go-passman add              Add password (manual)
  go-passman add -g           Add with generated password
  go-passman list             List entries (numbered)
  go-passman copy <name>      Copy password to clipboard (e.g. copy 2)
  go-passman update           Update an entry (select from list)
  go-passman remove           Remove an entry (select from list)
  go-passman status           Show vault status (encrypted or not)
  go-passman -w               Start web interface (http://127.0.0.1:8080)
  go-passman --help           Show all commands

WHERE IS THE VAULT?
------------------
The vault file (vault.json) is in the same directory as go-passman.
Run: go-passman path  to see the full path.

================================================================================
  go-passman — менеджер паролей (первый запуск и защита)
================================================================================

ПРИ ПЕРВОМ ЗАПУСКЕ
------------------
При первом запуске go-passman создаёт файл хранилища (vault.json) в той же
папке, где лежит программа. Хранилище создаётся БЕЗ шифрования — пароли
хранятся в открытом виде, пока вы его не зашифруете.

РЕКОМЕНДАЦИЯ: зашифруйте хранилище сразу после добавления первого пароля
------------------
  1. Добавьте пароль:   go-passman add
     (или: go-passman add --generate  для сгенерированного пароля)

  2. Зашифруйте хранилище: go-passman encrypt
     Задайте надёжный мастер-пароль и подтвердите его.

  3. Далее при любой команде (list, add, copy и т.д.) программа один раз
     запросит мастер-пароль — хранилище защищено.

ОСНОВНЫЕ КОМАНДЫ
----------------
  go-passman add              Добавить пароль (вручную)
  go-passman add -g            Добавить со сгенерированным паролем
  go-passman list             Список записей (с номерами)
  go-passman copy <имя>       Скопировать пароль (например copy 2)
  go-passman update           Изменить запись (выбор из списка)
  go-passman remove           Удалить запись (выбор из списка)
  go-passman status           Статус хранилища (зашифровано или нет)
  go-passman -w               Веб-интерфейс (http://127.0.0.1:8080)
  go-passman --help           Все команды

ГДЕ ХРАНИЛИЩЕ?
--------------
Файл хранилища (vault.json) лежит в той же папке, что и go-passman.
Команда: go-passman path  покажет полный путь.
