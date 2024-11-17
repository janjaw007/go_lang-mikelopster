package main

// Connection Database Successful
// fmt.Println("Connection Database Successful")

// ##### SELECT
// product, err := getProduct(2)

// fmt.Println("Get Successful!", product)

// ##### SELECT Mutiple
// products, err := getProducts()
// fmt.Println("Get Mutiple Successful!", products)

// Select Product with Supplier
// productWithSupplier, err := getProductAndSupplier(3)
// fmt.Println("GET Product with Supplier Successful!", productWithSupplier)

// Select Mutiple with Product with Supplier
// productsWithSupplier, err := getProductsAndSupplier()
// fmt.Println("Get Mutiple ProductSupplier Successful!", productsWithSupplier)

// ##### Create
// err = createProduct(&Product{Name: "Go product 6", Price: 423, Supplier_id: 1})
// if err != nil {
// 	log.Fatal(err)
// }

// fmt.Println("Create Successful")

// ##### Update
// product, err := updateProduct(1, &Product{Name: "UpdateTestProduct1", Price: 124})
// if err != nil {
// 	log.Fatal(err)
// }

// fmt.Println("Update Successful", product)

// ##### Delete
// product, err := deleteProduct(8)
// if err != nil {
// 	log.Fatal(err)
// }

// fmt.Println("Delete Successful", product)

func createProduct(product *Product) error {
	_, err := db.Exec(
		"INSERT INTO public.products(name, price) VALUES ($1, $2);",
		product.Name,
		product.Price,
	)
	return err
}

func getProduct(id int) (Product, error) {
	var p Product
	row := db.QueryRow("SELECT id, name, price FROM products where id=$1;", id)

	err := row.Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return Product{}, err
	}

	return p, nil
}

func getProducts() ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price from products")

	if err != nil {
		return nil, err
	}

	var products []Product

	for rows.Next() {
		var p Product

		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func getProductAndSupplier(id int) (ProductWithSupplier, error) {
	var p ProductWithSupplier

	row := db.QueryRow(`
	SELECT p.id, p.name, p.price, s.name
	FROM products p
	INNER JOIN suppliers s ON p.supplier_id = s.id
	WHERE p.id = $1;`, id)

	err := row.Scan(&p.ProductID, &p.ProductName, &p.ProductPrice, &p.SupplierName)

	if err != nil {
		return ProductWithSupplier{}, err
	}

	return p, nil

}

func getProductsAndSupplier() ([]ProductWithSupplier, error) {
	// Query Mutiple Rows
	rows, err := db.Query(`
	SELECT p.id, p.name, p.price, s.name
	FROM products p
	INNER JOIN suppliers s ON p.supplier_id = s.id
	`)
	// handles if query error
	if err != nil {
		return nil, err
	}

	// create Slice
	var productWithSupplier []ProductWithSupplier

	// loop through rows
	for rows.Next() {
		var p ProductWithSupplier
		//map with type
		err := rows.Scan(&p.ProductID, &p.ProductName, &p.ProductPrice, &p.SupplierName)
		// check if map type error
		if err != nil {
			return nil, err
		}
		//append rows into slice
		productWithSupplier = append(productWithSupplier, p)
	}

	if rows.Err(); err != nil {
		return nil, err
	}

	return productWithSupplier, nil

}

func updateProduct(id int, product *Product) (Product, error) {
	var p Product

	row := db.QueryRow(
		"UPDATE public.products SET name=$1, price=$2 WHERE id = $3 RETURNING id, name, price;",
		product.Name,
		product.Price,
		id,
	)

	err := row.Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return Product{}, err
	}

	return p, err
}

func deleteProduct(id int) (Product, error) {
	var p Product

	row := db.QueryRow("DELETE FROM public.products WHERE id = $1 RETURNING id, name, price;", id)

	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}

	return p, err
}
