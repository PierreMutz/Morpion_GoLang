package main

import "github.com/tadvi/winc"

type game struct{
  P1, P2, curPlayer, winner string
  cases[9] string
  finished bool
}

func newGame() *game {
  g := game{}
  g.P1 = "X"
  g.P2 = "O"
  g.winner = ""
  g.curPlayer = g.P1
  g.cases = [9]string{""}
  g.finished = false

  return &g
}

func updateText(text *winc.Label, g *game) {
  if g.finished == true {
    if g.winner == "X" {
      text.SetText("Le joueur 1 a gagné")
    } else if g.winner == "O" {
      text.SetText("Le joueur 2 a gagné")
    } else {
      text.SetText("Egalité")
    }
  } else {
    text.SetText("C'est au tour du joueur " + g.curPlayer)
  }
}

func (g *game) changePlayer() {
  if g.curPlayer == g.P1 {
    g.curPlayer = g.P2
  } else {
    g.curPlayer = g.P1
  }
}

func (g *game) checkVictory() {
    victory := false
    equality := true
    symbol := g.curPlayer
    if (g.cases[0] == symbol && g.cases[1] == symbol && g.cases[2] == symbol) || (g.cases[3] == symbol && g.cases[4] == symbol && g.cases[5] == symbol) || (g.cases[6] == symbol && g.cases[7] == symbol && g.cases[8] == symbol) || (g.cases[0] == symbol && g.cases[3] == symbol && g.cases[6] == symbol) || (g.cases[1] == symbol && g.cases[4] == symbol && g.cases[7] == symbol) || (g.cases[2] == symbol && g.cases[5] == symbol && g.cases[8] == symbol) || (g.cases[0] == symbol && g.cases[4] == symbol && g.cases[8] == symbol) || (g.cases[2] == symbol && g.cases[4] == symbol && g.cases[6] == symbol) {
      victory = true
    } else {
      for i := 0; i < 9; i++{
        if g.cases[i] != ""{
         equality = false
       }
      }
    }

    if victory == true {
      if symbol == g.P1 {
        g.winner = "X"
      } else {
        g.winner = "O"
      }
      g.finished = true
    } else if equality == true {
      g.finished = true
    }
}

func main() {
    mainWindow := winc.NewForm(nil)
    game := newGame()
    text := winc.NewLabel(mainWindow)
    text.SetPos(5, 5)
    text.SetSize(150, 25)
    text.SetText("C'est au tour du joueur " + game.curPlayer)
    createButton(mainWindow, game, text, 0, 50, 50)
    createButton(mainWindow, game, text, 1, 100, 50)
    createButton(mainWindow, game, text, 2, 150, 50)
    createButton(mainWindow, game, text, 3, 50, 100)
    createButton(mainWindow, game, text, 4, 100, 100)
    createButton(mainWindow, game, text, 5, 150, 100)
    createButton(mainWindow, game, text, 6, 50, 150)
    createButton(mainWindow, game, text, 7, 100, 150)
    createButton(mainWindow, game, text, 8, 150, 150)
    mainWindow.SetSize(270, 300)
    mainWindow.SetText("Morpion")
    mainWindow.Center()
    mainWindow.Show()
    mainWindow.OnClose().Bind(wndOnClose)
    winc.RunMainLoop()
}

func wndOnClose(arg *winc.Event) {
    winc.Exit()
}

func createButton(window *winc.Form, game *game, text *winc.Label, pos int, x int, y int) {
  btn := winc.NewPushButton(window)
  btn.SetPos(x, y)
  btn.SetSize(50, 50)
  btn.SetText("")

  btn.OnClick().Bind(func(e *winc.Event){
    if game.cases[pos] == "" && game.finished == false {
      btn.SetText(game.curPlayer)
      game.cases[pos] = game.curPlayer
      game.checkVictory()
      if game.finished == false {
        game.changePlayer()
      }
      updateText(text, game)
    }
  })
}
