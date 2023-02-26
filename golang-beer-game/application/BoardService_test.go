package application

//func TestCreateTask(t *testing.T) {
//	boardName := "test"
//
//	t.Run("a board should be created if name is not taken", func(t *testing.T) {
//		board, err := CreateBoard(boardName)
//		if err != nil {
//			t.Error("there should not be errors")
//		}
//
//		savedBoard, err := boardRepository.LoadBoardByName(boardName)
//		if err != nil {
//			t.Error("board not saved")
//		}
//
//		assert.Equal(t, savedBoard.Name, board.Name, "wrong name")
//		assert.Equal(t, savedBoard.State, board.State, "wrong state")
//		assert.Equal(t, savedBoard.Finished, board.Finished, "wrong finished")
//		assert.Equal(t, savedBoard.Full, board.Full, "wrong full")
//		assert.Equal(t, len(savedBoard.Players), len(board.Players), "wrong name")
//
//		boardRepository.DeleteAll()
//	})
//
//	t.Run("a board should not be created if name is taken", func(t *testing.T) {
//		board, err := CreateBoard(boardName)
//		if err != nil {
//			t.Error("there should not be errors")
//		}
//
//		secondBoard, err := CreateBoard(boardName)
//		if err == nil && secondBoard != nil {
//			t.Error("there should be an error and board should be nil")
//		}
//
//		savedBoard, err := boardRepository.LoadBoardByName(boardName)
//		if err != nil {
//			t.Error("board not saved")
//		}
//
//		assert.Equal(t, savedBoard.Name, board.Name, "wrong name")
//		assert.Equal(t, savedBoard.State, board.State, "wrong state")
//		assert.Equal(t, savedBoard.Finished, board.Finished, "wrong finished")
//		assert.Equal(t, savedBoard.Full, board.Full, "wrong full")
//		assert.Equal(t, len(savedBoard.Players), len(board.Players), "wrong name")
//	})
//}
