package models


type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func AllBooks() ([]*Book, error) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}

/*

drop table books;
create table books(isbn varchar(256), title varchar(256), author varchar(256), price float);

insert into books values('001','book1','lhy',1);
insert into books values('001','book1','lhy',1);
insert into books values('001','book1','lhy',1);
insert into books values('001','book1','lhy',1);
insert into books values('001','book1','lhy',1);

curl -i localhost:3000/books

*/