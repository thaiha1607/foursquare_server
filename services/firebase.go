package services

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/thaiha1607/foursquare_server/common/env"
)

type firebaseService struct {
	*firebase.App
}

func NewFirebaseService() firebaseService {
	config := &firebase.Config{
		StorageBucket: env.GetEnvOrDefault("FIREBASE_STORAGE_BUCKET", ""),
	}
	app, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return firebaseService{
		app,
	}
}
