package common

import (
	"fmt"

	"github.com/TSAP-Laval/models"
	"github.com/jinzhu/gorm"
)

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

// GetMatches obtient les match d'un joueur (pour une équipe, pour une saison)
func (d *Datasource) GetMatches(teamID uint, seasonID uint) ([]models.Partie, error) {
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

	return matches, err
}

// GetAllTeamsMatchs retourne tous matchs de toutes les équipes d'une saison donnée.
func (d *Datasource) GetAllTeamsMatchs(seasonID uint) ([]models.Partie, error) {
	var err error

	db, err := gorm.Open(d.dbType, d.dbConn)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	if seasonID <= 0 {
		return nil, fmt.Errorf("Invalid SeasonID value in GetAllTeamsMatchs function")
	}

	// Toutes les équipes.
	allTeams := []models.Equipe{}

	allMatches := []models.Partie{}

	matches := []models.Partie{}

	//Get all teams. NOT REALLY GOOD BUT THERE IS CURRENTLY NO WAY TO HAVE
	//A LIST OF TEAMS BASED ON A GIVEN SEASON.
	if err := db.Find(&allTeams).Error; err != nil {
		return nil, fmt.Errorf("No team found in GetAllTeamsMatchs function")
	}

	for _, team := range allTeams {
		matches, err = d.GetMatches(team.ID, seasonID)

		for _, match := range matches {
			allMatches = append(allMatches, match)
		}
	}

	return allMatches, err
}

// GetAllPositions retourne toutes les positions des joueurs.
func (d *Datasource) GetAllPositions() ([]models.Position, error) {
	var err error

	db, err := gorm.Open(d.dbType, d.dbConn)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	allPositions := []models.Position{}

	if err := db.Find(&allPositions).Error; err != nil {
		return nil, fmt.Errorf("No position found in GetAllPositions function")
	}

	return allPositions, err

}

// GetPositions retourne une liste de positions occupées par le joueur
func (d *Datasource) GetPositions(playerID int) ([]models.Position, error) {
	var err error

	db, err := gorm.Open(d.dbType, d.dbConn)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	positionsMap := make(map[models.Position]bool)
	positions := []models.Position{}

	positionParties := []models.JoueurPositionPartie{}

	db.Where(models.JoueurPositionPartie{JoueurID: playerID}).Find(&positionParties)

	for i := 0; i < len(positionParties); i++ {
		positionParties[i].Expand(db)
		positionsMap[positionParties[i].Position] = true
	}

	for k := range positionsMap {
		positions = append(positions, k)
	}

	return positions, nil
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

// GetLastMatch retourne le dernier match joué par une équipe.
func (d *Datasource) GetLastMatch(teamID uint) (*models.Partie, error) {
	var err error

	db, err := gorm.Open(d.dbType, d.dbConn)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	lastMatch := models.Partie{}

	tID := int(teamID)

	db.Where(models.Partie{EquipeMaisonID: tID}).Or(models.Partie{EquipeAdverseID: tID}).Last(&lastMatch)

	return &lastMatch, err

}

// GetCoach retourne l'instance de l'entraineur correspondant au ID
func (d *Datasource) GetCoach(coachID uint) (*models.Entraineur, error) {
	var err error

	db, err := gorm.Open(d.dbType, d.dbConn)

	defer db.Close()

	if err != nil {
		return nil, err
	}

	e := models.Entraineur{}

	db.First(&e, coachID)

	if e.ID != coachID {
		return nil, fmt.Errorf("Coach %d not found", coachID)
	}

	db.Model(&e).Association("Equipes").Find(&e.Equipes)

	return &e, err
}
