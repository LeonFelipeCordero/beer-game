package repositories

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/repositories/neo4j/entities"
	"github.com/mindstand/gogm/v2"
)

func ConfigureDatabase() gogm.SessionV2 {
	config := gogm.Config{
		IndexStrategy: gogm.IGNORE_INDEX,
		PoolSize:      50,
		Port:          7687,
		IsCluster:     false,
		Host:          "127.0.0.1",
		Password:      "12345678",
		Username:      "neo4j",
	}

	_gogm, err := gogm.New(&config, gogm.UUIDPrimaryKeyStrategy, &entities.BoardNode{})
	if err != nil {
		panic(err)
	}

	gogm.SetGlobalGogm(_gogm)

	sess, err := _gogm.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		panic(err)
	}

	defer sess.Close()

	return sess
}

func GlobalSession(ctx context.Context) gogm.SessionV2 {
	session, err := gogm.G().NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		panic(err)
	}
	session.Begin(ctx)
	return session
}
