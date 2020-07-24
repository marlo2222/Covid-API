package main

import (

	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"    
	"os"
	"encoding/json"
	"reflect"
	//"strings"
)	

var casos = []dadosUsuario{}

func abrirCSV() [][]string{

	csvFile, err := os.Open("csv/casos_coronavirus.csv")
	
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	
	if err != nil {
		fmt.Println(err)
	}

	return csvLines
}
//preencho a lit
func preencherListCasos(){

	csvLines := abrirCSV()

	for _, line := range csvLines {
		caso := dadosUsuario{}

			caso.bairroPaciente = line[0]
			caso.CLASSIFICACAOESTADOSIVEP = line[1]
			caso.CODIGOMUNICIPIOPACIENTE = line[2]
			caso.CODIGOPACIENTE = line[3]
			caso.COMORBIDADEASMASIVEP = line[4]
			caso.COMORBIDADECARDIOVASCULARSIVEP = line[5]
			caso.COMORBIDADEDIABETESSIVEP = line[6]
			caso.COMORBIDADEHEMATOLOGIASIVEP = line[7]
			caso.COMORBIDADEIMUNODEFICIENCIASIVEP = line[8]
			caso.COMORBIDADENEUROLOGIASIVEP = line[9]
			caso.COMORBIDADEOBESIDADESIVEP = line[10]
			caso.COMORBIDADEPNEUMOPATIASIVEP = line[11]
			caso.COMORBIDADEPUERPERASIVEP = line[12]
			caso.COMORBIDADERENALSIVEP = line[13]
			caso.COMORBIDADESINDROMEDOWNSIVEP = line[14]
			caso.DATACOLETAEXAME = line[15]
			caso.DATAENTRADAUTISSVEP = line[16]
			caso.DATAEVOLUCAOCASOSIVEP = line[17]
			caso.DATAINICIOSINTOMAS = line[18]
			caso.DATAINTERNACAOSIVEP = line[19]
			caso.DATANOTIFICACAO = line[20]
			caso.DATAOBITO = line[21]
			caso.DATARESULTADOEXAME = line[22]
			caso.DATASAIDAUTISSVEP = line[23]
			caso.DATASOLICITACAOEXAME = line[24]
			caso.ESTADOPACIENTE = line[25]
			caso.EVOLUCAOCASOSIVEP = line[26]
			caso.IDADEPACIENTE = line[27]
			caso.IDSIVEP = line[28]
			caso.MUNICIPIOPACIENTE = line[29]
			caso.OBITOCONFIRMADO = line[30]
			caso.PAISPACIENTE = line[31]
			caso.RESULTADOFINALEXAME = line[32]
			caso.SEXOPACIENTE = line[33]

			casos = append(casos, caso)
			
	}
}

//estou pegando somente os positivos por enquanto{a lista e muito grande}
func getAll(w http.ResponseWriter, r *http.Request) {

	if len(casos) == 0{
		preencherListCasos()
	}

	resultado := [] dadosUsuario{}
	for _, caso := range casos {
		resultado = append(resultado, caso)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}

//preciso ver se retorno somente a quantidade
func getNumeroExames(w http.ResponseWriter, r *http.Request) {
	
	if len(casos) == 0{
		preencherListCasos()
	}

	count := 0
	exames := []Exame{}
	for i, caso := range casos {
			if i == 1 || i == 3 || i == 100 {
				if (caso.DATACOLETAEXAME == "" || caso.DATASOLICITACAOEXAME == ""){
					continue
				}
			fmt.Println(caso.DATACOLETAEXAME, caso.DATASOLICITACAOEXAME )
			count ++
			exame := Exame{}
			exame.DATACOLETAEXAME = caso.DATACOLETAEXAME
			exame.DATASOLICITACAOEXAME = caso.DATASOLICITACAOEXAME
			exames = append(exames, exame)
		}	
	}
	
	resultado := NumeroExames{
		quantidade: count,
		exames: exames,
	}
	//mapD := mapmmmtring]int{"exames": count}
   // mapB, _ := json.Marshal(mapD)
	//resultado.quantidade = count
	fmt.Println(resultado.quantidade)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)

}

//casos confirmados
func getCasosConfirmados(w http.ResponseWriter, r *http.Request) {

	if len(casos) == 0{
		preencherListCasos()
	}
	resultado := [] dadosUsuario{}
	for _, caso := range casos {
		if caso.RESULTADOFINALEXAME == "Positivo" {
			resultado = append(resultado, caso)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}

//casos confimados por municipio
func getCasosConfirmadosMunicipio(w http.ResponseWriter, r *http.Request) {

	if len(casos) == 0{
		preencherListCasos()
	}

	municipio := mux.Vars(r)
	resultado := []dadosUsuario{}
	for _, caso := range casos {
		if (municipio["municipio"] == caso.MUNICIPIOPACIENTE){
			if (caso.RESULTADOFINALEXAME == "Positivo"){
				resultado = append(resultado, caso)
			}
		}
	}
	fmt.Println(reflect.TypeOf(resultado))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}

//casos confirmados por sexo.
func getCasosConfirmadosPorSexo(w http.ResponseWriter, r *http.Request) {

	if len(casos) == 0{
		preencherListCasos()
	}

	sexo := mux.Vars(r)
	resultado := []dadosUsuario{}
	for _, caso := range casos {
		fmt.Println(sexo["sexo"])
		fmt.Println(caso.SEXOPACIENTE)
		if (sexo["sexo"] == caso.SEXOPACIENTE){
			if (caso.RESULTADOFINALEXAME == "Positivo"){
				resultado = append(resultado, caso)
			}
		}
	}
	fmt.Println(reflect.TypeOf(resultado))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}

func getCasosInvestigacao(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Isto ainda não serve para nada \n")
}

//obitos geral
func getAllObitos(w http.ResponseWriter, r *http.Request) {
	if len(casos) == 0{
		preencherListCasos()
	}

	resultado := []dadosUsuario{}
	for _, caso := range casos {
		if (caso.OBITOCONFIRMADO == "true"){
			resultado = append(resultado, caso)
		}
	
	}
	fmt.Println(reflect.TypeOf(resultado))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}

//obitos municipio
func getObitosMunicipio(w http.ResponseWriter, r *http.Request) {
	if len(casos) == 0{
		preencherListCasos()
	}

	municipio := mux.Vars(r)
	resultado := []dadosUsuario{}
	for _, caso := range casos {
		if (municipio["municipio"] == caso.MUNICIPIOPACIENTE){
			if (caso.OBITOCONFIRMADO == "true"){
				resultado = append(resultado, caso)
			}
		}
	}
	fmt.Println(reflect.TypeOf(resultado))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}

func getLetalidade(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Category: %v\n")
}


func getCasosRecuperados(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Category: %v\n")
}

//so confirmados
func getCasosPorSexo(w http.ResponseWriter, r *http.Request) {

	if len(casos) == 0{
		preencherListCasos()
	}

	sexo :=mux.Vars(r)
	resultado := []dadosUsuario{}

	for i, caso := range casos {
		if caso.SEXOPACIENTE == sexo["sexo"] {
			resultado = append(resultado, caso)		
		}
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}

//so confirmados data -> passa a data no formato "aaaa-mm-dd"
func getCasosConfirmadosData(w http.ResponseWriter, r *http.Request) {
	if len(casos) == 0{
		preencherListCasos()
	}

	data :=mux.Vars(r)
   	data["data"] = data["data"]+" 00:00:00.0"

	resultado := []dadosUsuario{}
	for _, caso := range casos {
		if (data["data"] == caso.DATARESULTADOEXAME){
			if (caso.OBITOCONFIRMADO == "true"){
				resultado = append(resultado, caso)
			}
		}

	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}
type casosPorSexo struct{
	IDADEPACIENTE string
	MUNICIPIOPACIENTE string
	RESULTADOFINALEXAME string
	SEXOPACIENTE string
}

type NumeroExames struct{
	quantidade int `json:"qtd,omitempty"`
	exames []Exame `json:"exames,omitempty"`
}

type Exame struct{
	DATACOLETAEXAME string `json:"datacoletaexame,omitempty"`
	DATASOLICITACAOEXAME string `json:"datasolicitacaoexame,omitempty"`
}

type dadosUsuario struct {
	bairroPaciente string `json:"bairropaciente,omitempty"`
	CLASSIFICACAOESTADOSIVEP string `json:"classificacaoestadosivep,omitempty"`
	CODIGOMUNICIPIOPACIENTE string `json:"codigomunicipiopaciente,omitempty"`
	CODIGOPACIENTE string `json:"codigopaciente,omitempty"`
	COMORBIDADEASMASIVEP string `json:"comorbidadeasmasivep,omitempty"`
	COMORBIDADECARDIOVASCULARSIVEP string `json:"comorbidadecardiovascularsivep,omitempty"`
	COMORBIDADEDIABETESSIVEP string `json:"comorbidadediabetessivep,omitempty"`
	COMORBIDADEHEMATOLOGIASIVEP string `json:"comorbidadehematologiasivep,omitempty"`
	COMORBIDADEIMUNODEFICIENCIASIVEP string `json:"comorbidadeimunodeficienciasivep,omitempty"`
	COMORBIDADENEUROLOGIASIVEP string `json:"comorbidadeneurologiasivep,omitempty"`
	COMORBIDADEOBESIDADESIVEP string `json:"comorbidadeobesidadesivep,omitempty"`
	COMORBIDADEPNEUMOPATIASIVEP string `json:"comorbidadepneumopatiasivep,omitempty"`
	COMORBIDADEPUERPERASIVEP string `json:"comorbidaderenalsivep,omitempty"`
	COMORBIDADERENALSIVEP string `json:"comorbidaderenalsivep,omitempty"`
	COMORBIDADESINDROMEDOWNSIVEP string `json:"comorbidadesindromedownsivep,omitempty"`
	DATACOLETAEXAME string `json:"datacoletaexame,omitempty"`
	DATAENTRADAUTISSVEP string `json:"dataentradautissvep,omitempty"`
	DATAEVOLUCAOCASOSIVEP string `json:"dataevolucaocasosivep,omitempty"`
	DATAINICIOSINTOMAS string `json:"datainiciosintomas,omitempty"`
	DATAINTERNACAOSIVEP string `json:"datainternacaosivep,omitempty"`
	DATANOTIFICACAO string `json:"datanotificacao,omitempty"`
	DATAOBITO string `json:"dataobito,omitempty"`
	DATARESULTADOEXAME string `json:"dataresultadoexame,omitempty"`
	DATASAIDAUTISSVEP string `json:"datasaidautissvep,omitempty"`
	DATASOLICITACAOEXAME string `json:"datasolicitacaoexame,omitempty"`
	ESTADOPACIENTE string `json:"estadopaciente,omitempty"`
	EVOLUCAOCASOSIVEP string `json:"evolucaocasosivep,omitempty"`
	IDADEPACIENTE string `json:"idadepaciente,omitempty"`
	IDSIVEP string `json:"idsivep,omitempty"`
	MUNICIPIOPACIENTE string `json:"municipiopaciente,omitempty"`
	OBITOCONFIRMADO string `json:"obitoconfirmado,omitempty"`
	PAISPACIENTE string `json:"paispaciente,omitempty"`
	RESULTADOFINALEXAME string `json:"resultadofinalexame,omitempty"`
	SEXOPACIENTE string `json:"sexopaciente,omitempty"`
	//34
}




//var casosS casosPorSexo
//var casos  dadosUsuario
//casos := []dadosUsuario{}

func main() {
	rotas := mux.NewRouter().StrictSlash(true)

	rotas.HandleFunc("/covid", getAll).Methods("GET")
	rotas.HandleFunc("/covid/numeroExames", getNumeroExames).Methods("GET")
	rotas.HandleFunc("/covid/casosConfirmados", getCasosConfirmados).Methods("GET")
	rotas.HandleFunc("/covid/casosConfirmados/{municipio}", getCasosConfirmadosMunicipio).Methods("GET")
	rotas.HandleFunc("/covid/casosConfirmadosSexo/{sexo}", getCasosConfirmadosPorSexo).Methods("GET")
	rotas.HandleFunc("/covid/casosInvestigacao", getCasosInvestigacao).Methods("GET")
	rotas.HandleFunc("/covid/obitos", getAllObitos).Methods("GET")
	rotas.HandleFunc("/covid/obitos/{municipio}", getObitosMunicipio).Methods("GET")
	rotas.HandleFunc("/covid/letalidade", getLetalidade).Methods("GET")
	rotas.HandleFunc("/covid/casosRecuperados", getCasosRecuperados).Methods("GET")
	rotas.HandleFunc("/covid/casos/{sexo}", getCasosPorSexo).Methods("GEt")
	rotas.HandleFunc("/covid/casosConfimadosData/{data}", getCasosConfirmadosData).Methods("GEt")
	//casos por sexo, por idade, por municipio. falta corrigir a questão de organizar as
	//requests que recebem parametros. :-)

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, rotas))

}


