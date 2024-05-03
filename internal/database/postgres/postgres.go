package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"testWallet/internal/config"
	"testWallet/internal/models"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}
//NewPGConnection is create postgres driver
func NewPGConnection(cfg *config.Config) *PostgresRepository {
	const op = "internal.database.postgres"

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cfg.Database.User,
																			cfg.Database.Password,
																			cfg.Database.DbName)
	db, err := sql.Open("postgres", connStr)
	if err !=nil{
		panic(err)
	}

	fmt.Println(db.Stats().OpenConnections)

	return &PostgresRepository{
		db: db,
		}
}

func (p *PostgresRepository)Create(userName string) error{	
	var id int
	err := p.db.QueryRow("insert into users (name) values ($1) returning id", userName).Scan(&id)
	if err != nil {
		return err
	}
	if id > 0 {
		if _, err := p.db.Exec("insert into wallets (user_id) values ($1)", id); err != nil{
			return err
		}
		
	}
	return nil
}

func (p *PostgresRepository)Delete(name string) error{
	res, err := p.db.Exec("DELETE FROM users WHERE name=$1", name); 
	if err != nil{
		return err		
	}
	fmt.Println(res)
	

	return nil
}

func (p *PostgresRepository)Update(id string, u models.User) error{
	_, err := p.db.Query("UPDATE balance form wallets where user_id=$1", id)
	if err != nil{
		return err
	}
	return nil
}

func (p *PostgresRepository)GetByName(u *models.User)(*models.User, error){
	
	// selectStr := `select user_id, balance, users.id, users.name from wallets, users where user_id=(
	// 	select id from users where name='Stas4');`
	row := p.db.QueryRow("select * from users where name=$1", strings.ToLower(u.Name))
	
	err := row.Scan(&u.ID, &u.Name, &u.Time)	
	if err != nil{
		return nil, err
	}
	return u, nil	
	}

func (p *PostgresRepository)CloseDB() error{
	const op = "internal.database.postgres"
	err := p.db.Close()
	if err != nil{
		fmt.Println(op, err)
		return err
	}
	return nil
}
