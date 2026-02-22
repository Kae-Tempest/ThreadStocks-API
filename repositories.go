package main

import (
	"context"

	"gorm.io/gorm"
)

// --- Account Repository ---

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) UserRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) GetByID(ctx context.Context, id uint) (*User, error) {
	var user User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *accountRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	if err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *accountRepository) Create(ctx context.Context, user *User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *accountRepository) Update(ctx context.Context, user *User) error {
	return r.db.WithContext(ctx).Model(user).Where("id = ?", user.ID).Updates(user).Error
}

// --- Thread Repository ---

type threadRepository struct {
	db *gorm.DB
}

func NewThreadRepository(db *gorm.DB) ThreadRepository {
	return &threadRepository{db: db}
}

func (r *threadRepository) GetByID(ctx context.Context, id uint) (*Thread, error) {
	var thread Thread
	if err := r.db.WithContext(ctx).First(&thread, id).Error; err != nil {
		return nil, err
	}
	return &thread, nil
}

func (r *threadRepository) GetByUserID(ctx context.Context, userID uint) ([]Thread, error) {
	var threads []Thread
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&threads).Error; err != nil {
		return nil, err
	}
	return threads, nil
}

func (r *threadRepository) Create(ctx context.Context, thread *Thread) error {
	var existing Thread
	err := r.db.WithContext(ctx).Unscoped().Where("user_id = ? AND thread_id = ?", thread.UserID, thread.ThreadId).First(&existing).Error

	if err == nil {
		// Le thread existe déjà (peut-être supprimé)
		if existing.DeletedAt.Valid {
			// Il était supprimé, on le restaure
			thread.ID = existing.ID
			return r.db.WithContext(ctx).Unscoped().Model(&existing).Updates(map[string]any{
				"deleted_at":   nil,
				"is_e":         thread.IsE,
				"is_c":         thread.IsC,
				"is_s":         thread.IsS,
				"brand":        thread.Brand,
				"thread_count": thread.ThreadCount,
			}).Error
		}
		// Il n'est pas supprimé, on laisse GORM renvoyer l'erreur de contrainte unique
	}

	return r.db.WithContext(ctx).Create(thread).Error
}

func (r *threadRepository) Update(ctx context.Context, thread *Thread) error {
	return r.db.WithContext(ctx).Model(thread).Where("user_id = ?", thread.UserID).Updates(thread).Error
}

func (r *threadRepository) Delete(ctx context.Context, userID uint, id uint) error {
	return r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&Thread{}, id).Error
}

func (r *threadRepository) DeleteMultiple(ctx context.Context, userID uint, ids []string) error {
	return r.db.WithContext(ctx).Where("user_id = ? AND thread_id IN ?", userID, ids).Delete(&Thread{}).Error
}

// --- Password Reset Repository ---

type passwordResetRepository struct {
	db *gorm.DB
}

func NewPasswordResetRepository(db *gorm.DB) PasswordResetTokenRepository {
	return &passwordResetRepository{db: db}
}

func (r *passwordResetRepository) Create(ctx context.Context, token *PasswordResetToken) error {
	return r.db.WithContext(ctx).Create(token).Error
}

func (r *passwordResetRepository) GetByToken(ctx context.Context, token string) (*PasswordResetToken, error) {
	var prt PasswordResetToken
	if err := r.db.WithContext(ctx).Preload("User").First(&prt, "token = ?", token).Error; err != nil {
		return nil, err
	}
	return &prt, nil
}

func (r *passwordResetRepository) DeleteByUserID(ctx context.Context, userID uint) error {
	return r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&PasswordResetToken{}).Error
}
