## Создание и открытие файла

Используем пакет [os](https://golang.org/pkg/os/ "golang пакет os")

### Создать файл

```go
file, err := os.Create("имя файла")
```

### Открыть файл

```go
file, err := os.Open("имя файла")
```

### Открыть или создать при отсутствии файла

```go
file1, err := os.OpenFile("file1.txt", os.O_RDONLY, 0666)

file2, err := os.OpenFile("file2.txt", os.O_WRONLY, 0666)
```

После окончания работы с файлом - нужно закрыть его - метод ```Close()```
