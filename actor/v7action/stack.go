package v7action

type Stack struct {
	GUID        string
	Name        string
	Description string
}

func (actor Actor) GetStacks() ([]Stack, Warnings, error) {
	return nil, nil, nil
}
