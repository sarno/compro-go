package cmd

import (
	"compro/config"
	"compro/database/seeds"
	"log"

	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "db:seed",
	Short: "Seed the database with initial data.",
	Long:  `This command connects to the database and runs all the required seeders to populate the tables with initial data.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting database seeding...")
		// 1. Load application configuration
		cfg := config.NewConfig()

		postgres, err := cfg.ConnectDB()
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		db := postgres.DB
		// 3. Run the seeders
		log.Println("Seeding users...")
		seeds.SeedAdmin(db)

		log.Println("Database seeding completed successfully!")
	},
	
}

func init() {
	rootCmd.AddCommand(seedCmd)
}