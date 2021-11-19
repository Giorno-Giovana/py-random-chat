package db

import (
	"context"
	"time"

	"gitlab.com/procsy/attorney/api/internal/model/core"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *core.User) error

	GetUserByID(ctx context.Context, ID core.UserID) (*core.User, error)
	GetUserByEmail(ctx context.Context, email string) (*core.User, error)

	GetUsers(ctx context.Context) ([]core.User, error)
	GetUsersByID(ctx context.Context, IDs []core.UserID) ([]core.User, error)

	SelectUsers(ctx context.Context, selector string) ([]core.User, error)

	CheckUserEmailExistence(ctx context.Context, email string) (bool, error)

	UpdateUser(ctx context.Context, user *core.User) error

	DeleteUser(ctx context.Context, user *core.User) error
}

type userRepositoryImpl struct {
	db   *mongo.Database
	coll *mongo.Collection
}

// CreateUser tries to insert given user to the db:
// returns error if the email is already taken, otherwise inserts
func (repo *userRepositoryImpl) CreateUser(ctx context.Context, user *core.User) error {
	repo.InitUser(user)
	_, err := repo.coll.InsertOne(ctx, user)
	return err
}

func (repo *userRepositoryImpl) GetUserByID(ctx context.Context, ID core.UserID) (*core.User, error) {
	user := new(core.User)
	filter := bson.M{"_id": ID}
	err := repo.coll.FindOne(ctx, filter).Decode(user)
	return user, wrapError(err)
}

// GetUserByEmail looks up in the db for user with the provided email
func (repo *userRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*core.User, error) {
	user := new(core.User)
	filter := bson.M{"email": email}
	err := repo.coll.FindOne(ctx, filter).Decode(user)
	return user, wrapError(err)
}

// GetUsers returns all the users from the db
func (repo *userRepositoryImpl) GetUsers(ctx context.Context) ([]core.User, error) {
	var users []core.User
	cursor, err := repo.coll.Find(ctx, bson.D{})
	if err != nil {
		return users, err
	} else {
		err = cursor.All(ctx, &users)
	}
	return users, err
}

// GetUsersByID looks up in db for users with provided IDs
func (repo *userRepositoryImpl) GetUsersByID(ctx context.Context, IDs []core.UserID) ([]core.User, error) {
	var users []core.User
	filter := bson.M{"_id": bson.M{"$in": IDs}}
	cursor, err := repo.coll.Find(ctx, filter)
	if err != nil {
		return users, err
	} else {
		err = cursor.All(ctx, &users)
	}
	return users, err
}

func (repo *userRepositoryImpl) SelectUsers(ctx context.Context, selector string) ([]core.User, error) {
	users := []core.User{}

	fuzzy := bson.M{"$regex": selector, "$options": "i"}
	filter := bson.M{"$or": []bson.M{
		{"email": fuzzy},
		{"username": fuzzy}},
	}

	cursor, err := repo.coll.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return users, nil
		}
		return nil, err
	} else {
		err = cursor.All(ctx, &users)
	}

	return users, err
}

// CheckUserEmailExistence checks whether user with given email exists. Returns true if email is already taken.
func (repo *userRepositoryImpl) CheckUserEmailExistence(ctx context.Context, email string) (bool, error) {
	filter := bson.M{"email": email}
	err := repo.coll.FindOne(ctx, filter).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// UpdateUser updates user .
func (repo *userRepositoryImpl) UpdateUser(ctx context.Context, user *core.User) error {
	filter := bson.M{"_id": user.ID}
	_, err := repo.coll.ReplaceOne(ctx, filter, user)
	return err
}

// DeleteUser deletes from the db user with the provided email
func (repo *userRepositoryImpl) DeleteUser(ctx context.Context, user *core.User) error {
	filter := bson.M{"email": user.Email}
	_, err := repo.coll.DeleteOne(ctx, filter)
	return err
}

func (repo *userRepositoryImpl) InitUser(user *core.User) error {
	uid, err := core.GenUUID()
	if err != nil {
		return err
	}
	user.ID = core.UserID(uid)
	timeNow := time.Now().Unix()
	user.CreatedAt = timeNow
	user.UpdatedAt = timeNow
	return nil
}

// NewUserRepository creates a new instance of userRepositoryImpl
func NewUserRepository(db *mongo.Database) (*userRepositoryImpl, error) {
	return &userRepositoryImpl{db: db, coll: db.Collection("users")}, nil
}
