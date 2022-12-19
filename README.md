### Running application

```
go run main.go
```

### Get Movie Details

```
curl localhost:8080/movie/5584D1F9-D627-4205-BDF5-68A541F1BD85 | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   392  100   392    0     0   6758      0 --:--:-- --:--:-- --:--:--  7840
{
  "_id": "5584D1F9-D627-4205-BDF5-68A541F1BD85",
  "name": "Tidal Wave (English dub)",
  "title": "Tidal Wave (English dub)",
  "rating": "R",
  "slug": "tidal-wave-english-dub",
  "description": "A deep-sea earthquake occurs, creating a tidal wave that is headed straight for Haeundae, a popular vacation spot on the south coast of Korea, which draws visitors from all over the world.",
  "programming_type": "movie"
}
```

### Get Series Details

```
curl localhost:8080/series/227336A3-CA32-4DD2-BDC9-6221C36DF5B9 | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   296  100   296    0     0   1903      0 --:--:-- --:--:-- --:--:--  2000
{
  "_id": "1029E3AB-AE97-43BE-A0E9-180D9BA5E688",
  "name": "Tough Day at The Office",
  "description": "Think your day at work was hard? Well check out the nightmare scenarios facing these men and woman as their 9-5 shifts turn into a race to survive. Record-breaking snow, torrential rainfall, and even "
}
```

#### Quick Play Endpoints

[Movie](/content/urn/resource/catalog/movie/%s?reg=us&dt=androidmobile&client=amd-localnow-web)
[Series](/content/series/%s/episodes?reg=us&dt=androidmobile&client=amd-localnow-web&seasonId=00FFFEBA-9E34-4C3E-99F5-D6D814403FD5&pageNumber=1&pageSize=10&sortBy=ut&st=published)

#### Libraries Used

[Config](https://github.com/spf13/viper)
[Logging](https://github.com/sirupsen/logrus)
[Used To Map QuickPlay response keys to rails](https://github.com/tidwall/gjson)
[HTTP Requests](https://pkg.go.dev/net/http)
[gock Mock HTTP Requests](https://github.com/h2non/gock)

#### Test Coverage

```
❯ go test -cover ./...
?   	github.com/subbarao/transformer	[no test files]
?   	github.com/subbarao/transformer/pkg/api	[no test files]
ok  	github.com/subbarao/transformer/pkg/server	(cached)	coverage: 72.7% of statements
ok  	github.com/subbarao/transformer/pkg/transform	(cached)	coverage: 86.2% of statements
```
