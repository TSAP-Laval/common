package common

import (
	"github.com/jinzhu/gorm"
)

// Entraineur est une modelisation d'un entraineur
// comprenant son id, son nom, son prenom, son email
// et ses équipes
type Entraineur struct {
	gorm.Model
	Prenom   string
	Nom      string
	Email    string
	PassHash string
	Token    string
	Equipes  []Equipe `gorm:"many2many:entraineur_equipe;"`
}

// Joueur est une modèlisation d'un joueur
// d'une équipe sportive
type Joueur struct {
	gorm.Model
	Prenom                string
	Nom                   string
	Numero                int
	Email                 string
	PassHash              string
	TokenInvitation       string
	TokenReinitialisation string
	TokenConnexion        string
	Equipes               []Equipe `gorm:"many2many:joueur_equipe;"`
}

// Saison est un modèlisation d'une saison
// sportive, comprenant son ID et ses années
type Saison struct {
	gorm.Model
	Annees string `gorm:"size:10"`
}

// Lieu est une modélisation des endroits possibles
// pour un match
type Lieu struct {
	gorm.Model
	Nom     string
	Ville   string
	Adresse string
}

// TypeAction est une modélisation des types
// possibles  d'action
type TypeAction struct {
	gorm.Model
	Nom string
}

// Sport est une modélisation des noms possibles
// des sports pratiqués
type Sport struct {
	gorm.Model
	Nom string
}

// Zone est une modélisation de la zone de jeu
// de l'action
type Zone struct {
	gorm.Model
	Nom string
}

// Niveau est une modélisation du niveau d'un joueur
type Niveau struct {
	gorm.Model
	Nom string
}

// Equipe est une modélisation d'une équipe de Joueurs
// dirigiée par un ou plusieurs entraineurs
type Equipe struct {
	gorm.Model
	Nom         string
	Ville       string
	Sport       Sport
	SportID     int
	Niveau      Niveau
	NiveauID    int
	Entraineurs []Entraineur `gorm:"many2many:entraineur_equipe;"`
	Joueurs     []Joueur     `gorm:"many2many:joueur_equipe;"`
}

// Partie est une modélisation des informations
// sur une partie entre deux équipes
type Partie struct {
	gorm.Model
	EquipeMaison    Equipe
	EquipeMaisonID  int
	EquipeAdverse   Equipe
	EquipeAdverseID int
	Saison          Saison
	SaisonID        int
	Lieu            Lieu
	LieuID          int
	Video           Video
	VideoID         int
	Date            string
}

// Action est une modélisation des informations sur une
// action exécutée par un joueur
type Action struct {
	gorm.Model
	TypeAction      TypeAction
	ActionPositive  bool
	Zone            Zone
	ZoneID          int
	Partie          Partie
	PartieID        int
	X1              int
	Y1              int
	X2              int
	Y2              int
	Temps           string
	PointageMaison  int
	PointageAdverse int
	Joueur          Joueur
	JoueurID        int
}

// Administrateur est une modélisation des informations
// de connexion d'un Administrateur
type Administrateur struct {
	gorm.Model
	Email                 string
	PassHash              string
	TokenReinitialisation string
	TokenConnexion        string
}

// Position est une modélisation du nom de la Position
// du joueur
type Position struct {
	gorm.Model
	Nom string
}

// JoueurPositionPartie est une modélisation de la Position
// d'un joueur dans une partie
type JoueurPositionPartie struct {
	gorm.Model
	Joueur     Joueur
	JoueurID   int
	Partie     Partie
	PartieID   int
	Position   Position
	PositionID int
}

// Metrique est la modélisation d'une unité de calcul
type Metrique struct {
	gorm.Model
	Nom      string `gorm:"primary_key"`
	Equation string `gorm:"primary_key"`
	Equipe   Equipe
	EquipeID int
}

// Video est la modélisation des informations
// sur une Video
type Video struct {
	gorm.Model
	Path           string
	AnalyseTermine bool
}
