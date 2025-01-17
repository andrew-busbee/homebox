package services

import (
	"bytes"
	"encoding/csv"
	"reflect"
	"testing"
)

const CSV_DATA = `
Import Ref,Location,Labels,Quantity,Name,Description,Insured,Serial Number,Mode Number,Manufacturer,Notes,Purchase From,Purchased Price,Purchased Time,Lifetime Warranty,Warranty Expires,Warranty Details,Sold To,Sold Price,Sold Time,Sold Notes
A,Garage,IOT;Home Assistant; Z-Wave,1,Zooz Universal Relay ZEN17,"Zooz 700 Series Z-Wave Universal Relay ZEN17 for Awnings, Garage Doors, Sprinklers, and More | 2 NO-C-NC Relays (20A, 10A) | Signal Repeater | Hub Required (Compatible with SmartThings and Hubitat)",,,ZEN17,Zooz,,Amazon,39.95,10/13/2021,,,,,,,
B,Living Room,IOT;Home Assistant; Z-Wave,1,Zooz Motion Sensor,"Zooz Z-Wave Plus S2 Motion Sensor ZSE18 with Magnetic Mount, Works with Vera and SmartThings",,,ZSE18,Zooz,,Amazon,29.95,10/15/2021,,,,,,,
C,Office,IOT;Home Assistant; Z-Wave,1,Zooz 110v Power Switch,"Zooz Z-Wave Plus Power Switch ZEN15 for 110V AC Units, Sump Pumps, Humidifiers, and More",,,ZEN15,Zooz,,Amazon,39.95,10/13/2021,,,,,,,
D,Downstairs,IOT;Home Assistant; Z-Wave,1,Ecolink Z-Wave PIR Motion Sensor,"Ecolink Z-Wave PIR Motion Detector Pet Immune, White (PIRZWAVE2.5-ECO)",,,PIRZWAVE2.5-ECO,Ecolink,,Amazon,35.58,10/21/2020,,,,,,,
E,Entry,IOT;Home Assistant; Z-Wave,1,Yale Security Touchscreen Deadbolt,"Yale Security YRD226-ZW2-619 YRD226ZW2619 Touchscreen Deadbolt, Satin Nickel",,,YRD226ZW2619,Yale,,Amazon,120.39,10/14/2020,,,,,,,
F,Kitchen,IOT;Home Assistant; Z-Wave,1,Smart Rocker Light Dimmer,"UltraPro Z-Wave Smart Rocker Light Dimmer with QuickFit and SimpleWire, 3-Way Ready, Compatible with Alexa, Google Assistant, ZWave Hub Required, Repeater/Range Extender, White Paddle Only, 39351",,,39351,Honeywell,,Amazon,65.98,09/30/0202,,,,,,,`

func loadcsv() [][]string {
	reader := csv.NewReader(bytes.NewBuffer([]byte(CSV_DATA)))

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return records
}

func Test_csvRow_getLabels(t *testing.T) {
	type fields struct {
		LabelStr string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "basic test",
			fields: fields{
				LabelStr: "IOT;Home Assistant;Z-Wave",
			},
			want: []string{"IOT", "Home Assistant", "Z-Wave"},
		},
		{
			name: "no labels",
			fields: fields{
				LabelStr: "",
			},
			want: []string{},
		},
		{
			name: "single label",
			fields: fields{
				LabelStr: "IOT",
			},
			want: []string{"IOT"},
		},
		{
			name: "trailing semicolon",
			fields: fields{
				LabelStr: "IOT;",
			},
			want: []string{"IOT"},
		},

		{
			name: "whitespace",
			fields: fields{
				LabelStr: " IOT;		Home Assistant;   Z-Wave ",
			},
			want: []string{"IOT", "Home Assistant", "Z-Wave"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := csvRow{
				LabelStr: tt.fields.LabelStr,
			}
			if got := c.getLabels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("csvRow.getLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
