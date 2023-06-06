package transactions

import (
	"encoding/json"
	"os"
)

type JSONFileRepo struct {
	filename string
	cache    InMemoryRepository
}

func CreateJSONRepo(file string) JSONFileRepo {
	cacheRepo := CreateInMemoryRepo()
	previousData, err := readPreviousData(file)
	if err == nil {
		cacheRepo.Ts = previousData
	}

	return JSONFileRepo{file, cacheRepo}
}

func readPreviousData(file string) (ts map[int]Transaction, err error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &ts)
	return
}

func (r *JSONFileRepo) commit() (err error) {
	data, err := json.MarshalIndent(&r.cache.Ts, "", "    ")
	if err != nil {
		return
	}
	err = os.WriteFile(r.filename, data, 0644)
	return
}

func (r *JSONFileRepo) GetAll() ([]Transaction, error) {
	return r.cache.GetAll()
}

func (r *JSONFileRepo) Store(t Transaction) (err error) {
	err = r.cache.Store(t)
	r.commit()
	return
}

func (r *JSONFileRepo) Delete(id int) (err error) {
	err = r.cache.Delete(id)
	r.commit()
	return
}

func (r *JSONFileRepo) Replace(t Transaction) (err error) {
	err = r.cache.Replace(t)
	r.commit()
	return
}

func (r *JSONFileRepo) Find(id int) (Transaction, error) {
	return r.cache.Find(id)
}
