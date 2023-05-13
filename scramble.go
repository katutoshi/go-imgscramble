package imgscramble

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"math/rand"
)

func Seed(file io.Reader) (int64, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return 0, fmt.Errorf("failed to hash: %w", err)
	}

	value := hash.Sum(nil)
	return int64(binary.LittleEndian.Uint64(value)), nil
}

func Scramble(file io.Reader, seed int64) ([]byte, error) {
	s, format, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w", err)
	}

	raw := image.NewNRGBA(s.Bounds())
	index := shuffledIndex(len(raw.Pix), seed)
	Shuffle(raw.Pix, index)

	var b bytes.Buffer
	switch format {
	case "jpeg":
		err = jpeg.Encode(&b, raw, &jpeg.Options{Quality: 100})
	case "png":
		err = png.Encode(&b, raw)
	default:
		err = fmt.Errorf("unknown format: %s", format)
	}
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func Unscramble(file io.Reader, seed int64) ([]byte, error) {
	s, format, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %w", err)
	}

	raw := image.NewNRGBA(s.Bounds())
	// seedを元にシャッフル用のindexを作成する
	index := shuffledIndex(len(raw.Pix), seed)

	// シャッフル用のindexを元に画像をシャッフルする
	Unshuffle(raw.Pix, index)

	var b bytes.Buffer
	switch format {
	case "jpeg":
		err = jpeg.Encode(&b, raw, &jpeg.Options{Quality: 100})
	case "png":
		err = png.Encode(&b, raw)
	default:
		err = fmt.Errorf("unknown format: %s", format)
	}
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func Shuffle(slice []uint8, index []uint64) {
	shuffledPix := make([]uint8, len(index))
	// Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
	for i, v := range index {
		shuffledPix[i] = slice[v]
	}

	copy(slice, shuffledPix)
}

func Unshuffle(slice []uint8, index []uint64) {
	unshuffled := make([]uint8, len(index))
	for i, v := range index {
		unshuffled[i] = slice[v]
	}

	copy(slice, unshuffled)
}

func shuffledIndex(len int, seed int64) []uint64 {
	index := makeRange(len)
	r := rand.New(rand.NewSource(seed))
	r.Shuffle(len, func(i, j int) {
		index[i], index[j] = index[j], index[i]
	})
	return index
}

func makeRange(len int) []uint64 {
	s := make([]uint64, len)
	for i := range s {
		s[i] = uint64(i)
	}
	return s
}
