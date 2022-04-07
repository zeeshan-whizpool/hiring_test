package servicelayer

import (
	"os"
	"testing"
)

func TestGetProjetList(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"zee test", "dummy-whizpool-hiring", `{"forks":1,"names":"dummy-whizpool-hiring-1, dummy-whizpool-hiring-2"}`},
	}
	os.Chdir("../")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetProjetList(tt.args); got != tt.want {
				t.Errorf("GetProjetList() = %v, want %v", got, tt.want)
			}
		})
	}

	os.Chdir("./")

}
