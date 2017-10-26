package yac

/*
import "image"

func resize(img image.Image) image.Image {
	bounds := img.Bounds()
	if img.Bounds().Dx() == 2 &&
		img.Bounds().Dy() == 2 {
		return img
	}

	scaleX := img.Bounds().Dx() / 2
	scaleY := img.Bounds().Dy() / 2

	switch src := img.(type) {
	case *image.RGBA:
		temp := image.NewRGBA(image.Rect(0, 0, src.Bounds().Dy(), 2))
		result := image.NewRGBA(image.Rect(0, 0, 2, 2))
		return result
	case *image.NRGBA:
		temp := image.NewRGBA(image.Rect(0, 0, src.Bounds().Dy(), 2))
		result := image.NewRGBA(image.Rect(0, 0, 2, 2))
		return result
	case *image.YCbCr:
	case *image.RGBA64:
		temp := image.NewRGBA64(image.Rect(0, 0, src.Bounds().Dy(), 2))
		result := image.NewRGBA64(image.Rect(0, 0, 2, 2))
		return result
	case *image.NRGBA64:
		temp := image.NewNRGBA64(image.Rect(0, 0, src.Bounds().Dy(), 2))
		result := image.NewNRGBA64(image.Rect(0, 0, 2, 2))
		return result
	case *image.Gray:
		temp := image.NewGray(image.Rect(0, 0, src.Bounds().Dy(), 2))
		result := image.NewGray(image.Rect(0, 0, 2, 2))
		return result
	case *image.Gray16:
		temp := image.NewGray16(image.Rect(0, 0, src.Bounds().Dy(), 2))
		result := image.NewGray16(image.Rect(0, 0, 2, 2))
		return result
	default:
		temp := image.NewRGBA64(image.Rect(0, 0, img.Bounds().Dy(), 2))
		result := image.NewRGBA64(image.Rect(0, 0, 2, 2))
		return result
	}
	return nil
}
*/
