package game

import "testing"

func TestWrongTurn(t *testing.T) {
	board := NewNoughtCross()

	err := board.AddNought(0, 0)
	if err == nil {
		t.Error("We were expecting a wrong turn error")
	}
}

func TestWrongMove(t *testing.T) {
	board := NewNoughtCross()

	err := board.AddCross(0, 0)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}
	err = board.AddNought(0, 0)
	if err == nil {
		t.Error("We were expecting a wrong move error")
	}

	err = board.AddNought(0, 3)
	if err == nil {
		t.Error("We were expecting a wrong turn error")
	}

	err = board.AddNought(-1, 0)
	if err == nil {
		t.Error("We were expecting a wrong turn error")
	}
}

func TestWinHoriz(t *testing.T) {
	board := NewNoughtCross()

	err := board.AddCross(0, 0)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}
	err = board.AddNought(1, 0)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}

	err = board.AddCross(0, 1)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}
	err = board.AddNought(1, 1)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}

	err = board.AddCross(0, 2)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}
	err = board.AddNought(1, 2)
	if err == nil {
		t.Error("We were expecting a game finished error")
		return
	}

	if err.Error() != "Game finished" {
		t.Errorf("We received: %v instead of game finished", err)
	}

}

func TestWinVert(t *testing.T) {
	board := NewNoughtCross()

	err := board.AddCross(0, 0)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}
	err = board.AddNought(0, 1)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}

	err = board.AddCross(1, 0)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}
	err = board.AddNought(1, 1)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}

	err = board.AddCross(2, 0)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}
	err = board.AddNought(2, 1)
	if err == nil {
		t.Error("We were expecting a game finished error")
		return
	}

	if err.Error() != "Game finished" {
		t.Errorf("We received: %v instead of game finished", err)
	}

}

func TestWinVertNuts(t *testing.T) {
	board := NewNoughtCross()

	err := board.AddCross(0, 2)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}

	err = board.AddNought(0, 1)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}

	err = board.AddCross(0, 0)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}

	err = board.AddNought(1, 1)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}

	err = board.AddCross(1, 0)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}

	err = board.AddNought(2, 1)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}

	err = board.AddCross(2, 0)
	if err == nil {
		t.Error("We were expecting a game finished error")
		return
	}

	if err.Error() != "Game finished" {
		t.Errorf("We received: %v instead of game finished", err)
	}

}

func TestWinDiagLftR(t *testing.T) {
	board := NewNoughtCross()

	err := board.AddCross(0, 0)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}
	err = board.AddNought(0, 1)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}

	err = board.AddCross(1, 1)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}
	err = board.AddNought(1, 2)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}

	err = board.AddCross(2, 2)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}
	err = board.AddNought(1, 0)
	if err == nil {
		t.Error("We were expecting a game finished error")
		return
	}

	if err.Error() != "Game finished" {
		t.Errorf("We received: %v instead of game finished", err)
	}

}

func TestWinDiagRgtLft(t *testing.T) {
	board := NewNoughtCross()

	err := board.AddCross(0, 2)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}
	err = board.AddNought(0, 1)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}

	err = board.AddCross(1, 1)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}
	err = board.AddNought(1, 2)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}

	err = board.AddCross(2, 0)
	if err != nil {
		t.Errorf("We were not expecting an error but we got %v", err)
	}
	err = board.AddNought(1, 0)
	if err == nil {
		t.Error("We were expecting a game finished error")
		return
	}

	if err.Error() != "Game finished" {
		t.Errorf("We received: %v instead of game finished", err)
	}

}
