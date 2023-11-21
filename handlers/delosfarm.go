package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/marcellinoknl/-BE-Mini-Projects---Delosaqua/database"
    "github.com/marcellinoknl/-BE-Mini-Projects---Delosaqua/models"
    "gorm.io/gorm"
)

type ErrorResponse struct {
    Message string `json:"message"`
}

func handleError(c *fiber.Ctx, status int, message string) error {
    return c.Status(status).JSON(ErrorResponse{Message: message})
}

//Create Farms Function
func CreateFarm(c *fiber.Ctx) error {
    farm := new(models.Delosfarm)
    if err := c.BodyParser(farm); err != nil {
        return handleError(c, fiber.StatusBadRequest, "Invalid request body")
    }

    // Create a new farm
    if err := database.DB.Db.Create(farm).Error; err != nil {
        return handleError(c, fiber.StatusInternalServerError, "Error creating farm: "+err.Error())
    }

    return c.Status(fiber.StatusCreated).JSON(farm)
}
//create function for update farm
func UpdateFarm(c *fiber.Ctx) error {
	farmID := c.Params("id")
	farm := new(models.Delosfarm)
	if err := c.BodyParser(farm); err != nil {
		return handleError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Check if the farm exists
	var existingFarm models.Delosfarm
	if err := database.DB.Db.First(&existingFarm, farmID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return handleError(c, fiber.StatusNotFound, "Farm not found")
		}
		return handleError(c, fiber.StatusInternalServerError, "Error finding farm: "+err.Error())
	}

	// Update the farm
	if err := database.DB.Db.Model(&existingFarm).Updates(farm).Error; err != nil {
		return handleError(c, fiber.StatusInternalServerError, "Error updating farm: "+err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(existingFarm)
}


//view all row data from farm 
func ListFarms(c *fiber.Ctx) error {
    farms := []models.Delosfarm{}
    database.DB.Db.Unscoped().Find(&farms) // Gunakan .Unscoped() di sini

    // Preload pond data for each farm, including soft deleted ponds
    for i := range farms {
        var pondsWithFarm []models.Pond
        if err := database.DB.Db.Where("farm_id = ?", farms[i].ID).Find(&pondsWithFarm).Error; err != nil {
            return handleError(c, fiber.StatusInternalServerError, "Error fetching ponds with farm data: "+err.Error())
        }

        // Update farm's ponds with the correct data
        for j := range pondsWithFarm {
            pondsWithFarm[j].Farm = farms[i] // Set the farm field of each pond to the current farm
        }
        farms[i].Ponds = pondsWithFarm
    }

    return c.Status(fiber.StatusOK).JSON(farms)
}

// GetFarmByID returns a specific farm by its ID with associated pond data, including soft deleted farms and ponds.
func GetFarmByID(c *fiber.Ctx) error {
    farmID := c.Params("id")
    var farm models.Delosfarm

    // Check if the farm exists and preload ponds, including soft deleted ponds
    if err := database.DB.Db.Unscoped().Preload("Ponds").First(&farm, farmID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return handleError(c, fiber.StatusNotFound, "Farm not found")
        }
        return handleError(c, fiber.StatusInternalServerError, "Error finding farm: "+err.Error())
    }

    // Preload farm data for each pond, including soft deleted ponds
    var pondsWithFarm []models.Pond
    if err := database.DB.Db.Unscoped().Preload("Farm").Where("farm_id = ?", farm.ID).Find(&pondsWithFarm).Error; err != nil {
        return handleError(c, fiber.StatusInternalServerError, "Error fetching ponds with farm data: "+err.Error())
    }

    // Update farm's ponds with the correct data
    farm.Ponds = pondsWithFarm

    return c.Status(fiber.StatusOK).JSON(farm)
}

// CreatePond creates a new pond.
func CreatePond(c *fiber.Ctx) error {
    pond := new(models.Pond)
    if err := c.BodyParser(pond); err != nil {
        return handleError(c, fiber.StatusBadRequest, "Invalid request body")
    }

    // Create a new pond
    if err := database.DB.Db.Create(pond).Error; err != nil {
        return handleError(c, fiber.StatusInternalServerError, "Error creating pond: "+err.Error())
    }

    // Fetch the farm data associated with the new pond using pond.FarmID
    var farm models.Delosfarm
    if err := database.DB.Db.First(&farm, pond.FarmID).Error; err != nil {
        return handleError(c, fiber.StatusInternalServerError, "Error fetching associated farm data: "+err.Error())
    }
    pond.Farm = farm

    return c.Status(fiber.StatusCreated).JSON(pond)
}

// UpdatePond updates an existing pond.
func UpdatePond(c *fiber.Ctx) error {
    pondID := c.Params("id")
    pond := new(models.Pond)
    if err := c.BodyParser(pond); err != nil {
        return handleError(c, fiber.StatusBadRequest, "Invalid request body")
    }

    // Check if the pond exists
    var existingPond models.Pond
    if err := database.DB.Db.First(&existingPond, pondID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return handleError(c, fiber.StatusNotFound, "Pond not found")
        }
        return handleError(c, fiber.StatusInternalServerError, "Error finding pond: "+err.Error())
    }

    // Update the pond
    if err := database.DB.Db.Model(&existingPond).Updates(pond).Error; err != nil {
        return handleError(c, fiber.StatusInternalServerError, "Error updating pond: "+err.Error())
    }

    // Fetch the farm data associated with the updated pond using existingPond.FarmID
    var farm models.Delosfarm
    if err := database.DB.Db.First(&farm, existingPond.FarmID).Error; err != nil {
        return handleError(c, fiber.StatusInternalServerError, "Error fetching associated farm data: "+err.Error())
    }
    existingPond.Farm = farm

    return c.Status(fiber.StatusOK).JSON(existingPond)
}

// ListPonds returns a list of all ponds with associated farm information, including soft deleted ponds.
func ListPonds(c *fiber.Ctx) error {
    ponds := []models.Pond{}
    database.DB.Db.Unscoped().Find(&ponds) // Use .Unscoped() here

    // Preload farm data for each pond
    for i := range ponds {
        var farm models.Delosfarm
        if err := database.DB.Db.Where("id = ?", ponds[i].FarmID).First(&farm).Error; err != nil {
            if err == gorm.ErrRecordNotFound {
                // Farm not found, continue to the next pond
                continue
            }
            return handleError(c, fiber.StatusInternalServerError, "Error fetching farm data for pond: "+err.Error())
        }
        ponds[i].Farm = farm
    }

    return c.Status(fiber.StatusOK).JSON(ponds)
}

// GetPondByID returns a specific pond by its ID with associated farm information, including soft deleted farms and ponds.
func GetPondByID(c *fiber.Ctx) error {
    pondID := c.Params("id")
    var pond models.Pond

    // Check if the pond exists
    if err := database.DB.Db.Unscoped().Where("id = ?", pondID).First(&pond).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return handleError(c, fiber.StatusNotFound, "Pond not found")
        }
        return handleError(c, fiber.StatusInternalServerError, "Error finding pond: "+err.Error())
    }

    // Check if the associated farm exists, including soft deleted farms
    var farm models.Delosfarm
    if err := database.DB.Db.Unscoped().Where("id = ?", pond.FarmID).First(&farm).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return handleError(c, fiber.StatusInternalServerError, "Farm not found for pond")
        }
        return handleError(c, fiber.StatusInternalServerError, "Error fetching farm data for pond: "+err.Error())
    }
    pond.Farm = farm

    return c.Status(fiber.StatusOK).JSON(pond)
}

//soft delete
// DeleteFarm soft deletes an existing farm.
func DeleteFarm(c *fiber.Ctx) error {
    farmID := c.Params("id")
    var farm models.Delosfarm

    // Check if the farm exists
    if err := database.DB.Db.First(&farm, farmID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return handleError(c, fiber.StatusNotFound, "Farm not found")
        }
        return handleError(c, fiber.StatusInternalServerError, "Error finding farm: "+err.Error())
    }

    // Soft delete the farm by updating DeletedAt field
    if err := database.DB.Db.Delete(&farm).Error; err != nil {
        return handleError(c, fiber.StatusInternalServerError, "Error soft deleting farm: "+err.Error())
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Farm soft deleted successfully",
    })
}

// DeletePond soft deletes an existing pond.
func DeletePond(c *fiber.Ctx) error {
    pondID := c.Params("id")
    var pond models.Pond

    // Check if the pond exists
    if err := database.DB.Db.First(&pond, pondID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return handleError(c, fiber.StatusNotFound, "Pond not found")
        }
        return handleError(c, fiber.StatusInternalServerError, "Error finding pond: "+err.Error())
    }

    // Soft delete the pond by using the Delete method with the instance
    if err := database.DB.Db.Delete(&pond).Error; err != nil {
        return handleError(c, fiber.StatusInternalServerError, "Error soft deleting pond: "+err.Error())
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Pond soft deleted successfully",
    })
}
