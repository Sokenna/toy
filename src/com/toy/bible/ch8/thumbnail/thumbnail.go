package thumbnail

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Image 将源图像缩放为缩略图
// 缩略图尺寸为128x128，保持原始图像的宽高比
// 使用最近邻插值算法进行缩放
func Image(src image.Image) image.Image {
	// 获取原图尺寸
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	// 初始化目标尺寸为128x128
	width, height := 128, 128
	// 根据宽高比调整目标尺寸，保持比例
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect)
	} else {
		height = int(128 / aspect)
	}
	// 计算缩放比例
	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	// 创建目标图像并进行最近邻采样
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst
}

func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(src)
	return jpeg.Encode(w, dst, nil)
}

// ImageFile2 从输入文件读取图像并生成缩略图写入到输出文件
//
// 参数:
//   - outfile: 输出文件路径
//   - infile: 输入文件路径
//
// 返回:
//   - error: 操作过程中出现的错误
func ImageFile2(outfile, infile string) (err error) {
	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(outfile)
	if err != nil {
		return err
	}
	if err := ImageStream(out, in); err != nil {
		out.Close()
		return fmt.Errorf("scaling %s to %s %s", infile, outfile, err)
	}
	return out.Close()
}

func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile)
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext
	return outfile, ImageFile2(outfile, infile)
}

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go ImageFile(f) //忽略错误
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {

		go func(f string) {
			ImageFile(f)
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		<-ch
	}
}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			errors <- err
		}(f)
	}
	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}

	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			ch <- it
		}(f)
	}
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

func makeThumbnails6(filenames []string) int64 {
	sizes := make(chan int64)
	wg := sync.WaitGroup{}
	for _, f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
