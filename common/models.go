package common

import (
	"github.com/jinzhu/gorm"
)

// Entraineur est une modelisation d'un entraineur
// comprenant son id, son nom, son prenom, son email
// et ses équipes
type Entraineur struct {
	gorm.Model
	Prenom   string `gorm:"size:45"`
	Nom      string `gorm:"size:45"`
	Email    string `gorm:"size:45"`
	PassHash string
	Token    string
	Equipes  []Equipe `gorm:"many2many:entraineur_equipe;"`
}

// Joueur est une modèlisation d'un joueur
// d'une équipe sportive
type Joueur struct {
	gorm.Model
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
	Nom     string `gorm:"size:45"`
	Ville   string `gorm:"size:45"`
	Adresse string `gorm:"size:45"`
}

// TypeAction est une modélisation des types
// possibles  d'action
type TypeAction struct {
	gorm.Model
	Nom string `gorm:"size:45"`
}

// Sport est une modélisation des noms possibles
// des sports pratiqués
type Sport struct {
	gorm.Model
	Nom string `gorm:"size:45"`
}

// Zone est une modélisation de la zone de jeu
// de l'action
type Zone struct {
	gorm.Model
	Nom string `gorm:"size:45"`
}

// Niveau est une modélisation du niveau d'un joueur
type Niveau struct {
	gorm.Model
	Nom string `gorm:"size:45"`
}

// Equipe est une modélisation d'une équipe de Joueurs
// dirigiée par un ou plusieurs entraineurs
type Equipe struct {
	gorm.Model
	Nom         string `gorm:"size:45"`
	Ville       string `gorm:"size:45"`
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
	Lieux           int
	Video           Video
	VideoID         int
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
	X               int
	Y               int
	Temps           string
	PointageMaison  int
	PointageAdverse int
	Joueur          Joueur
	JoueurID        int
}

type Administrateur struct {
	gorm.Model
	Email                 string
	PassHash              string
	TokenReinitialisation string
	TokenConexion         string
}

type Position struct {
	gorm.Model
	Nom string `gorm:"size:45"`
}

type JoueurPositionPartie struct {
	gorm.Model
	Joueur     Joueur
	JoueurID   int
	Partie     Partie
	PartieID   int
	Position   Position
	PositionID int
}

type Metrique struct {
	gorm.Model
	Nom      string `gorm:"primary_key"`
	Equation string `gorm:"primary_key"`
	Equipe   Equipe
	EquipeID int
}

type Video struct {
	gorm.Model
	Path           string
	AnalyseTermine bool
}
