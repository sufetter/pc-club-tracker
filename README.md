<!DOCTYPE html>
<html lang="ru">
<body>
<div class="container">
    <h1>Инструкция по запуску</h1>
    <h2>Команда для запуска приложения с тестовыми данными</h2>
    <table>
        <tr>
            <th>Команда</th>
            <th>Описание</th>
        </tr>
        <tr>
            <td><code>go run cmd/main.go storage/club.txt</code></td>
            <td>
                    <p>Запуск с указанием пути до файла (текстовый файл содержит данные из ТЗ)</p>
                    <p>Указываются пути к "главному" go файлу и файлу с данными</p>
            </td>
        </tr>
        <tr>
            <td><code>go run cmd/main.go storage/club_huge.txt</code></td>
            <td>
                    <p>Запуск программы, что вместо данных из ТЗ используются сгенерированные</p>
            </td>
        </tr>
        <tr>
            <td><code>go test github.com/sufetter/pc-club-tracker/pkg/queue</code></td>
            <td>Запуск юнит тестов приложения для queue</td>
        </tr>
        <tr>
            <td><code>go test github.com/sufetter/pc-club-tracker/internal/club</code></td>
            <td>Запуск юнит тестов приложения для club</td>
        </tr>
    </table>
    <h2>Запуск с использованием Docker</h2>
    <p>Рассматривается использование docker-compose, т.к это база в наши дни.</p>
    <table>
        <tr>
            <th>Команда</th>
            <th>Описание</th>
        </tr>
        <tr>
            <td><code>docker-compose up</code></td>
            <td>Будет множество из 3-х контейнеров, не рекомендуется</td>
        </tr>
        <tr>
            <td><code>docker-compose up tracker</code></td>
            <td>
                    <p>Запуск с использованием конфигурационного файла docker-compose.yml</p>
                    <p>Запустится контейнер с тестовыми данными из ТЗ</p>
            </td>
        </tr>
        <tr>
            <td><code>docker-compose up hugeTracker</code></td>
            <td>Запуск с использованием других тестовых данных</td>
        </tr>
        <tr>
            <td><code>docker-compose up test</code></td>
            <td>Запуск юнит тестов приложения</td>
        </tr>
    </table>
<p>P.S. Внесены некоторые изменения после отправки письма на почту HR (фактически код в репозитории никак не изменился вообще, просто коммитов больше стало), это связано с <strike>моей криворукостью</strike> небольшой ошибкой при работе с git. Перепутал стрелочки на клавиатуре (arrowUp + enter) и выполнил команду из другого проекта - git reset --hard HEAD~4. Постарался максимально быстро все откатить в предыдущее состояние, еще раз просветился возможностями git.</p>
</div>
</body>
</html>
