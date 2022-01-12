package ssmcli

import (
	"fmt"
	"testing"

	"github.com/manifoldco/promptui"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/shearn89/ssmcli/utils"
)

type promptMock struct {
	mock.Mock
}

func (p promptMock) Run() (int, string, error) {
	args := p.Called()
	return args.Int(0), args.String(1), args.Error(2)
}

func Test_promptRunner(t *testing.T) {
	emptyPrompt := new(promptMock)
	emptyPrompt.On("Run").Return(0, "", nil)

	errorPrompt := new(promptMock)
	errorPrompt.On("Run").Return(1, "", fmt.Errorf("woops"))

	type args struct {
		runner Runner
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{"empty case", args{runner: emptyPrompt}, "", assert.NoError},
		{"error case", args{runner: errorPrompt}, "", assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PromptRunner(tt.args.runner)
			if !tt.wantErr(t, err, fmt.Sprintf("PromptRunner(%v)", tt.args.runner)) {
				return
			}
			assert.Equalf(t, tt.want, got, "PromptRunner(%v)", tt.args.runner)
		})
	}
	emptyPrompt.AssertExpectations(t)
}

func Test_buildPrompt(t *testing.T) {
	type args struct {
		sessionMap map[string]string
	}
	tests := []struct {
		name string
		args args
		want Runner
	}{
		{"empty map", args{map[string]string{}},
			&promptui.Select{Label: SessionLabel, Items: []string{SkipPrompt}},
		},
		{"single item", args{map[string]string{"foo": "bar"}},
			&promptui.Select{Label: SessionLabel, Items: []string{"foo", SkipPrompt}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := BuildPrompt(SessionLabel, utils.MapKeysToSlice(tt.args.sessionMap))
			assert.EqualValues(t, tt.want, actual)
		})
	}
}

func Test_selectFromMap(t *testing.T) {
	emptyPrompt := new(promptMock)
	emptyPrompt.On("Run").Return(0, "", nil)

	errorPrompt := new(promptMock)
	errorPrompt.On("Run").Return(1, "", fmt.Errorf("woops"))

	skipPrompt := new(promptMock)
	skipPrompt.On("Run").Return(0, SkipPrompt, nil)

	goodPrompt := new(promptMock)
	goodPrompt.On("Run").Return(0, "someSession", nil)

	type args struct {
		runner     Runner
		sessionMap map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{"empty case", args{
			runner:     emptyPrompt,
			sessionMap: map[string]string{},
		}, "", assert.NoError},
		{"error case", args{
			runner:     errorPrompt,
			sessionMap: map[string]string{},
		}, "", assert.Error},
		{"skip case", args{
			runner:     skipPrompt,
			sessionMap: map[string]string{},
		}, "", assert.NoError},
		{"normal case", args{
			runner:     goodPrompt,
			sessionMap: map[string]string{"someSession": "foo"},
		}, "foo", assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectFromMap(tt.args.runner, tt.args.sessionMap)
			if !tt.wantErr(t, err, fmt.Sprintf("SelectFromMap(%v, %v)", tt.args.runner, tt.args.sessionMap)) {
				return
			}
			assert.Equalf(t, tt.want, got, "SelectFromMap(%v, %v)", tt.args.runner, tt.args.sessionMap)
		})
	}
}
