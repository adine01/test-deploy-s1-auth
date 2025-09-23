package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

// InitDB initializes the database connection pool

// InitDB initializes the database connection pool
func InitDB() error {
	var err error
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return fmt.Errorf("DATABASE_URL environment variable is required")
	}

	log.Println("Connecting to database...")

	// Parse the database URL to get host details
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Printf("Failed to parse database URL: %v", err)
		return err
	}

	// Force IPv4 resolution by resolving hostname to IPv4 only
	host := config.ConnConfig.Host
	log.Printf("Resolving hostname: %s", host)

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Printf("Failed to resolve hostname: %v", err)
		return err
	}

	// Find the first IPv4 address
	var ipv4Addr string
	for _, ip := range ips {
		if ip.To4() != nil {
			ipv4Addr = ip.String()
			log.Printf("Found IPv4 address: %s", ipv4Addr)
			break
		}
	}

	if ipv4Addr == "" {
		return fmt.Errorf("no IPv4 address found for hostname: %s", host)
	}

	// Replace hostname with IPv4 address in config
	config.ConnConfig.Host = ipv4Addr
	log.Printf("Updated database host to IPv4: %s:%d", ipv4Addr, config.ConnConfig.Port)

	// Configure connection with timeouts
	config.ConnConfig.DialFunc = func(ctx context.Context, network, addr string) (net.Conn, error) {
		d := &net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}
		return d.DialContext(ctx, "tcp4", addr)
	}

	log.Println("Creating database connection pool...")
	// Create connection pool
	db, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Printf("❌ Failed to create database pool: %v", err)
		return err
	}
	log.Println("✅ Database pool created successfully")

	log.Println("Testing database connection...")
	// Test the connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	if err := db.Ping(ctx); err != nil {
		log.Printf("❌ Failed to ping database: %v", err)
		return err
	}
	log.Println("✅ Database ping successful")

	log.Println("✅ Successfully connected to database")

	log.Println("Creating users table if needed...")
	// Create users table if it doesn't exist
	if err := createUsersTable(); err != nil {
		log.Printf("❌ Failed to create users table: %v", err)
		return err
	}

	log.Println("✅ Database initialization completed successfully")
	return nil
}

// CloseDB closes the database connection pool
func CloseDB() {
	if db != nil {
		db.Close()
	}
}

// IsDBConnected checks if the database connection is available
func IsDBConnected() bool {
	if db == nil {
		return false
	}

	// Create context with timeout to prevent hanging
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := db.Ping(ctx)
	return err == nil
}

// createUsersTable creates the users table if it doesn't exist
func createUsersTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		name VARCHAR(255) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);

	-- Create index on email for faster lookups
	CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
	`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	
	_, err := db.Exec(ctx, query)
	if err != nil {
		log.Printf("Error creating users table: %v", err)
		return err
	}

	log.Println("Users table created successfully")
	return nil
}

// Database operations for users

// CreateUser inserts a new user into the database
func CreateUser(user *User) error {
	query := `
		INSERT INTO users (email, password, name)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	err := db.QueryRow(context.Background(), query, user.Email, user.Password, user.Name).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	return err
}

// GetUserByEmail retrieves a user by email
func GetUserByEmail(email string) (*User, error) {
	user := &User{}
	query := `
		SELECT id, email, password, name, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	err := db.QueryRow(context.Background(), query, email).
		Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return user, nil
}

// GetUserByID retrieves a user by ID
func GetUserByID(id int) (*User, error) {
	user := &User{}
	query := `
		SELECT id, email, password, name, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	err := db.QueryRow(context.Background(), query, id).
		Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return user, nil
}

// UpdateUser updates an existing user
func UpdateUser(user *User) error {
	query := `
		UPDATE users
		SET email = $1, password = $2, name = $3, updated_at = CURRENT_TIMESTAMP
		WHERE id = $4
		RETURNING updated_at
	`

	err := db.QueryRow(context.Background(), query, user.Email, user.Password, user.Name, user.ID).
		Scan(&user.UpdatedAt)

	return err
}

// DeleteUser deletes a user by ID
func DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(context.Background(), query, id)
	return err
}
