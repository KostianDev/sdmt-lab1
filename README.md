# Quadratic Equation Solver

Цей застосунок розв'язує квадратні рівняння виду ax^2 + bx + c = 0. Підтримує два режими: інтерактивний для ручного вводу параметрів, неінтерактивний для зчитування параметрів з файлу.

## Як зібрати та запустити проект

1. [Встановіть Go](https://golang.org/doc/install)
2. Клонуйте репозиторій за допомогою [git](https://git-scm.com/downloads):
  ```sh
  $ git clone <repo-url>
  ```
3. Перейдіть до директорії проєкту:
  ```sh
  $ cd sdmt-lab1
  ```
5. Зберіть проєкт:
  ```sh
  $ go build .
  ```
6. Запустіть програму:
   - Для інтерактивного режиму:
     ```sh
     $ ./lab1
     ```
   - Для неінтерактивного режиму:
     ```sh
     $ ./lab1 <file-path>
     ```

## Формат файлу для неінтерактивного режиму

Файл повинен містити три числа, розділені символом пробілу(\s), що закінчуються переходом на нову строку(\n), наприклад: `1\s5\s3\n`

## Revert-коміт

Revert-коміт можна знайти за ідентифікатором: `34bccc7be70392e7d737c3039c8b3d6298a13f0f`