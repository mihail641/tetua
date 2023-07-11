package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/ngocphuongnb/tetua/app/config"
	jadecmd "github.com/ngocphuongnb/tetua/packages/jade"
)

type BuildCache struct {
	ViewMTime map[string]int64
	EntMTime  int64
}

func getDirLastMTime(root string) (int64, error) {
	var lastModTime int64
	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if info.ModTime().UnixMicro() > lastModTime {
				lastModTime = info.ModTime().UnixMicro()
			}
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return lastModTime, nil
}

func getBuildCache(buildCacheFile string) (BuildCache, error) {
	buildCache := BuildCache{}
	buildCacheBytes, err := ioutil.ReadFile(buildCacheFile)
	if err == nil && len(buildCacheBytes) > 0 {
		err = json.Unmarshal(buildCacheBytes, &buildCache)
		if err != nil {
			return buildCache, err
		}
	}

	if buildCache.ViewMTime == nil {
		buildCache.ViewMTime = make(map[string]int64)
	}
	return buildCache, nil
}

func BuildViewAssets(viewsDir, viewsOutputDir, buildCacheFile string, force bool) error {
	buildCache, err := getBuildCache(buildCacheFile)
	if err != nil {
		return err
	}
	cachedViewMTimes := buildCache.ViewMTime
	changedPageFiles := make([]string, 0)
	changedPartialsFiles := make([]string, 0)
	changedFiles := make([]string, 0)
	if !force {
		if err := filepath.Walk(viewsDir, func(viewFilePath string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				isPageFile := strings.HasPrefix(viewFilePath, path.Join(viewsDir, "pages"))
				fileMTime := info.ModTime().UnixMicro()

				if viewFileLastMTime, ok := cachedViewMTimes[viewFilePath]; ok && viewFileLastMTime >= fileMTime {

				} else {
					if isPageFile {
						changedPageFiles = append(changedPageFiles, viewFilePath)
					} else {
						changedPartialsFiles = append(changedPartialsFiles, viewFilePath)
					}
					cachedViewMTimes[viewFilePath] = fileMTime
				}
			}
			return nil
		}); err != nil {
			return err
		}
	}

	// If has partials change, rebuild all
	if force || len(changedPartialsFiles) > 0 {
		fmt.Println("> Has partials changed")
		changedFiles = append(changedFiles, viewsDir+"/pages")
	} else {
		// If has page change, rebuild only changed page
		if len(changedPageFiles) > 0 {
			fmt.Println("> Has pages changed")
			changedFiles = changedPageFiles
		}
	}

	if len(changedFiles) > 0 {
		jadecmd.CompilePaths(changedFiles, jadecmd.CompileConfig{
			Writer:  true,
			PkgName: "views",
			OutDir:  viewsOutputDir,
		})

		buildCache.ViewMTime = cachedViewMTimes
		b, err := json.Marshal(buildCache)
		if err != nil {
			return err
		}

		if err = os.WriteFile(buildCacheFile, b, 0644); err != nil {
			return err
		}
	} else {
		fmt.Println("> No views changed")
	}

	return nil
}

//Функция `GenerateEnt(buildCacheFile string, force bool)` принимает путь к файлу `cache.json` и флаг `force`.
//Она проверяет, были ли изменены файлы схемы Ent с момента последней генерации кода Ent, используя информацию из
//`cache.json`. Если были изменены файлы, функция вызывает команду `go generate` для генерации кода Ent. Затем функция
//обновляет информацию о времени последнего изменения файлов схемы Ent в `cache.json`.
func GenerateEnt(buildCacheFile string, force bool) error {
	lastEntSchemaMTime, err := getDirLastMTime("./packages/entrepository/ent/schema")
	if err != nil {
		return err
	}
	buildCache, err := getBuildCache(buildCacheFile)
	if err != nil {
		return err
	}
	if force || lastEntSchemaMTime > buildCache.EntMTime {
		fmt.Println("> Has ent schema changed")
		buildCache.EntMTime = lastEntSchemaMTime
		cmd := exec.Command("go", "generate", "./packages/entrepository/ent", "-vvvv")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}

		b, err := json.Marshal(buildCache)
		if err != nil {
			return err
		}

		if err = os.WriteFile(buildCacheFile, b, 0644); err != nil {
			return err
		}
		return nil
	}

	fmt.Println("> No ent schema changed")
	return nil
}

//Основная функция `main()` сначала проверяет наличие флага `force`, который указывает, нужно ли принудительно
//перекомпилировать представления и перегенерировать код Ent. Затем она вызывает функцию `createDir()`, чтобы
//создать директорию и файл `cache.json`, если они не существуют. Затем вызывает функцию `BuildViewAssets()`
//для компиляции файлов представлений и функцию `GenerateEnt()` для генерации кода Ent. Если во время выполнения
//получены ошибки, они выводятся на экран и программа завершается с ошибкой.
func main() {
	force := false
	if len(os.Args) > 1 {
		force = os.Args[1] == "--force"
	}

	if err := BuildViewAssets(
		path.Join("./app/themes", config.APP_THEME, "views"),
		"./views",
		"./private/tmp/cache.json",
		force,
	); err != nil {
		log.Fatal("Ошибка в функции BuildViewAssets", err)

	}
	if err := GenerateEnt("./private/tmp/cache.json", force); err != nil {
		log.Fatal("Ошибка в функции GenerateEnt", err)

	}
}
