package main

import (
	_ "embed"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

//go:embed img/icon.png
var icon []byte

func main() {
	a := app.NewWithID("io.github.jacalz.focalplane")
	a.SetIcon(fyne.NewStaticResource("icon.png", icon))
	w := a.NewWindow("Focalplane")

	// Validators for entry contents.
	uintValidator := validation.NewRegexp(`^\d+$`, "Must be a positive integer type value.")
	floatValidator := validation.NewRegexp(`\d+(\.\d+)?$`, "Must be a valid decimal number.")

	// Camera and lens parameters.
	focal := &widget.Entry{PlaceHolder: "mm", Validator: uintValidator}
	distance := &widget.Entry{PlaceHolder: "m", Validator: floatValidator}
	aperture := &widget.Entry{PlaceHolder: "f-stops", Validator: floatValidator}
	sensor := &widget.Select{Options: sensors}

	// Building blocks for the user interface.
	dofText := &widget.Label{Text: "Depth of field:", TextStyle: fyne.TextStyle{Bold: true}}
	dofValue := &widget.Label{}
	dofData := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Focal length", Widget: focal, HintText: "Given in millimeters."},
			{Text: "Distance to subject", Widget: distance, HintText: "Given in meters."},
			{Text: "Aperture", Widget: aperture, HintText: "Given in f-stops."},
			{Text: "Sensor type", Widget: sensor, HintText: "Digital and analog formats."},
		},
	}

	// Calculate the depth of field from values above.
	// The validators guarantee that we don't get invalid inputs.
	recalculateDOF := func(_ string) {
		focallengh, _ := strconv.ParseUint(focal.Text, 10, 64)
		distance, _ := strconv.ParseFloat(distance.Text, 64)
		fstop, _ := strconv.ParseFloat(aperture.Text, 64)
		circle := sensorToCoC[sensor.Selected]

		dof := depthOfField(float64(focallengh)/1000, distance, fstop, circle)
		dofValue.SetText(strconv.FormatFloat(dof, 'f', 6, 64) + " m")
	}

	// Hook up widgets to recalculate when something changes.
	focal.OnChanged = recalculateDOF
	distance.OnChanged = recalculateDOF
	aperture.OnChanged = recalculateDOF
	sensor.OnChanged = recalculateDOF

	w.SetContent(container.NewBorder(nil, container.NewHBox(dofText, dofValue), nil, nil, dofData))
	w.ShowAndRun()
}
