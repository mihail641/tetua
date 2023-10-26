package image_utils

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"mime/multipart"
	"regexp"
	"runtime"
)

type OptimizedFile struct {
	io.ReadCloser
}

func (of *OptimizedFile) Seek(offset int64, whence int) (int64, error) {
	return 0, errors.New("Seek is not supported for OptimizedFile")
}

func (of *OptimizedFile) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, errors.New("ReadAt is not supported for OptimizedFile")
}

// StartImageProcessing запускает worker pool для обработки изображений
func StartImageProcessing() {
	totalWorkers := 2 * runtime.NumCPU()
	//канал для передачи картинки для обработки
	job = make(chan *imageCompression)
	// Запускаем  горутину для обработки задач с бесконечным циклом
	for i := 0; i < totalWorkers; i++ {
		go func() {
			for {
				select {
				case task := <-job:
					file, err := optimizeImage(&task.imageFile)
					imageCompression := &result{file, err}
					task.resultChan <- imageCompression
				}
			}
		}()
	}

}

// IsValidImageType проверка регулярного выражения для image/ с поддерживаемыми расширениями
func IsValidImageType(fileType string) bool {
	supportedExtensions := fmt.Sprintf("(%s|%s|%s|%s)",
		string(BMP), string(JPEG), string(JPG), string(PNG))
	pattern := regexp.MustCompile(fmt.Sprintf("^%s/(%s)$", string(FileType), supportedExtensions))
	return pattern.MatchString(fileType)
}

//optimizeImage функция по оттимизации изображения
func optimizeImage(src *multipart.File) (multipart.File, error) {
	img, _, err := image.Decode(*src)
	if err != nil {
		return nil, err
	}
	// Оптимизация размера изображения до full hd пропорционально
	img = imaging.Fit(img, 1920, 1080, imaging.Lanczos)
	//dir, _ := filepath.Split(dst)

	// Создаем временный буфер для хранения оптимизированных данных
	var optimizedBuffer bytes.Buffer
	if err := jpeg.Encode(&optimizedBuffer, img, nil); err != nil {
		return nil, err
	}

	// Создаем новый OptimizedFile на основе оптимизированных данных
	optimizedFile := &OptimizedFile{
		ReadCloser: ioutil.NopCloser(bytes.NewReader(optimizedBuffer.Bytes())),
	}

	return optimizedFile, nil
}
