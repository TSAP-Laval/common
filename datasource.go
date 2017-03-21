package common

import (
	"fmt"

	"github.com/TSAP-Laval/models"
	"github.com/jinzhu/gorm"
)

// IDatasource représente l'interface abastraite
// d'une struct capable de servir de source de données
// pour l'application
type IDatasource interface {
	GetCurrentSeason() (*models.Saison, error)
	GetSeasons() (*[]models.Saison, error)
	GetTeam(teamID uint) (*models.Equipe, error)
	GetPlayer(playerID uint) (*models.Joueur, error)
	GetMatches(teamID uint, seasonID uint) (*[]models.Partie, error)
	GetMatchPosition(playerID uint, matchID uint) (*models.Position, error)
	GetPositions(playerID uint) (*[]models.Position, error)
	GetMatch(matchID uint) (*models.Partie, error)
	GetLatestMatch(teamID uint) (*models.Partie, error)
	GetCoach(coachID uint) (*models.Entraineur, error)
	CreateMetric(name string, formula string, description string, teamID uint) error
	UpdateMetric(metricID uint, name string, formula string, description string) error
	DeleteMetric(metricID uint) error
	GetMetrics(teamID uint) (*[]models.Metrique, error)
}

// Datasource représente une connexion à une base de
// données
type Datasource struct {
	dbType string
	dbConn string
}

// NewDatasource retourne une nouvelle datasource
func NewDatasource(dbType string, dbConnString string) *Datasource {
	return &Datasource{dbType: dbType, dbConn: dbConnString}
}

// GetCurrentSeason retourne la saison en cours
func (d *Datasource) GetCurrentSeason() (*models.Saison, error) {
	db, err := gorm.Open(d.dbType, d.dbConn)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	s := models.Saison{}

	db.Last(&s)

	return &s, err
}

// GetSeasons retourne la liste des saisons
func (d *Datasource) GetSeasons() (*[]models.Saison, error) {
	db, err := gorm.Open(d.dbType, d.dbConn)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	s := []models.Saison{}

	db.Find(&s)

	return &s, nil
}

// GetTeam retourne l'instance d'Equipe correspondant au ID
func (d *Datasource) GetTeam(teamID uint) (*models.Equipe, error) {
	var err error

	db, err := gorm.Open(d.dbType, d.dbConn)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	t := models.Equipe{}

	db.First(&t, teamID)

	if t.ID != teamID {
		return nil, fmt.Errorf("Team %d not found", teamID)
	}
	db.Model(&t).Association("Joueurs").Find(&t.Joueurs)

	return &t, err
}

// GetPlayer retourne l'instance de player correspondant au ID
func (d *Datasource) GetPlayer(playerID uint) (*models.Joueur, error) {
	var err error

	db, err := gorm.Open(d.dbType, d.dbConn)

	defer db.Close()

	if err != nil {
		return nil, err
	}

	j := models.Joueur{}

	db.First(&j, playerID)

	if j.ID != playerID {
		return nil, fmt.Errorf("Player %d not found", playerID)
	}

	db.Model(&j).Association("Equipes").Find(&j.Equipes)

	return &j, err
}

// GetMatches obtient les match d'un joueur (pour une équipe, pour une saison, pour une position?)
func (d *Datasource) GetMatches(teamID uint, seasonID uint) (*[]models.Partie, error) {
	var err error

	db, err := gorm.Open(d.dbType, d.dbConn)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	matches := []models.Partie{}

	t := int(teamID)
	s := int(seasonID)

	db.Where(models.Partie{EquipeMaisonID: t, SaisonID: s}).Or(models.Partie{EquipeAdverseID: t, SaisonID: s}).Find(&matches)

	for i := 0; i < len(matches); i++ {
		matches[i].Expand(db)
		for j := 0; j < len(matches[i].Actions); j++ {
			matches[i].Actions[j].Expand(db)
		}
	}

	return &matches, err
}

// GetMatchPosition retourne la position occupée par un joueur pour un match
func (d *Datasource) GetMatchPosition(playerID uint, matchID uint) (*models.Position, error) {
	var err error

	db, err := gorm.Open(d.dbType, d.dbConn)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	posPartie := models.JoueurPositionPartie{}

	db.Where(models.JoueurPositionPartie{JoueurID: int(playerID), PartieID: int(matchID)}).First(&posPartie)

	if posPartie.JoueurID != int(playerID) {
		return nil, fmt.Errorf("No position for player %d", playerID)
	}

	posPartie.Expand(db)

	pos := posPartie.Position

	return &pos, nil
}

// GetPositions retourne une liste de positions occupées par le joueur
func (d *Datasource) GetPositions(playerID uint) (*[]models.Position, error) {
	var err error

	db, err := gorm.Open(d.dbType, d.dbConn)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	positionsMap := make(map[models.Position]bool)
	positions := []models.Position{}

	positionParties := []models.JoueurPositionPartie{}

	db.Where(models.JoueurPositionPartie{JoueurID: int(playerID)}).Find(&positionParties)

	for i := 0; i < len(positionParties); i++ {
		positionParties[i].Expand(db)
		positionsMap[positionParties[i].Position] = true
	}

	for k := range positionsMap {
		alreadyAdded := false

		// On vérifie l'absence de la position de la liste
		// (Il n'y a pas de .indexOf() en Go)
		for iP := 0; iP < len(positions) && !alreadyAdded; iP++ {
			p := positions[iP]
			if k.ID == p.ID {
				alreadyAdded = true
			}
		}

		if !alreadyAdded {
			positions = append(positions, k)
		}
	}

	return &positions, nil
}

// GetMatch obtient toutes les informations sur un match
func (d *Datasource) GetMatch(matchID uint) (*models.Partie, error) {
	var err error

	db, err := gorm.Open(d.dbType, d.dbConn)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	match := models.Partie{}

	db.First(&match, matchID)

	if match.ID != matchID {
		return nil, fmt.Errorf("Match %d not found", matchID)
	}

	match.Expand(db)
	for i := 0; i < len(match.Actions); i++ {
		match.Actions[i].Expand(db)
	}

	return &match, err
}

// GetLatestMatch retourne le dernie match d'une equipe dont l'ID est reçu en paramètre.
func (d *Datasource) GetLatestMatch(teamID uint) (*models.Partie, error) {
	var err error

	db, err := gorm.Open(d.dbType, d.dbConn)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	match := models.Partie{}

	tID := int(teamID)

	db.Where(models.Partie{EquipeMaisonID: tID}).Or(models.Partie{EquipeAdverseID: tID}).Last(&match)

	if match.EquipeMaisonID != tID && match.EquipeAdverseID != tID {
		return nil, fmt.Errorf("Last match for team %d not found", teamID)
	}

	match.Expand(db)

	for i := 0; i < len(match.Actions); i++ {
		match.Actions[i].Expand(db)
	}

	return &match, err
}

// GetCoach retourne l'instance de l'entraineur correspondant au ID
func (d *Datasource) GetCoach(coachID uint) (*models.Entraineur, error) {
	var err error

	db, err := gorm.Open(d.dbType, d.dbConn)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	e := models.Entraineur{}

	db.First(&e, coachID)

	if e.ID != coachID {
		return nil, fmt.Errorf("Coach %d not found", coachID)
	}

	db.Model(&e).Association("Equipes").Find(&e.Equipes)

	return &e, err
}

// CreateMetric crée une nouvelle métrique
func (d *Datasource) CreateMetric(name string, formula string, description string, teamID uint) error {

	db, err := gorm.Open(d.dbType, d.dbConn)

	if err != nil {
		return err
	}

	defer db.Close()

	metric := models.Metrique{Nom: name, Equation: formula, Description: description, EquipeID: int(teamID)}

	db.Create(&metric)

	return nil
}

// UpdateMetric modifie une métrique existante
func (d *Datasource) UpdateMetric(metricID uint, name string, formula string, description string) error {

	db, err := gorm.Open(d.dbType, d.dbConn)

	if err != nil {
		return err
	}

	defer db.Close()

	m := models.Metrique{}

	db.First(&m, metricID)

	if m.ID != metricID {
		return fmt.Errorf("Metric %d not found", metricID)
	}

	m.Nom = name
	m.Equation = formula
	m.Description = description

	db.Save(&m)

	return nil
}

// DeleteMetric supprime une métrique existante
func (d *Datasource) DeleteMetric(metricID uint) error {

	db, err := gorm.Open(d.dbType, d.dbConn)

	if err != nil {
		return err
	}

	defer db.Close()

	m := models.Metrique{}

	db.First(&m, metricID)

	if m.ID != metricID {
		return fmt.Errorf("Metric %d not found", metricID)
	}

	db.Delete(&m)

	return nil
}

// GetMetrics retourne une liste de toutes les métriques d'une équipe
func (d *Datasource) GetMetrics(teamID uint) (*[]models.Metrique, error) {
	db, err := gorm.Open(d.dbType, d.dbConn)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	metrics := []models.Metrique{}

	db.Where(&models.Metrique{EquipeID: int(teamID)}).Find(&metrics)

	return &metrics, nil
}
