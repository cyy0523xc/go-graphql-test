package main

func init() {
	// 初始化用户
	var user User
	users = make([]User, 0)
	user = User{
		Id:     1,
		Name:   "userA",
		Status: 1,
	}
	users = append(users, user)
	user = User{
		Id:     2,
		Name:   "userB",
		Status: 1,
	}
	users = append(users, user)
	user = User{
		Id:     3,
		Name:   "userC",
		Status: 0,
	}
	users = append(users, user)

	// 初始化books
	books = make([]Book, 0)
	var book Book
	book = Book{
		Id:     1,
		Name:   "bookA",
		UserId: 1,
	}
	books = append(books, book)
	book = Book{
		Id:     2,
		Name:   "bookB",
		UserId: 1,
	}
	books = append(books, book)
	book = Book{
		Id:     3,
		Name:   "bookC",
		UserId: 1,
	}
	books = append(books, book)
	book = Book{
		Id:     4,
		Name:   "bookC",
		UserId: 3,
	}
	books = append(books, book)

	comments = make([]Comment, 0)
	var comment Comment
	comment = Comment{
		Id:      1,
		UserId:  1,
		BookId:  1,
		Content: "content 1",
	}
	comments = append(comments, comment)
	comment = Comment{
		Id:      2,
		UserId:  2,
		BookId:  1,
		Content: "content 2",
	}
	comments = append(comments, comment)
	comment = Comment{
		Id:      3,
		UserId:  3,
		BookId:  2,
		Content: "content 3",
	}
	comments = append(comments, comment)
	comment = Comment{
		Id:      4,
		UserId:  2,
		BookId:  3,
		Content: "content 4",
	}
	comments = append(comments, comment)
	comment = Comment{
		Id:      5,
		UserId:  3,
		BookId:  1,
		Content: "content 5",
	}
	comments = append(comments, comment)
}
