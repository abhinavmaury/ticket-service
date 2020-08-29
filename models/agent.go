package models

import (
	"abhinav/ticket-service/db"
	"abhinav/ticket-service/utils"
	"context"
	"errors"
)

type Agent struct {
	ID         string `json:"id,omitempty"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	IsLoggedIn bool   `json:"is_logged_in,omitempty"`
}

func IsUsernamePasswordValid(ctx *context.Context, agent Agent) error {
	query := `select username, password, is_login from users where username = ?`
	conn, err := db.UserTicketDB.Conn(*ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	row := conn.QueryRowContext(*ctx, query, agent.Username)
	var ag Agent
	if err := row.Scan(&ag.Username, &ag.Password, ag.IsLoggedIn); err != nil {
		return err
	}
	if ag.IsLoggedIn {
		return errors.New("user is already logged in")
	}
	if err = utils.VerifyPassword(ag.Password, agent.Password); err != nil {
		return errors.New("Invalid Password!")
	}
	return nil
}

func SetUserToActiveState(ctx *context.Context, agent Agent) error {
	query := `UPDATE users set is_online = true where username = ?`
	conn, err := db.UserTicketDB.Conn(*ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.ExecContext(*ctx, query, agent.Username)
	if err != nil {
		return err
	}
	return nil
}

func GetCurrentActiveAgents(ctx context.Context) (int, error) {
	query := `select count(*) from users where user_type = ? and is_online = true`
	conn, err := db.UserTicketDB.Conn(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	row := conn.QueryRowContext(ctx, query, `agent`)
	if err != nil {
		return 0, err
	}
	var count int
	row.Scan(&count)
	return count, nil
}
