package repository

import (
	"ankigen/internal/model"
	"errors"

	"gorm.io/gorm"
)

type LLMRepository interface {
	// GetActiveConfig returns the active LLM configuration.
	GetActiveConfig() (*model.LLMConfig, error)
	// SaveConfig saves the LLM configuration.
	SaveConfig(cfg *model.LLMConfig) error
	// ListConfigs lists all LLM configurations.
	ListConfigs() ([]model.LLMConfig, error)
	// GetByID gets an LLM configuration by ID.
	GetByID(id uint) (*model.LLMConfig, error)
	// DeleteConfig deletes an LLM configuration by ID.
	DeleteConfig(id uint) error
}

type llmRepoImpl struct {
	db *gorm.DB
}

// NewLLMRepository creates a repository for LLM configs.
func NewLLMRepository(db *gorm.DB) LLMRepository {
	return &llmRepoImpl{db: db}
}

// GetActiveConfig looks up the active LLM configuration.
func (r *llmRepoImpl) GetActiveConfig() (*model.LLMConfig, error) {
	var cfg model.LLMConfig
	err := r.db.Where("is_active = ?", true).First(&cfg).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no active llm configuration found")
		}
		return nil, err
	}
	return &cfg, nil
}

// SaveConfig saves cfg and, if active, deactivates others.
func (r *llmRepoImpl) SaveConfig(cfg *model.LLMConfig) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if cfg.IsActive {
			// Update all other records to is_active = false.
			err := tx.Model(&model.LLMConfig{}).
				Where("id != ?", cfg.ID).
				Update("is_active", false).Error
			if err != nil {
				return err
			}
		}

		// Perform Save
		return tx.Save(cfg).Error
	})
}

// ListConfigs returns configs ordered by update time.
func (r *llmRepoImpl) ListConfigs() ([]model.LLMConfig, error) {
	var configs []model.LLMConfig
	err := r.db.Order("updated_at DESC").Find(&configs).Error
	return configs, err
}

// GetByID finds a configuration by ID.
func (r *llmRepoImpl) GetByID(id uint) (*model.LLMConfig, error) {
	var cfg model.LLMConfig
	err := r.db.First(&cfg, id).Error
	return &cfg, err
}

// DeleteConfig removes a configuration by ID.
func (r *llmRepoImpl) DeleteConfig(id uint) error {
	return r.db.Delete(&model.LLMConfig{}, id).Error
}
