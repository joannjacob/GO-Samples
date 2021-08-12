package post

// import (
// 	"context"
// 	"fmt"

// 	models "chi_post_sample/models"
// 	pRepo "chi_post_sample/repository"

// 	"github.com/jinzhu/gorm"
// )

// // NewSQLPostRepo retunrs implement of post repository interface
// func NewSQLPostRepo(Conn *gorm.DB) pRepo.PostRepo {
// 	return &mysqlPostRepo{
// 		Conn: Conn,
// 	}
// }

// type mysqlPostRepo struct {
// 	Conn *gorm.DB
// }

// func (m *mysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Post, error) {
// 	rows, err := m.Conn.Raw(query).Rows()
// 	fmt.Println(rows)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	payload := make([]*models.Post, 0)
// 	for rows.Next() {
// 		data := new(models.Post)

// 		err := rows.Scan(
// 			&data.ID,
// 			&data.Title,
// 			&data.Content,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		payload = append(payload, data)
// 	}
// 	return payload, nil
// }

// func (m *mysqlPostRepo) Fetch(ctx context.Context, num int64) ([]*models.Post, error) {
// 	query := "Select id, title, content From posts limit ?"
// 	fmt.Println("HELLOO")

// 	return m.fetch(ctx, query, num)
// }

// func (m *mysqlPostRepo) GetByID(ctx context.Context, id int64) (*models.Post, error) {
// 	query := "Select id, title, content From posts where id=?"

// 	rows, err := m.fetch(ctx, query, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	payload := &models.Post{}
// 	if len(rows) > 0 {
// 		payload = rows[0]
// 	} else {
// 		return nil, models.ErrNotFound
// 	}

// 	return payload, nil
// }
