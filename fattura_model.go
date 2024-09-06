package main

import "encoding/xml"

type FatturaElettronica struct {
	XMLName                  xml.Name `xml:"FatturaElettronica"`
	Text                     string   `xml:",chardata"`
	P                        string   `xml:"p,attr"`
	Ds                       string   `xml:"ds,attr"`
	Xsi                      string   `xml:"xsi,attr"`
	Versione                 string   `xml:"versione,attr"`
	FatturaElettronicaHeader struct {
		Text             string `xml:",chardata"`
		DatiTrasmissione struct {
			Text           string `xml:",chardata"`
			IdTrasmittente struct {
				Text     string `xml:",chardata"`
				IdPaese  string `xml:"IdPaese"`
				IdCodice string `xml:"IdCodice"`
			} `xml:"IdTrasmittente"`
			ProgressivoInvio     string `xml:"ProgressivoInvio"`
			FormatoTrasmissione  string `xml:"FormatoTrasmissione"`
			CodiceDestinatario   string `xml:"CodiceDestinatario"`
			ContattiTrasmittente struct {
				Text     string `xml:",chardata"`
				Telefono string `xml:"Telefono"`
				Email    string `xml:"Email"`
			} `xml:"ContattiTrasmittente"`
		} `xml:"DatiTrasmissione"`
		CedentePrestatore struct {
			Text           string `xml:",chardata"`
			DatiAnagrafici struct {
				Text         string `xml:",chardata"`
				IdFiscaleIVA struct {
					Text     string `xml:",chardata"`
					IdPaese  string `xml:"IdPaese"`
					IdCodice string `xml:"IdCodice"`
				} `xml:"IdFiscaleIVA"`
				CodiceFiscale string `xml:"CodiceFiscale"`
				Anagrafica    struct {
					Text          string `xml:",chardata"`
					Nome          string `xml:"Nome"`
					Cognome       string `xml:"Cognome"`
					Denominazione string `xml:"Denominazione"`
				} `xml:"Anagrafica"`
				RegimeFiscale string `xml:"RegimeFiscale"`
			} `xml:"DatiAnagrafici"`
			Sede struct {
				Text      string `xml:",chardata"`
				Indirizzo string `xml:"Indirizzo"`
				CAP       string `xml:"CAP"`
				Comune    string `xml:"Comune"`
				Provincia string `xml:"Provincia"`
				Nazione   string `xml:"Nazione"`
			} `xml:"Sede"`
		} `xml:"CedentePrestatore"`
		CessionarioCommittente struct {
			Text           string `xml:",chardata"`
			DatiAnagrafici struct {
				Text          string `xml:",chardata"`
				CodiceFiscale string `xml:"CodiceFiscale"`
				IdFiscaleIVA  struct {
					Text     string `xml:",chardata"`
					IdPaese  string `xml:"IdPaese"`
					IdCodice string `xml:"IdCodice"`
				} `xml:"IdFiscaleIVA"`
				Anagrafica struct {
					Text          string `xml:",chardata"`
					Nome          string `xml:"Nome"`
					Cognome       string `xml:"Cognome"`
					Denominazione string `xml:"Denominazione"`
				} `xml:"Anagrafica"`
			} `xml:"DatiAnagrafici"`
			Sede struct {
				Text      string `xml:",chardata"`
				Indirizzo string `xml:"Indirizzo"`
				CAP       string `xml:"CAP"`
				Comune    string `xml:"Comune"`
				Provincia string `xml:"Provincia"`
				Nazione   string `xml:"Nazione"`
			} `xml:"Sede"`
		} `xml:"CessionarioCommittente"`
		TerzoIntermediarioOSoggettoEmittente struct {
			Text           string `xml:",chardata"`
			DatiAnagrafici struct {
				Text         string `xml:",chardata"`
				IdFiscaleIVA struct {
					Text     string `xml:",chardata"`
					IdPaese  string `xml:"IdPaese"`
					IdCodice string `xml:"IdCodice"`
				} `xml:"IdFiscaleIVA"`
				CodiceFiscale string `xml:"CodiceFiscale"`
				Anagrafica    struct {
					Text          string `xml:",chardata"`
					Denominazione string `xml:"Denominazione"`
				} `xml:"Anagrafica"`
			} `xml:"DatiAnagrafici"`
		} `xml:"TerzoIntermediarioOSoggettoEmittente"`
		SoggettoEmittente string `xml:"SoggettoEmittente"`
	} `xml:"FatturaElettronicaHeader"`
	FatturaElettronicaBody struct {
		Text         string `xml:",chardata"`
		DatiGenerali struct {
			Text                  string `xml:",chardata"`
			DatiGeneraliDocumento struct {
				Text                   string  `xml:",chardata"`
				TipoDocumento          string  `xml:"TipoDocumento"`
				Divisa                 string  `xml:"Divisa"`
				Data                   string  `xml:"Data"`
				Numero                 string  `xml:"Numero"`
				ImportoTotaleDocumento float64 `xml:"ImportoTotaleDocumento"`
				Causale                string  `xml:"Causale"`
			} `xml:"DatiGeneraliDocumento"`
		} `xml:"DatiGenerali"`
		DatiBeniServizi struct {
			Text           string `xml:",chardata"`
			DettaglioLinee []struct {
				Text                string  `xml:",chardata"`
				NumeroLinea         string  `xml:"NumeroLinea"`
				Descrizione         string  `xml:"Descrizione"`
				Quantita            float64 `xml:"Quantita"`
				UnitaMisura         string  `xml:"UnitaMisura"`
				PrezzoUnitario      float64 `xml:"PrezzoUnitario"`
				PrezzoTotale        float64 `xml:"PrezzoTotale"`
				AliquotaIVA         float64 `xml:"AliquotaIVA"`
				ScontoMaggiorazione struct {
					Text        string  `xml:",chardata"`
					Tipo        string  `xml:"Tipo"`
					Importo     float64 `xml:"Importo"`
					Percentuale float64 `xml:"Percentuale"`
				}
				CodiceArticolo struct {
					Text         string `xml:",chardata"`
					CodiceTipo   string `xml:"CodiceTipo"`
					CodiceValore string `xml:"CodiceValore"`
				}
			} `xml:"DettaglioLinee"`
			DatiRiepilogo []struct {
				Text              string  `xml:",chardata"`
				AliquotaIVA       float64 `xml:"AliquotaIVA"`
				ImponibileImporto float64 `xml:"ImponibileImporto"`
				Imposta           float64 `xml:"Imposta"`
				EsigibilitaIVA    string  `xml:"EsigibilitaIVA"`
			} `xml:"DatiRiepilogo"`
		} `xml:"DatiBeniServizi"`
		DatiPagamento struct {
			Text                string `xml:",chardata"`
			CondizioniPagamento string `xml:"CondizioniPagamento"`
			DettaglioPagamento  []struct {
				Text                  string  `xml:",chardata"`
				Beneficiario          string  `xml:"Beneficiario"`
				ModalitaPagamento     string  `xml:"ModalitaPagamento"`
				DataScadenzaPagamento string  `xml:"DataScadenzaPagamento"`
				ImportoPagamento      float64 `xml:"ImportoPagamento"`
				IstitutoFinanziario   string  `xml:"IstitutoFinanziario"`
				IBAN                  string  `xml:"IBAN"`
			} `xml:"DettaglioPagamento"`
		} `xml:"DatiPagamento"`
		Allegati []struct {
			Text           string `xml:",chardata"`
			NomeAttachment string `xml:"NomeAttachment"`
			Attachment     string `xml:"Attachment"`
		} `xml:"Allegati"`
	} `xml:"FatturaElettronicaBody"`
	Signature struct {
		Text       string `xml:",chardata"`
		ID         string `xml:"Id,attr"`
		SignedInfo struct {
			Text                   string `xml:",chardata"`
			CanonicalizationMethod struct {
				Text      string `xml:",chardata"`
				Algorithm string `xml:"Algorithm,attr"`
			} `xml:"CanonicalizationMethod"`
			SignatureMethod struct {
				Text      string `xml:",chardata"`
				Algorithm string `xml:"Algorithm,attr"`
			} `xml:"SignatureMethod"`
			Reference []struct {
				Text       string `xml:",chardata"`
				ID         string `xml:"Id,attr"`
				URI        string `xml:"URI,attr"`
				Type       string `xml:"Type,attr"`
				Transforms struct {
					Text      string `xml:",chardata"`
					Transform struct {
						Text      string `xml:",chardata"`
						Algorithm string `xml:"Algorithm,attr"`
						XPath     struct {
							Text   string `xml:",chardata"`
							Xmlns  string `xml:"xmlns,attr"`
							Filter string `xml:"Filter,attr"`
						} `xml:"XPath"`
					} `xml:"Transform"`
				} `xml:"Transforms"`
				DigestMethod struct {
					Text      string `xml:",chardata"`
					Algorithm string `xml:"Algorithm,attr"`
				} `xml:"DigestMethod"`
				DigestValue string `xml:"DigestValue"`
			} `xml:"Reference"`
		} `xml:"SignedInfo"`
		SignatureValue struct {
			Text string `xml:",chardata"`
			ID   string `xml:"Id,attr"`
		} `xml:"SignatureValue"`
		KeyInfo struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"Id,attr"`
			X509Data struct {
				Text            string `xml:",chardata"`
				X509Certificate string `xml:"X509Certificate"`
			} `xml:"X509Data"`
		} `xml:"KeyInfo"`
		Object struct {
			Text                 string `xml:",chardata"`
			QualifyingProperties struct {
				Text             string `xml:",chardata"`
				Xades            string `xml:"xades,attr"`
				Target           string `xml:"Target,attr"`
				SignedProperties struct {
					Text                      string `xml:",chardata"`
					ID                        string `xml:"Id,attr"`
					SignedSignatureProperties struct {
						Text               string `xml:",chardata"`
						SigningTime        string `xml:"SigningTime"`
						SigningCertificate struct {
							Text string `xml:",chardata"`
							Cert struct {
								Text       string `xml:",chardata"`
								CertDigest struct {
									Text         string `xml:",chardata"`
									DigestMethod struct {
										Text      string `xml:",chardata"`
										Algorithm string `xml:"Algorithm,attr"`
									} `xml:"DigestMethod"`
									DigestValue string `xml:"DigestValue"`
								} `xml:"CertDigest"`
								IssuerSerial struct {
									Text             string `xml:",chardata"`
									X509IssuerName   string `xml:"X509IssuerName"`
									X509SerialNumber string `xml:"X509SerialNumber"`
								} `xml:"IssuerSerial"`
							} `xml:"Cert"`
						} `xml:"SigningCertificate"`
					} `xml:"SignedSignatureProperties"`
					SignedDataObjectProperties struct {
						Text             string `xml:",chardata"`
						DataObjectFormat []struct {
							Text            string `xml:",chardata"`
							ObjectReference string `xml:"ObjectReference,attr"`
							MimeType        string `xml:"MimeType"`
						} `xml:"DataObjectFormat"`
					} `xml:"SignedDataObjectProperties"`
				} `xml:"SignedProperties"`
			} `xml:"QualifyingProperties"`
		} `xml:"Object"`
	} `xml:"Signature"`
}
