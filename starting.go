package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"io/ioutil"	
	 s "strings"
)


type Type struct {
	Name string `json:"name"`
	WeakAgainst []string `json:"weakAgainst"`
	EffectiveAgainst []string `json:"effectiveAgainst"`
}

type Pokemon struct {
	Number         string   `json:"Number"`
	Name           string   `json:"Name"`
	Classification string   `json:"Classification"`
	TypeI          []string `json:"Type I"`
	TypeII         []string `json:"Type II,omitempty"`
	Weaknesses     []string `json:"Weaknesses"`
	FastAttackS    []string `json:"Fast Attack(s)"`
	Weight         string   `json:"Weight"`
	Height         string   `json:"Height"`
	Candy          struct {
		Name     string `json:"Name"`
		FamilyID int    `json:"FamilyID"`
	} `json:"Candy"`
	NextEvolutionRequirements struct {
		Amount int    `json:"Amount"`
		Family int    `json:"Family"`
		Name   string `json:"Name"`
	} `json:"Next Evolution Requirements,omitempty"`
	NextEvolutions []struct {
		Number string `json:"Number"`
		Name   string `json:"Name"`
	} `json:"Next evolution(s),omitempty"`
	PreviousEvolutions []struct {
		Number string `json:"Number"`
		Name   string `json:"Name"`
	} `json:"Previous evolution(s),omitempty"`
	SpecialAttacks      []string `json:"Special Attack(s)"`
	BaseAttack          int      `json:"BaseAttack"`
	BaseDefense         int      `json:"BaseDefense"`
	BaseStamina         int      `json:"BaseStamina"`
	CaptureRate         float64  `json:"CaptureRate"`
	FleeRate            float64  `json:"FleeRate"`
	BuddyDistanceNeeded int      `json:"BuddyDistanceNeeded"`
}

type Move struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Damage int `json:"damage"`
	Energy int `json:"energy"`
	Dps float64 `json:"dps"`
	Duration int `json:"duration"`
}

type BaseData struct {
	Types    []Type    `json:"types"`
	Pokemons []Pokemon `json:"pokemons"`
	Moves    []Move    `json:"moves"`
}
func Listele(w http.ResponseWriter, r *http.Request){
   dosya,err := ioutil.ReadFile("./data.json")
    if err != nil {
        fmt.Println("Hata")
       os.Exit(1)
    }
	vars :=mux.Vars(r)
	tip:=vars["tip"]
    
	deg:=&BaseData{}
    json.Unmarshal(dosya, &deg)
	tipler := deg.Pokemons
	json.Unmarshal(dosya, &tipler)

 	for k,v:= range tipler {
		if deg.Pokemons[k].TypeI[0]==tip {
       fmt.Fprintf(w,"<html><h7><br>%s<br>&nbsp Weight:%s<br>&nbsp Height:%s<br>&nbsp BaseAttack:%d<br>&nbsp BaseDefense:%d<br>&nbsp BaseStamina:%d<br>&nbsp FastAttackS:<br>&nbsp&nbsp  %v<br>&nbsp&nbsp</h7></html",v.Name,v.Weight,v.Height,v.BaseAttack,v.BaseDefense,v.BaseStamina,s.Join( v.FastAttackS,"<br>&nbsp&nbsp  "))	
	    }
    }
	 fmt.Fprintf(w,"Sayfa bulunanamadı")
}
func Pokiler(w http.ResponseWriter, r *http.Request){
   dosya,err := ioutil.ReadFile("./data.json")
    if err != nil {
        fmt.Println("Hata")
       os.Exit(1)
    }
	vars :=mux.Vars(r)
	ad:=vars["ad"]

	deg:=&BaseData{}
    json.Unmarshal(dosya, &deg)
	pokiler := deg.Pokemons

 	for k,v:= range pokiler {
		if deg.Pokemons[k].Name==ad {
          fmt.Fprintf(w,"<html><h7><br>%s<br>&nbsp Weight:%s<br>&nbsp Height:%s<br>&nbsp BaseAttack:%d<br>&nbsp BaseDefense:%d<br>&nbsp BaseStamina:%d<br>&nbsp FastAttackS:<br>&nbsp&nbsp  %v<br>&nbsp&nbsp</h7></html",v.Name,v.Weight,v.Height,v.BaseAttack,v.BaseDefense,v.BaseStamina,s.Join( v.FastAttackS,"<br>&nbsp&nbsp  "))
		}
	}
	fmt.Fprintf(w,"Sayfa bulunanamadı")	 
}
func Tipler(w http.ResponseWriter, r *http.Request){
    dosya, err := ioutil.ReadFile("./data.json")
    if err != nil {
        fmt.Println("Hata")
        os.Exit(1)
    }
	deg:=&BaseData{}
    json.Unmarshal(dosya, &deg)
	t :=deg.Types
	vars :=mux.Vars(r)
	pt:=vars["pt"]
	for k,v:= range t {
		if deg.Types[k].Name==pt{
		  fmt.Fprintf(w,"<html>Pokemon Type: %s<br> WeakAgainst:<br>-%s<br>EffectiveAgainst:<br>-%s</html",v.Name,s.Join( v.WeakAgainst, "<br>-"),s.Join(v.EffectiveAgainst, "<br>-"))
		}			  		
	 }	
	fmt.Fprintf(w,"Sayfa bulunanamadı")
}
 func Oyun(w http.ResponseWriter, r *http.Request){
    dosya, err := ioutil.ReadFile("./data.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
	deg:=&BaseData{}
    json.Unmarshal(dosya, &deg)
	m := deg.Moves
	vars :=mux.Vars(r)
	isim:=vars["isim"]
	
 	for k,v:= range m {
		if deg.Moves[k].Name==isim{
			fmt.Fprintf(w,"<html><h7>ID: %d<br>Name: %s<br> Type: %s<br> Damage: %d<br>Energy: %d<br> Dps: %g<br> Duration: %d<br></h7><html" ,v.ID,v.Name,v.Type,v.Damage,v.Energy,v.Dps,v.Duration)
		}			
		 }
	      fmt.Fprintf(w,"Sayfa bulunanamadı")
}

func main() {
   gorillar := mux.NewRouter()
   gorillar.HandleFunc("/{isim}", Oyun)               //localhost:8080/Wrap 
   gorillar.HandleFunc("/list/type={tip}", Listele)   //localhost:8080/list/type=Grass
   gorillar.HandleFunc("/get/{pt}",Tipler)            //localhost:8080/get/Bug
   gorillar.HandleFunc("/name/{ad}",Pokiler)          //localhost:8080/name/Bulbasaur
   log.Println("Başlatıldı:8080")
   http.ListenAndServe(":8080",gorillar)
   
}
