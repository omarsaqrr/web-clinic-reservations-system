package main

import (
	"log"
	"net/http"

	"example/hello/models"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*type User struct {
	gorm.Model
	Username string
	Password string
	Role     string // "doctor" or "patient"
}
*/

var doctors []models.Doctor
var patients []models.Patient

//var doctorSlots []models.DoctorSlots

var DB *gorm.DB

func main() {
	initDB()

	// Create a Gin router
	router := gin.Default()

	// Define API endpoints

	router.POST("/sign-up/doctor", signUpDoctor)
	router.POST("/sign-up/patient", signUpPatient)
	router.POST("/sign-in", signIn)
	router.POST("/insertDoctorSlot", insertDoctorSlot)
	router.GET("/getDoctorSlot", getDoctorSlots)
	router.POST("/creatAppoinment", chooseSlot)
	router.GET("/getallapponment", getAllappoinment)
	router.GET("/getallapponmentforpatient", getAppointmentsForPatient)
	router.PUT("/updatetheappointment", updateAppointment)
	router.DELETE("/cancelappointment", cancelAppointment)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
func initDB() {
	// Replace these with your actual PostgreSQL connection details

	dsn := "host=castor.db.elephantsql.com user=pqztdjdw password=VWvbG6Aig5zP-HuLo3YBtGyELLNa40GH dbname=pqztdjdw port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Auto-migrate the models to create the database table
	DB.AutoMigrate(&models.Doctor{})
	DB.AutoMigrate(&models.Patient{})
	DB.AutoMigrate(&models.DoctorSlots{})
	DB.AutoMigrate(&models.Appointment{})

}

func signIn(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var doctor models.Doctor
	var patient models.Patient

	// Check if the user is a doctor
	result := DB.Where("email = ? AND password = ? AND role = ?", input.Email, input.Password, "doctor").First(&doctor)
	if result.Error != nil {
		// If not, check if the user is a patient
		result = DB.Where("email = ? AND password = ? AND role = ?", input.Email, input.Password, "patient").First(&patient)
		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sign-in successful"})
}

func signUpDoctor(c *gin.Context) {
	var input struct {
		Email    string `json:"Email"`
		Password string `json:"password"`
		Role     string `json:"role"` // "doctor" or "patient"
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if the email is already taken

	if err := DB.Where("email = ?", input.Email).First(&doctors).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already taken"})
		return
	}

	/*for _, doctor := range doctors {
		if doctor.Email == input.Email {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already taken "})
			return
		}
	}
	*/

	// Create a new doctor
	newDoctor := models.Doctor{
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
	}

	doctors = append(doctors, newDoctor)
	doctor := models.Doctor{Email: input.Email, Password: input.Password, Role: input.Role}
	result := DB.Create(&doctor)
	//result := DB.Where("email = ?", doctor.Email).First(&doctor)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, doctors)
}

func signUpPatient(c *gin.Context) {
	var input struct {
		Email    string `json:"Email"`
		Password string `json:"password"`
		Role     string `json:"role"` // "doctor" or "patient"
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if the email is already taken

	if err := DB.Where("email = ?", input.Email).First(&patients).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already taken"})
		return
	}
	/*for _, patient := range patients {
		if patient.Email == input.Email {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already taken "})
			return
		}
	}
	*/

	// Create a new patient
	newPatient := models.Patient{
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
	}

	patients = append(patients, newPatient)
	patient := models.Patient{Email: input.Email, Password: input.Password, Role: input.Role}
	result := DB.Create(&patient)
	//result := DB.Where("email = ?", patient.Email).First(&patient)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, patients)

}

func insertDoctorSlot(c *gin.Context) {

	var input struct {
		Email string `json:"Email"`

		/*Slots []struct {
			StartTime time.Time `json:"start_time"`
			EndTime   time.Time `json:"end_time"`
		} `json:"slots"`
		*/

		Date string `json:"Date"`
		Hour string `json:"Hour"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if the user is a doctor
	var doctor models.Doctor
	result := DB.Where("email = ? AND role = ?", input.Email, "doctor").First(&doctor)
	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "doctor not found"})
		return
	}

	// Create a new doctor slot

	newdoctorSlot := models.DoctorSlots{
		DoctorID:    doctor.ID,
		DoctorEmail: input.Email,
		Date:        input.Date,
		Hour:        input.Hour,
	}

	/*for _, slot := range input.Slots {
	newDoctorSlot := models.DoctorSlots{
		DoctorID:  doctor.ID,
		StartTime: slot.StartTime,
		EndTime:   slot.EndTime,
	}
	*/

	result = DB.Create(&newdoctorSlot)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create slot"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Slots created successfully"})

	/*doctorSlots = append(doctorSlots, newdoctorSlot)
	doctorslot := models.DoctorSlots{DoctorEmail: input.Email, StartTime: input.StartTime, EndTime: input.EndTime}
	result = DB.Create(&doctorslot)
	//result := DB.Where("email = ?", doctor.Email).First(&doctor)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create slot"})
		return
	}

	c.JSON(http.StatusOK, doctorSlots)
	//c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	*/

}

func getDoctorSlots(c *gin.Context) {
	//doctorEmail := c.Param("doctor_email")
	var input struct {
		DoctorEmail string `json:"doctor_email"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var slots []models.DoctorSlots
	result := DB.Where("doctor_email = ?", input.DoctorEmail).Find(&slots)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found or no slots available"})
		return
	}

	c.JSON(http.StatusOK, slots)
}

/*func getAllDoctors(c *gin.Context) {
	var doctors []models.Doctor
	result := DB.Find(&doctors)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch doctors"})
		return
	}

	c.JSON(http.StatusOK, doctors)
}
*/

func chooseSlot(c *gin.Context) {
	var input struct {
		SlotID       uint   `json:"slot_id"`
		PatientEmail string `json:"patient_email"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if the slot exists and is available
	var slot models.DoctorSlots
	var patient models.Patient
	result := DB.First(&slot, input.SlotID)
	if result.Error != nil || slot.DoctorID == 0 || slot.ISBooked {
		c.JSON(http.StatusNotFound, gin.H{"error": "Slot not found or already booked"})
		return
	}

	// Fetch patient information by email
	result = DB.Where("email = ? AND role = ?", input.PatientEmail, "patient").First(&patient)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	// Create an appointment
	appointment := models.Appointment{
		//DoctorID:     slot.DoctorID,
		//SlotID:       slot.ID,
		PatientID:    patient.ID,
		PatientEmail: input.PatientEmail,
		DoctorEmail:  slot.DoctorEmail,
		Date:         slot.Date,
		Hour:         slot.Hour,
		// Set the appointment time based on the slot's start time
		// For example, you can set it as slot.StartTime
		// Add any other fields as needed
	}

	// Save the appointment to the database

	result = DB.Create(&appointment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appoinment"})
		return
	}

	// Mark the slot as booked
	slot.ISBooked = true
	DB.Save(&slot)

	c.JSON(http.StatusOK, gin.H{"message": "Appointment booked successfully"})
}

func updateAppointment(c *gin.Context) {
	var input struct {
		AppointmentID uint   `json:"appointment_id"`
		DoctorEmail   string `json:"doctor_email"`
		//Date          string `json:"Date"`
		//Hour          string `json:"Hour"`
		SlotID uint `json:"slot_id"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var slots models.DoctorSlots
	result := DB.Where("id = ?", input.SlotID).First(&slots)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": " no slots with this id"})
		return
	}

	result = DB.Where("doctor_email = ?", input.DoctorEmail).First(&slots)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found or no slots available"})
		return
	}

	if slots.ISBooked {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Doctor slot is already booked"})
		return
	}

	var appointment models.Appointment
	//var slot models.DoctorSlots
	result = DB.Where("id = ?", input.AppointmentID).First(&appointment)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	appointment.DoctorEmail = input.DoctorEmail
	appointment.Date = slots.Date
	appointment.Hour = slots.Hour
	slots.ISBooked = true
	DB.Save(&slots)

	DB.Save(&appointment)

	c.JSON(http.StatusOK, appointment)

}

func getAllappoinment(c *gin.Context) {
	var appointment []models.Appointment
	result := DB.Find(&appointment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch doctors"})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func getAppointmentsForPatient(c *gin.Context) {
	var input struct {
		PatientEmail string `json:"patient_email"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var appointment []models.Appointment
	result := DB.Where("patient_email = ?", input.PatientEmail).Find(&appointment)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found or no slots available"})
		return
	}

	c.JSON(http.StatusOK, appointment)

}

func cancelAppointment(c *gin.Context) {
	var input struct {
		AppointmentID uint `json:"appointment_id"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var appointment models.Appointment
	result := DB.Where("id = ?", input.AppointmentID).First(&appointment)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	// Fetch the associated doctor slot
	var doctorSlot models.DoctorSlots
	result = DB.Where("doctor_email = ?", appointment.DoctorEmail).First(&doctorSlot)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch doctor slot"})
		return
	}

	// Update the doctor slot to mark it as available
	doctorSlot.ISBooked = false
	DB.Save(&doctorSlot)

	// Delete the appointment from the database
	if result = DB.Delete(&appointment, input.AppointmentID); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel appointment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment canceled successfully"})
}
