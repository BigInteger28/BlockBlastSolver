import (
	"fmt"
	"strings"
)

type Zet struct {
	Figuur [][]bool
	X, Y   int
}

type ZetScore struct {
	Zet   Zet
	Score int
}

var LShape = [][]bool{
	{true, false},
	{true, false},
	{true, true},
}

var LShapeHorizontal = [][]bool{
	{true, false, false},
	{true, true, true},
}

var LShape3Long = [][]bool{
	{true, true, true},
	{false, false, true},
	{false, false, true},
}

var PiramideShape = [][]bool{
	{false, true, false},
	{true, true, true},
}

var HorizontalLine3 = [][]bool{
	{true, true, true},
}

var VerticalLine3 = [][]bool{
	{true},
	{true},
	{true},
}

var Square2x2 = [][]bool{
	{true, true},
	{true, true},
}

var HorizontalLine5 = [][]bool{
	{true, true, true, true, true},
}

var HorizontalLine4 = [][]bool{
	{true, true, true, true},
}

var VerticalLine4 = [][]bool{
	{true},
	{true},
	{true},
	{true},
}

var Square3x3 = [][]bool{
	{true, true, true},
	{true, true, true},
	{true, true, true},
}

func PlaatsFiguur(bord *[8][8]bool, figuur [][]bool, x, y int) bool {
	for i := 0; i < len(figuur); i++ {
		for j := 0; j < len(figuur[i]); j++ {
			if figuur[i][j] {
				if x+i >= 8 || y+j >= 8 || bord[x+i][y+j] {
					// De figuur past niet op het bord op deze positie
					return false
				}
			}
		}
	}
	// Plaats de figuur op het bord
	for i := 0; i < len(figuur); i++ {
		for j := 0; j < len(figuur[i]); j++ {
			if figuur[i][j] {
				bord[x+i][y+j] = true
			}
		}
	}
	return true
}

func ControleerEnVerwijderVolleRijenKolommen(bord *[8][8]bool) {
	// Controleer en verwijder volle rijen
	for i := 0; i < 8; i++ {
		vol := true
		for j := 0; j < 8; j++ {
			if !bord[i][j] {
				vol = false
				break
			}
		}
		if vol {
			for j := 0; j < 8; j++ {
				bord[i][j] = false
			}
		}
	}
	// Controleer en verwijder volle kolommen
	for j := 0; j < 8; j++ {
		vol := true
		for i := 0; i < 8; i++ {
			if !bord[i][j] {
				vol = false
				break
			}
		}
		if vol {
			for i := 0; i < 8; i++ {
				bord[i][j] = false
			}
		}
	}
}

func PrintBord(bord *[8][8]bool) {
	fmt.Println("\n  A B C D E F G H")
	for i := 0; i < 8; i++ {
		fmt.Printf("%d ", i+1)
		for j := 0; j < 8; j++ {
			if bord[i][j] {
				fmt.Print("X ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func CoordinatenNaarIndex(coordinaten string) (int, int, error) {
	if len(coordinaten) != 2 {
		return 0, 0, fmt.Errorf("ongeldige coördinaten")
	}

	kolom := strings.ToUpper(coordinaten[:1])[0]
	if kolom < 'A' || kolom > 'H' {
		return 0, 0, fmt.Errorf("ongeldige kolom")
	}

	rij := int(coordinaten[1] - '1')
	if rij < 0 || rij > 7 {
		return 0, 0, fmt.Errorf("ongeldige rij")
	}
	return rij, int(kolom - 'A'), nil
}

func StelBordOp(bord *[8][8]bool, bordLayout string) error {
	if len(bordLayout) != 64 {
		return fmt.Errorf("ongeldige bordlayout: de layout moet precies 64 karakters lang zijn")
	}

	for i := 0; i < 64; i++ {
		char := bordLayout[i]
		rij := i / 8
		kolom := i % 8
		if char == 'x' {
			bord[rij][kolom] = true
		} else if char == '.' {
			bord[rij][kolom] = false
		} else {
			return fmt.Errorf("ongeldige karakter '%c' in bordlayout op index %d", char, i)
		}
	}
	return nil
}

func KanPlaatsen(bord *[8][8]bool, figuur [][]bool, x, y int) bool {
	for i := 0; i < len(figuur); i++ {
		for j := 0; j < len(figuur[i]); j++ {
			if figuur[i][j] {
				if x+i >= 8 || y+j >= 8 || bord[x+i][y+j] {
					// De figuur past niet op het bord op deze positie
					return false
				}
			}
		}
	}
	return true
}

func GenereerZetten(bord *[8][8]bool, figuren [][][]bool) []Zet {
	var zetten []Zet
	for _, figuur := range figuren {
		for i := 0; i <= 8-len(figuur); i++ {
			for j := 0; j <= 8-len(figuur[0]); j++ {
				if KanPlaatsen(bord, figuur, i, j) {
					zetten = append(zetten, Zet{Figuur: figuur, X: i, Y: j})
				}
			}
		}
	}
	return zetten
}

func EvalueerBord(bord *[8][8]bool) int {
	score := 0
	// Voeg logica toe om de score te berekenen
	// Bijvoorbeeld, controleer op volledige rijen/kolommen
	// Tel het aantal lege cellen voor ruimtebeheer
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if !bord[i][j] {
				score++
			}
		}
	}

	// Tel het aantal volledige rijen of kolommen
	for i := 0; i < 8; i++ {
		volleRij := true
		volleKolom := true
		for j := 0; j < 8; j++ {
			if !bord[i][j] {
				volleRij = false
			}
			if !bord[j][i] {
				volleKolom = false
			}
		}
		if volleRij {
			score += 10 // Geef extra punten voor een volle rij
		}
		if volleKolom {
			score += 10 // Geef extra punten voor een volle kolom
		}
	}

	// Beoordeel de compactheid van de figuren
	// Evalueer open ruimte, vooral in het midden van het bord

	return score
}

func VindBesteZetParallel(bord *[8][8]bool, figuren [][][]bool) (besteZet Zet, besteScore int) {
	zetten := GenereerZetten(bord, figuren)
	scoreChannel := make(chan ZetScore)

	for _, zet := range zetten {
		go func(z Zet) {
			kopieBord := bord
			PlaatsFiguur(kopieBord, z.Figuur, z.X, z.Y)
			score := EvalueerBord(kopieBord)
			scoreChannel <- ZetScore{Zet: z, Score: score}
		}(zet)
	}

	besteScore = -1
	for range zetten {
		zetScore := <-scoreChannel
		if zetScore.Score > besteScore {
			besteScore = zetScore.Score
			besteZet = zetScore.Zet
		}
	}
	close(scoreChannel)
	return
}

func kopieVanBord(bord *[8][8]bool) *[8][8]bool {
	kopie := new([8][8]bool)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			kopie[i][j] = bord[i][j]
		}
	}
	return kopie
}

func VindBesteZetVoorDrieBlokkenParallel(bord *[8][8]bool, figuren [][][]bool) (besteZet Zet, besteScore int) {
	zetten := GenereerZetten(bord, figuren)
	scoreChannel := make(chan ZetScore)

	for _, zet := range zetten {
		go func(z Zet) {
			kopieBord := kopieVanBord(bord)
			if PlaatsFiguur(kopieBord, z.Figuur, z.X, z.Y) {
				// Plaatsing geslaagd, nu beoordeel het bord
				score := EvalueerBord(kopieBord)
				scoreChannel <- ZetScore{Zet: z, Score: score}
			}
		}(zet)
	}

	besteScore = -1
	for range zetten {
		zetScore := <-scoreChannel
		if zetScore.Score > besteScore {
			besteScore = zetScore.Score
			besteZet = zetScore.Zet
		}
	}
	close(scoreChannel)
	return
}

func PrintBordMetZet(bord *[8][8]bool, zet Zet) {
	fmt.Println("\n  A B C D E F G H")
	for i := 0; i < 8; i++ {
		fmt.Printf("%d ", i+1)
		for j := 0; j < 8; j++ {
			if zet.X <= i && i < zet.X+len(zet.Figuur) && zet.Y <= j && j < zet.Y+len(zet.Figuur[0]) {
				if zet.Figuur[i-zet.X][j-zet.Y] {
					fmt.Print("X ")
					continue
				}
			}
			if bord[i][j] {
				fmt.Print("X ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func BordToString(bord *[8][8]bool) string {
	var builder strings.Builder
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if bord[i][j] {
				builder.WriteString("x")
			} else {
				builder.WriteString(".")
			}
		}
	}
	return builder.String()
}

func PrintFiguur(figuur [][]bool) {
	for i := 0; i < len(figuur); i++ {
		for j := 0; j < len(figuur[i]); j++ {
			if figuur[i][j] {
				fmt.Print("X ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func KiesDrieFiguren(figuren [][][]bool) [][][]bool {
	chosenFiguren := make([][][]bool, 0)

	for i := 0; i < 3; i++ {
		for j, figuur := range figuren {
			fmt.Printf("\n%d. Figuur %d\n", j+1, j+1)
			PrintFiguur(figuur)
		}

		var keuze int
		fmt.Println("Kies een figuur (1-9):")
		fmt.Scan(&keuze)

		if keuze < 1 || keuze > len(figuren) {
			fmt.Println("Ongeldige keuze. Probeer opnieuw.")
			i-- // Retry the same choice
			continue
		}

		chosenFiguren = append(chosenFiguren, figuren[keuze-1])
	}

	return chosenFiguren
}

// .....xxx.......x.......xxx......xx.....x.x.xxxxx........xxxxxxx
func main() {
	var bord [8][8]bool
	var bordS string
	fmt.Print("Bord string: ")
	fmt.Scanln(&bordS)
	StelBordOp(&bord, bordS)
	for {
		figuren := [][][]bool{
			LShape,
			LShapeHorizontal,
			LShape3Long,
			PiramideShape,
			HorizontalLine3,
			VerticalLine3,
			Square2x2,
			HorizontalLine5,
			HorizontalLine4,
			VerticalLine4,
			Square3x3,
		}
		gekozenFiguren := KiesDrieFiguren(figuren)
		bestZet, bestScore := VindBesteZetVoorDrieBlokkenParallel(&bord, gekozenFiguren)
		fmt.Printf("Best Score: %d\n", bestScore)
		PrintBord(&bord)
		fmt.Println("Best Placement:")
		PrintBordMetZet(&bord, bestZet)
		newBoard := BordToString(&bord)
		StelBordOp(&bord, newBoard) // Updated to use the correct 'newBoard'
	}
}
