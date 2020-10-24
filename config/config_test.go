package config

import "testing"

func TestConfig(t *testing.T) {
	server := Config.Database.Server
	if server != "localhost" {
		t.Errorf("server = %s; want localhost", server)
	} else {
		t.Logf("Server %s\n", server)
	}

	user := Config.Database.User
	if user != "lw" {
		t.Errorf("user = %s; want lw", user)
	} else {
		t.Logf("user %s\n", user)
	}

	password := Config.Database.Password
	if password != "lw" {
		t.Errorf("password = %s; want lw", password)
	} else {
		t.Logf("password %s\n", password)
	}

	databaseName := Config.Database.Database
	if databaseName != "stock" {
		t.Errorf("database name = %s; want stock", databaseName)
	} else {
		t.Logf("database name %s\n", databaseName)
	}

	maxPoolsize := Config.ConnectionPool.MaxPoolSize
	if maxPoolsize != 2 {
		t.Errorf("maxPoolsize = %d; want 2", maxPoolsize)
	} else {
		t.Logf("maxPoolsize %d\n", maxPoolsize)
	}
}
