package secret

import (
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"os"
	"sync"

	"github.com/opxyc/secret/cipher"
)

func File(encodingKey, filepath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filePath:    filepath,
		keyValues:   make(map[string]string),
	}
}

type Vault struct {
	encodingKey string
	filePath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

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

func (v *Vault) readKeyValues(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&v.keyValues)
}

func (v *Vault) save() error {
	writeFile, err := os.OpenFile(v.filePath, os.O_RDWR|os.O_CREATE, fs.ModeExclusive)
	if err != nil {
		return err
	}
	defer writeFile.Close()

	w, err := cipher.EncryptWriter(v.encodingKey, writeFile)
	if err != nil {
		return err
	}
	return v.writeKeyValues(w)
}

func (v *Vault) writeKeyValues(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v.keyValues)
}

func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return "", err
	}

	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("secret: no value for the key")
	}
	return value, nil
}

func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return err
	}

	v.keyValues[key] = value
	return v.save()
}
