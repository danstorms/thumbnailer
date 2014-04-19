package nailer

import (
    "github.com/nfnt/resize"
    "image/jpeg"
    "log"
    "os"
    "strings"
)

func Process(path string) string {
    // open "test.jpg"
    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }

    // decode jpeg into image.Image
    img, err := jpeg.Decode(file)
    if err != nil {
        log.Fatal(err)
    }
    file.Close()

    // resize to width 1000 using Lanczos resampling
    // and preserve aspect ratio
    m := resize.Resize(100, 0, img, resize.Lanczos3)

    output := path + "-resized.jpg"

    out, err := os.Create(output)
    if err != nil {
        log.Fatal(err)
    }
    defer out.Close()

    // write new image to file
    jpeg.Encode(out, m, nil)

    return strings.TrimPrefix(output, "public/")
}
