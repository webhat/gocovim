/*
		Copyright (C) 2019  Daniël W. Crompton

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
		along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_addToCoverageMap(t *testing.T) {
	type args struct {
		cover Cover
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := addToCoverageMap(tt.args.cover); (err != nil) != tt.wantErr {
				t.Errorf("addToCoverageMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createCoverageFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createCoverageFile()
		})
	}
}

func Test_parseCoverLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name     string
		args     args
		wantItem Cover
		wantErr  bool
	}{
		{name: "Test", args: args{line: "bitbucket.org/specialbrands/master-control-unit/controller/config.go:122.13,151.2 20 1"}, wantItem: Cover{"bitbucket.org/specialbrands/master-control-unit/controller/config.go", 122, 151, 13, 2, 20, 1}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem, err := parseCoverLine(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCoverLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("parseCoverLine() = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}
