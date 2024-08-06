# makves-test

## Сборка образов
> <span style="font-size:20px">sudo docker compose build --no-cache <br></span>
## Запуск контейнеров
> <span style="font-size:20px">sudo docker compose up -d <br></span>
## Список запущенных контейнеров
> <span style="font-size:20px">sudo docker ps <br></span>

## API
> <span style="font-size:20px">http://localhost:8084/get-items <br></span>
> > <span style="font-size:20px">Метод: GET <br></span>
> > <span style="font-size:20px">Параметры URI: id - массив чисел <br></span>
> > <span style="font-size:20px">Пример: http://localhost:8084/get-items?id=[804] <br></span>
> > <span style="font-size:20px">Удачный ответ: ```json {"success":1,"data": ...}``` <br></span>
> > <span style="font-size:20px">Неудачный ответ: ```json {"success":0,"error":"Ошибка при получении данных"}``` <br></span>

