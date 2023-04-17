package views

import (
	"strings"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	//"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"

)
func GUI(firsts map[string][]string, follows map[string][]string, solution map[string]map[string]bool){
	myApp := app.New()
    myWindow := myApp.NewWindow("LL1")

	// Crear dos rejillas
	grid1 := fyne.NewContainerWithLayout(layout.NewGridLayout(2))
	grid2 := fyne.NewContainerWithLayout(layout.NewGridLayout(2))
	grid3 := fyne.NewContainerWithLayout(layout.NewGridLayout(2))


	// Agregar etiquetas de clave-valor a la primera grid
	keys1 := make([]string, 0, len(firsts))
	for k := range firsts {
		keys1 = append(keys1, k)
	}
	sort.Strings(keys1)
	for _, key := range keys1 {
		values := firsts[key]
		valueStr := strings.Join(values, ", ")
		keyLabel := widget.NewLabel(key)
		valueLabel := widget.NewLabel(valueStr)
		grid1.Add(keyLabel)
		grid1.Add(valueLabel)
	}

	// Agregar etiquetas de clave-valor a la segunda grid
	keys2 := make([]string, 0, len(follows))
	for k := range follows {
		keys2 = append(keys2, k)
	}
	sort.Strings(keys2)
	for _, key := range keys2 {
		values := follows[key]
		valueStr := strings.Join(values, ", ")
		keyLabel := widget.NewLabel(key)
		valueLabel := widget.NewLabel(valueStr)
		grid2.Add(keyLabel)
		grid2.Add(valueLabel)
	}

	keys3 := make([]string, 0, len(solution))
for k := range solution {
    keys3 = append(keys3, k)
}
sort.Strings(keys3)
for _, key := range keys3 {
    subMap := solution[key]
    keyLabel := widget.NewLabel(key)
    valueLabel := widget.NewLabel("")
    for subKey, subValue := range subMap {
        if subValue {
            valueLabel.SetText(subKey)
            break
        }
    }
    grid3.Add(keyLabel)
    grid3.Add(valueLabel)
}



    // Agregar la grid a un scroll container
    //scrollContainer := container.NewScroll(grid1)
	gridContainer := fyne.NewContainerWithLayout(layout.NewGridLayout(2))
	
	gridContainer.Add(grid1)
	gridContainer.Add(grid2)
	//gridContainer.Add(grid3)



    // Establecer el contenido de la ventana como el scroll container
    //myWindow.SetContent(scrollContainer)
	myWindow.SetContent(gridContainer)
    // Mostrar la ventana
    myWindow.ShowAndRun()
}