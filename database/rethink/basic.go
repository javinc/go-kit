package rethink

import r "gopkg.in/gorethink/gorethink.v3"

// Find basic query
func Find(term r.Term, result interface{}) error {
	r, err := Run(term)
	if err != nil {
		return err
	}

	err = r.All(result)
	if err != nil {
		return err
	}

	return nil
}

// FindOne basic query
func FindOne(term r.Term, result interface{}) error {
	r, err := Run(term)
	if err != nil {
		return err
	}

	err = r.One(result)
	if err != nil {
		return err
	}

	return nil
}

// Create basic query
func Create(term r.Term) (string, error) {
	r, err := RunWrite(term)
	if err != nil {
		return "", err
	}

	// it should return atleast 1 key
	return r.GeneratedKeys[0], nil
}

// Update basic query
func Update(term r.Term) error {
	_, err := RunWrite(term)

	return err
}

// Remove basic query
func Remove(term r.Term) error {
	_, err := RunWrite(term)

	return err
}

// Count basic query
func Count(term r.Term, count *int) error {
	r, err := Run(term.Count())
	if err != nil {
		return err
	}

	err = r.One(count)
	if err != nil {
		return err
	}

	return nil
}

// CreateTable create table
func CreateTable(name string) error {
	_, err := RunWrite(r.TableCreate(name).Wait())

	return err
}
