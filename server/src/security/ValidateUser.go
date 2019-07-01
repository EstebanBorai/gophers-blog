package security

import (
	"fmt"

	data "github.com/estebanborai/go-server-sample/server/src/data"
	models "github.com/estebanborai/go-server-sample/server/src/models"
	"github.com/gin-gonic/gin"
)

// ValidateUser makes sure the plainPassword matches with the userName given
func ValidateUser(c *gin.Context, userName string, plainPassword string) (models.User, error) {
	var user models.User
	var userSecret models.UserSecret

	db := data.Connection(c)

	if userResult := db.Where(&models.User{UserName: userName}).First(&user); userResult.Error == nil {
		if credsResult := db.Where(&models.UserSecret{UserID: user.ID}).First(&userSecret); credsResult.Error == nil {
			if ValidatePassword(userSecret.Hash, plainPassword) {
				return user, nil
			}

			return user, fmt.Errorf("Invalid password")
		}

		return user, fmt.Errorf("Invalid username")
	}

	return user, fmt.Errorf(fmt.Sprintf("%s doesn't exist.", userName))
}
