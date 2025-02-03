package models

// id uuid primary key default uuid_generate_v4(),
// name text not null,
// password text not null,
// specialisation text not null,
// phone text,
// email text,
// dob date,
// address text

type Doctor struct {
	Person      Person        `json:"Person"`
	Departments []*Department `json:"department"`
}
