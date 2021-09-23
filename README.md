<p align="center">
  <img alt="GoReleaser Logo" src="https://www.pngkey.com/png/detail/548-5489775_finance-gopher-go-lang.png" height="200" />
  <h3 align="center">GoFin</h3>
  <p align="center">Controle financeiro pessoal</p>
</p>

---

## API Requests

add transaction

```
curl -i -X "POST" "http://localhost:3000/v1/transaction" -H 'Content-Type: application/json' -H 'Accept: application/json' -d $'{
  "title": "Supermarket",
  "value":100
}'
```

list transactions

```
curl -i "http://localhost:3000/v1/transaction" -H 'Content-Type: application/json' -H 'Accept: application/json'
```

search transaction

```
curl -i "http://localhost:3000/v1/transaction?title=Supermarket" -H 'Content-Type: application/json' -H 'Accept: application/json'
```

get transaction

```
curl -i "http://localhost:3000/v1/transaction/96516731-63bc-42a8-b0e9-3263cdb29677" -H 'Content-Type: application/json' -H 'Accept: application/json'
```