package main

import (
	_ "embed"
	"encoding/xml"
	"flag"
	"html/template"
	"log"
	"math"
	"os"
	"os/exec"
	"runtime"
)

//go:embed template.html
var AssoModel string

func main() {
	file_path := flag.String("f", "fattura.xml", "File XML fattura da interpretare")
	output_path := flag.String("o", "", "File HTML di esportazione")
	flag.Parse()
	bts, err := os.ReadFile(*file_path)
	if err != nil {
		log.Fatal(err)
	}
	fatt := FatturaElettronica{}
	err = xml.Unmarshal(bts, &fatt)
	totale := 0.0
	for _, v := range fatt.FatturaElettronicaBody.DatiBeniServizi.DatiRiepilogo {
		totale += v.ImponibileImporto + v.Imposta
	}

	if math.Abs(fatt.FatturaElettronicaBody.DatiGenerali.DatiGeneraliDocumento.ImportoTotaleDocumento-totale) > 1 {
		fatt.FatturaElettronicaBody.DatiGenerali.DatiGeneraliDocumento.ImportoTotaleDocumento = totale
	}

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
	var tmpFile *os.File
	if *output_path == "" {
		tmpFile, err = os.CreateTemp(os.TempDir(), "*.html")
	} else {
		tmpFile, err = os.Create(*output_path)
	}
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(tmpFile, "T", fatt)
	if err != nil {
		log.Fatal(err)
	}
	if *output_path == "" {
		TryOpenFile(tmpFile.Name())
	}
}
func run_cmd(args ...string) error {
	cmnd := exec.Command(args[0], args[1:]...)
	if err := cmnd.Start(); err != nil {
		return err
	}
	if err := cmnd.Wait(); err != nil {
		return err
	}
	return nil
}
func TryOpenFile(file string) error {
	if runtime.GOOS == "windows" {
		return run_cmd("explorer.exe", file)
	}
	if runtime.GOOS == "darwin" {
		return run_cmd("open", file)
	}
	return run_cmd("xdg-open", file)
}
