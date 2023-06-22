package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"time"
	"fmt"
)

type User struct {
	gorm.Model
	FirstName      string    `gorm:"uniqueIndex"`
	LastName       string    `gorm:"uniqueIndex"`
	Email          string    `gorm:"not null"`
	Country        string    `gorm:"not null"`
	Role           string    `gorm:"not null"`
	Age            int       `gorm:"not null;size:3"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

func main() {
    //Create a new Postgresql database connection
    dsn := "host=<your_host> user=<your_user> password=<your_password> dbname=<your_dbname> port=<your_port>"

    // Open a connection to the database
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database: " + err.Error())
    }

	// AutoMigrate will create the necessary tables based on the defined models/structs
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}

	// ... Define a new post instance ...

	newUser := User{
	    FirstName: "Jane",
	    LastName: "Doe",
	    Email: "janedoe@gmail.com",
	    Country: "Spain",
	    Role: "Chef",
	    Age: 30,
	}

	// ... Create a new user record...
		result := db.Create(&newUser)
		if result.Error != nil {
			panic("failed to create user: " + result.Error.Error())
		}
		
		// ... Handle successful creation ...
		fmt.Printf("New user %s %s was created successfully!\\n", newUser.FirstName, newUser.LastName)

	// ... RETRIEVING RECORDS WITHs with GORM ...

		// ... Retrieve the first user from the database ...
			//     var user User
			//     result := db.First(&user)
			//     if result.Error != nil {
			//         panic("failed to retrieve user: " + result.Error.Error())
			// }

			// Use the user record
			// fmt.Printf("User ID: %d, Name: %s %s, Email: %s\n", user.ID, user.FirstName, user.LastName, user.Email)


		// Retrieve the first user from the database
			// 	var user User
			// 	result := db.First(&user)
			// 	if result.Error != nil {
			// 		panic("failed to retrieve user: " + result.Error.Error())
			// }

			// ... Use the user record ...
			// fmt.Printf("User ID: %d, Name: %s %s, Email: %s\n", user.ID, user.FirstName, user.LastName, user.Email)


		// Retirve records based on some certain conditions
			// var users []User
			// result := db.Where("ID = ?", 1).Find(&users)
			// if result.Error != nil {
			// 	panic("failed to retrieve user: " + result.Error.Error())
			// }
			
			// // iterate over the users slice and print the details of each user
			// for _, user := range users {
			// 	fmt.Printf("User ID: %d, Name: %s %s, Email: %s\n", user.ID, user.FirstName, user.LastName, user.Email)
			// }
		
		// Chainning multiple conditions
			// var users []User
			// result := db.Where("FirstName = ?", "Jane").Where("Country = ?", "Spain").Find(&users)
			// if result.Error != nil {
			// 	panic("failed to retrieve users: " + result.Error.Error())
			// }
			
			// // Use the user records
			// for _, user := range users {
			// 	fmt.Printf("User ID: %d, Name: %s %s, Email: %s\n", user.ID, user.FirstName, user.LastName, user.Email)
			// }
		
		// Multiple records with no specific condition
			// var users []User
			// result := db.Find(&users)
			// if result.Error != nil {
			// 	// handle error
			// 	panic("failed to retrieve users: " + result.Error.Error())
			// }
			
			// // Iterate over the users slice and print the details of each user
			// for _, user := range users {
			// 	fmt.Printf("User ID: %d, Name: %s %s, Email: %s\n", user.ID, user.FirstName, user.LastName, user.Email)
			// }

	// ... UPDATING RECORDS WITH GORM ...
		// Retrieve the record you want to update
			// var user User
			// result := db.First(&user, 1)
			// if result.Error != nil {
			// 	panic("failed to retrieve user: " + result.Error.Error())
			// }
		
			// // Modify the attributes of the retrieved record, in this case, the first three columns
			// user.FirstName = "Agnes"
			// user.LastName = "Doe"
			// user.Email = "agnesdoe@example.com"
		
			// // Save the changes back to the database
			// result = db.Save(&user)
			// if result.Error != nil {
			// 	panic("failed to update user: " + result.Error.Error())
			// }
		
			// fmt.Println("User updated successfully")
		// Update the 'role' column of the record with ID 1
			// result := db.Model(&User{}).Where("id = ?", 1).Update("role", "admin")
			// if result.Error != nil {
			// 	panic("failed to update user: " + result.Error.Error())
			// }

			// fmt.Println("User role updated successfully")
		// Update the record with ID 1 using the `User` struct
			// result := db.Model(&User{}).Where("id = ?", 1).Updates(User{
			// 	FirstName: "John",
			// 	LastName:  "Doe",
			// 	Email:     "johndoe@example.com",
			// })
			// if result.Error != nil {
			// 	panic("failed to update user: " + result.Error.Error())
			// }
		
			// fmt.Println("User updated successfully")

	// ... DELETING RECORDS WITH GORM
		// Deleting a single user record
			// var user User
			// result := db.First(&user)
			// if result.Error != nil {
			// 	panic("failed to retrieve user: " + result.Error.Error())
			// }
			
			// result = db.Delete(&user)
			// if result.Error != nil {
			// 	panic("failed to delete user: " + result.Error.Error())
			// } else if result.RowsAffected == 0 {
			// 	panic("no user record was deleted")
			// } else {
			// 	fmt.Println("User record deleted successfully")
			// }
		// Delete multiple records with a condition
			// Delete the record where the country is "Spain"
			// record := db.Where("country = ?", "Spain").Delete(&User{})
			// if record.Error != nil {
			// 	panic("failed to delete user: " + record.Error.Error())
			// }
			// fmt.Println(record.RowsAffected, "user record(s) deleted successfully")

	// ...TRANSACTIONS WITH GORM...
				
			// tx := db.Begin())

			// 	// Create a new user
			// 	newUser := User{
			// 		FirstName: "Billy",
			// 		LastName: "John",
			// 		Email: "billy56@gmail.com",
			// 		Country: "Germany",
			// 		Role: "Developer Advocate",
			// 		Age: 40,
			// 	}

			// 	// Perform database operations within the transaction
			// 	if err := tx.Create(&newUser).Error; err != nil {
			// 		tx.Rollback() // Rollback the transaction if an error occurs
			// 		panic("failed to create user: " + err.Error())
			// 	}

			// 	// Update the user's profile
			// 	newUser.Country = "Morocco"
			// 	if err := tx.Save(&newUser).Error; err != nil {
			// 		tx.Rollback() // Rollback the transaction if an error occurs
			// 		panic("failed to update user: " + err.Error())
			// 	}

			// 	// Commit the transaction if all operations succeed
			// 	tx.Commit()

			// 	fmt.Println("User created and updated successfully")

	// ...Hooks...
		// // ... User Struct ...
		// func (u *User) BeforeCreate(tx *gorm.DB) error {
		// 	// Perform some actions before creating a user
		// 	fmt.Println("Preparing to create user:", u.FirstName, u.LastName)
		// 	return nil
		// }
		
		// func (u *User) AfterCreate(tx *gorm.DB) error {
		// 	// Perform some actions after creating a user
		// 	fmt.Println("User created successfully:", u.FirstName, u.LastName)
		// 	return nil
		// }
		
		// // ... main function ...
		// func main() {
		// 		// ... database connection code ...
		
		// 	// ... Auto migration code ...
		
		// 	// Begin a transaction
		// 	tx := db.Begin()
		
		// 	// Create a new user
		// 	newUser := User{
		// 		FirstName: "John",
		// 		LastName: "mark",
		// 		Email: "john@gmail.com",
		// 		Country: "Argentina",
		// 		Role: "Technical Writer",
		// 		Age: 35,
		// 	}
		
		// 	// Perform database operations within the transaction
		// 	...
		// 	// Update the user's profile
		// 	...
		// 	// Commit the transaction if all operations succeed
		// 	tx.Commit()
			
		// 	fmt.Println("User created and updated successfully")
		// }
						
}
