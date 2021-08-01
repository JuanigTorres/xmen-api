## X-Men API

App para detectar ADN mutante a peticion de Magneto

#### Dominio
El servicio es hosteado por heroku bajo el dominio
```https://xmen-finder.herokuapp.com```

#### Servicios

`/mutant` - Recibe un body con el adn como parametro. Si es correcto, retornara un status 200. Caso contrario, retornara un status 403

```bash
curl --request POST \
  --url https://xmen-finder.herokuapp.com/mutant \
  --header 'Content-Type: application/json' \
  --data '{
	"dna": ["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
  }'
```

`/stats`  - Devuelve un JSON con la cantidad de mutantnes y no-mutantes encontrados. Ademas devolvera el ratio de mutantes sobre los no-mutantes.

```bash
curl --request GET --url http://localhost:8080/stats
```

#### Testing + Covertura

Para testear el servicio inspector:
```bash
go test ./service -v
```

Para correr tests de covertura
```bash
go test ./service -cover
```

#### JMeter

Se agregaron benchmarks realizados con JMeter para ver si soportaba varias peticiones en simultaneo.
La configuracion se encuentra en `contrib/jmeter`
