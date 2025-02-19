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
}

func FetchConnections() ([]Instance, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to find database: %v", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT name, instance_id, limite, is_evo FROM instances WHERE maturador = TRUE")
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	var instances []Instance
	for rows.Next() {
		var inst Instance
		if err := rows.Scan(&inst.Name, &inst.InstanceId, &inst.Limit, &inst.IsEvo); err != nil {
			fmt.Fprintf(os.Stderr, "Row scan with issues: %v\n", err)
			continue
		}
		instances = append(instances, inst)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return instances, nil
}
