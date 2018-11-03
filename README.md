# animacy
Make Animation GIF easily in golang

:o: Automatic median cut and dithering

## Install
```
go get -u github.com/soniakeys/quant/median
go get -u github.com/yokoe/animacy
```

## Usage
```
gifWriter := animacy.NewWriter()
gifWriter.AppendImage(img1, 1000) // display img1 for 1000ms
gifWriter.AppendImage(img2, 1000) 
gifWriter.AppendImage(img3, 1000)

f, err := os.OpenFile("sample.gif")
defer f.Close()

if err := gifWriter.Write(f); err != nil {
    log.Fatalf(err)
}
```
