package image_utils

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
	"testing"
)

// TestStartImageProcessing тестирование колличества worker-ов обрабатывающих картинку
func TestStartImageProcessing(t *testing.T) {
	StartImageProcessing()
	var image imageCompression
	expectedGoroutines := runtime.GOMAXPROCS(0) * 2
	uploadFilePath := "test_file.jpg"
	sampleFile, err := os.Open(uploadFilePath)
	image.resultChan = make(chan *result)
	if err != nil {
		fmt.Println("Error open file", err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		job <- &imageCompression{imageFile: sampleFile, resultChan: image.resultChan}

	}()
	wg.Wait()

	activeGoroutines := runtime.NumGoroutine()
	expectedGoroutines = expectedGoroutines + 3
	if activeGoroutines >= expectedGoroutines {
		t.Errorf("Expected no more than %d goroutines, but got %d", expectedGoroutines, activeGoroutines)
	}

}

//TestRunImageCompression тестирует, что файлы различных размеров будут преобразованы пропорционально до
//размеров full hd, а так же что отправитель получит свой файл после преобразования
func TestRunImageCompression(t *testing.T) {
	StartImageProcessing()
	//создаем слайс структур файла
	type fileForTest struct {
		width    int
		height   int
		fileName string
	}
	var filesAfterOptimized []fileForTest
	var imageTest imageCompression
	imageTest.resultChan = make(chan *result)
	files := []fileForTest{{
		width:    1920,
		height:   5000,
		fileName: "firstFile.jpg",
	}, {width: 4000,
		height:   1080,
		fileName: "secondFile.bmp"}, {
		width:    1850,
		height:   1000,
		fileName: "thirdFile.jpg"}, {
		width:    4000,
		height:   5000,
		fileName: "forthFile.png",
	}}

	for _, file := range files {
		fileForJob := createImage(file.width, file.height)
		var optimizedBuffer bytes.Buffer
		if err := jpeg.Encode(&optimizedBuffer, fileForJob, nil); err != nil {
			fmt.Println("Error Encode file", err)
		}

		// Создаем новый OptimizedFile на основе оптимизированных данных
		fileForOptimized := &OptimizedFile{
			ReadCloser: ioutil.NopCloser(bytes.NewReader(optimizedBuffer.Bytes())),
		}

		fileAfterOptimization, err := RunImageCompression(fileForOptimized)
		if err != nil {
			fmt.Println("Error RunImageCompression", err)
		}
		// Создать изображение из данных
		img, _, err := image.Decode(fileAfterOptimization)
		if err != nil {
			fmt.Println("Error Decode file", err)
		}

		width := img.Bounds().Dx()
		height := img.Bounds().Dy()
		fileParametrAfterOptimization := &fileForTest{width: width, height: height, fileName: file.fileName}
		filesAfterOptimized = append(filesAfterOptimized, *fileParametrAfterOptimization)
	}
	expectedResult := []fileForTest{{
		width:    414,
		height:   1080,
		fileName: "firstFile.jpg",
	}, {width: 1920,
		height:   518,
		fileName: "secondFile.bmp"}, {
		width:    1850,
		height:   1000,
		fileName: "thirdFile.jpg"}, {
		width:    864,
		height:   1080,
		fileName: "forthFile.png",
	}}

	assert.Equal(t, expectedResult, filesAfterOptimized)
}

// CreateImage создает изображение и возвращает его в виде image.Image.
func createImage(width int, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	return img
}
