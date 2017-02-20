package common

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"

	"path"

	"github.com/TSAP-Laval/models"
	"github.com/jinzhu/gorm"

	// Import global pour utiliser sqlite avec gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// SeedData seed des données contenues dans un dossier passé en paramètres
func SeedData(dbType string, connString string, dataFolder string) error {
	var err error

	db, err := gorm.Open(dbType, connString)

	if err != nil {
		return err
	}
	defer db.Close()

	db.CreateTable(&models.Entraineur{}, &models.Joueur{},
		&models.Saison{}, &models.Lieu{}, &models.TypeAction{},
		&models.Sport{}, &models.Zone{}, &models.Niveau{},
		&models.Equipe{}, &models.Partie{}, &models.Action{},
		&models.Administrateur{}, &models.Position{},
		&models.JoueurPositionPartie{}, &models.Metrique{},
		&models.Video{})

	var joueursData []models.Joueur
	err = jsonLoad(path.Join(dataFolder, "joueurs.json"), &joueursData)
	if err != nil {
		return err
	}
	for _, joueur := range joueursData {
		db.Create(&joueur)
	}

	var entraineursData []models.Entraineur
	err = jsonLoad(path.Join(dataFolder, "entraineurs.json"), &entraineursData)
	if err != nil {
		return err
	}
	for _, entraineur := range entraineursData {
		db.Create(&entraineur)
	}

	var lieuxData []models.Lieu
	err = jsonLoad(path.Join(dataFolder, "lieux.json"), &lieuxData)
	if err != nil {
		return err
	}
	for _, lieu := range lieuxData {
		db.Create(&lieu)
	}

	var saisonData []models.Saison
	err = jsonLoad(path.Join(dataFolder, "saisons.json"), &saisonData)
	if err != nil {
		return err
	}
	for _, lieu := range saisonData {
		db.Create(&lieu)
	}

	var typeActionData []models.TypeAction
	err = jsonLoad(path.Join(dataFolder, "typesactions.json"), &typeActionData)
	if err != nil {
		return err
	}
	for i := 0; i < len(typeActionData); i++ {
		db.Create(&typeActionData[i])
	}

	var sportData []models.Sport
	err = jsonLoad(path.Join(dataFolder, "sports.json"), &sportData)
	if err != nil {
		return err
	}
	for _, sport := range sportData {
		db.Create(&sport)
	}

	var niveauData []models.Niveau
	err = jsonLoad(path.Join(dataFolder, "niveaux.json"), &niveauData)
	if err != nil {
		return err
	}
	for _, niveau := range niveauData {
		db.Create(&niveau)
	}

	var equipeData []models.Equipe
	err = jsonLoad(path.Join(dataFolder, "equipes.json"), &equipeData)
	if err != nil {
		return err
	}

	playerIndex := 0

	for i := 0; i < len(equipeData); i++ {
		equipe := &equipeData[i]
		x := &models.Sport{}
		if i%2 == 0 {
			db.First(x)
			equipe.Sport = *x
		} else {
			db.Last(x)
			equipe.Sport = *x
		}

		y := &models.Niveau{}
		db.First(y, rand.Intn(6)+1)
		equipe.Niveau = *y

		for j := 0; j < 9; j++ {
			equipe.Joueurs = append(equipe.Joueurs, joueursData[playerIndex])
			db.Model(&joueursData[playerIndex]).Association("Joueurs").Append(equipe)
			playerIndex++
		}

		db.Create(&equipe)
	}

	admin := models.Administrateur{Email: "admin@admin.com", PassHash: "admin"}
	db.Create(&admin)

	var positionData []models.Position
	err = jsonLoad(path.Join(dataFolder, "positions.json"), &positionData)
	if err != nil {
		return err
	}
	for i := 0; i < len(positionData); i++ {
		db.Create(&positionData[i])
	}

	var metriqueData []models.Metrique
	err = jsonLoad(path.Join(dataFolder, "metriques.json"), &metriqueData)
	if err != nil {
		return err
	}
	for _, metrique := range metriqueData {
		db.Create(&metrique)
	}

	video := models.Video{Path: "aucun video", AnalyseTermine: false}
	db.Create(&video)

	var partieData []models.Partie
	err = jsonLoad(path.Join(dataFolder, "parties.json"), &partieData)
	if err != nil {
		return err
	}
	for i := 0; i < len(partieData); i++ {
		partie := &partieData[i]
		nb1 := 1
		nb2 := 1
		for nb1 == nb2 {
			nb1 = rand.Intn(len(equipeData)) + 1
			nb2 = rand.Intn(len(equipeData)) + 1
		}

		equipe1 := &models.Equipe{}
		db.First(equipe1, nb1)
		db.Model(equipe1).Association("Joueurs").Find(&(equipe1.Joueurs))

		partie.EquipeMaison = *equipe1
		equipe2 := &models.Equipe{}
		db.First(equipe2, nb2)
		db.Model(equipe2).Association("Joueurs").Find(&(equipe2.Joueurs))
		partie.EquipeAdverse = *equipe2
		saison := &models.Saison{}

		db.First(saison, rand.Intn(3)+1)
		partie.Saison = *saison
		lieu := &models.Lieu{}
		db.First(lieu, rand.Intn(100)+1)
		partie.Lieu = *lieu
		video := &models.Video{}
		db.First(video)
		partie.Video = *video

		db.Create(partie)
	}

	zoneOff := models.Zone{Nom: "offensive"}
	db.Create(&zoneOff)
	zoneDef := models.Zone{Nom: "defensive"}
	db.Create(&zoneDef)

	for i := 0; i < len(partieData); i++ {
		// On pige deux équipes
		ind1 := 1
		ind2 := 1
		for ind1 == ind2 {
			ind1 = rand.Intn(len(equipeData))
			ind2 = rand.Intn(len(equipeData))
		}

		team1 := &partieData[i].EquipeMaison
		team2 := &partieData[i].EquipeAdverse

		for k := 0; k < len(team1.Joueurs); k++ {
			db.Save(&models.JoueurPositionPartie{
				Joueur:   team1.Joueurs[k],
				Partie:   partieData[i],
				Position: positionData[rand.Intn(len(positionData))],
				Equipe:   *team1,
			})
		}

		for k := 0; k < len(team2.Joueurs); k++ {
			db.Save(&models.JoueurPositionPartie{
				Joueur:   team2.Joueurs[k],
				Partie:   partieData[i],
				Position: positionData[rand.Intn(len(positionData))],
				Equipe:   *team2,
			})
		}

	}

	pickAction := func(team *models.Equipe, match *models.Partie) *models.Action {
		joueur := team.Joueurs[rand.Intn(len(team.Joueurs))]
		typeAction := typeActionData[rand.Intn(len(typeActionData))]

		var pos bool
		if typeAction.Nom == "BC" || typeAction.Nom == "PO" || typeAction.Nom == "TB" {
			pos = true
		}

		var zone models.Zone

		X1 := rand.Float64()
		Y1 := rand.Float64()

		var X2 float64
		var Y2 float64

		if rand.Int()%2 == 0 {
			zone = zoneOff
			X2 = -1
			Y2 = -1
		} else {
			zone = zoneDef
			X2 = rand.Float64()
			Y2 = rand.Float64()
		}

		t := time.Duration(40)

		return &models.Action{
			TypeAction:      typeAction,
			ActionPositive:  pos,
			Zone:            zone,
			Partie:          *match,
			X1:              X1,
			X2:              X2,
			Y1:              Y1,
			Y2:              Y2,
			Temps:           t,
			PointageMaison:  0,
			PointageAdverse: 0,
			Joueur:          joueur,
		}
	}

	for i := 0; i < len(partieData); i++ {
		var actionCount int
		if i < 4 {
			actionCount = 1000
		} else {
			actionCount = 50
		}

		for j := 0; j <= actionCount; j++ {
			if j%2 == 0 {
				db.Save(pickAction(&partieData[i].EquipeMaison, &partieData[i]))
			} else {
				// Action visiteur
				db.Save(pickAction(&partieData[i].EquipeAdverse, &partieData[i]))
			}
		}
	}

	return err
}

// jsonLoad est une fonction générique qui charge un fichier JSON
// dans une struct passée en référence
func jsonLoad(path string, out interface{}) error {
	raw, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, out)

	if err != nil {
		return err
	}

	return nil
}
