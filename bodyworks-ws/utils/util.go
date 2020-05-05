package utils

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"io/ioutil"
	"encoding/hex"
)

const MsgErrorNoValidStatus string = "No se puede eliminar, el estado no es valido"

type File struct {
	Name, Body string
}
type SpaHandler struct {
	StaticPath string
	IndexPath  string
}
func (h SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path = filepath.Join(h.StaticPath, path)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.StaticPath, h.IndexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.FileServer(http.Dir(h.StaticPath)).ServeHTTP(w, r)
}
func ConvertObjToJson(obj interface{}) string {
	var str []byte
	str, err := json.Marshal(obj)
	if err != nil {
		log.Print(err)
		return ""
	}
	return string(str)
}
func ConvertObjToXml(obj interface{}) []byte {
	var strXml []byte
	strXml, err := xml.Marshal(obj)
	if err != nil {
		log.Print(err)
		return strXml
	}
	return strXml
}
func StreamZip(file File) []byte {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	f, err := w.Create(file.Name)
	if err != nil {
		log.Print(err)
		return nil
	}
	_, err = f.Write([]byte(file.Body))
	if err != nil {
		log.Print(err)
		return nil
	}
	w.Close()
	return buf.Bytes()
}
func StreamUnzip(inFile []byte) (File, error) {
	zipReader, err := zip.NewReader(bytes.NewReader(inFile), int64(len(inFile)))
	if err != nil {
        log.Print(err)
	}
	var outFile File
	for _, zipFile := range zipReader.File {
		outFile.Name = zipFile.Name
		f, err := zipFile.Open()
		if err != nil {
			return File{}, err
		}
		defer f.Close()
		readFile, err := ioutil.ReadAll(f)
		if err != nil {
			return File{}, err
		}
		log.Print(hex.EncodeToString(readFile))
		outFile.Body = string(readFile)
    }
	return outFile, nil
}