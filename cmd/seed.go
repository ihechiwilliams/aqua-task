package cmd

import (
	"fmt"
	"log"

	"aqua-backend/internal/appbase"
	customers2 "aqua-backend/internal/repositories/customers"
	"aqua-backend/pkg/postgres"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"gorm.io/plugin/dbresolver"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database with initial data",
	Long:  `This command seeds the database with predefined cloud resources.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := lo.Must(appbase.LoadConfig())

		db, err := postgres.InitDB(
			cfg.DatabaseURL,
		)
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}

		// Seed customers
		customers := []struct {
			Name  string
			Email string
		}{
			{Name: "Customer One", Email: "customer1@example.com"},
			{Name: "Customer Two", Email: "customer2@example.com"},
		}

		for _, c := range customers {
			customerID := uuid.New()
			err := db.Clauses(dbresolver.Write).Table(`"customers"`).Exec(
				`INSERT INTO customers (id, name, email) VALUES ($1, $2, $3) 
     ON CONFLICT (email) DO NOTHING`,
				customerID, c.Name, c.Email,
			).Error
			if err != nil {
				log.Printf("Failed to seed customer %s: %v", c.Name, err)
			} else {
				log.Printf("Seeded customer: %s", c.Name)
			}
		}

		// Retrieve the customer ID for "customer1@example.com"
		var customer customers2.DBCustomer
		if err := db.Table(`"customers"`).Raw(
			`SELECT id FROM customers WHERE email = ?`,
			"customer1@example.com",
		).Scan(&customer).Error; err != nil {
			log.Fatalf("Failed to retrieve customer ID: %v", err)
		}

		fmt.Printf("Customer ID: %s\n", customer.ID)

		// Seed resources
		resources := []struct {
			Name   string
			Type   string
			Region string
		}{
			{Name: "Resource1", Type: "Compute", Region: "us-west-1"},
			{Name: "Resource2", Type: "Storage", Region: "us-east-1"},
		}

		for _, r := range resources {
			resourceID := uuid.New()
			if err := db.Exec(
				`INSERT INTO resources (id, name, type, region, customer_id) VALUES ($1, $2, $3, $4, $5) 
                 ON CONFLICT (name) DO NOTHING`,
				resourceID, r.Name, r.Type, r.Region, customer.ID,
			).Error; err != nil {
				log.Printf("Failed to seed resource %s: %v", r.Name, err)
			} else {
				log.Printf("Seeded resource: %s", r.Name)
			}
		}

		log.Println("Seeding completed successfully.")
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
