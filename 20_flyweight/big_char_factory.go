package main

type bigCharFactory struct {
	pool map[string]*bigChar
}

var factory = &bigCharFactory{
	pool: map[string]*bigChar{},
}

func getFactory() *bigCharFactory {
	return factory
}

func (f *bigCharFactory) getBigChar(charname string) (*bigChar, error) {
	bc := f.pool[charname]
	if bc == nil {
		bc, err := newBigChar(charname)
		if err != nil {
			return nil, err
		}
		f.pool[charname] = bc
		return bc, nil
	}
	return bc, nil
}
