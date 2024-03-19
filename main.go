package main

import (
	"fmt"
	"slices"
	"strconv"
)

type MinimaxState struct {
  board [SIZE]int
  player int
}

type MinimaxCache map[MinimaxState]MinimaxRes

func (s MinimaxState) done() bool {
  for _, l := range s.board {
    if l != 0 { return false }
  }
  return true
}

type MinimaxRes struct {
  line, count, score int
}

func minimax(s MinimaxState, c MinimaxCache) MinimaxRes {
  res, ok := c[s]
  if ok { return res }
  res = MinimaxRes{}
  for i, l := range s.board {
    for _, m := range MOVES {
      if l >= m {
        newboard := s.board
        newboard[i] -= m
        lres := minimax(MinimaxState{
          newboard,
          -s.player,
        }, c)
        if lres.score < res.score * s.player {
          res.score = lres.score
          res.line = i
          res.count = m
        }
      }
    }
  }
  if res.score == 0 {
    res.score = s.player
  }
  c[s] = res
  return res
}

const (
  SIZE = 4
  LEN = 4
)
var (
  MOVES = []int{1,2}
)

func main() {
  state := MinimaxState{
    board: [SIZE]int{3,2,3,3},
    player: 1,
  }
  cache := make(MinimaxCache)
  minimax(state, cache)
  for !state.done() {
    var l, c int
    fmt.Printf("%v\n", state.board)
    switch state.player {
    case 1:
      lr, cr := "", ""
      fmt.Print("move: ")
      fmt.Scanln(&lr, &cr)
      l, _ = strconv.Atoi(lr)
      c, _ = strconv.Atoi(cr)
    case -1:
      res, ok := cache[state]
      if !ok { res = minimax(state, cache) }
      l, c = res.line, res.count
      fmt.Printf("computer move: %v %v\n", l, c)
    }
    if l >= SIZE || c > state.board[l] || !slices.Contains(MOVES, c) {
      fmt.Print("invalid move\n")
      continue
    }
    state.board[l] -= c
    state.player *= -1
  }
  fmt.Printf("%v\n", state.board)
  switch state.player {
  case 1:
    fmt.Print("computer won\n")
  case -1:
    fmt.Print("player won\n")
  }
}
