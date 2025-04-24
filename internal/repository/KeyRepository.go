package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"

	"github.com/jmoiron/sqlx"
)

type KeyRepository struct {
	db *sqlx.DB
}

func NewKeyRepository(db *sqlx.DB) *KeyRepository {
	return &KeyRepository{db: db}
}

// Crear un configuraciones del sistema
func (r *KeyRepository) CreateKey(systemSetting *domain.Keys) error {
	query := `
		INSERT INTO keys (
			name, system_id, description, config_key, config_value, created_by, updated_by, status
		) VALUES (
			:name, :system_id, :description, :config_key, :config_value, :created_by, :updated_by, :status
		)
	`
	_, err := r.db.NamedExec(query, systemSetting)
	if err != nil {
		fmt.Printf("Error creando la configuraciones del sistema: %v\n", err)
		return err
	}
	return nil
}

// Obtener todas las áreas de trabajo
func (r *KeyRepository) GetAllKeys() ([]domain.Keys, error) {
	var systemSettings []domain.Keys
	query := `
		SELECT 
			k.id, k.system_id, k.name, k.description, k.config_key, k.config_value, k.created_by, k.updated_by, k.status,
			COALESCE(s.id, 0) AS "system.id", COALESCE(s.name, '') AS "system.name"
		FROM keys k
		LEFT JOIN systems s ON k.system_id = s.id
		ORDER BY k.id DESC
	`
	err := r.db.Select(&systemSettings, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return systemSettings, nil
}

// Obtener un configuraciones del sistema por Id
func (r *KeyRepository) GetKeyById(systemSettingId int) (*domain.Keys, error) {
	var systemSetting domain.Keys
	query := `
		SELECT 
			k.id, k.system_id, k.name, k.description, k.config_key, k.config_value, k.created_by, k.updated_by, k.status,
			COALESCE(s.id, 0) AS "system.id", COALESCE(s.name, '') AS "system.name"
		FROM keys k
		LEFT JOIN systems s ON k.system_id = s.id
		WHERE k.id = $1
	`

	// Usar Get para mapear automáticamente los resultados
	err := r.db.Get(&systemSetting, query, systemSettingId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Área de trabajo no encontrado
		}
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	return &systemSetting, nil
}

// Actualizar un configuraciones del sistema
func (r *KeyRepository) UpdateKey(systemSetting *domain.Keys) error {
	query := `
		UPDATE keys SET
			system_id = :system_id,
			name = :name,
			description = :description,
			config_key = :config_key,
			config_value = :config_value,
			created_by = :created_by,
			updated_by = :updated_by,
			status = :status,
			updated_at = NOW()
		WHERE id = :id
	`
	_, err := r.db.NamedExec(query, systemSetting)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}