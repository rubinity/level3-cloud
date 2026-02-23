curl -i -c cookies.txt \
  -H "Content-Type: application/json" \
  -d '{"namespace":"test2","password":"pass"}' \
  http://188.34.109.28/auth


  curl -v -c cookie.txt \
  -H "Content-Type: application/json" \
  -X POST http://188.34.109.28/auth \
  -d '{"namespace":"test2","password":"pass"}'


  curl -i -b cookies.txt \
  -H "Content-Type: application/json" \
  -d '{"title":"write secure JWT blog"}' \
  http://localhost:8080/todo

  curl -i -b cookies.txt \
  http://188.34.109.28/connection/test2/repl2

curl -v -c cookie.txt -X POST http://188.34.109.28/auth
