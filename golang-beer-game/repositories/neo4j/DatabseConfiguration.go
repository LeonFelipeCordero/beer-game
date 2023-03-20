package neo4j

import (
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
		Password:      "123456789",
		Username:      "neo4j",
	}

	_gogm, err := gogm.New(&config, gogm.DefaultPrimaryKeyStrategy, &entities.BoardNode{}, &entities.PlayerNode{}, &entities.OrderNode{})
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

func GlobalSession() gogm.SessionV2 {
	session, err := gogm.G().NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	return session
}
