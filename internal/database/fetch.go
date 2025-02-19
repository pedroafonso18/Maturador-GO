package database

import (
	"context"
	"fmt"
	"os"
)

type Instance struct {
	Name       string
	InstanceId string
	Limit      uint
	IsEvo      bool
	Numero     string
}

func FetchConnections() ([]Instance, error) {
	conn, err := GetConnection()
	if err != nil {
		fmt.Println("Database not found, attempting to create...")
		if createErr := createDatabase(); createErr != nil {
			return nil, fmt.Errorf("failed to create database: %v", createErr)
		}
		conn, err = GetConnection()
		if err != nil {
			return nil, fmt.Errorf("failed to establish connection after database creation: %v", err)
		}
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT name, instance_id, limite, is_evo, numero FROM instances WHERE maturador = TRUE")
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	var instances []Instance
	for rows.Next() {
		var inst Instance
		if err := rows.Scan(&inst.Name, &inst.InstanceId, &inst.Limit, &inst.IsEvo, &inst.Numero); err != nil {
			fmt.Fprintf(os.Stderr, "Row scan error: %v\n", err)
			continue
		}
		instances = append(instances, inst)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return instances, nil
}

func createDatabase() error {

	conn, err := GetConnection()
	if err != nil {
		return fmt.Errorf("failed to get admin connection: %v", err)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "CREATE DATABASE maturador")
	if err != nil {
		return fmt.Errorf("failed to execute create database: %v", err)
	}

	fmt.Println("Database created successfully.")
	return nil
}
