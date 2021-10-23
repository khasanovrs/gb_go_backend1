https://www.statista.com/statistics/1083219/worldwide-api-performance/

- почему нужны фреймворки http-роутеров:
  1. неудобно разделять методы
  2. неудобно обрабатывать параметры
  3. неудобно организовывать мидлвары
  4. скорость маршрутизации, количество аллокаций
  5. скорость разработки

Посмотрим: gorilla/mux, go-chi, gin, fasthttp, fiber, 

https://github.com/mingrammer/go-web-framework-stars
https://github.com/smallnest/go-web-framework-benchmark

- переведем пример с прошлого урока на chi

- go get -u github.com/deepmap/oapi-codegen/cmd/oapi-codegen
//go:generate oapi-codegen -generate chi-server,spec -package api -o ./api/api.go ./api.oapi3.yaml
- https://goswagger.io/ - генерация swagger по комментариям

- gin-swagger : code first vs schema first
- https://github.com/swaggo
