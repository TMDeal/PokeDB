package db

import "log"

type Type struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type TypeEfficacy struct {
	From        []Type
	FromMult    float32
	Against     []Type
	AgainstMult float32
}

func (db DB) FindTypeByID(id int) (*Type, error) {
	var t Type

	err := db.QueryRowx(`
	select id, name from types where id = $1
	`, id).StructScan(&t)

	if err != nil {
		log.Println("Unable to execute query!")
		return nil, err
	}

	return &t, nil
}

func (db DB) GetTypeWeaknessList(t Type) (*TypeEfficacy, error) {
	var from []Type
	var against []Type

	rows, err := db.Queryx(`
	select t2.id, t2.name
	from types as t1, types as t2, type_efficacies as te
	where
	t1.id = te.type_id
	and t2.id = te.target_type_id
	and te.damage_multiplier = 0.5
	and t1.name = $1
	order by t2.id
	`, t.Name)
	if err != nil {
		log.Println("Unable to execute query")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Type
		err = rows.StructScan(&t)
		if err != nil {
			log.Println("Unable to marshal into struct!")
			return nil, err
		}
		against = append(against, t)
	}

	rows, err = db.Queryx(`
	select t2.id, t2.name
	from types as t1, types as t2, type_efficacies as te
	where
	t2.id = te.type_id
	and t1.id = te.target_type_id
	and te.damage_multiplier = 2
	and t1.name = $1
	order by t2.id
	`, t.Name)
	if err != nil {
		log.Println("Unable to execute query")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Type
		err = rows.StructScan(&t)
		if err != nil {
			log.Println("Unable to marshal into struct!")
			return nil, err
		}
		from = append(from, t)
	}

	return &TypeEfficacy{
		From:        from,
		FromMult:    2,
		Against:     against,
		AgainstMult: 0.5,
	}, nil
}
