### Running application

```
go run main.go
```

### Get Movie Details

```
 curl localhost:8080/movie/5584D1F9-D627-4205-BDF5-68A541F1BD85 | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   285  100   285    0     0  11322      0 --:--:-- --:--:-- --:--:-- 17812
{
  "_id": "5584D1F9-D627-4205-BDF5-68A541F1BD85",
  "name": "Tidal Wave (English dub)",
  "description": "A deep-sea earthquake occurs, creating a tidal wave that is headed straight for Haeundae, a popular vacation spot on the south coast of Korea, which draws visitors from all over the world."
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
