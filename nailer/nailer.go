package nailer

import (
    "github.com/nfnt/resize"
    "image/jpeg"
    "log"
    "os"
    "strings"
    "strconv"
    "net/http"
)

func Process(url string, width string) string {
    // convert the width string to a uint
    i, err := strconv.Atoi(width)
    if err != nil {
        log.Fatal(err)
    }

    iWidth := uint(i)
    //iWidth := uint(width[1])

    // get the file from remote
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }

    // decode jpeg into image.Image
    img, err := jpeg.Decode(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    resp.Body.Close()

    // resize to width 1000 using Lanczos resampling
    // and preserve aspect ratio
    m := resize.Resize(iWidth, 0, img, resize.Lanczos3)

    output := "public/storms-resized.jpg"

    out, err := os.Create(output)
    if err != nil {
        log.Fatal(err)
    }
    defer out.Close()

    // write new image to file
    jpeg.Encode(out, m, nil)

    return "http://localhost:3000/" + strings.TrimPrefix(output, "public/")
}
