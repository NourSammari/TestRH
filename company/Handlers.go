package company

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitializeDB(database *gorm.DB) {
	db = database
}
func GetCompanies(c *gin.Context) {
	var companies []Company
	db.Find(&companies)
	c.JSON(200, companies)
}

/*func CreateCompany(c *gin.Context) {
	var company Company
	stmt, err := db.Prepare("INSERT INTO companies (name, location, industry, emp_number, emp_fax, emp_email, founder, date_founded, website, description) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		c.JSON(404, gin.H{"error": "error occurred while preparing statement"})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(company.NameCompany, company.LocationCompany, company.Industry, company.EmpNumber, company.EmpFax, company.EmpEmail, company.Founder, company.DateFounded, company.Website, company.Description)
	if err != nil {
		c.JSON(404, gin.H{"error": "error occurred while executing statement"})
		return
	}
}*/

func CreateCompany(c *gin.Context) {
	var company Company
	if err := c.BindJSON(&company); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	 existingCompany := Company{}
	if err := db.Where("name_company = ?", company.NameCompany).First(&existingCompany).Error; err == nil {
		c.JSON(400, gin.H{"error": "Company already exists"})
		return
	} 

	if err := db.Create(&company).Error; err != nil {
		c.JSON(500, gin.H{"error": "error occurred while creating company"})
		return
	}

	c.JSON(200, company)
}

func UpdateCompany(c *gin.Context) {
	id := c.Param("id")
	var company Company
	if err := db.Where("id_company = ?", id).First(&company).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Company not found"})
		return
	}
	if err := c.BindJSON(&company); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Save(&company)
	c.JSON(200, company)
}

func DeleteCompany(c *gin.Context) {
	id := c.Param("id")

	var company Company
	db.Where("id_company = ?", id).Delete(&company)

	c.JSON(200, gin.H{"message": "Company deleted"})

}
func GetCompanyByID(c *gin.Context) {
	CompanyID := c.Param("id")

	var company Company

	if err := db.First(&company, "id_company = ?", CompanyID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(200, company)
}
