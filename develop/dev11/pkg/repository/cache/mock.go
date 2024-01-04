package cache

import "dev11/pkg/models"

func (o *UserCacheRepo) addTestUser() {
	testUser := models.NewUser("1")

	o.PutUser(testUser.Id, testUser)
}
