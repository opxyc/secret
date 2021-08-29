package secret

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/opxyc/secret/cipher"
)

// New gives you a new Vault. The Vault will read and write to filepath
// after the contents are encoded/decoded using the encodingKey.
func New(encodingKey, filepath string) *Vault {
	v := &Vault{
		encodingKey: encodingKey,
		filePath:    filepath,
		keyValues:   make(map[string]string),
	}
	return v
}

type Vault struct {
	encodingKey string
	filePath    string
	keyValues   map[string]string
}

// load loads the secrets from a given file to Vault.
func (v *Vault) load() error {
	f, err := os.Open(v.filePath)
	if err != nil {
		v.keyValues = map[string]string{}
		return nil
	}
	defer f.Close()

	r, err := cipher.DecryptReader(v.encodingKey, f)
	if err != nil {
		return err
	}

	err = v.readKeyValues(r)
	if err != nil {
		return errors.New("invalid encoding key")
	}
	return nil
}

// readKeyValues decodes file contents to json format.
func (v *Vault) readKeyValues(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&v.keyValues)
}

// save is used to save the secrets to file.
func (v *Vault) save() error {
	// Create a temporary file and write to it.
	f, err := os.CreateTemp(filepath.Dir(v.filePath), filepath.Base(v.filePath))
	if err != nil {
		fmt.Println("createTemp", err)
		return err
	}

	w, err := cipher.EncryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}

	err = v.writeKeyValues(w)
	if err != nil {
		return err
	}
	f.Close()

	return os.Rename(f.Name(), v.filePath)
}

func (v *Vault) writeKeyValues(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v.keyValues)
}

// Get retrieves the value corresponding to the key.
func (v *Vault) Get(key string) (string, error) {
	if len(v.keyValues) == 0 {
		err := v.load()
		if err != nil {
			return "", err
		}
	}

	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("secret: no value for the key")
	}
	return value, nil
}

// Set adds key,value to the secrets file.
func (v *Vault) Set(key, value string) error {
	err := v.load()
	if err != nil {
		return err
	}

	v.keyValues[key] = value
	return v.save()
}

func (v *Vault) List() ([]string, error) {
	err := v.load()
	if err != nil {
		return nil, err
	}

	var keys []string
	for k := range v.keyValues {
		keys = append(keys, k)
	}
	return keys, nil
}

func (v *Vault) ChangeEncodingKey(newEncodingKey string) error {
	err := v.load()
	if err != nil {
		return err
	}

	v.encodingKey = newEncodingKey

	return v.save()
}

func (v *Vault) Remove(key string) error {
	err := v.load()
	if err != nil {
		return err
	}
	_, ok := v.keyValues[key]
	if !ok {
		return errors.New(fmt.Sprintf("no value set for '%s'", key))
	}
	delete(v.keyValues, key)

	return v.save()
}
