package main

import (
	"Electric_Characteristics/config"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"testing"
)

func TestDropCollection(t *testing.T) {
	dropCollection(mongoConnection, config.SpiceDB, config.DataTable, ctx)
	res, err := mongoConnection.Database(config.SpiceDB).Collection(config.DataTable).Find(ctx, bson.D{{"_id", "A767KN186M1HLAE050"}})
	if err != nil {
		t.Fail()
	}
	if res.Current != nil {
		t.Fail()
	}
}

func TestInsertData(t *testing.T) {
	type args struct {
		partNumber string
	}
	tests := []struct {
		name    string
		args    args
		want    Fulldata
		wantErr bool
	}{
		{
			name: "A767KN186M1HLAE050", args: struct{ partNumber string }{partNumber: "A767KN186M1HLAE050"},
			wantErr: false,
			want: Fulldata{
				Part_number: "A767KN186M1HLAE050",
				Freq_min:    "0.0001",
				Freq_max:    "10",
				Z_Min:       map[string]string{"Freq (MHz)": "5.04E-01", "Z (Ohm)": "2.67E-02"},
				Frequency_values: []Equivalent_circuit_data{{
					Freq_MHZ: "0.0001",
					Z_Ohm:    "80.0754200506633",
					Re_Z_Ohm: "2.97870571613715",
					IM_Z_Ohm: "80.0199988037167",
					Cap_F:    "1.99E-05",
					Phase:    "-87.87",
					L_H:      "6.40E-10",
				}},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := insertData(mongoConnection, config.SpiceDB, config.DataTable, ctx, tt.want)
			if err != nil {
				log.Fatal(err)
			}
			expectID := "A767KN186M1HLAE050"
			if expectID != actual.InsertedID {
				t.Fail()
			}

		})

	}
}

func TestFindDataById(t *testing.T) {
	type args struct {
		partNumber string
	}
	tests := []struct {
		name    string
		args    args
		want    Fulldata
		wantErr bool
	}{
		{
			name: "A767KN186M1HLAE050", args: struct{ partNumber string }{partNumber: "A767KN186M1HLAE050"},
			wantErr: false,
			want: Fulldata{
				Part_number: "A767KN186M1HLAE050",
				Freq_min:    "0.0001",
				Freq_max:    "10",
				Z_Min:       map[string]string{"Freq (MHz)": "5.04E-01", "Z (Ohm)": "2.67E-02"},
				Frequency_values: []Equivalent_circuit_data{{
					Freq_MHZ: "0.0001",
					Z_Ohm:    "80.0754200506633",
					Re_Z_Ohm: "2.97870571613715",
					IM_Z_Ohm: "80.0199988037167",
					Cap_F:    "1.99E-05",
					Phase:    "-87.87",
					L_H:      "6.40E-10",
				}},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := findDataById(mongoConnection, config.SpiceDB, config.DataTable, tt.want.Part_number)
			if err != nil {
				log.Fatal(err)
			}
			fd := &Fulldata{}
			err = actual.Decode(fd)
			if err != nil {
				t.Fail()
			}

			if fd.Part_number != tt.want.Part_number {
				t.Fail()
			}
		})
	}
}
