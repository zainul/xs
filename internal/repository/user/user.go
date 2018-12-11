package user

import (
	"fmt"
	"strconv"

	"github.com/zainul/xs/internal/constant"
	"github.com/zainul/xs/internal/delivery/database"
	"github.com/zainul/xs/internal/entity"
	"github.com/zainul/xs/internal/pkg/cache"
	"github.com/zainul/xs/internal/pkg/converter"
	"github.com/zainul/xs/internal/pkg/counter"
	"github.com/zainul/xs/internal/repository"
)

type userRepository struct {
	db      database.UserStore
	cached  cache.Cache
	counter counter.Counter
}

// NewUserRepository ...
func NewUserRepository(
	db database.UserStore,
	cached cache.Cache,
	counter counter.Counter,
) repository.UserRepository {
	return &userRepository{
		db,
		cached,
		counter,
	}
}

func (u *userRepository) Save(user entity.User) error {
	err := u.db.Save(user)

	if err != nil {
		return err
	}

	key := fmt.Sprintf("%v_%v", constant.CachedPrefix, constant.UserInfo)

	// because is optional process
	go func(key string, v interface{}) {

		strData, err := converter.ToJSONString(user)

		if err != nil {
			fmt.Println("error while saved to cached ", err)
		}

		u.cached.Set(key, strData)
	}(key, user)

	return nil
}
func (u *userRepository) Edit(user entity.User, email string) error {
	return nil
}
func (u *userRepository) GetByField(field interface{}, fieldName string) *entity.User {
	return nil
}
func (u *userRepository) GenerateAccountNumber() (string, error) {
	val, err := u.counter.Increament(constant.CounterAccountNumber)
	return strconv.Itoa(int(val)), err
}
