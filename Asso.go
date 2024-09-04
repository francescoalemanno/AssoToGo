package main

import (
	_ "embed"
	"encoding/xml"
	"flag"
	"html/template"
	"log"
	"os"
	"os/exec"
	"runtime"
)

//go:embed template.html
var AssoModel string

func main() {
	file_path := flag.String("f", "fattura.xml", "File fattura da interpretare")
	flag.Parse()
	bts, err := os.ReadFile(*file_path)
	if err != nil {
		log.Fatal(err)
	}
	fatt := FatturaElettronica{}
	err = xml.Unmarshal(bts, &fatt)
	if err != nil {
		log.Fatal(err)
	}
	/*
		fmt.Println(fatt.FatturaElettronicaHeader.CedentePrestatore)
		fmt.Println(fatt.FatturaElettronicaHeader.CessionarioCommittente)
		fmt.Println(fatt.FatturaElettronicaBody.DatiGenerali)
		for _, v := range fatt.FatturaElettronicaBody.DatiBeniServizi.DettaglioLinee {
			fmt.Println(v.Descrizione+"\n", v.PrezzoTotale, " + ", v.AliquotaIVA, "% = ", v.PrezzoTotale*(1+v.AliquotaIVA/100))
		}
		fmt.Println(fatt.FatturaElettronicaBody.DatiPagamento)
	*/
	t, err := template.New("foo").Parse(AssoModel)
	if err != nil {
		log.Fatal(err)
	}
	tmpFile, err := os.CreateTemp(os.TempDir(), "*.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(tmpFile, "T", fatt)
	if err != nil {
		log.Fatal(err)
	}
	TryOpenFile(tmpFile.Name())
}

func TryOpenFile(file string) error {
	if runtime.GOOS == "windows" {
		cmnd := exec.Command("explorer.exe", file)
		err := cmnd.Start()
		return err
	}
	if runtime.GOOS == "darwin" {
		cmnd := exec.Command("open", file)
		err := cmnd.Start()
		return err
	}
	cmnd := exec.Command("xdg-open", file)
	err := cmnd.Start()
	return err
}
