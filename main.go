package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println("Wybierz pozycję:")
	fmt.Println(1, 2, 3)
	fmt.Println(4, 5, 6)
	fmt.Println(7, 8, 9)
	fmt.Println("Gracz 1 = X")
	fmt.Println("Gracz 2 = O")
	graj(board)
	fmt.Println("Koniec!!!")
}

func odnowa(board [][]string) {
	fmt.Println("y- graj od nowa, n-stop")
	var again string
	fmt.Scan(&again)
	if again == "y" {
		board = [][]string{
			[]string{"_", "_", "_"},
			[]string{"_", "_", "_"},
			[]string{"_", "_", "_"},
		}
		graj(board)
	}
}

func graj(board [][]string) [][]string {
	player1 := "X"
	player2 := "O"
	pos := wybierzLokal()
	zlapozycja(pos)
	turn(pos, board, player1)
	sukces := wygrany(board)
	if sukces == true {
		odnowa(board)
		return board
	}
	pos = wybierzLokal()
	zlapozycja(pos)
	turn(pos, board, player2)
	sukces = wygrany(board)
	if sukces == true {
		odnowa(board)
		return board
	}

	return graj(board)
}
func zlapozycja(pos int) {
	if pos > 9 || pos < 1 {
		pos = sprobujPonownie(pos)
	}
}
func sprobujPonownie(pos int) int {
	if pos > 9 || pos < 1 {
		fmt.Println("Wybierz liczbę pomiedzy 1 a 9")
		return sprobujPonownie(wybierzLokal())
	}
	return pos
}
func pozycja(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func wybierzLokal() int {
	fmt.Println("Wybierz liczbę z tablicy.")
	var pos int
	fmt.Scan(&pos)
	fmt.Println("Pozycja ", pos, " nie znajduje sie w mozliwosci wyboru")
	return pos
}

func turn(pos int, board [][]string, player string) {
	if pos == (1) {
		board[0][0] = player
	}
	if pos == (2) {
		board[0][1] = player
	}
	if pos == (3) {
		board[0][2] = player
	}
	if pos == (4) {
		board[1][0] = player
	}
	if pos == (5) {
		board[1][1] = player
	}
	if pos == (6) {
		board[1][2] = player
	}
	if pos == (7) {
		board[2][0] = player
	}
	if pos == (8) {
		board[2][1] = player
	}
	if pos == (9) {
		board[2][2] = player
	}
	clear()
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func wygrany(board [][]string) bool {
	xWinsrow := []string{"X", "X", "X"}
	oWinsrow := []string{"O", "O", "O"}
	sukces := false
	switch {
	case reflect.DeepEqual(board[0], xWinsrow):
		fmt.Println("Gracz 1 wygral")
		sukces = true
		return sukces
	case reflect.DeepEqual(board[1], xWinsrow):
		fmt.Println("Gracz 1 wygral")
		sukces = true
		return sukces
	case reflect.DeepEqual(board[2], xWinsrow):
		fmt.Println("Gracz 1 wygral")
		sukces = true
		return sukces
	case reflect.DeepEqual(board[0], oWinsrow):
		fmt.Println("Gracz 2 wygral")
		sukces = true
		return sukces
	case reflect.DeepEqual(board[1], oWinsrow):
		fmt.Println("Gracz 2 wygral")
		sukces = true
		return sukces
	case reflect.DeepEqual(board[2], oWinsrow):
		fmt.Println("Gracz 2 wygral")
		sukces = true
		return sukces
	case board[0][0] == "X" && board[1][0] == "X" && board[2][0] == "X":
		sukces = true
		fmt.Println("Gracz 1 wygral")
		return sukces
	case board[0][1] == "X" && board[1][1] == "X" && board[2][1] == "X":
		sukces = true
		fmt.Println("Gracz 1 wygral")
		return sukces
	case board[0][2] == "X" && board[1][2] == "X" && board[2][2] == "X":
		sukces = true
		fmt.Println("Gracz 1 wygral")
		return sukces
	case board[0][0] == "O" && board[1][0] == "O" && board[2][0] == "O":
		sukces = true
		fmt.Println("Gracz 2 wygral")
		return sukces
	case board[0][1] == "O" && board[1][1] == "O" && board[2][1] == "O":
		sukces = true
		fmt.Println("Gracz 2 wygral")
		return sukces
	case board[0][2] == "O" && board[1][2] == "O" && board[2][2] == "O":
		sukces = true
		fmt.Println("Gracz 2 wygral")
		return sukces
	case board[0][0] == "X" && board[1][1] == "X" && board[2][2] == "X":
		sukces = true
		fmt.Println("Gracz 1 wygral")
	case board[0][2] == "X" && board[1][1] == "X" && board[2][1] == "X":
		sukces = true
		fmt.Println("Gracz 1 wygral")
	case board[0][0] == "O" && board[1][1] == "O" && board[2][2] == "O":
		sukces = true
		fmt.Println("Gracz 2 wygral")
	case board[0][2] == "O" && board[1][1] == "O" && board[2][1] == "O":
		sukces = true
		fmt.Println("Gracz 2 wygral!")
	}
	return sukces
	return sukces
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
