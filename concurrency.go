package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

func channels_1() {
	// channels_1 START OMIT
	type F struct {
		Path string
		Data []byte
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	loader := func(paths <-chan string, datas chan<- F) {
		for path := range paths {
			b, err := ioutil.ReadFile(path)
			if err != nil {
				log.Println("file", path, ":", err)
			} else {
				log.Println("read", path, "", len(b)/(1<<20), "MB")
				datas <- F{path, b}
				wg.Add(1)
			}
		}
		close(datas)
		wg.Done()
	}
	// channels_1 END OMIT
	// channels_2 START OMIT
	compresser := func(datas <-chan F, compresseds chan<- F) {
		for f := range datas {
			log.Println("compressing", f.Path)
			out := bytes.NewBuffer(nil)
			w := gzip.NewWriter(out)

			if _, err := w.Write(f.Data); err != nil {
				log.Println("compress:", err)
				wg.Done()
				continue
			}
			if err := w.Close(); err != nil {
				log.Println("file", f.Path, ":", err)
				wg.Done()
				continue
			}
			compresseds <- F{f.Path, out.Bytes()}
		}
	}
	// channels_2 END OMIT
	// channels_3 START OMIT
	saver := func(compresseds <-chan F) {
		for f := range compresseds {
			log.Println("'saving'", f.Path, "", len(f.Data)/(1<<20), "MB")

			// mark done
			wg.Done()
		}
		// nop
		wg.Done()
	}
	// channels_3 END OMIT
	// channels_4 START OMIT

	paths := make(chan string, 16)
	datas := make(chan F, 16)
	compresseds := make(chan F, 16)

	go loader(paths, datas)
	for i := 0; i < runtime.NumCPU(); i++ {
		go compresser(datas, compresseds)
	}
	go saver(compresseds)

	filepath.Walk("testdir", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		log.Println("queueing", path)
		paths <- path
		return err
	})
	close(paths)
	wg.Wait()
	// channels_4 END OMIT
}

func main() {
	channels_1()
}
