package types

type TestErr struct{ S string }

func (e *TestErr) Error() string { return e.S }
