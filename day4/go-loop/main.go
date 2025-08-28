package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"golang.org/x/image/draw"
)

// luminance returns grayscale value 0..255
func luminance(c color.Color) uint8 {
	r, g, b, _ := c.RGBA()
	// r,g,b are 0..65535, convert to 0..255
	rr := float64(r) / 257.0
	gg := float64(g) / 257.0
	bb := float64(b) / 257.0
	// standard luminance
	l := 0.2126*rr + 0.7152*gg + 0.0722*bb
	if l < 0 {
		l = 0
	}
	if l > 255 {
		l = 255
	}
	return uint8(l)
}

func main() {
	inPath := flag.String("in", "", "input image path (png/jpg/gif)")
	outPath := flag.String("out", "", "optional: save ASCII output to file (e.g. out.txt)")
	width := flag.Int("w", 80, "output width in characters")
	threshold := flag.Int("t", -1, "threshold 0..255 (if -1 use auto mean)")
	invert := flag.Bool("invert", false, "invert black/white")
	aspect := flag.Float64("aspect", 0.5, "character aspect ratio (height/width), default 0.5")
	flag.Parse()

	if *inPath == "" {
		fmt.Fprintln(os.Stderr, "Usage: go run img2bw.go -in image.png [-w 80] [-t 128] [-invert]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	f, err := os.Open(*inPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot open input:", err)
		os.Exit(1)
	}
	defer f.Close()

	src, _, err := image.Decode(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Decode error:", err)
		os.Exit(1)
	}

	// compute target size keeping aspect ratio
	b := src.Bounds()
	srcW := b.Dx()
	srcH := b.Dy()

	targetW := *width
	if targetW <= 0 {
		targetW = 80
	}

	// Character cells are usually taller than wide; apply aspect ratio correction
	scale := float64(targetW) / float64(srcW)
	targetH := int(float64(srcH)*scale*(*aspect) + 0.5)
	if targetH < 1 {
		targetH = 1
	}

	// Resize using high-quality scaler
	dst := image.NewRGBA(image.Rect(0, 0, targetW, targetH))
	draw.CatmullRom.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)

	// compute mean luminance if threshold not set
	var mean float64
	if *threshold < 0 {
		var sum float64
		for y := 0; y < targetH; y++ {
			for x := 0; x < targetW; x++ {
				sum += float64(luminance(dst.At(x, y)))
			}
		}
		mean = sum / float64(targetW*targetH)
	} else {
		mean = float64(*threshold)
	}

	// produce ASCII lines
	lines := make([]string, 0, targetH)
	for y := 0; y < targetH; y++ {
		line := make([]rune, targetW)
		for x := 0; x < targetW; x++ {
			L := float64(luminance(dst.At(x, y)))
			var on bool
			if *threshold < 0 {
				// auto mean: pixel darker than mean -> black
				on = L < mean
			} else {
				on = L < mean
			}
			if *invert {
				on = !on
			}
			if on {
				line[x] = '█' // black pixel
			} else {
				line[x] = ' ' // white pixel
			}
		}
		lines = append(lines, string(line))
	}

	// Print to stdout
	for _, ln := range lines {
		fmt.Println(ln)
	}

	// Optionally save to file
	if *outPath != "" {
		if err := os.WriteFile(*outPath, []byte(fmt.Sprintln(stringsJoin(lines, "\n"))), 0644); err != nil {
			fmt.Fprintln(os.Stderr, "Cannot write out:", err)
			os.Exit(1)
		}
	}
}

// small helper to join lines without importing strings in top-level (keeps code explicit)
func stringsJoin(a []string, sep string) string {
	if len(a) == 0 {
		return ""
	}
	out := a[0]
	for i := 1; i < len(a); i++ {
		out += sep + a[i]
	}
	return out
}
