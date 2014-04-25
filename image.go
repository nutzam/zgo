package z

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

// 读取JPEG图片返回image.Image对象
func ImageJPEG(ph string) (image.Image, error) {
	// 打开图片文件
	f, fileErr := os.Open(ph)
	if fileErr != nil {
		return nil, fileErr
	}
	// 退出时关闭文件
	defer f.Close()
	// 解码
	j, jErr := jpeg.Decode(f)
	if jErr != nil {
		return nil, jErr
	}
	// 返回解码后的图片
	return j, nil
}

// 读取PNG图片返回image.Image对象
func ImagePNG(ph string) (image.Image, error) {
	// 打开图片文件
	f, fileErr := os.Open(ph)
	if fileErr != nil {
		return nil, fileErr
	}
	// 退出时关闭文件
	defer f.Close()
	// 解码
	p, pErr := png.Decode(f)
	if pErr != nil {
		return nil, pErr
	}
	// 返回解码后的图片
	return p, nil
}

// 按照分辨率创建一张空白图片对象
func ImageRGBA(width, height int) *image.RGBA {
	// 建立图像,image.Rect(最小X,最小Y,最大X,最小Y)
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

// 将图片绘制到图片
func ImageDrawRGBA(img *image.RGBA, imgcode image.Image, x, y int) {
	// 绘制图像
	// image.Point A点的X,Y坐标,轴向右和向下增加{0,0}
	// image.ZP ZP is the zero Point
	// image.Pt Pt is shorthand for Point{X, Y}
	draw.Draw(img, img.Bounds(), imgcode, image.Pt(x, y), draw.Over)
}

// 将图片绘制到图片
func ImageDrawRGBAOffSet(img *image.RGBA, imgcode image.Image, r image.Rectangle, x, y int) {
	// 绘制图像
	// image.Point A点的X,Y坐标,轴向右和向下增加{0,0}
	// image.ZP ZP is the zero Point
	// image.Pt Pt is shorthand for Point{X, Y}
	// r image.Rectangle img.Bounds() or img.Bounds().Add(offset)
	draw.Draw(img, r, imgcode, image.Pt(x, y), draw.Over)
}

// JPEG将编码生成图片
// 选择编码参数,质量范围从1到100,更高的是更好 &jpeg.Options{90}
func ImageEncodeJPEG(ph string, img image.Image, option int) error {
	// 确保文件父目录存在
	FcheckParents(ph)
	// 打开文件等待写入
	f := FileW(ph)
	// 保证文件正常关闭
	defer f.Close()
	// 写入文件
	return jpeg.Encode(f, img, &jpeg.Options{option})
}

// PNG将编码生成图片
func ImageEncodePNG(ph string, img image.Image) error {
	// 确保文件父目录存在
	FcheckParents(ph)
	// 打开文件等待写入
	f := FileW(ph)
	// 保证文件正常关闭
	defer f.Close()
	// 写入文件
	return png.Encode(f, img)
}
